package handlers

import (
	"encoding/json"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, _ *http.Request) {
	tasks, err := h.Service.GetTasks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdTask, err := h.Service.CreateTask(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) UpdateTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	formattedId, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid id sent"})
		return
	}

	var updatedTask taskService.TaskUpdate

	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedTask, err := h.Service.UpdateTaskById(uint(formattedId), updatedTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnedTask)
}

func (h *Handler) DeleteTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	formattedId, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid id sent"})
		return
	}

	err = h.Service.DeleteTask(uint(formattedId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
