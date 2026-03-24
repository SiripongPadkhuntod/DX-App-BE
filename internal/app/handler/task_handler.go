package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerdto "github.com/youruser/dexter-transport/internal/app/handler/dto"
	"github.com/youruser/dexter-transport/pkg/v1/dto"
)

// CreateTask godoc
// @Summary      Create a new task
// @Description  Create a new task with title and description
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        request body handlerdto.CreateTaskRequest true "Create Task Request"
// @Success      201 {object} handlerdto.TaskResponse
// @Router       /api/v1/tasks [post]
func (h *handler) CreateTask(c *gin.Context) func(ctx context.Context, req handlerdto.CreateTaskRequest) (*handlerdto.TaskResponse, error) {
	return func(ctx context.Context, req handlerdto.CreateTaskRequest) (*handlerdto.TaskResponse, error) {
		return h.svc.CreateTask(ctx, req)
	}
}

// GetTask godoc
// @Summary      Get a task by ID
// @Description  Get detailed information about a specific task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Success      200 {object} handlerdto.TaskResponse
// @Router       /api/v1/tasks/{id} [get]
func (h *handler) GetTask(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskResponse, error) {
	return func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskResponse, error) {
		id, _ := strconv.Atoi(c.Param("id"))
		return h.svc.GetTask(ctx, id)
	}
}

// ListTasks godoc
// @Summary      List all tasks
// @Description  Get a list of all tasks
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200 {object} handlerdto.TaskListResponse
// @Router       /api/v1/tasks [get]
func (h *handler) ListTasks(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskListResponse, error) {
	return func(ctx context.Context, _ dto.EmptyStruct) (*handlerdto.TaskListResponse, error) {
		return h.svc.ListTasks(ctx)
	}
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update a task's title, description, or status
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Param        request body handlerdto.UpdateTaskRequest true "Update Task Request"
// @Success      200 {object} handlerdto.TaskResponse
// @Router       /api/v1/tasks/{id} [put]
func (h *handler) UpdateTask(c *gin.Context) func(ctx context.Context, req handlerdto.UpdateTaskRequest) (*handlerdto.TaskResponse, error) {
	return func(ctx context.Context, req handlerdto.UpdateTaskRequest) (*handlerdto.TaskResponse, error) {
		id, _ := strconv.Atoi(c.Param("id"))
		return h.svc.UpdateTask(ctx, id, req)
	}
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Remove a task from the system
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Success      204 "No Content"
// @Router       /api/v1/tasks/{id} [delete]
func (h *handler) DeleteTask(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*dto.EmptyStruct, error) {
	return func(ctx context.Context, _ dto.EmptyStruct) (*dto.EmptyStruct, error) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := h.svc.DeleteTask(ctx, id); err != nil {
			return nil, err
		}
		return &dto.EmptyStruct{}, nil
	}
}
