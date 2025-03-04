package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func GetTasksHandler(w http.ResponseWriter, _ *http.Request) {
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskInput)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	formattedId, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid id sent"})
		return
	}

	var updatedTask TaskUpdate
	var returnedTask Task

	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := DB.Model(&Task{}).
		Clauses(clause.Returning{}).
		Where("id = ?", formattedId)

	if updatedTask.Task != nil {
		query = query.Update("task", *updatedTask.Task)
	}
	if updatedTask.IsDone != nil {
		query = query.Update("is_done", *updatedTask.IsDone)
	}

	result := query

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{Message: "task with the specified id does not exist"})
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{Message: "task with the specified id does not exist"})
		return
	}

	if err := result.Scan(&returnedTask).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "failed to retrieve updated task"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnedTask)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	formattedId, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid id sent"})
		return
	}

	if err := DB.Delete(&Task{}, formattedId).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{Message: "task with the specified id does not exist"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	r := mux.NewRouter()

	r.HandleFunc("/api/tasks", GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks", CreateTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks/{id:[0-9]+}", UpdateTaskHandler).Methods(http.MethodPatch)
	r.HandleFunc("/api/tasks/{id:[0-9]+}", DeleteTaskHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
