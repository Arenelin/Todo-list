package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
