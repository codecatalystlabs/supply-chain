package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create stock adjustment
// @Tags StockAdjustment
// @Accept json
// @Produce json
// @Param payload body dto.StockAdjustmentCreateDTO true "Stock adjustment payload"
// @Success 201 {object} dto.StockAdjustmentResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/adjustment [post]
func CreateStockAdjustment(c *gin.Context) {
	var payload dto.StockAdjustmentCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := models.StockAdjustment{
		AdjSystemCode:       payload.AdjSystemCode,
		AdjFacilityCode:     payload.AdjFacilityCode,
		AdjAdjustmentDate:   payload.AdjAdjustmentDate,
		AdjAdjustmentType:   payload.AdjAdjustmentType,
		AdjAdjustmentReason: payload.AdjAdjustmentReason,
		AdjProductCode:      payload.AdjProductCode,
		AdjBatchNumber:      payload.AdjBatchNumber,
		AdjQuantity:         payload.AdjQuantity,
		AdjExpiryDate:       payload.AdjExpiryDate,
	}
	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapToStockAdjustmentResponse(record))
}

// @Summary List stock adjustments
// @Tags StockAdjustment
// @Produce json
// @Success 200 {array} dto.StockAdjustmentResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/adjustment [get]
func ListStockAdjustments(c *gin.Context) {
	var records []models.StockAdjustment
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.StockAdjustmentResponseDTO
	for _, r := range records {
		resp = append(resp, mapToStockAdjustmentResponse(r))
	}
	c.JSON(200, resp)
}

// @Summary Get stock adjustment by id
// @Tags StockAdjustment
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.StockAdjustmentResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/adjustment/{id} [get]
func GetStockAdjustment(c *gin.Context) {
	id := c.Param("id")
	var record models.StockAdjustment
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToStockAdjustmentResponse(record))
}

// @Summary Update stock adjustment
// @Tags StockAdjustment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.StockAdjustmentUpdateDTO true "Update payload"
// @Success 200 {object} dto.StockAdjustmentResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/adjustment/{id} [put]
func UpdateStockAdjustment(c *gin.Context) {
	id := c.Param("id")
	var payload dto.StockAdjustmentUpdateDTO
	var record models.StockAdjustment
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.AdjQuantity != nil {
		record.AdjQuantity = *payload.AdjQuantity
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToStockAdjustmentResponse(record))
}

// @Summary Delete stock adjustment
// @Tags StockAdjustment
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/adjustment/{id} [delete]
func DeleteStockAdjustment(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.StockAdjustment{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToStockAdjustmentResponse(m models.StockAdjustment) dto.StockAdjustmentResponseDTO {
	return dto.StockAdjustmentResponseDTO{
		ID:                  m.ID,
		AdjSystemCode:       m.AdjSystemCode,
		AdjFacilityCode:     m.AdjFacilityCode,
		AdjTimestamp:        m.AdjTimestamp,
		AdjAdjustmentDate:   m.AdjAdjustmentDate,
		AdjAdjustmentType:   m.AdjAdjustmentType,
		AdjAdjustmentReason: m.AdjAdjustmentReason,
		AdjProductCode:      m.AdjProductCode,
		AdjBatchNumber:      m.AdjBatchNumber,
		AdjQuantity:         m.AdjQuantity,
		AdjExpiryDate:       m.AdjExpiryDate,
		ValidationStatus:    m.ValidationStatus,
		SyncStatus:          m.SyncStatus,
	}
}
