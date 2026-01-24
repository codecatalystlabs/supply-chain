package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"supply-chain/internals/config"
	"supply-chain/internals/handlers"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB
	config.ConnectDatabase()

	// Gin
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/stock/on-hand", handlers.CreateStockOnHand)
	}

	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
