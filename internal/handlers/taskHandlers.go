package handlers

import (
	"context"
	"github.com/Arenelin/Todo-list/internal/taskService"
	"github.com/Arenelin/Todo-list/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) GetTasks(_ context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	taskRequest := request.Body

	allTasks, err := h.Service.GetTasks(taskRequest.UserId)
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
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
		UserId: &returnedTask.UserID,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	idParam := request.Id

	err := h.Service.DeleteTask(idParam)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId204Response{}

	return response, nil
}
