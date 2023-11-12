package dto

import (
	"strconv"

	"github.com/alaust/foodiee/backend/internal/entities"
)

type RecipeResponse struct {
	Id              string               `json:"id"`
	Name            string               `json:"name"`
	Author          string               `json:"author"`
	Thumbnail       string               `json:"thumbnail"`
	PreperationTime int                  `json:"preperationTime"`
	Ingredients     []IngredientResponse `json:"ingredients"`
	Preperation     string               `json:"preperation"`
}

type RecipePost struct {
	Name            string           `json:"name"`
	Author          string           `json:"author"`
	Thumbnail       string           `json:"thumbnail"`
	PreperationTime int              `json:"preperationTime"`
	Ingredients     []IngredientPost `json:"ingredients"`
	Preperation     string           `json:"preperation"`
}

func ToRecipeResponse(r *entities.Recipe, ings *[]entities.Ingredient) *RecipeResponse {
	ingR := []IngredientResponse{}
	for _, i := range *ings {
		ingR = append(ingR, *ToIngredientResponse(&i))
	}

	return &RecipeResponse{
		Id:              "r_" + strconv.FormatInt(r.Id, 10),
		Author:          "u_" + strconv.Itoa(r.Author),
		Name:            r.Name,
		Thumbnail:       r.Thumbnail,
		PreperationTime: r.PreperationTime,
		Ingredients:     ingR,
		Preperation:     r.Preperation,
	}
}
