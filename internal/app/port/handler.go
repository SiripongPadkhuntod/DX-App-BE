package port

import (
	"context"

	"github.com/gin-gonic/gin"
	handlerdto "github.com/youruser/dexter-transport/internal/app/handler/dto"
	"github.com/youruser/dexter-transport/pkg/v1/dto"
)

type Handler interface {

	// health check
	Health(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.HealthResponse, error)

	// Task CRUD examples
	CreateTask(c *gin.Context) func(ctx context.Context, req handlerdto.CreateTaskRequest) (*handlerdto.TaskResponse, error)
	GetTask(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskResponse, error)
	ListTasks(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskListResponse, error)
	UpdateTask(c *gin.Context) func(ctx context.Context, req handlerdto.UpdateTaskRequest) (*handlerdto.TaskResponse, error)
	DeleteTask(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*dto.EmptyStruct, error)
}
