package models

type Recipe struct {
	Id              string
	Name            string
	Author          string
	Thumbnail       string
	PreperationTime int
	Ingredients     []Ingredient
	Preperation     string
}
