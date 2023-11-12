package dto

import "github.com/alaust/foodiee/backend/internal/entities"

type UserResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func ToUserResponse(u *entities.User) *UserResponse {
	return &UserResponse{
		Id:      "u_" + u.Id,
		Name:    u.Name,
		Picture: u.Picture,
	}
}
