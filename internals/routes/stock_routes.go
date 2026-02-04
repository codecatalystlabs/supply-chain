package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupStockRoutes configures all stock-related routes
func SetupStockRoutes(api *gin.RouterGroup) {
	stock := api.Group("/stock")
	{
		// Stock on-hand
		stock.POST("/on-hand", handlers.CreateStockOnHand)
		stock.GET("/on-hand", handlers.ListStockOnHand)
		stock.GET("/on-hand/:id", handlers.GetStockOnHand)
		stock.PUT("/on-hand/:id", handlers.UpdateStockOnHand)
		stock.DELETE("/on-hand/:id", handlers.DeleteStockOnHand)

		// Stock dispensed
		stock.POST("/dispensed", handlers.CreateStockDispensed)
		stock.GET("/dispensed", handlers.ListStockDispensed)
		stock.GET("/dispensed/:id", handlers.GetStockDispensed)
		stock.PUT("/dispensed/:id", handlers.UpdateStockDispensed)
		stock.DELETE("/dispensed/:id", handlers.DeleteStockDispensed)

		// Stock returns
		stock.POST("/return", handlers.CreateStockReturn)
		stock.GET("/return", handlers.ListStockReturns)
		stock.GET("/return/:id", handlers.GetStockReturn)
		stock.PUT("/return/:id", handlers.UpdateStockReturn)
		stock.DELETE("/return/:id", handlers.DeleteStockReturn)

		// Stock adjustments
		stock.POST("/adjustment", handlers.CreateStockAdjustment)
		stock.GET("/adjustment", handlers.ListStockAdjustments)
		stock.GET("/adjustment/:id", handlers.GetStockAdjustment)
		stock.PUT("/adjustment/:id", handlers.UpdateStockAdjustment)
		stock.DELETE("/adjustment/:id", handlers.DeleteStockAdjustment)

		// Stock transfers
		stock.POST("/transfers", handlers.CreateStockTransfer)
		stock.GET("/transfers", handlers.ListStockTransfers)
		stock.GET("/transfers/:id", handlers.GetStockTransfer)
		stock.PUT("/transfers/:id", handlers.UpdateStockTransfer)
		stock.POST("/transfers/:id/approve", handlers.ApproveStockTransfer)
		stock.POST("/transfers/:id/receive", handlers.ReceiveStockTransfer)
	}
}
