package service

import (
	"context"
	"errors"
	"strings"

	"github.com/nightmaker00/go-tasks-api/internal/domain"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrInvalidStatus = errors.New("invalid status")
	ErrInvalidTitle  = errors.New("invalid title")
	ErrInvalidLimit  = errors.New("invalid limit")
	ErrInvalidOffset = errors.New("invalid offset")
	maxListLimit     = 1000
	defaultListLimit = 100
)

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(ctx context.Context, title string, description string) (int64, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return 0, ErrInvalidTitle
	}

	desc := normalizeDescription(description)
	return s.repo.Create(ctx, title, desc, domain.TaskStatusNew)
}

func (s *taskService) GetByID(ctx context.Context, id int64) (*domain.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (s *taskService) Update(ctx context.Context, id int64, title string, description *string, status string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrInvalidTitle
	}
	if !isValidStatus(status) {
		return ErrInvalidStatus
	}

	desc := normalizeDescriptionPtr(description)
	updated, err := s.repo.Update(ctx, id, title, desc, status)
	if err != nil {
		return err
	}
	if !updated {
		return ErrTaskNotFound
	}
	return nil
}

func (s *taskService) Delete(ctx context.Context, id int64) error {
	deleted, err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrTaskNotFound
	}
	return nil
}

func (s *taskService) List(ctx context.Context, status string, limit, offset int) ([]domain.TaskListItem, error) {
	if status != "" && !isValidStatus(status) {
		return nil, ErrInvalidStatus
	}
	if limit == 0 {
		limit = defaultListLimit
	}
	if limit < 0 || limit > maxListLimit {
		return nil, ErrInvalidLimit
	}
	if offset < 0 {
		return nil, ErrInvalidOffset
	}

	return s.repo.List(ctx, status, limit, offset)
}

func isValidStatus(status string) bool {
	switch status {
	case domain.TaskStatusNew, domain.TaskStatusInProgress, domain.TaskStatusDone:
		return true
	default:
		return false
	}
}

func normalizeDescription(description string) *string {
	description = strings.TrimSpace(description)
	if description == "" {
		return nil
	}
	return &description
}

func normalizeDescriptionPtr(description *string) *string {
	if description == nil {
		return nil
	}
	desc := strings.TrimSpace(*description)
	if desc == "" {
		return nil
	}
	return &desc
}
