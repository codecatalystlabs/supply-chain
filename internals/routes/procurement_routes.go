package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupProcurementRoutes configures all procurement-related routes
func SetupProcurementRoutes(api *gin.RouterGroup) {
	// Purchase Orders
	po := api.Group("/purchase-orders")
	{
		po.POST("/", handlers.CreatePurchaseOrder)
		po.GET("/", handlers.ListPurchaseOrders)
		po.GET("/:id", handlers.GetPurchaseOrder)
		po.PUT("/:id", handlers.UpdatePurchaseOrder)
		po.DELETE("/:id", handlers.DeletePurchaseOrder)
	}

	// Procurement Plans
	pp := api.Group("/procurement-plans")
	{
		pp.POST("/", handlers.CreateProcurementPlan)
		pp.GET("/", handlers.ListProcurementPlans)
		pp.GET("/:id", handlers.GetProcurementPlan)
		pp.DELETE("/:id", handlers.DeleteProcurementPlan)
	}
}
