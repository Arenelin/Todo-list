package userService

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUserById(id uint, user UserUpdate) (User, error)
	DeleteUserById(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUserById(id uint, user UserUpdate) (User, error) {
	var returnedUser User

	query := r.db.Model(&User{}).
		Clauses(clause.Returning{}).
		Where("id = ?", id)

	if user.Email != nil {
		query = query.Update("email", *user.Email)
	}
	if user.Password != nil {
		query = query.Update("password", *user.Password)
	}

	result := query

	if result.Error != nil || result.RowsAffected == 0 {
		return User{}, result.Error
	}

	err := result.Scan(&returnedUser).Error
	return returnedUser, err
}

func (r *userRepository) DeleteUserById(id uint) error {
	err := r.db.Delete(&User{}, id).Error
	return err
}
