package dto

import "time"

type PharmacyStockCreateDTO struct {
	PhaSystemCode   string    `json:"pha_system_code" binding:"required"`
	PhaFacilityCode string    `json:"pha_facility_code" binding:"required"`
	PhaProductCode  string    `json:"pha_product_code" binding:"required"`
	PhaBatchNumber  string    `json:"pha_batch_number" binding:"required"`
	PhaQuantity     int       `json:"pha_quantity" binding:"required"`
	PhaExpiryDate   time.Time `json:"pha_expiry_date" binding:"required"`
}

type PharmacyStockUpdateDTO struct {
	PhaQuantity *int `json:"pha_quantity,omitempty" binding:"omitempty"`
}

type PharmacyStockResponseDTO struct {
	ID               uint64    `json:"id"`
	PhaSystemCode    string    `json:"pha_system_code"`
	PhaFacilityCode  string    `json:"pha_facility_code"`
	PhaTimestamp     time.Time `json:"pha_timestamp"`
	PhaProductCode   string    `json:"pha_product_code"`
	PhaBatchNumber   string    `json:"pha_batch_number"`
	PhaQuantity      int       `json:"pha_quantity"`
	PhaExpiryDate    time.Time `json:"pha_expiry_date"`
	ValidationStatus int16     `json:"validation_status"`
	SyncStatus       int16     `json:"sync_status"`
}
