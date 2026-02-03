package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

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
// @Router /stock/adjustment [post]
func CreateStockAdjustment(c *gin.Context) {
	var payload dto.StockAdjustmentCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Verify facility exists
	var facility models.Facility
	if err := config.DB.Where("facility_code = ?", payload.AdjFacilityCode).First(&facility).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Facility not found"})
		return
	}

	// Verify pharmacy exists if provided
	if payload.AdjPharmacyID != nil {
		var pharmacy models.Pharmacy
		if err := config.DB.Where("id = ? AND facility_id = ?", *payload.AdjPharmacyID, facility.ID).First(&pharmacy).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pharmacy not found or doesn't belong to facility"})
			return
		}
	}

	record := models.StockAdjustment{
		AdjSystemCode:       payload.AdjSystemCode,
		AdjFacilityCode:     payload.AdjFacilityCode,
		AdjPharmacyID:       payload.AdjPharmacyID,
		AdjTimestamp:        time.Now(),
		AdjAdjustmentDate:   payload.AdjAdjustmentDate,
		AdjAdjustmentType:   payload.AdjAdjustmentType,
		AdjAdjustmentReason: payload.AdjAdjustmentReason,
		AdjProductCode:      payload.AdjProductCode,
		AdjBatchNumber:      payload.AdjBatchNumber,
		AdjQuantity:         payload.AdjQuantity,
		AdjExpiryDate:       payload.AdjExpiryDate,
		AdjReferenceNumber:  payload.AdjReferenceNumber,
		AdjApprovedBy:       payload.AdjApprovedBy,
		AdjNotes:            payload.AdjNotes,
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
// @Param facility_code query string false "Filter by facility code"
// @Param pharmacy_id query int false "Filter by pharmacy ID"
// @Success 200 {array} dto.StockAdjustmentResponseDTO
// @Failure 500 {object} map[string]string
// @Router /stock/adjustment [get]
func ListStockAdjustments(c *gin.Context) {
	var records []models.StockAdjustment
	query := config.DB.Preload("Pharmacy")

	// Apply filters
	if facilityCode := c.Query("facility_code"); facilityCode != "" {
		query = query.Where("adj_facility_code = ?", facilityCode)
	}
	if pharmacyID := c.Query("pharmacy_id"); pharmacyID != "" {
		query = query.Where("adj_pharmacy_id = ?", pharmacyID)
	}

	if err := query.Find(&records).Error; err != nil {
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
// @Router /stock/adjustment/{id} [get]
func GetStockAdjustment(c *gin.Context) {
	id := c.Param("id")
	var record models.StockAdjustment
	if err := config.DB.Preload("Pharmacy").First(&record, id).Error; err != nil {
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
// @Router /stock/adjustment/{id} [put]
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
// @Router /stock/adjustment/{id} [delete]
func DeleteStockAdjustment(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.StockAdjustment{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToStockAdjustmentResponse(m models.StockAdjustment) dto.StockAdjustmentResponseDTO {
	var pharmacyCode, pharmacyName *string
	if m.Pharmacy != nil && m.Pharmacy.ID != 0 {
		pharmacyCode = &m.Pharmacy.PharmacyCode
		pharmacyName = &m.Pharmacy.PharmacyName
	}

	return dto.StockAdjustmentResponseDTO{
		ID:                  m.ID,
		AdjSystemCode:       m.AdjSystemCode,
		AdjFacilityCode:     m.AdjFacilityCode,
		AdjPharmacyID:       m.AdjPharmacyID,
		AdjPharmacyCode:     pharmacyCode,
		AdjPharmacyName:     pharmacyName,
		AdjTimestamp:        m.AdjTimestamp,
		AdjAdjustmentDate:   m.AdjAdjustmentDate,
		AdjAdjustmentType:   m.AdjAdjustmentType,
		AdjAdjustmentReason: m.AdjAdjustmentReason,
		AdjProductCode:      m.AdjProductCode,
		AdjBatchNumber:      m.AdjBatchNumber,
		AdjQuantity:         m.AdjQuantity,
		AdjExpiryDate:       m.AdjExpiryDate,
		AdjReferenceNumber:  m.AdjReferenceNumber,
		AdjApprovedBy:       m.AdjApprovedBy,
		AdjNotes:            m.AdjNotes,
		ValidationStatus:    m.ValidationStatus,
		SyncStatus:          m.SyncStatus,
	}
}
