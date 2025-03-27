package taskService

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	GetTasks(userId *uint) ([]Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTaskById(id uint, task TaskUpdate) (Task, error)
	DeleteTaskById(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(userId *uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) UpdateTaskById(id uint, task TaskUpdate) (Task, error) {
	var returnedTask Task

	query := r.db.Model(&Task{}).
		Clauses(clause.Returning{}).
		Where("id = ?", id)

	if task.Task != nil {
		query = query.Update("task", *task.Task)
	}
	if task.IsDone != nil {
		query = query.Update("is_done", *task.IsDone)
	}

	result := query

	if result.Error != nil || result.RowsAffected == 0 {
		return Task{}, result.Error
	}

	err := result.Scan(&returnedTask).Error
	return returnedTask, err
}

func (r *taskRepository) DeleteTaskById(id uint) error {
	err := r.db.Delete(&Task{}, id).Error
	return err
}
