package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GET HANDLER
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

//POST HANDLER
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

//PUT HANDLER
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

//DELETE HANDLER
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