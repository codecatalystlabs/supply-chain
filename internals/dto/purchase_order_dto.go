package dto

import (
	"time"
)

/* ========= CREATE ========= */
type PurchaseOrderCreateDTO struct {
	OrdSystemCode      string    `json:"ord_system_code" binding:"required"`
	OrdFacilityCode    string    `json:"ord_facility_code" binding:"required"`
	OrdOrderDate       time.Time `json:"ord_order_date" binding:"required"`
	OrdOrderRefNumber  string    `json:"ord_order_ref_number" binding:"required"`
	OrdOrderNumber     string    `json:"ord_order_number" binding:"required"`
	OrdProductCode     string    `json:"ord_product_code" binding:"required"`
	OrdOrderedQuantity int       `json:"ord_ordered_quantity" binding:"required,gt=0"`
}

/* ========= UPDATE ========= */
type PurchaseOrderUpdateDTO struct {
	OrdOrderedQuantity *int `json:"ord_ordered_quantity" binding:"omitempty,gt=0"`
}

/* ========= RESPONSE ========= */
type PurchaseOrderResponseDTO struct {
	ID                 uint64    `json:"id"`
	OrdSystemCode      string    `json:"ord_system_code"`
	OrdFacilityCode    string    `json:"ord_facility_code"`
	OrdTimestamp       time.Time `json:"ord_timestamp"`
	OrdOrderDate       time.Time `json:"ord_order_date"`
	OrdOrderRefNumber  string    `json:"ord_order_ref_number"`
	OrdOrderNumber     string    `json:"ord_order_number"`
	OrdProductCode     string    `json:"ord_product_code"`
	OrdOrderedQuantity int       `json:"ord_ordered_quantity"`
	ValidationStatus   int16     `json:"validation_status"`
	SyncStatus         int16     `json:"sync_status"`
}
