package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	"strconv"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)


var tasks []Task

func main() {
	
	loadTasks()
	//fmt.Printf("Nmr de tarefas carregadas: %d\n", len(tasks)) 

	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

    c := cors.New(cors.Options{
            AllowedOrigins:   []string{"*"},
            AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            AllowedHeaders:   []string{"Content-Type", "Authorization"},
            AllowCredentials: true,
        })

    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":8080", handler))
	
	nextID := findNextID()
	//fmt.Printf("Próximo ID: %d\n", nextID)
	
	newTask := Task{
		ID:        nextID,
		Titulo:    "Teste persistência",
		Descricao: "Verificar se JSON está funcionando",
		Status:    "todo",
	}
	tasks = append(tasks, newTask)
	
	saveTasks()
	//fmt.Println("Tarefa salva")
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func findNextID() int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func loadTasks() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Arquivo não encontrado, criando novo...")
		tasks = []Task{}
		return
	}

	var store TaskStore
	json.Unmarshal(file, &store)
	tasks = store.Tasks
	fmt.Printf("Tarefas no arquivo: %+v\n", tasks)
}

func saveTasks() {
	store := TaskStore{Tasks: tasks}
	
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar:", err)
		return
	}
	
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var newTask Task
    err := json.NewDecoder(r.Body).Decode(&newTask)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    
    if newTask.Titulo == "" {
        http.Error(w, "Título é obrigatório", http.StatusBadRequest)
        return
    }
    
    validStatus := map[string]bool{
        "todo":       true,
        "inprogress": true,
        "done":       true,
    }
    if !validStatus[newTask.Status] {
        http.Error(w, "Status inválido. Use: todo, inprogress ou done", http.StatusBadRequest)
        return
    }
    
    newTask.ID = findNextID()
    tasks = append(tasks, newTask)
    
    saveTasks()
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }
    
    var updatedTask Task
    err = json.NewDecoder(r.Body).Decode(&updatedTask)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    
    if updatedTask.Titulo == "" {
        http.Error(w, "Título é obrigatório", http.StatusBadRequest)
        return
    }
    
    validStatus := map[string]bool{
        "todo":       true,
        "inprogress": true,
        "done":       true,
    }
    if !validStatus[updatedTask.Status] {
        http.Error(w, "Status inválido", http.StatusBadRequest)
        return
    }
    
    for i, task := range tasks {
        if task.ID == id {
            updatedTask.ID = id 
            tasks[i] = updatedTask
            saveTasks()
            json.NewEncoder(w).Encode(updatedTask)
            return
        }
    }
    
    http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}

// Handler DELETE /tasks/{id}
func deleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }
    
    for i, task := range tasks {
        if task.ID == id {
            
            tasks = append(tasks[:i], tasks[i+1:]...)
            saveTasks()
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    
    http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}
