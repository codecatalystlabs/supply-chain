package handlers

import (
	"time"

	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

func CreateStockOnHand(c *gin.Context) {
	var dto StockOnHandCreateDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	record := models.StockOnHand{
		SrcSystemCode:   dto.SrcSystemCode,
		SrcFacilityCode: dto.SrcFacilityCode,
		SrcProductCode:  dto.SrcProductCode,
		SrcBatchNumber:  dto.SrcBatchNumber,
		SrcQuantity:     dto.SrcQuantity,
		SrcExpiryDate:   dto.SrcExpiryDate,
		SrcTimestamp:    time.Now(),
	}

	if err := db.Create(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, record)
}
