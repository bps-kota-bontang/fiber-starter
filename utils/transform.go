package utils

import (
	"github.com/bps-kota-bontang/fiber-starter/dto"
	"github.com/bps-kota-bontang/fiber-starter/models"
)

func ToResponses[T any, R any](items []T, convertFunc func(*T) R) []R {
	if len(items) == 0 {
		return []R{} // Explicitly return an empty slice
	}

	responses := make([]R, 0, len(items))
	for i := range items {
		responses = append(responses, convertFunc(&items[i]))
	}
	return responses
}

func ToUserResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
