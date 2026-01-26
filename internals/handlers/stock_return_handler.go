package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create stock return
// @Tags StockReturn
// @Accept json
// @Produce json
// @Param payload body dto.StockReturnCreateDTO true "Stock return payload"
// @Success 201 {object} dto.StockReturnResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/return [post]
func CreateStockReturn(c *gin.Context) {
	var payload dto.StockReturnCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.StockReturn{
		RtnSystemCode:   payload.RtnSystemCode,
		RtnFacilityCode: payload.RtnFacilityCode,
		RtnReturnDate:   payload.RtnReturnDate,
		RtnReturnNumber: payload.RtnReturnNumber,
		RtnProductCode:  payload.RtnProductCode,
		RtnBatchNumber:  payload.RtnBatchNumber,
		RtnUnitCode:     payload.RtnUnitCode,
		RtnQuantity:     payload.RtnQuantity,
	}

	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapToStockReturnResponse(record))
}

// @Summary List stock returns
// @Tags StockReturn
// @Produce json
// @Success 200 {array} dto.StockReturnResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/return [get]
func ListStockReturns(c *gin.Context) {
	var records []models.StockReturn
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.StockReturnResponseDTO
	for _, r := range records {
		resp = append(resp, mapToStockReturnResponse(r))
	}
	c.JSON(200, resp)
}

// @Summary Get stock return by id
// @Tags StockReturn
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.StockReturnResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/return/{id} [get]
func GetStockReturn(c *gin.Context) {
	id := c.Param("id")
	var record models.StockReturn
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToStockReturnResponse(record))
}

// @Summary Update stock return
// @Tags StockReturn
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.StockReturnUpdateDTO true "Update payload"
// @Success 200 {object} dto.StockReturnResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/return/{id} [put]
func UpdateStockReturn(c *gin.Context) {
	id := c.Param("id")
	var payload dto.StockReturnUpdateDTO
	var record models.StockReturn
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.RtnQuantity != nil {
		record.RtnQuantity = *payload.RtnQuantity
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToStockReturnResponse(record))
}

// @Summary Delete stock return
// @Tags StockReturn
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/stock/return/{id} [delete]
func DeleteStockReturn(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.StockReturn{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToStockReturnResponse(m models.StockReturn) dto.StockReturnResponseDTO {
	return dto.StockReturnResponseDTO{
		ID:               m.ID,
		RtnSystemCode:    m.RtnSystemCode,
		RtnFacilityCode:  m.RtnFacilityCode,
		RtnTimestamp:     m.RtnTimestamp,
		RtnReturnNumber:  m.RtnReturnNumber,
		RtnReturnDate:    m.RtnReturnDate,
		RtnProductCode:   m.RtnProductCode,
		RtnBatchNumber:   m.RtnBatchNumber,
		RtnUnitCode:      m.RtnUnitCode,
		RtnQuantity:      m.RtnQuantity,
		ValidationStatus: m.ValidationStatus,
		SyncStatus:       m.SyncStatus,
	}
}
