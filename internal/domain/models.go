package domain

import "time"

const (
	TaskStatusNew        = "new"
	TaskStatusInProgress = "in_progress"
	TaskStatusDone       = "done"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskListItem struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
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
	ID int64 `json:"id"`
}

type UpdateTaskResponse struct {
	Status string `json:"status"`
}
