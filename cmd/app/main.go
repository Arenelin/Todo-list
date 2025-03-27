package main

import (
	"github.com/Arenelin/Todo-list/internal/database"
	"github.com/Arenelin/Todo-list/internal/handlers"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/Arenelin/Todo-list/internal/userService"
	"github.com/Arenelin/Todo-list/internal/web/tasks"
	"github.com/Arenelin/Todo-list/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":9090"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
