package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupEMRRoutes configures all EMR integration routes
func SetupEMRRoutes(api *gin.RouterGroup) {
	// EMR Integrations
	emrIntegrations := api.Group("/emr-integrations")
	{
		emrIntegrations.POST("/", handlers.CreateEMRIntegration)
		emrIntegrations.GET("/", handlers.ListEMRIntegrations)
		emrIntegrations.GET("/:id", handlers.GetEMRIntegration)
		emrIntegrations.PUT("/:id", handlers.UpdateEMRIntegration)
		emrIntegrations.POST("/:id/verify", handlers.VerifyEMRIntegration)
		emrIntegrations.GET("/:id/sync-logs", handlers.GetEMRSyncLogs)
	}
}
