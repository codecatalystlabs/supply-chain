package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
)

/* ========= CREATE ========= */
// @Summary Create stock on hand
// @Tags StockOnHand
// @Accept json
// @Produce json
// @Param payload body dto.StockOnHandCreateDTO true "Stock payload"
// @Success 201 {object} dto.StockOnHandResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/on-hand [post]
func CreateStockOnHand(c *gin.Context) {
	var payload dto.StockOnHandCreateDTO

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.StockOnHand{
		SrcSystemCode:   payload.SrcSystemCode,
		SrcFacilityCode: payload.SrcFacilityCode,
		SrcProductCode:  payload.SrcProductCode,
		SrcBatchNumber:  payload.SrcBatchNumber,
		SrcQuantity:     payload.SrcQuantity,
		SrcExpiryDate:   payload.SrcExpiryDate,
		SrcTimestamp:    time.Now(),
	}

	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapToStockOnHandResponse(record))
}

/* ========= LIST ========= */
// @Summary List stock on hand
// @Tags StockOnHand
// @Produce json
// @Success 200 {array} dto.StockOnHandResponseDTO
// @Failure 500 {object} map[string]string
// @Router /stock/on-hand [get]
func ListStockOnHand(c *gin.Context) {
	var records []models.StockOnHand

	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var response []dto.StockOnHandResponseDTO
	for _, r := range records {
		response = append(response, mapToStockOnHandResponse(r))
	}

	c.JSON(200, response)
}

/* ========= GET BY ID ========= */
// @Summary List stock on hand
// @Tags StockOnHand
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.StockOnHandResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/on-hand/{id} [get]
func GetStockOnHand(c *gin.Context) {
	id := c.Param("id")
	var record models.StockOnHand

	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, mapToStockOnHandResponse(record))
}

/* ========= UPDATE ========= */
// @Summary Update stock on hand
// @Tags StockOnHand
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.StockOnHandUpdateDTO true "Stock payload"
// @Success 200 {object} dto.StockOnHandResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/on-hand/{id} [put]
func UpdateStockOnHand(c *gin.Context) {
	id := c.Param("id")
	var payload dto.StockOnHandUpdateDTO
	var record models.StockOnHand

	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if payload.SrcQuantity > 0 {
		record.SrcQuantity = payload.SrcQuantity
	}
	if !payload.SrcExpiryDate.IsZero() {
		record.SrcExpiryDate = payload.SrcExpiryDate
	}

	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, mapToStockOnHandResponse(record))
}

/* ========= DELETE ========= */
// @Summary Delete stock on hand
// @Tags StockOnHand
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/on-hand/{id} [delete]
func DeleteStockOnHand(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.StockOnHand{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

/* ========= MAPPER ========= */
func mapToStockOnHandResponse(m models.StockOnHand) dto.StockOnHandResponseDTO {
	return dto.StockOnHandResponseDTO{
		ID:               m.ID,
		SrcSystemCode:    m.SrcSystemCode,
		SrcFacilityCode:  m.SrcFacilityCode,
		SrcTimestamp:     m.SrcTimestamp,
		SrcProductCode:   m.SrcProductCode,
		SrcBatchNumber:   m.SrcBatchNumber,
		SrcQuantity:      m.SrcQuantity,
		SrcExpiryDate:    m.SrcExpiryDate,
		ValidationStatus: m.ValidationStatus,
		SyncStatus:       m.SyncStatus,
	}
}
