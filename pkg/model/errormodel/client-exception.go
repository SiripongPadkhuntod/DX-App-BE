package errormodel

type ClientError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func ClientErrorDefaultCode(description string) *ClientError {
	return &ClientError{
		Code:        GenericErrorCode,
		Description: description,
	}
}

func (e *ClientError) Error() string {
	return e.Description
}
