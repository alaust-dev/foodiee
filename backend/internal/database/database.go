package database

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/alaust/foodiee/backend/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(path string) *Database {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	return &Database{db: db}
}

func (d *Database) GetUsers() ([]models.User, error) {
	rows, err := squirrel.Select("*").From("user").RunWith(d.db).Query()

	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Picture)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

/*func (d *Database) GetRecipes() ([]models.Recipe, error) {
	rows, err := squirrel.Select("*").From("recipe").RunWith(d.db).Query()
	if err != nil {
		return []models.Recipe{}, err
	}

	var recipes []models.Recipe
	for rows.Next() {
		recipe := models.Recipe{}
		err := rows.Scan(&recipe.Id, &recipe.Author, &recipe.Name, &recipe.Thumbnail, &recipe.PreperationTime, &recipe.Preperation)
		if err != nil {
			return []models.Recipe{}, err
		}
		recipe.Ingredients, _ = d.getIngredientsOfRecipe(recipe.Id)
		if err != nil {
			return []models.Recipe{}, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (d *Database) getIngredientsOfRecipe(recipeId int) ([]models.Ingredient, error) {
	ir, err := squirrel.Select("ingredient_id, i.name, i.unit, amount").
		From("recipe_ingredient").
		Join("ingredient i on i.id = recipe_ingredient.ingredient_id").
		Where(squirrel.Eq{"recipe_id": recipeId}).
		RunWith(d.db).Query()

	if err != nil {
		return []models.Ingredient{}, err
	}

	var ingredients []models.Ingredient
	for ir.Next() {
		ingredient := models.Ingredient{}
		err := ir.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Unit, &ingredient.Ammount)
		if err != nil {
			return []models.Ingredient{}, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}*/
