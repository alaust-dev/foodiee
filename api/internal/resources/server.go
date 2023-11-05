package resources

import (
	"fmt"
	"net/http"

	api_models "github.com/alaust/foodiee/api/internal/api/models"
	"github.com/alaust/foodiee/api/internal/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB database.Database
}

func (s *Server) GetRecipes(c *gin.Context) {
	var recipe api_models.PostRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, api_models.Error{Msg: err.Error()})
		return
	}
	for ingredient, _ := range recipe.Ingredients {
		fmt.Print(ingredient)
	}
}

func (s *Server) PostRecipes(c *gin.Context) {}

func (s *Server) GetRecipesId(c *gin.Context, id string) {}

func (s *Server) DeleteShoppingListUserId(c *gin.Context, userId string) {}

func (s *Server) GetShoppingListUserId(c *gin.Context, userId string) {}

func (s *Server) PutShoppingListUserId(c *gin.Context, userId string) {}

func (s *Server) GetUsers(c *gin.Context) {
	users, err := s.DB.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api_models.Error{Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &api_models.ListWrapper{Data: users})
}
