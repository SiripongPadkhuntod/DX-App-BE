package dto

import servicesconstant "github.com/youruser/dexter-transport/internal/constant"

type EmptyStruct struct{}

type SuccessResponse struct {
	Code    string                          `json:"code" default:"0000"`
	Message servicesconstant.SuccessMessage `json:"description"`
}

type ErrorResponse struct {
	Code    string `json:"code" default:"E500"`
	Message string `json:"description"`
}

type BaseResponse[T any] struct {
	Code    string `json:"code"`
	Message string `json:"description"`
	Data    T      `json:"data,omitempty"`
}
