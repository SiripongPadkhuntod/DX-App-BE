package handlerdto

import modelV1 "github.com/youruser/dexter-transport/pkg/model/v1"

type TaskListResponse struct {
	Tasks []modelV1.Task `json:"tasks"`
}
