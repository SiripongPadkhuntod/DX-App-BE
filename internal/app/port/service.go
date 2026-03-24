package port

import (
	"context"

	handlerdto "github.com/youruser/dexter-transport/internal/app/handler/dto"
)

type Service interface {
	Health(ctx context.Context) (id string, service string, err error)

	// Task CRUD examples
	CreateTask(ctx context.Context, req handlerdto.CreateTaskRequest) (*handlerdto.TaskResponse, error)
	GetTask(ctx context.Context, id int) (*handlerdto.TaskResponse, error)
	ListTasks(ctx context.Context) (*handlerdto.TaskListResponse, error)
	UpdateTask(ctx context.Context, id int, req handlerdto.UpdateTaskRequest) (*handlerdto.TaskResponse, error)
	DeleteTask(ctx context.Context, id int) error
}
