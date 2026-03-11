package routes

import (
	"log"
	"os"
	"path/filepath"
	"supply-chain/internals/middleware"
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
		// Authentication routes (no auth required)
		portal.GET("/login", controllers.ShowLoginPage)
		portal.POST("/login", controllers.HandleLogin)
		portal.GET("/logout", controllers.Logout) // Logout should be accessible without auth
		portal.POST("/logout", controllers.Logout) // Support POST logout as well

		// Protected routes (require authentication)
		protected := portal.Group("")
		protected.Use(middleware.AuthRequired())
		{
			// Home/Dashboard
			protected.GET("/home", controllers.ShowHome)
			protected.GET("/dashboard", controllers.ShowDashboard)

			// Inventory Management
			protected.GET("/inventory", middleware.RequirePermission("stock.read"), controllers.ShowInventory)
			protected.GET("/stock", middleware.RequirePermission("stock.read"), controllers.ShowStockList)

			// Facilities
			protected.GET("/facilities", middleware.RequirePermission("facilities.read"), controllers.ShowFacilities)
			protected.GET("/facility-orders", middleware.RequirePermission("purchase_orders.read"), controllers.ShowFacilityOrders)

			// Warehouses
			protected.GET("/warehouses", middleware.RequirePermission("warehouses.read"), controllers.ShowWarehouses)
			protected.GET("/warehouse-orders", middleware.RequirePermission("purchase_orders.read"), controllers.ShowWarehouseOrders)
			protected.GET("/goods-receipt", middleware.RequirePermission("stock.read"), controllers.ShowGoodsReceipt)

			// Procurement
			protected.GET("/procurement", middleware.RequirePermission("procurement_plans.read"), controllers.ShowProcurement)
			protected.GET("/purchase-orders", middleware.RequirePermission("purchase_orders.read"), controllers.ShowPurchaseOrders)
			protected.GET("/procurement-plans", middleware.RequirePermission("procurement_plans.read"), controllers.ShowProcurementPlans)
			protected.GET("/procurement-plan-import", middleware.RequirePermission("procurement_plans.read"), controllers.ShowProcurementPlanImport)

			// Pharmacies
			protected.GET("/pharmacies", middleware.RequirePermission("pharmacies.read"), controllers.ShowPharmacies)
			protected.GET("/pharmacy-stock", middleware.RequirePermission("stock.read"), controllers.ShowPharmacyStock)

			// Stock Management
			protected.GET("/stock-on-hand", middleware.RequirePermission("stock.read"), controllers.ShowStockOnHand)
			protected.GET("/stock-dispensed", middleware.RequirePermission("stock.read"), controllers.ShowStockDispensed)
			protected.GET("/stock-adjustments", middleware.RequirePermission("stock.adjust"), controllers.ShowStockAdjustments)
			protected.GET("/stock-returns", middleware.RequirePermission("stock.read"), controllers.ShowStockReturns)
			protected.GET("/stock-transfers", middleware.RequirePermission("stock.transfer"), controllers.ShowStockTransfers)

		// Patient Management
		protected.GET("/patient-visits", middleware.RequirePermission("patient_visits.read"), controllers.ShowPatientVisits)
		protected.GET("/product-amc", middleware.RequirePermission("reports.read"), controllers.ShowProductAMC)

		// User Management (Admin only)
		protected.GET("/users", middleware.RequirePermission("admin.users"), controllers.ShowUsers)
		protected.GET("/roles", middleware.RequirePermission("admin.roles"), controllers.ShowRoles)
		}
	}
}
