package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusNew        TaskStatus = "new"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskListItem struct {
	ID     uuid.UUID  `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}

type CreateTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateTaskResponse struct {
	Status string `json:"status"`
}
