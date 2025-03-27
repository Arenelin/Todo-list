package userService

import (
	"github.com/Arenelin/Todo-list/internal/web/tasks"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []tasks.Task
}

type UserUpdate struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
