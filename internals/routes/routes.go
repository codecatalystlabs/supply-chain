package routes

import (
	"supply-chain/internals/handlers"
	webRoutes "supply-chain/internals/web/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Supply Chain API is running",
			"version": "2.0",
		})
	})

	// API routes group
	api := router.Group("/api/v1")

	// Setup feature-specific routes
	SetupStockRoutes(api)
	SetupProcurementRoutes(api)
	SetupFacilityRoutes(api)
	SetupPharmacyRoutes(api)
	SetupWarehouseRoutes(api)
	SetupOrderRoutes(api)
	SetupPatientRoutes(api)
	SetupEMRRoutes(api)
	SetupUserRoutes(api)
	
	// Roles endpoint
	api.GET("/roles", handlers.ListRoles)
	api.GET("/roles/:id", handlers.GetRole)

	// Web routes (for future dashboard)
	webRoutes.SetupWebRoutes(router)
}
