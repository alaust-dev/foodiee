package models

type ApiRecipe struct {
	Id              string
	Name            string
	Author          string
	Thumbnail       string
	PreperationTime int
	Ingredients     []ApiIngredient
	Preperation     string
}
