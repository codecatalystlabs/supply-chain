package dto

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
)

func CreateStockOnHand(c *gin.Context) {
	var payload dto.StockOnHandCreateDTO

	// Validate JSON payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	// Map DTO → Model
	record := models.StockOnHand{
		SrcSystemCode:   payload.SrcSystemCode,
		SrcFacilityCode: payload.SrcFacilityCode,
		SrcProductCode:  payload.SrcProductCode,
		SrcBatchNumber:  payload.SrcBatchNumber,
		SrcQuantity:     payload.SrcQuantity,
		SrcExpiryDate:   payload.SrcExpiryDate,
		SrcTimestamp:    time.Now(),

		BaseModel: models.BaseModel{
			ValidationStatus: 0,
			SyncStatus:       0,
		},
	}

	//Persist
	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create stock record",
			"details": err.Error(),
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Stock on hand recorded successfully",
		"data":    record,
	})
}
