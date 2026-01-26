package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create goods receipt
// @Tags GoodsReceipt
// @Accept json
// @Produce json
// @Param payload body dto.GoodsReceiptCreateDTO true "Goods receipt payload"
// @Success 201 {object} dto.GoodsReceiptResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/goods-receipt [post]
func CreateGoodsReceipt(c *gin.Context) {
	var payload dto.GoodsReceiptCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := models.GoodsReceipt{
		GrnSystemCode:            payload.GrnSystemCode,
		GrnFacilityCode:          payload.GrnFacilityCode,
		GrnReceiptDate:           payload.GrnReceiptDate,
		GrnFacilityReceiptNumber: payload.GrnFacilityReceiptNumber,
		GrnWarehouseRefNumber:    payload.GrnWarehouseRefNumber,
		GrnOrderNumber:           payload.GrnOrderNumber,
		GrnProductCode:           payload.GrnProductCode,
		GrnBatchNumber:           payload.GrnBatchNumber,
		GrnQuantity:              payload.GrnQuantity,
		GrnExpiryDate:            payload.GrnExpiryDate,
		GrnSupplierCode:          payload.GrnSupplierCode,
	}
	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapToGoodsReceiptResponse(record))
}

// @Summary List goods receipts
// @Tags GoodsReceipt
// @Produce json
// @Success 200 {array} dto.GoodsReceiptResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/v1/goods-receipt [get]
func ListGoodsReceipts(c *gin.Context) {
	var records []models.GoodsReceipt
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.GoodsReceiptResponseDTO
	for _, r := range records {
		resp = append(resp, mapToGoodsReceiptResponse(r))
	}
	c.JSON(200, resp)
}

// @Summary Get goods receipt by id
// @Tags GoodsReceipt
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.GoodsReceiptResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/goods-receipt/{id} [get]
func GetGoodsReceipt(c *gin.Context) {
	id := c.Param("id")
	var record models.GoodsReceipt
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToGoodsReceiptResponse(record))
}

// @Summary Update goods receipt
// @Tags GoodsReceipt
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param payload body dto.GoodsReceiptUpdateDTO true "Update payload"
// @Success 200 {object} dto.GoodsReceiptResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/goods-receipt/{id} [put]
func UpdateGoodsReceipt(c *gin.Context) {
	id := c.Param("id")
	var payload dto.GoodsReceiptUpdateDTO
	var record models.GoodsReceipt
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.GrnQuantity != nil {
		record.GrnQuantity = *payload.GrnQuantity
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToGoodsReceiptResponse(record))
}

// @Summary Delete goods receipt
// @Tags GoodsReceipt
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/goods-receipt/{id} [delete]
func DeleteGoodsReceipt(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.GoodsReceipt{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToGoodsReceiptResponse(m models.GoodsReceipt) dto.GoodsReceiptResponseDTO {
	return dto.GoodsReceiptResponseDTO{
		ID:                       m.ID,
		GrnSystemCode:            m.GrnSystemCode,
		GrnFacilityCode:          m.GrnFacilityCode,
		GrnTimestamp:             m.GrnTimestamp,
		GrnReceiptDate:           m.GrnReceiptDate,
		GrnFacilityReceiptNumber: m.GrnFacilityReceiptNumber,
		GrnWarehouseRefNumber:    m.GrnWarehouseRefNumber,
		GrnOrderNumber:           m.GrnOrderNumber,
		GrnProductCode:           m.GrnProductCode,
		GrnBatchNumber:           m.GrnBatchNumber,
		GrnQuantity:              m.GrnQuantity,
		GrnExpiryDate:            m.GrnExpiryDate,
		GrnSupplierCode:          m.GrnSupplierCode,
		ValidationStatus:         m.ValidationStatus,
		SyncStatus:               m.SyncStatus,
	}
}
