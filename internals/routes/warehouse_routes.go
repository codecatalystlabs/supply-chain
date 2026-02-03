package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupWarehouseRoutes configures all warehouse-related routes
func SetupWarehouseRoutes(api *gin.RouterGroup) {
	// Warehouses
	warehouses := api.Group("/warehouses")
	{
		warehouses.POST("/", handlers.CreateWarehouse)
		warehouses.GET("/", handlers.ListWarehouses)
		warehouses.GET("/:id", handlers.GetWarehouse)
		warehouses.PUT("/:id", handlers.UpdateWarehouse)
		warehouses.DELETE("/:id", handlers.DeleteWarehouse)
	}

	// Warehouse Orders
	warehouseOrders := api.Group("/warehouse-orders")
	{
		warehouseOrders.POST("/", handlers.ReceiveWarehouseOrder)
		warehouseOrders.GET("/", handlers.ListWarehouseOrders)
		warehouseOrders.GET("/:id", handlers.GetWarehouseOrder)
	}
}
