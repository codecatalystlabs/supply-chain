package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// Disable Gin's debug logging to reduce noise
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Increase multipart upload limit (default is too small for Excel files).
	// Configure via env var MAX_UPLOAD_MB (defaults to 50).
	maxUploadMB := 50
	if v := os.Getenv("MAX_UPLOAD_MB"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxUploadMB = n
		}
	}
	router.MaxMultipartMemory = int64(maxUploadMB) << 20

	// Initialize sessions
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "supply-chain-secret-key-change-in-production"
		log.Println("⚠️  Using default session secret. Set SESSION_SECRET in .env for production!")
	}
	store := cookie.NewStore([]byte(sessionSecret))
	router.Use(sessions.Sessions("supply_chain_session", store))

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
