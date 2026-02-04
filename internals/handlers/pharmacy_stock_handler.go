package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create pharmacy stock (bulk)
// @Tags PharmacyStock
// @Accept json
// @Produce json
// @Param payload body []dto.PharmacyStockCreateDTO true "List of pharmacy stock payloads"
// @Success 201 {array} dto.PharmacyStockResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /pharmacy-stock [post]
func CreatePharmacyStock(c *gin.Context) {
	var payloads []dto.PharmacyStockCreateDTO
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(payloads) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload must be a non-empty array"})
		return
	}

	var created []dto.PharmacyStockResponseDTO
	for _, payload := range payloads {
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
		created = append(created, mapToPharmacyStockResponse(record))
	}

	c.JSON(http.StatusCreated, created)
}

// @Summary List pharmacy stock
// @Tags PharmacyStock
// @Produce json
// @Success 200 {array} dto.PharmacyStockResponseDTO
// @Failure 500 {object} map[string]string
// @Router /pharmacy-stock [get]
func ListPharmacyStock(c *gin.Context) {
	var records []models.PharmacyStock
	query := config.DB

	// Apply facility filter if user is facility-scoped
	if user, exists := c.Get("user"); exists {
		u := user.(*models.User)
		if u.FacilityID != nil && !u.HasRole("super_admin") {
			// Get facility code
			var facility models.Facility
			if err := config.DB.First(&facility, *u.FacilityID).Error; err == nil {
				query = query.Where("pha_facility_code = ?", facility.FacilityCode)
			}
		}
	}

	if err := query.Find(&records).Error; err != nil {
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
// @Router /pharmacy-stock/{id} [get]
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
// @Router /pharmacy-stock/{id} [put]
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
// @Router /pharmacy-stock/{id} [delete]
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
