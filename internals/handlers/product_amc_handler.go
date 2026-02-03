package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create product AMC
// @Tags ProductAmc
// @Accept json
// @Produce json
// @Param payload body []dto.ProductAmcCreateDTO true "Product AMC payloads"
// @Success 201 {array} dto.ProductAmcResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-amc [post]
func CreateProductAmc(c *gin.Context) {
	var payloads []dto.ProductAmcCreateDTO
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var responses []dto.ProductAmcResponseDTO
	for _, payload := range payloads {
		record := models.ProductAmc{
			AmcSystemCode:   payload.AmcSystemCode,
			AmcFacilityCode: payload.AmcFacilityCode,
			AmcDate:         payload.AmcDate,
			AmcProductCode:  payload.AmcProductCode,
			AmcProductName:  payload.AmcProductName,
			AmcMonth:        payload.AmcMonth,
			AmcYear:         payload.AmcYear,
			AmcValue:        payload.AmcValue,
		}

		if err := config.DB.Create(&record).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responses = append(responses, mapToProductAmcResponse(record))
	}

	c.JSON(http.StatusCreated, responses)
}

// @Summary List product AMC
// @Tags ProductAmc
// @Produce json
// @Success 200 {array} dto.ProductAmcResponseDTO
// @Failure 500 {object} map[string]string
// @Router /product-amc [get]
func ListProductAmc(c *gin.Context) {
	var records []models.ProductAmc
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ProductAmcResponseDTO
	for _, r := range records {
		resp = append(resp, mapToProductAmcResponse(r))
	}
	c.JSON(200, resp)
}

// @Summary Get product AMC by id
// @Tags ProductAmc
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.ProductAmcResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-amc/{id} [get]
func GetProductAmc(c *gin.Context) {
	id := c.Param("id")
	var record models.ProductAmc
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToProductAmcResponse(record))
}

// @Summary Update product AMC
// @Tags ProductAmc
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.ProductAmcUpdateDTO true "Update payload"
// @Success 200 {object} dto.ProductAmcResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-amc/{id} [put]
func UpdateProductAmc(c *gin.Context) {
	id := c.Param("id")
	var payload dto.ProductAmcUpdateDTO
	var record models.ProductAmc
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.AmcValue != nil {
		record.AmcValue = *payload.AmcValue
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToProductAmcResponse(record))
}

// @Summary Delete product AMC
// @Tags ProductAmc
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-amc/{id} [delete]
func DeleteProductAmc(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.ProductAmc{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToProductAmcResponse(m models.ProductAmc) dto.ProductAmcResponseDTO {
	return dto.ProductAmcResponseDTO{
		ID:               m.ID,
		AmcSystemCode:    m.AmcSystemCode,
		AmcFacilityCode:  m.AmcFacilityCode,
		AmcTimestamp:     m.AmcTimestamp,
		AmcProductCode:   m.AmcProductCode,
		AmcProductName:   m.AmcProductName,
		AmcDate:          m.AmcDate,
		AmcMonth:         m.AmcMonth,
		AmcYear:          m.AmcYear,
		AmcValue:         m.AmcValue,
		AmcDaysOutStock:  m.AmcDaysOutStock,
		ValidationStatus: m.ValidationStatus,
		SyncStatus:       m.SyncStatus,
	}
}
