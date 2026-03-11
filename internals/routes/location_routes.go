package routes

import (
	"supply-chain/internals/handlers"

	"github.com/gin-gonic/gin"
)

// SetupLocationRoutes configures location lookup routes (regions, zones, districts, levels of care)
func SetupLocationRoutes(api *gin.RouterGroup) {
	api.GET("/regions", handlers.ListRegionsFromTable)
	api.GET("/zones", handlers.ListZones)
	api.GET("/districts", handlers.ListDistrictsFromTable)
	api.GET("/levels-of-care", handlers.ListLevelsOfCare)
}
