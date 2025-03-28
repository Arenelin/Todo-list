package userService

import "github.com/Arenelin/Todo-list/internal/taskService"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetTasksByUserID(id uint) ([]taskService.Task, error) {
	return s.repo.GetTasksByUserId(id)
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUserById(id uint, user UserUpdate) (User, error) {
	return s.repo.UpdateUserById(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUserById(id)
}
