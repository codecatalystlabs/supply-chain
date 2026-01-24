package dto

import (
	"time"
)

/* ========= CREATE ========= */
type StockOnHandCreateDTO struct {
	SrcSystemCode   string    `json:"src_system_code" binding:"required"`
	SrcFacilityCode string    `json:"src_facility_code" binding:"required"`
	SrcProductCode  string    `json:"src_product_code" binding:"required"`
	SrcBatchNumber  string    `json:"src_batch_number" binding:"required"`
	SrcQuantity     int       `json:"src_quantity" binding:"required,gt=0"`
	SrcExpiryDate   time.Time `json:"src_expiry_date" binding:"required"`
}

/* ========= UPDATE ========= */
type StockOnHandUpdateDTO struct {
	SrcQuantity   int       `json:"src_quantity" binding:"omitempty,gt=0"`
	SrcExpiryDate time.Time `json:"src_expiry_date" binding:"omitempty"`
}

/* ========= RESPONSE ========= */
type StockOnHandResponseDTO struct {
	ID               uint64    `json:"id"`
	SrcSystemCode    string    `json:"src_system_code"`
	SrcFacilityCode  string    `json:"src_facility_code"`
	SrcTimestamp     time.Time `json:"src_timestamp"`
	SrcProductCode   string    `json:"src_product_code"`
	SrcBatchNumber   string    `json:"src_batch_number"`
	SrcQuantity      int       `json:"src_quantity"`
	SrcExpiryDate    time.Time `json:"src_expiry_date"`
	ValidationStatus int16     `json:"validation_status"`
	SyncStatus       int16     `json:"sync_status"`
}
