package service

import (
	"context"

	"github.com/youruser/dexter-transport/internal/app/domain"
	handlerdto "github.com/youruser/dexter-transport/internal/app/handler/dto"
)

func (s *service) CreateTask(ctx context.Context, req handlerdto.CreateTaskRequest) (*handlerdto.TaskResponse, error) {
	t := &domain.Task{
		Title:       req.Title,
		Description: req.Description,
	}
	res, err := s.repo.Sql.CreateTask(ctx, t)
	if err != nil {
		return nil, err
	}
	return mapTaskDomainToDTO(res), nil
}

func (s *service) GetTask(ctx context.Context, id int) (*handlerdto.TaskResponse, error) {
	res, err := s.repo.Sql.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapTaskDomainToDTO(res), nil
}

func (s *service) ListTasks(ctx context.Context) (*handlerdto.TaskListResponse, error) {
	tasks, err := s.repo.Sql.ListTasks(ctx)
	if err != nil {
		return nil, err
	}
	
	dtoTasks := make([]handlerdto.TaskResponse, len(tasks))
	for i, t := range tasks {
		dtoTasks[i] = *mapTaskDomainToDTO(&t)
	}
	return &handlerdto.TaskListResponse{Tasks: dtoTasks}, nil
}

func (s *service) UpdateTask(ctx context.Context, id int, req handlerdto.UpdateTaskRequest) (*handlerdto.TaskResponse, error) {
	// First get current task to handle partial updates
	current, err := s.repo.Sql.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		current.Title = req.Title
	}
	if req.Description != "" {
		current.Description = req.Description
	}
	if req.Status != "" {
		current.Status = req.Status
	}

	res, err := s.repo.Sql.UpdateTask(ctx, current)
	if err != nil {
		return nil, err
	}
	return mapTaskDomainToDTO(res), nil
}

func (s *service) DeleteTask(ctx context.Context, id int) error {
	return s.repo.Sql.DeleteTask(ctx, id)
}

func mapTaskDomainToDTO(d *domain.Task) *handlerdto.TaskResponse {
	return &handlerdto.TaskResponse{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
