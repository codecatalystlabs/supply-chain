package dto

import "time"

type GoodsReceiptCreateDTO struct {
	GrnSystemCode            string    `json:"grn_system_code" binding:"required"`
	GrnFacilityCode          string    `json:"grn_facility_code" binding:"required"`
	GrnReceiptDate           time.Time `json:"grn_receipt_date" binding:"required"`
	GrnFacilityReceiptNumber string    `json:"grn_facility_receipt_number" binding:"required"`
	GrnWarehouseRefNumber    string    `json:"grn_warehouse_ref_number" binding:"required"`
	GrnOrderNumber           string    `json:"grn_order_number" binding:"required"`
	GrnProductCode           string    `json:"grn_product_code" binding:"required"`
	GrnBatchNumber           string    `json:"grn_batch_number" binding:"required"`
	GrnQuantity              int       `json:"grn_quantity" binding:"required"`
	GrnExpiryDate            time.Time `json:"grn_expiry_date" binding:"required"`
	GrnSupplierCode          string    `json:"grn_supplier_code" binding:"required"`
}

type GoodsReceiptUpdateDTO struct {
	GrnQuantity *int `json:"grn_quantity,omitempty" binding:"omitempty"`
}

type GoodsReceiptResponseDTO struct {
	ID                       uint64    `json:"id"`
	GrnSystemCode            string    `json:"grn_system_code"`
	GrnFacilityCode          string    `json:"grn_facility_code"`
	GrnTimestamp             time.Time `json:"grn_timestamp"`
	GrnReceiptDate           time.Time `json:"grn_receipt_date"`
	GrnFacilityReceiptNumber string    `json:"grn_facility_receipt_number"`
	GrnWarehouseRefNumber    string    `json:"grn_warehouse_ref_number"`
	GrnOrderNumber           string    `json:"grn_order_number"`
	GrnProductCode           string    `json:"grn_product_code"`
	GrnBatchNumber           string    `json:"grn_batch_number"`
	GrnQuantity              int       `json:"grn_quantity"`
	GrnExpiryDate            time.Time `json:"grn_expiry_date"`
	GrnSupplierCode          string    `json:"grn_supplier_code"`
	ValidationStatus         int16     `json:"validation_status"`
	SyncStatus               int16     `json:"sync_status"`
}
