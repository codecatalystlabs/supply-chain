package routes

import (
	"log"
	"os"
	"path/filepath"
	"supply-chain/internals/web/controllers"

	"github.com/gin-gonic/gin"
)

// findStaticPath finds the static files directory by checking multiple possible locations
func findStaticPath() string {
	// Possible paths relative to different working directories
	possiblePaths := []string{
		"internals/web/static",                    // From project root
		"../internals/web/static",                 // From cmd/server
		"../../internals/web/static",              // From cmd/server/server
		"./internals/web/static",                 // Current directory
	}

	// Get current working directory
	wd, err := os.Getwd()
	if err == nil {
		// Also try absolute paths
		possiblePaths = append(possiblePaths,
			filepath.Join(wd, "internals/web/static"),
			filepath.Join(wd, "../internals/web/static"),
			filepath.Join(wd, "../../internals/web/static"),
		)
	}

	// Check each path
	for _, path := range possiblePaths {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err == nil {
				return absPath
			}
			return path
		}
	}

	// Default fallback
	return "internals/web/static"
}

// SetupWebRoutes configures all web portal routes under /cp/* prefix
func SetupWebRoutes(router *gin.Engine) {
	// Find static files directory dynamically
	staticPath := findStaticPath()
	log.Printf("📁 Serving static files from: %s\n", staticPath)
	
	// Serve static files from two paths for convenience
	router.Static("/static", staticPath)
	router.Static("/cp/static", staticPath)

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

		// Stock Management
		portal.GET("/stock-on-hand", controllers.ShowStockOnHand)
		portal.GET("/stock-dispensed", controllers.ShowStockDispensed)
		portal.GET("/stock-adjustments", controllers.ShowStockAdjustments)
		portal.GET("/stock-returns", controllers.ShowStockReturns)
		portal.GET("/stock-transfers", controllers.ShowStockTransfers)

		// Reports
		portal.GET("/patient-visits", controllers.ShowPatientVisits)
		portal.GET("/product-amc", controllers.ShowProductAMC)
	}
}
