package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create pharmacy stock
// @Tags PharmacyStock
// @Accept json
// @Produce json
// @Param payload body dto.PharmacyStockCreateDTO true "Pharmacy stock payload"
// @Success 201 {object} dto.PharmacyStockResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/pharmacy-stock [post]
func CreatePharmacyStock(c *gin.Context) {
	var payload dto.PharmacyStockCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := models.PharmacyStock{
		PhaSystemCode:   payload.PhaSystemCode,
		PhaFacilityCode: payload.PhaFacilityCode,
		PhaProductCode:  payload.PhaProductCode,
		PhaBatchNumber:  payload.PhaBatchNumber,
		PhaQuantity:     payload.PhaQuantity,
		PhaExpiryDate:   payload.PhaExpiryDate,
	}
	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapToPharmacyStockResponse(record))
}

// @Summary List pharmacy stock
// @Tags PharmacyStock
// @Produce json
// @Success 200 {array} dto.PharmacyStockResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/v1/pharmacy-stock [get]
func ListPharmacyStock(c *gin.Context) {
	var records []models.PharmacyStock
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.PharmacyStockResponseDTO
	for _, r := range records {
		resp = append(resp, mapToPharmacyStockResponse(r))
	}
	c.JSON(200, resp)
}

// @Summary Get pharmacy stock by id
// @Tags PharmacyStock
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.PharmacyStockResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/pharmacy-stock/{id} [get]
func GetPharmacyStock(c *gin.Context) {
	id := c.Param("id")
	var record models.PharmacyStock
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToPharmacyStockResponse(record))
}

// @Summary Update pharmacy stock
// @Tags PharmacyStock
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.PharmacyStockUpdateDTO true "Update payload"
// @Success 200 {object} dto.PharmacyStockResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/pharmacy-stock/{id} [put]
func UpdatePharmacyStock(c *gin.Context) {
	id := c.Param("id")
	var payload dto.PharmacyStockUpdateDTO
	var record models.PharmacyStock
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.PhaQuantity != nil {
		record.PhaQuantity = *payload.PhaQuantity
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToPharmacyStockResponse(record))
}

// @Summary Delete pharmacy stock
// @Tags PharmacyStock
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/pharmacy-stock/{id} [delete]
func DeletePharmacyStock(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.PharmacyStock{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToPharmacyStockResponse(m models.PharmacyStock) dto.PharmacyStockResponseDTO {
	return dto.PharmacyStockResponseDTO{
		ID:               m.ID,
		PhaSystemCode:    m.PhaSystemCode,
		PhaFacilityCode:  m.PhaFacilityCode,
		PhaTimestamp:     m.PhaTimestamp,
		PhaProductCode:   m.PhaProductCode,
		PhaBatchNumber:   m.PhaBatchNumber,
		PhaQuantity:      m.PhaQuantity,
		PhaExpiryDate:    m.PhaExpiryDate,
		ValidationStatus: m.ValidationStatus,
		SyncStatus:       m.SyncStatus,
	}
}
