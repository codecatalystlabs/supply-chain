package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create stock dispensed (bulk)
// @Tags StockDispensed
// @Accept json
// @Produce json
// @Param payload body []dto.StockDispensedCreateDTO true "List of stock dispensed payloads"
// @Success 201 {array} dto.StockDispensedResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/dispensed [post]
func CreateStockDispensed(c *gin.Context) {
	var payloads []dto.StockDispensedCreateDTO
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(payloads) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload must be a non-empty array"})
		return
	}

	var created []dto.StockDispensedResponseDTO
	for _, payload := range payloads {
		record := models.StockDispensed{
			DspSystemCode:        payload.DspSystemCode,
			DspFacilityCode:      payload.DspFacilityCode,
			DspDispenseDate:      payload.DspDispenseDate,
			DspProductCode:       payload.DspProductCode,
			DspBatchNumber:       payload.DspBatchNumber,
			DspDispensedQuantity: payload.DspDispensedQuantity,
			DspPatientHash:       payload.DspPatientHash,
			DspExpiryDate:        payload.DspExpiryDate,
		}

		if err := config.DB.Create(&record).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		created = append(created, mapToStockDispensedResponse(record))
	}

	c.JSON(http.StatusCreated, created)
}

// @Summary List stock dispensed
// @Tags StockDispensed
// @Produce json
// @Success 200 {array} dto.StockDispensedResponseDTO
// @Failure 500 {object} map[string]string
// @Router /stock/dispensed [get]
func ListStockDispensed(c *gin.Context) {
	var records []models.StockDispensed
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.StockDispensedResponseDTO
	for _, r := range records {
		resp = append(resp, mapToStockDispensedResponse(r))
	}

	c.JSON(200, resp)
}

// @Summary Get stock dispensed by id
// @Tags StockDispensed
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.StockDispensedResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/dispensed/{id} [get]
func GetStockDispensed(c *gin.Context) {
	id := c.Param("id")
	var record models.StockDispensed
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToStockDispensedResponse(record))
}

// @Summary Update stock dispensed
// @Tags StockDispensed
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.StockDispensedUpdateDTO true "Update payload"
// @Success 200 {object} dto.StockDispensedResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/dispensed/{id} [put]
func UpdateStockDispensed(c *gin.Context) {
	id := c.Param("id")
	var payload dto.StockDispensedUpdateDTO
	var record models.StockDispensed

	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if payload.DspDispensedQuantity != nil {
		record.DspDispensedQuantity = *payload.DspDispensedQuantity
	}

	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, mapToStockDispensedResponse(record))
}

// @Summary Delete stock dispensed
// @Tags StockDispensed
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/dispensed/{id} [delete]
func DeleteStockDispensed(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.StockDispensed{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToStockDispensedResponse(m models.StockDispensed) dto.StockDispensedResponseDTO {
	return dto.StockDispensedResponseDTO{
		ID:                   m.ID,
		DspSystemCode:        m.DspSystemCode,
		DspFacilityCode:      m.DspFacilityCode,
		DspTimestamp:         m.DspTimestamp,
		DspDispenseDate:      m.DspDispenseDate,
		DspProductCode:       m.DspProductCode,
		DspBatchNumber:       m.DspBatchNumber,
		DspDispensedQuantity: m.DspDispensedQuantity,
		DspPatientHash:       m.DspPatientHash,
		DspExpiryDate:        m.DspExpiryDate,
		ValidationStatus:     m.ValidationStatus,
		SyncStatus:           m.SyncStatus,
	}
}
