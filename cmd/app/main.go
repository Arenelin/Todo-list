package main

import (
	"github.com/Arenelin/Todo-list/internal/database"
	"github.com/Arenelin/Todo-list/internal/handlers"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks/{id:[0-9]+}", handler.UpdateTaskByIdHandler).Methods(http.MethodPatch)
	r.HandleFunc("/api/tasks/{id:[0-9]+}", handler.DeleteTaskByIdHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
