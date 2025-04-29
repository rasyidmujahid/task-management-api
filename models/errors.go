package models

import "errors"

var (
	ErrEmptyTitle         = errors.New("title cannot be empty")
	ErrTitleTooLong       = errors.New("title cannot be longer than 100 characters")
	ErrDescriptionTooLong = errors.New("description cannot be longer than 500 characters")
	ErrInvalidStatus      = errors.New("status must be one of: pending, in_progress, completed")
	ErrTaskNotFound       = errors.New("task not found")
) 