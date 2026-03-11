package handlers

import (
	"net/http"

	"supply-chain/internals/config"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary List regions
// @Tags Location
// @Produce json
// @Success 200 {array} models.Region
// @Router /regions [get]
func ListRegionsFromTable(c *gin.Context) {
	var regions []models.Region
	if err := config.DB.Order("name").Find(&regions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)
}

// @Summary List zones (optionally by region_id)
// @Tags Location
// @Produce json
// @Param region_id query int false "Filter by region ID"
// @Success 200 {array} models.Zone
// @Router /zones [get]
func ListZones(c *gin.Context) {
	var zones []models.Zone
	query := config.DB.Order("name")
	if regionID := c.Query("region_id"); regionID != "" {
		query = query.Where("region_id = ?", regionID)
	}
	if err := query.Find(&zones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, zones)
}

// @Summary List districts (by zone_id or region_id)
// @Tags Location
// @Produce json
// @Param zone_id query int false "Filter by zone ID"
// @Param region_id query int false "Filter by region ID (returns districts in that region)"
// @Success 200 {array} models.District
// @Router /districts [get]
func ListDistrictsFromTable(c *gin.Context) {
	var districts []models.District
	query := config.DB.Model(&models.District{}).Order("name")
	if zoneID := c.Query("zone_id"); zoneID != "" {
		query = query.Where("zone_id = ?", zoneID)
	} else if regionID := c.Query("region_id"); regionID != "" {
		subQuery := config.DB.Model(&models.Zone{}).Select("id").Where("region_id = ?", regionID)
		query = query.Where("zone_id IN (?)", subQuery)
	}
	if err := query.Find(&districts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, districts)
}

// @Summary List levels of care
// @Tags Location
// @Produce json
// @Success 200 {array} models.LevelOfCare
// @Router /levels-of-care [get]
func ListLevelsOfCare(c *gin.Context) {
	var levels []models.LevelOfCare
	if err := config.DB.Order("code").Find(&levels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, levels)
}
