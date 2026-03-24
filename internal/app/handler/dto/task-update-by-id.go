package handlerdto

import modelV1 "github.com/youruser/dexter-transport/pkg/model/v1"

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTaskResponse struct {
	modelV1.Task
}
