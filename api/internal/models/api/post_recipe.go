package models

type ApiPostRecipe struct {
	Author          string
	Thumbnail       string
	PreperationTime int
	Ingredients     []ApiIngredient
	Preperation     string
}
