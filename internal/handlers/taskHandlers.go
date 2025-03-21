package handlers

import (
	"context"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/Arenelin/Todo-list/internal/web/tasks"
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

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	idParam := request.Id

	taskRequest := request.Body

	taskToUpdate := taskService.TaskUpdate{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}

	returnedTask, err := h.Service.UpdateTaskById(idParam, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &returnedTask.ID,
		Task:   &returnedTask.Task,
		IsDone: &returnedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	idParam := request.Id

	err := h.Service.DeleteTask(idParam)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId204Response{}

	return response, nil
}
