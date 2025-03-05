package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TaskUpdate struct {
	Task   *string `json:"task,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
}
