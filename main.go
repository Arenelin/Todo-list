package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func GetTasksHandler(w http.ResponseWriter, _ *http.Request) {
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskInput Task

	err := json.NewDecoder(r.Body).Decode(&taskInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := DB.Create(&taskInput).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Task successfully created"})
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	r := mux.NewRouter()

	r.HandleFunc("/api/tasks", GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/create-task", CreateTaskHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
