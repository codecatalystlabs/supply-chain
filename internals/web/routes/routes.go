package routes

import (
	"supply-chain/internals/web/controllers"

	"github.com/gin-gonic/gin"
)

// SetupWebRoutes configures all web portal routes under /cp/* prefix
func SetupWebRoutes(router *gin.Engine) {
	// Serve static files from two paths for convenience
	router.Static("/static", "internals/web/static")
	router.Static("/cp/static", "internals/web/static")

	// Home/Landing page (JSON API info)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Supply Chain Management System",
			"version": "2.0",
			"links": gin.H{
				"api_docs":   "/swagger/index.html",
				"health":     "/health",
				"api_v1":     "/api/v1",
				"web_portal": "/cp/login",
				"dashboard":  "/cp/dashboard",
			},
		})
	})

	// Web portal routes - all under /cp/* prefix
	portal := router.Group("/cp")
	{
		// Authentication routes
		portal.GET("/login", controllers.ShowLoginPage)
		portal.POST("/login", controllers.HandleLogin)
		portal.GET("/logout", controllers.Logout)

		// Dashboard
		portal.GET("/dashboard", controllers.ShowDashboard)

		// Inventory Management
		portal.GET("/inventory", controllers.ShowInventory)
		portal.GET("/stock", controllers.ShowStockList)

		// Facilities
		portal.GET("/facilities", controllers.ShowFacilities)
		portal.GET("/facility-orders", controllers.ShowFacilityOrders)

		// Warehouses
		portal.GET("/warehouses", controllers.ShowWarehouses)
		portal.GET("/warehouse-orders", controllers.ShowWarehouseOrders)
		portal.GET("/goods-receipt", controllers.ShowGoodsReceipt)

		// Procurement
		portal.GET("/procurement", controllers.ShowProcurement)
		portal.GET("/purchase-orders", controllers.ShowPurchaseOrders)
		portal.GET("/procurement-plans", controllers.ShowProcurementPlans)

		// Pharmacies
		portal.GET("/pharmacies", controllers.ShowPharmacies)
		portal.GET("/pharmacy-stock", controllers.ShowPharmacyStock)

		// Reports
		portal.GET("/patient-visits", controllers.ShowPatientVisits)
		portal.GET("/product-amc", controllers.ShowProductAMC)
	}
}
