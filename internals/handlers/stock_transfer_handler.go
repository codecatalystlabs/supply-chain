package handlers

import (
	"fmt"
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create stock transfers (bulk)
// @Tags StockTransfer
// @Accept json
// @Produce json
// @Param payload body []dto.StockTransferCreateDTO true "List of transfer payloads"
// @Success 201 {array} dto.StockTransferResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/transfers [post]
func CreateStockTransfer(c *gin.Context) {
	var payloads []dto.StockTransferCreateDTO
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(payloads) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload must be a non-empty array"})
		return
	}

	var created []dto.StockTransferResponseDTO

	for _, payload := range payloads {
		// Validate transfer type requirements
		if payload.TransferType == "intra_facility" {
			if payload.FromPharmacyID == nil || payload.ToPharmacyID == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "FromPharmacyID and ToPharmacyID are required for intra-facility transfers"})
				return
			}
			if payload.FromFacilityID != payload.ToFacilityID {
				c.JSON(http.StatusBadRequest, gin.H{"error": "FromFacilityID and ToFacilityID must be the same for intra-facility transfers"})
				return
			}
		}

		// Verify facilities exist
		var fromFacility, toFacility models.Facility
		if err := config.DB.First(&fromFacility, payload.FromFacilityID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "From facility not found"})
			return
		}
		if err := config.DB.First(&toFacility, payload.ToFacilityID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "To facility not found"})
			return
		}

		// Verify pharmacies exist if provided
		if payload.FromPharmacyID != nil {
			var fromPharmacy models.Pharmacy
			if err := config.DB.Where("id = ? AND facility_id = ?", *payload.FromPharmacyID, payload.FromFacilityID).First(&fromPharmacy).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "From pharmacy not found or does not belong to from facility"})
				return
			}
		}
		if payload.ToPharmacyID != nil {
			var toPharmacy models.Pharmacy
			if err := config.DB.Where("id = ? AND facility_id = ?", *payload.ToPharmacyID, payload.ToFacilityID).First(&toPharmacy).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "To pharmacy not found or does not belong to to facility"})
				return
			}
		}

		// Generate transfer reference
		transferRef := fmt.Sprintf("TRF-%s-%s-%d", fromFacility.FacilityCode, toFacility.FacilityCode, time.Now().Unix())

		transfer := models.StockTransfer{
			TransferRef:    transferRef,
			TransferType:   payload.TransferType,
			FromFacilityID: payload.FromFacilityID,
			FromPharmacyID: payload.FromPharmacyID,
			ToFacilityID:   payload.ToFacilityID,
			ToPharmacyID:   payload.ToPharmacyID,
			ProductCode:    payload.ProductCode,
			BatchNumber:    payload.BatchNumber,
			Quantity:       payload.Quantity,
			ExpiryDate:     payload.ExpiryDate,
			TransferDate:   payload.TransferDate,
			Status:         "pending",
			RequestedBy:    payload.RequestedBy,
			Notes:          payload.Notes,
			CreatedAt:      time.Now(),
		}

		if err := config.DB.Create(&transfer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Reload with relationships
		config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy").First(&transfer, transfer.ID)
		created = append(created, mapToStockTransferResponse(transfer))
	}

	c.JSON(http.StatusCreated, created)
}

