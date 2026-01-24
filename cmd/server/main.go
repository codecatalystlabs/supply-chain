package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"supply-chain/internals/config"
	"supply-chain/internals/handlers"

	_ "supply-chain/docs" // 👈 REQUIRED for Swagger
)

// @title MoH Emergency Dispatch API
// @version 1.0
// @description Backend API for emergency call & dispatch system
// @termsOfService https://health.go.ug

// @contact.name ICT Division
// @contact.email ict@moh.go.ug

// @host localhost:8080
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1/stock")
	{
		api.POST("/on-hand", handlers.CreateStockOnHand)
		api.GET("/on-hand", handlers.ListStockOnHand)
		api.GET("/on-hand/:id", handlers.GetStockOnHand)
		api.PUT("/on-hand/:id", handlers.UpdateStockOnHand)
		api.DELETE("/on-hand/:id", handlers.DeleteStockOnHand)
	}

	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
