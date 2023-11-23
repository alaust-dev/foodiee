package resources

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/alaust/foodiee/backend/api/dto"
	"github.com/alaust/foodiee/backend/internal/database"
	"github.com/alaust/foodiee/backend/internal/entities"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB database.Database
}

func (s *Server) GetRecipes(c *gin.Context) {
	result := []dto.RecipeResponse{}
	recipes, err := s.DB.GetRecipes()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}

	for _, r := range recipes {
		ings, err := s.DB.GetIngredientsOfRecipe(&r.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
		}
		result = append(result, *dto.ToRecipeResponse(&r, &ings))
	}
	c.JSON(http.StatusOK, &dto.ListWrapper{Data: result})
}

func (s *Server) PostRecipes(c *gin.Context) {
	var recipe dto.RecipePost
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Msg: err.Error()})
	}

	a_id, err := strconv.Atoi(strings.TrimPrefix(recipe.Author, "u_"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &dto.Error{Msg: err.Error()})
	}
	r_id, err := s.DB.CreateRecipe(&entities.Recipe{
		Author:          a_id,
		Name:            recipe.Name,
		Thumbnail:       recipe.Thumbnail,
		PreperationTime: recipe.PreperationTime,
		Preperation:     recipe.Preperation,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}
	ings := []entities.Ingredient{}
	for _, i := range recipe.Ingredients {
		ings = append(ings, entities.Ingredient{
			Name:    i.Name,
			Unit:    i.Unit,
			Ammount: i.Amount,
		})
	}
	err = s.DB.CreateIngredientsForRecipe(&r_id, ings)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}

	r_result, err := s.DB.GetRecipe(r_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Msg: err.Error()})
	}

	ing_result, err := s.DB.GetIngredientsOfRecipe(&r_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Error{Msg: err.Error()})
	}

	c.AbortWithStatusJSON(http.StatusOK, dto.ToRecipeResponse(r_result, &ing_result))
}

func (s *Server) GetRecipesId(c *gin.Context, id string) {
	r_id, err := strconv.Atoi(strings.TrimPrefix(id, "r_"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &dto.Error{Msg: err.Error()})
	}

	recipe, err := s.DB.GetRecipe(int64(r_id))
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatusJSON(http.StatusNotFound, &dto.Error{Msg: err.Error()})
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}

	ings, err := s.DB.GetIngredientsOfRecipe(&recipe.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}

	c.JSON(http.StatusOK, dto.ToRecipeResponse(recipe, &ings))
}

func (s *Server) GetUsers(c *gin.Context) {
	result := []dto.UserResponse{}
	users, err := s.DB.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{Msg: err.Error()})
	}

	for _, u := range users {
		result = append(result, *dto.ToUserResponse(&u))
	}

	c.JSON(http.StatusOK, &dto.ListWrapper{Data: result})
}
