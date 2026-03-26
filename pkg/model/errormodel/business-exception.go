package errormodel

import "fmt"

type BusinessError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Description)
}

func RaiseBusinessError(code, description string) *BusinessError {
	return &BusinessError{
		Code:        code,
		Description: description,
	}
}
