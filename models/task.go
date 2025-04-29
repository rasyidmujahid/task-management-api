package models

import (
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskInput struct {
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=500"`
	Status      string `json:"status" validate:"required,oneof=pending in_progress completed"`
}

func (t *TaskInput) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	if len(t.Title) > 100 {
		return ErrTitleTooLong
	}
	if len(t.Description) > 500 {
		return ErrDescriptionTooLong
	}
	if t.Status != "pending" && t.Status != "in_progress" && t.Status != "completed" {
		return ErrInvalidStatus
	}
	return nil
} 