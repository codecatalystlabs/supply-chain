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

// @contact.name DHI
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

	stock := router.Group("/api/v1/stock")
	{
		stock.POST("/on-hand", handlers.CreateStockOnHand)
		stock.GET("/on-hand", handlers.ListStockOnHand)
		stock.GET("/on-hand/:id", handlers.GetStockOnHand)
		stock.PUT("/on-hand/:id", handlers.UpdateStockOnHand)
		stock.DELETE("/on-hand/:id", handlers.DeleteStockOnHand)
	}
	// Purchase Order routes
	po := router.Group("/api/v1/purchase-order")
	{
		po.POST("/", handlers.CreatePurchaseOrder)
		po.GET("/:id", handlers.GetPurchaseOrder)
		po.PUT("/:id", handlers.UpdatePurchaseOrder)
		po.DELETE("/:id", handlers.DeletePurchaseOrder)
		po.GET("/", handlers.ListPurchaseOrders)
	}

	// Run server
	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
