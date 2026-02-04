package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures all user management routes
func SetupUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/", handlers.CreateUser)
		users.GET("/", handlers.ListUsers)
		users.GET("/:id", handlers.GetUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}