// @Summary List stock transfers
// @Tags StockTransfer
// @Produce json
// @Param transfer_type query string false "Filter by transfer type"
// @Param from_facility_id query int false "Filter by from facility ID"
// @Param to_facility_id query int false "Filter by to facility ID"
// @Param status query string false "Filter by status"
// @Success 200 {array} dto.StockTransferResponseDTO
// @Failure 500 {object} map[string]string
// @Router /stock/transfers [get]
func ListStockTransfers(c *gin.Context) {
	var transfers []models.StockTransfer
	query := config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy")

	// Apply filters
	if transferType := c.Query("transfer_type"); transferType != "" {
		query = query.Where("transfer_type = ?", transferType)
	}
	if fromFacilityID := c.Query("from_facility_id"); fromFacilityID != "" {
		query = query.Where("from_facility_id = ?", fromFacilityID)
	}
	if toFacilityID := c.Query("to_facility_id"); toFacilityID != "" {
		query = query.Where("to_facility_id = ?", toFacilityID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&transfers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.StockTransferResponseDTO
	for _, t := range transfers {
		resp = append(resp, mapToStockTransferResponse(t))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get stock transfer by ID
// @Tags StockTransfer
// @Produce json
// @Param id path int true "Transfer ID"
// @Success 200 {object} dto.StockTransferResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/transfers/{id} [get]
func GetStockTransfer(c *gin.Context) {
	id := c.Param("id")
	var transfer models.StockTransfer
	if err := config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy").First(&transfer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}
	c.JSON(http.StatusOK, mapToStockTransferResponse(transfer))
}

// @Summary Update stock transfer
// @Tags StockTransfer
// @Accept json
// @Produce json
// @Param id path int true "Transfer ID"
// @Param payload body dto.StockTransferUpdateDTO true "Update payload"
// @Success 200 {object} dto.StockTransferResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/transfers/{id} [put]
func UpdateStockTransfer(c *gin.Context) {
	id := c.Param("id")
	var payload dto.StockTransferUpdateDTO
	var transfer models.StockTransfer

	if err := config.DB.First(&transfer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.Status != nil {
		transfer.Status = *payload.Status
	}
	if payload.ApprovedBy != nil {
		transfer.ApprovedBy = payload.ApprovedBy
	}
	if payload.ReceivedBy != nil {
		transfer.ReceivedBy = payload.ReceivedBy
	}
	if payload.ReceivedAt != nil {
		transfer.ReceivedAt = payload.ReceivedAt
	}
	if payload.Notes != nil {
		transfer.Notes = payload.Notes
	}
	now := time.Now()
	transfer.UpdatedAt = &now

	if err := config.DB.Save(&transfer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy").First(&transfer, transfer.ID)
	c.JSON(http.StatusOK, mapToStockTransferResponse(transfer))
}

// @Summary Approve stock transfer
// @Tags StockTransfer
// @Param id path int true "Transfer ID"
// @Param payload body map[string]string true "Approve payload" SchemaExample({"approved_by": "Pharmacy Manager"})
// @Success 200 {object} dto.StockTransferResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/transfers/{id}/approve [post]
func ApproveStockTransfer(c *gin.Context) {
	id := c.Param("id")
	var transfer models.StockTransfer
	var payload struct {
		ApprovedBy string `json:"approved_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&transfer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	transfer.Status = "in_transit"
	transfer.ApprovedBy = &payload.ApprovedBy
	now := time.Now()
	transfer.UpdatedAt = &now

	if err := config.DB.Save(&transfer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy").First(&transfer, transfer.ID)
	c.JSON(http.StatusOK, mapToStockTransferResponse(transfer))
}

// @Summary Receive stock transfer
// @Tags StockTransfer
// @Param id path int true "Transfer ID"
// @Param payload body map[string]string true "Receive payload" SchemaExample({"received_by": "Store Keeper"})
// @Success 200 {object} dto.StockTransferResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/transfers/{id}/receive [post]
func ReceiveStockTransfer(c *gin.Context) {
	id := c.Param("id")
	var transfer models.StockTransfer
	var payload struct {
		ReceivedBy string `json:"received_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&transfer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	transfer.Status = "completed"
	transfer.ReceivedBy = &payload.ReceivedBy
	now := time.Now()
	transfer.ReceivedAt = &now
	transfer.UpdatedAt = &now

	if err := config.DB.Save(&transfer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("FromFacility").Preload("ToFacility").Preload("FromPharmacy").Preload("ToPharmacy").First(&transfer, transfer.ID)
	c.JSON(http.StatusOK, mapToStockTransferResponse(transfer))
}

// Helper function
func mapToStockTransferResponse(t models.StockTransfer) dto.StockTransferResponseDTO {
	var fromPharmacyCode, toPharmacyCode *string
	if t.FromPharmacy != nil {
		fromPharmacyCode = &t.FromPharmacy.PharmacyCode
	}
	if t.ToPharmacy != nil {
		toPharmacyCode = &t.ToPharmacy.PharmacyCode
	}

	var fromFacilityCode, toFacilityCode string
	if t.FromFacility.ID != 0 {
		fromFacilityCode = t.FromFacility.FacilityCode
	}
	if t.ToFacility.ID != 0 {
		toFacilityCode = t.ToFacility.FacilityCode
	}

	return dto.StockTransferResponseDTO{
		ID:               t.ID,
		TransferRef:      t.TransferRef,
		TransferType:     t.TransferType,
		FromFacilityID:   t.FromFacilityID,
		FromFacilityCode: fromFacilityCode,
		FromPharmacyID:   t.FromPharmacyID,
		FromPharmacyCode: fromPharmacyCode,
		ToFacilityID:     t.ToFacilityID,
		ToFacilityCode:   toFacilityCode,
		ToPharmacyID:     t.ToPharmacyID,
		ToPharmacyCode:   toPharmacyCode,
		ProductCode:      t.ProductCode,
		BatchNumber:      t.BatchNumber,
		Quantity:         t.Quantity,
		ExpiryDate:       t.ExpiryDate,
		TransferDate:     t.TransferDate,
		Status:           t.Status,
		RequestedBy:      t.RequestedBy,
		ApprovedBy:       t.ApprovedBy,
		ReceivedBy:       t.ReceivedBy,
		ReceivedAt:       t.ReceivedAt,
		Notes:            t.Notes,
		CreatedAt:        t.CreatedAt,
		UpdatedAt:        t.UpdatedAt,
	}
}
