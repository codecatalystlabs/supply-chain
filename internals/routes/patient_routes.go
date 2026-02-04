package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupPatientRoutes configures all patient-related routes
func SetupPatientRoutes(api *gin.RouterGroup) {
	// Patient Visits
	patientVisit := api.Group("/patient-visit")
	{
		patientVisit.POST("/", handlers.CreatePatientVisit)
		patientVisit.GET("/", handlers.ListPatientVisits)
		patientVisit.GET("/:id", handlers.GetPatientVisit)
		patientVisit.PUT("/:id", handlers.UpdatePatientVisit)
		patientVisit.DELETE("/:id", handlers.DeletePatientVisit)
	}

	// Product AMC
	productAmc := api.Group("/product-amc")
	{
		productAmc.POST("/", handlers.CreateProductAmc)
		productAmc.GET("/", handlers.ListProductAmc)
		productAmc.GET("/:id", handlers.GetProductAmc)
		productAmc.PUT("/:id", handlers.UpdateProductAmc)
		productAmc.DELETE("/:id", handlers.DeleteProductAmc)
	}
}
