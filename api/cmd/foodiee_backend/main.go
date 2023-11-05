package main

import (
	"log"

	"github.com/alaust/foodiee/api/internal/api"
	"github.com/alaust/foodiee/api/internal/database"
	"github.com/alaust/foodiee/api/internal/resources"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.NewDatabase("/home/alaust/Development/alaust/foodiee/api/foodiee.db")

	server := &resources.Server{
		DB: *db,
	}

	api.RegisterHandlers(router, server)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
