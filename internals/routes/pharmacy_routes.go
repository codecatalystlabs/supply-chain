package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupPharmacyRoutes configures all pharmacy-related routes
func SetupPharmacyRoutes(api *gin.RouterGroup) {
	// Pharmacies
	pharmacies := api.Group("/pharmacies")
	{
		pharmacies.POST("/", handlers.CreatePharmacy)
		pharmacies.GET("/", handlers.ListPharmacies)
		pharmacies.GET("/:id", handlers.GetPharmacy)
		pharmacies.PUT("/:id", handlers.UpdatePharmacy)
		pharmacies.DELETE("/:id", handlers.DeletePharmacy)
	}

	// Pharmacy Stock
	pharmacyStock := api.Group("/pharmacy-stock")
	{
		pharmacyStock.POST("/", handlers.CreatePharmacyStock)
		pharmacyStock.GET("/", handlers.ListPharmacyStock)
		pharmacyStock.GET("/:id", handlers.GetPharmacyStock)
		pharmacyStock.PUT("/:id", handlers.UpdatePharmacyStock)
		pharmacyStock.DELETE("/:id", handlers.DeletePharmacyStock)
	}
}
