package service

import (
	"context"

	"github.com/nightmaker00/go-tasks-api/internal/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, title string, description *string, status string) (int64, error)
	GetByID(ctx context.Context, id int64) (*domain.Task, error)
	Update(ctx context.Context, id int64, title string, description *string, status string) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	List(ctx context.Context, status string, limit, offset int) ([]domain.TaskListItem, error)
}
