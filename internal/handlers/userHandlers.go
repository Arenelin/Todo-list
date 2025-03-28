package handlers

import (
	"context"
	"github.com/Arenelin/Todo-list/internal/userService"
	"github.com/Arenelin/Todo-list/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) GetUsersUserIdTasks(_ context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	idParam := request.UserId

	allUserTasks, err := h.Service.GetTasksByUserID(idParam)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range allUserTasks {
		task := users.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	idParam := request.Id

	userRequest := request.Body

	userToUpdate := userService.UserUpdate{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	returnedUser, err := h.Service.UpdateUserById(idParam, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &returnedUser.ID,
		Email:    &returnedUser.Email,
		Password: &returnedUser.Password,
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	idParam := request.Id

	err := h.Service.DeleteUser(idParam)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUsersId204Response{}

	return response, nil
}
