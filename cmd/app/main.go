package main

import (
	"github.com/Arenelin/Todo-list/internal/database"
	"github.com/Arenelin/Todo-list/internal/handlers"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/Arenelin/Todo-list/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":9090"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
