package models

import models "github.com/alaust/foodiee/api/internal/models/api"

type Ingredient struct {
	Id      int
	Name    string
	Unit    string
	Ammount int
}

func ToApiModel(ing Ingredient) models.ApiIngredient {
	return &ApiIngredient{
		Id: "",
	}
}
