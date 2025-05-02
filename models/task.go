package models

import (
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	Assignee    string    `json:"assignee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskInput struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	Status      string    `json:"status" validate:"required,oneof=pending in_progress completed"`
	Priority    string    `json:"priority" validate:"oneof=low medium high"`
	DueDate     time.Time `json:"due_date"`
	Category    string    `json:"category" validate:"max=50"`
	Tags        []string  `json:"tags" validate:"max=5"`
	Assignee    string    `json:"assignee" validate:"max=100"`
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
	if t.Priority != "" && t.Priority != "low" && t.Priority != "medium" && t.Priority != "high" {
		return ErrInvalidPriority
	}
	if len(t.Tags) > 5 {
		return ErrTooManyTags
	}
	return nil
} 