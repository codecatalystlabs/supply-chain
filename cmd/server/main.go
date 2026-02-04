package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"supply-chain/internals/config"
	"supply-chain/internals/middleware"
	"supply-chain/internals/routes"

	_ "supply-chain/docs" // 👈 REQUIRED for Swagger
)

// @title MoH Supply Chain Management API
// @version 2.0
// @description Backend API for supply chain management system
// @termsOfService https://health.go.ug

// @contact.name DHI
// @contact.email ict@moh.go.ug

// @host localhost:5500
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		// try loading .env from the cmd/server folder (when run from repo root)
		if err2 := godotenv.Load("cmd/server/.env"); err2 != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.ConnectDatabase()
	config.SeedDatabase()

	router := gin.Default()

	// Apply global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	// Register all routes
	routes.SetupRoutes(router)

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "5500"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📚 Swagger: http://localhost:%s/swagger/index.html", port)
	log.Printf("💊 Health check: http://localhost:%s/health", port)
	log.Printf("📦 API Base: http://localhost:%s/api/v1", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
