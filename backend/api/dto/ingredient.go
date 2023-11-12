package dto

import (
	"strconv"

	"github.com/alaust/foodiee/backend/internal/entities"
)

type IngredientResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

type IngredientPost struct {
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

func ToIngredientResponse(ing *entities.Ingredient) *IngredientResponse {
	return &IngredientResponse{
		Id:     "ing_" + strconv.Itoa(ing.Id),
		Name:   ing.Name,
		Unit:   ing.Unit,
		Amount: ing.Ammount,
	}
}
