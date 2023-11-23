package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/alaust/foodiee/backend/internal/entities"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CreateNew(path *string) *Database {
	db, err := sql.Open("sqlite3", *path)
	if err != nil {
		log.Fatal("Could not open database: ", err)
	}

	return &Database{db: db}
}

func (d *Database) GetUsers() ([]entities.User, error) {
	rows, err := squirrel.Select("*").From("user").RunWith(d.db).Query()
	if err != nil {
		return []entities.User{}, err
	}

	var users []entities.User
	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Picture)
		if err != nil {
			return []entities.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *Database) GetRecipes() ([]entities.Recipe, error) {
	rows, err := squirrel.Select("*").From("recipe").RunWith(d.db).Query()
	if err != nil {
		return []entities.Recipe{}, err
	}

	var recipes []entities.Recipe
	for rows.Next() {
		recipe := entities.Recipe{}
		err := rows.Scan(&recipe.Id,
			&recipe.Author,
			&recipe.Name,
			&recipe.Thumbnail,
			&recipe.PreperationTime,
			&recipe.Preperation)
		if err != nil {
			return []entities.Recipe{}, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (d *Database) GetRecipe(id int64) (*entities.Recipe, error) {
	var recipe entities.Recipe
	err := squirrel.Select("*").
		From("recipe").
		Where(squirrel.Eq{"id": id}).
		RunWith(d.db).QueryRow().
		Scan(&recipe.Id,
			&recipe.Author,
			&recipe.Name,
			&recipe.Thumbnail,
			&recipe.PreperationTime,
			&recipe.Preperation)
	if err != nil {
		return &entities.Recipe{}, err
	}

	return &recipe, nil
}

func (d *Database) CreateRecipe(r *entities.Recipe) (int64, error) {
	result, err := squirrel.Insert("recipe").
		Columns("author", "name", "thumbnail", "preperation_time", "preperation").
		Values(r.Author, r.Name, r.Thumbnail, r.PreperationTime, r.Preperation).
		RunWith(d.db).Exec()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (d *Database) GetIngredientsOfRecipe(recipeId *int64) ([]entities.Ingredient, error) {
	ir, err := squirrel.Select("ingredient_id, i.name, i.unit, amount").
		From("recipe_ingredient").
		Join("ingredient i on i.id = recipe_ingredient.ingredient_id").
		Where(squirrel.Eq{"recipe_id": recipeId}).
		RunWith(d.db).Query()
	if err != nil {
		return []entities.Ingredient{}, err
	}

	var ingredients []entities.Ingredient
	for ir.Next() {
		ingredient := entities.Ingredient{}
		err := ir.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Unit, &ingredient.Ammount)

		if err != nil {
			return []entities.Ingredient{}, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (d *Database) CreateIngredientsForRecipe(recipeId *int64, ing []entities.Ingredient) error {
	for _, i := range ing {
		var id int64
		err := squirrel.Select("id").
			From("ingredient").
			Where(squirrel.Eq{"name": strings.ToLower(i.Name)}).
			RunWith(d.db).QueryRow().Scan(&id)
		if errors.Is(err, sql.ErrNoRows) {
			r, err := squirrel.Insert("ingredient").
				Columns("name", "unit").
				Values(strings.ToLower(i.Name), i.Unit).
				RunWith(d.db).Exec()
			if err != nil {
				return err
			}

			id, _ = r.LastInsertId()
		} else if err != nil {
			return err
		}

		_, err = squirrel.Insert("recipe_ingredient").
			Columns("recipe_id", "ingredient_id", "amount").
			Values(*recipeId, id, i.Ammount).
			RunWith(d.db).Exec()
	}
	return nil
}
