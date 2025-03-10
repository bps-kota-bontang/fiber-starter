package utils

import (
	"github.com/bps-kota-bontang/fiber-starter/dto"
	"github.com/bps-kota-bontang/fiber-starter/models"
)

func ToUserResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []models.User) []dto.UserResponse {
	if len(users) == 0 {
		return []dto.UserResponse{}
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}
