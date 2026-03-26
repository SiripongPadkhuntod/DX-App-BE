package servicesconstant

type SuccessMessage string

const (
	SUCCESS_MESSAGE_VALUE SuccessMessage = "The operation was successful"
	HEALTHY_STATUS        SuccessMessage = "healthy"
	HEALTHY_MESSAGE       SuccessMessage = "Service is running smoothly"

	SUCCESS_MESSAGE_SUCCESS        SuccessMessage = "successfully"
	SUCCESS_MESSAGE_CREATED        SuccessMessage = "successfully created"
	SUCCESS_MESSAGE_UPDATED        SuccessMessage = "successfully updated"
	SUCCESS_MESSAGE_DELETED        SuccessMessage = "successfully deleted"
	SUCCESS_MESSAGE_RESET_PASSWORD SuccessMessage = "A password reset email has already been sent. Please check your inbox and follow the instructions to reset your password."
	SUCCESS_MESSAGE_LOGOUT         SuccessMessage = "Logout Successfully"
)

const (
	SUCCESS_STATUS = "success"
	SUCCESS_CODE   = "0000"
)
