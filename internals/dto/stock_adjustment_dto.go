package dto

import "time"

type StockAdjustmentCreateDTO struct {
	AdjSystemCode       string    `json:"adj_system_code" binding:"required"`
	AdjFacilityCode     string    `json:"adj_facility_code" binding:"required"`
	AdjAdjustmentDate   time.Time `json:"adj_adjustment_date" binding:"required"`
	AdjAdjustmentType   string    `json:"adj_adjustment_type" binding:"required"`
	AdjAdjustmentReason string    `json:"adj_adjustment_reason" binding:"required"`
	AdjProductCode      string    `json:"adj_product_code" binding:"required"`
	AdjBatchNumber      string    `json:"adj_batch_number" binding:"required"`
	AdjQuantity         int       `json:"adj_quantity" binding:"required"`
	AdjExpiryDate       time.Time `json:"adj_expiry_date" binding:"required"`
}

type StockAdjustmentUpdateDTO struct {
	AdjQuantity *int `json:"adj_quantity,omitempty" binding:"omitempty"`
}

type StockAdjustmentResponseDTO struct {
	ID                  uint64    `json:"id"`
	AdjSystemCode       string    `json:"adj_system_code"`
	AdjFacilityCode     string    `json:"adj_facility_code"`
	AdjTimestamp        time.Time `json:"adj_timestamp"`
	AdjAdjustmentDate   time.Time `json:"adj_adjustment_date"`
	AdjAdjustmentType   string    `json:"adj_adjustment_type"`
	AdjAdjustmentReason string    `json:"adj_adjustment_reason"`
	AdjProductCode      string    `json:"adj_product_code"`
	AdjBatchNumber      string    `json:"adj_batch_number"`
	AdjQuantity         int       `json:"adj_quantity"`
	AdjExpiryDate       time.Time `json:"adj_expiry_date"`
	ValidationStatus    int16     `json:"validation_status"`
	SyncStatus          int16     `json:"sync_status"`
}
