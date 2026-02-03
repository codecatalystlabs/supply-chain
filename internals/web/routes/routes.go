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

		// Protected web routes (dashboard, etc)
		// TODO: Add auth middleware to protect these routes
		portal.GET("/dashboard", controllers.ShowDashboard)
		portal.GET("/logout", controllers.Logout)
	}
}
