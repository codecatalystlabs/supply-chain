package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
)

/* ========== CREATE PURCHASE ORDER ========== */
// @Summary Create purchase order
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param payload body dto.PurchaseOrderCreateDTO true "Purchase order payload"
// @Success 201 {object} dto.PurchaseOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchase-order [post]
func CreatePurchaseOrder(c *gin.Context) {
	var input dto.PurchaseOrderCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.PurchaseOrder{
		OrdSystemCode:      input.OrdSystemCode,
		OrdFacilityCode:    input.OrdFacilityCode,
		OrdTimestamp:       time.Now(),
		OrdOrderDate:       input.OrdOrderDate,
		OrdOrderRefNumber:  input.OrdOrderRefNumber,
		OrdOrderNumber:     input.OrdOrderNumber,
		OrdProductCode:     input.OrdProductCode,
		OrdOrderedQuantity: input.OrdOrderedQuantity,
	}

	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, mapToPurchaseOrderResponse(record))
}

/* ========== GET PURCHASE ORDER BY ID ========== */
// @Summary Get purchase order by id
// @Tags PurchaseOrder
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.PurchaseOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchase-order/{id} [get]
func GetPurchaseOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var record models.PurchaseOrder
	if err := config.DB.First(&record, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, mapToPurchaseOrderResponse(record))
}

/* ========== MAP TO RESPONSE DTO ========== */
func mapToPurchaseOrderResponse(record models.PurchaseOrder) dto.PurchaseOrderResponseDTO {
	return dto.PurchaseOrderResponseDTO{
		ID:                 record.ID,
		OrdSystemCode:      record.OrdSystemCode,
		OrdFacilityCode:    record.OrdFacilityCode,
		OrdTimestamp:       record.OrdTimestamp,
		OrdOrderDate:       record.OrdOrderDate,
		OrdOrderRefNumber:  record.OrdOrderRefNumber,
		OrdOrderNumber:     record.OrdOrderNumber,
		OrdProductCode:     record.OrdProductCode,
		OrdOrderedQuantity: record.OrdOrderedQuantity,
		ValidationStatus:   record.ValidationStatus,
		SyncStatus:         record.SyncStatus,
	}
}

// Get purchase order by id
func GetPurchaseOrderByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var record models.PurchaseOrder
	if err := config.DB.First(&record, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, mapToPurchaseOrderResponse(record))
}

// Update purchase order
// @Summary Update purchase order
// @Tags PurchaseOrder
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.PurchaseOrderUpdateDTO true "Update payload"
// @Success 200 {object} dto.PurchaseOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchase-order/{id} [put]
func UpdatePurchaseOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var payload dto.PurchaseOrderUpdateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var record models.PurchaseOrder
	if err := config.DB.First(&record, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if payload.OrdOrderedQuantity != nil {
		record.OrdOrderedQuantity = *payload.OrdOrderedQuantity
	}

	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapToPurchaseOrderResponse(record))
}

// Delete purchase order
// @Summary Delete purchase order
// @Tags PurchaseOrder
// @Param id path int true "ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /purchase-order/{id} [delete]
func DeletePurchaseOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := config.DB.Delete(&models.PurchaseOrder{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// List purchase orders
// @Summary List purchase orders
// @Tags PurchaseOrder
// @Produce json
// @Success 200 {array} dto.PurchaseOrderResponseDTO
// @Failure 500 {object} map[string]string
// @Router /purchase-order [get]
func ListPurchaseOrders(c *gin.Context) {
	var records []models.PurchaseOrder
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.PurchaseOrderResponseDTO
	for _, record := range records {
		response = append(response, mapToPurchaseOrderResponse(record))
	}

	c.JSON(http.StatusOK, response)
}
