package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes configures all order-related routes
func SetupOrderRoutes(api *gin.RouterGroup) {
	// Goods Receipt
	goodsReceipt := api.Group("/goods-receipt")
	{
		goodsReceipt.POST("/", handlers.CreateGoodsReceipt)
		goodsReceipt.GET("/", handlers.ListGoodsReceipts)
		goodsReceipt.GET("/:id", handlers.GetGoodsReceipt)
		goodsReceipt.PUT("/:id", handlers.UpdateGoodsReceipt)
		goodsReceipt.DELETE("/:id", handlers.DeleteGoodsReceipt)
	}
}
