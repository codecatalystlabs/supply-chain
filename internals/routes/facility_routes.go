package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupFacilityRoutes configures all facility-related routes
func SetupFacilityRoutes(api *gin.RouterGroup) {
	// Facilities
	facilities := api.Group("/facilities")
	{
		facilities.POST("/", handlers.CreateFacility)
		facilities.GET("/regions", handlers.ListRegions)
		facilities.GET("/districts", handlers.ListDistricts)
		facilities.GET("/", handlers.ListFacilities)
		facilities.GET("/:id", handlers.GetFacility)
		facilities.PUT("/:id", handlers.UpdateFacility)
		facilities.DELETE("/:id", handlers.DeleteFacility)
	}

	// Facility Orders
	facilityOrders := api.Group("/facility-orders")
	{
		facilityOrders.POST("/", handlers.CreateFacilityOrder)
		facilityOrders.GET("/", handlers.ListFacilityOrders)
		facilityOrders.GET("/:id", handlers.GetFacilityOrder)
		facilityOrders.PUT("/:id", handlers.UpdateFacilityOrder)
		facilityOrders.POST("/:id/submit", handlers.SubmitFacilityOrder)
		facilityOrders.POST("/:id/approve", handlers.ApproveFacilityOrder)
		facilityOrders.POST("/:id/deliveries", handlers.CreateFacilityDelivery)
	}
}
