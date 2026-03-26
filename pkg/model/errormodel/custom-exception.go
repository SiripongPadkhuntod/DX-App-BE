package errormodel

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Status      int    `json:"-"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Description)
}

func BadRequestError(err error) *CustomError {
	return &CustomError{
		Status:      http.StatusBadRequest,
		Code:        InvalidRequest,
		Description: err.Error(),
	}
}

func BadRequestErrorCdp(err error) *CustomError {
	return &CustomError{
		Status:      http.StatusBadRequest,
		Code:        InvalidRequestCdp,
		Description: err.Error(),
	}
}

func UnauthorizedError(err error) *CustomError {
	return &CustomError{
		Status:      http.StatusUnauthorized,
		Code:        Unauthorized,
		Description: err.Error(),
	}
}

func AuthenticationError() *CustomError {
	return &CustomError{
		Status:      http.StatusForbidden,
		Code:        Forbidden,
		Description: ForbiddenMsgEn,
	}
}

func InvalidJwtError() *CustomError {
	return &CustomError{
		Status:      http.StatusForbidden,
		Code:        InvalidJwt,
		Description: InvalidJwtMsgEn,
	}
}
