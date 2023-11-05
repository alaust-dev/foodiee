package models

type Recipe struct {
	Id              int
	Name            string
	Author          string
	Thumbnail       string
	PreperationTime int
	Ingredients     []Ingredient
	Preperation     string
}
