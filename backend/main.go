package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)


var tasks []Task

func main() {
	
	loadTasks()
	//fmt.Printf("Nmr de tarefas carregadas: %d\n", len(tasks)) DEBUG

	router := mux.NewRouter()

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
}

//cria novo ID
func findNextID() int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

//funções de persistencia do json
func loadTasks() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		// cria um novo json caso não exista
		//fmt.Println("Arquivo não encontrado, criando novo...") DEBUG
		tasks = []Task{}
		return
	}

	var store TaskStore
	json.Unmarshal(file, &store)
	tasks = store.Tasks
	//fmt.Printf("Tarefas no arquivo: %+v\n", tasks) DEBUG
	fmt.Printf("RODANDO BABY 😎")
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