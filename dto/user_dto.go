package dto

import "time"

type CreateUserRequest struct {
	Name string `form:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name *string `json:"name"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
