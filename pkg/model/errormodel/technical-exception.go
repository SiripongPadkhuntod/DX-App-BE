package errormodel

import (
	"errors"
	"fmt"
)

type TechnicalError struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Inner       error       `json:"-"`
	StackTrace  []TraceInfo `json:"-"`
}

func (e *TechnicalError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Description)
}

func (e *TechnicalError) GetStackTrace() string {
	stackTrace := ""
	for _, trace := range e.StackTrace {
		stackTrace += fmt.Sprintf("%s\n\t%s:%d\n", trace.Name, trace.File, trace.Line)
	}
	return stackTrace
}

func (e *TechnicalError) Unwrap() error {
	return e.Inner
}

func WrapTechnicalError(e error) *TechnicalError {
	return &TechnicalError{
		Code:        GenericErrorCode,
		Description: e.Error(),
		Inner:       e,
		StackTrace:  getStackTrace(),
	}
}

func NewTechnicalError(message string, err error) *TechnicalError {
	return &TechnicalError{
		Code:        GenericErrorCode,
		Description: message,
		Inner:       err,
		StackTrace:  getStackTrace(),
	}
}

func RaiseTechnicalError(message string) *TechnicalError {
	return &TechnicalError{
		Code:        GenericErrorCode,
		Description: message,
		Inner:       errors.New(message),
		StackTrace:  getStackTrace(),
	}
}

// Deprecated: use RaiseTechnicalError instead
func RaiseTechnicalgRPCError(code string, message string, err error) *TechnicalError {
	return RaiseTechnicalError(message)
}
