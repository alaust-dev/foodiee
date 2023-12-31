package main

import (
	"log"
	"os"

	"github.com/alaust/foodiee/backend/api"
	"github.com/alaust/foodiee/backend/internal/database"
	"github.com/alaust/foodiee/backend/internal/resources"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	path := os.Getenv("DB_FILE_PATH")
	db := database.CreateNew(&path)

	server := &resources.Server{
		DB: *db,
	}

	api.RegisterHandlers(router, server)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
