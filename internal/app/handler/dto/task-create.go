package handlerdto

import modelV1 "github.com/youruser/dexter-transport/pkg/model/v1"

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	modelV1.Task
}
