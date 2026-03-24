package servicesconstant

const (
	ErrorStatus     = "error"
	UnhealthyStatus = "unhealthy"

	// Error Codes
	InternalServerErrorCode = "E500"
	InvalidRequestCode      = "E400"
	ResourceNotFoundCode    = "E404"
	DatabaseErrorCode       = "E501"

	// Error Messages
	InternalServerErrorMessage = "Internal server error"
	InvalidRequestBodyMessage  = "Invalid request body"
	ResourceNotFoundMessage    = "Resource not found"
	DatabaseErrorMessage       = "Database error"
)
