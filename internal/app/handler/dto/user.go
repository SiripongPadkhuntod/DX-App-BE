package handlerdto

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
	Name  string `json:"name"`
}
