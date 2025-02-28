package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

var task string

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello, %s", task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskInput RequestBody

	err := json.NewDecoder(r.Body).Decode(&taskInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task = taskInput.Task

	response := "Task: " + task + " successfully created"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/task", HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/create-task", CreateTaskHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
