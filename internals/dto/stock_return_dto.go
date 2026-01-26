package dto

import "time"

type StockReturnCreateDTO struct {
	RtnSystemCode   string    `json:"rtn_system_code" binding:"required"`
	RtnFacilityCode string    `json:"rtn_facility_code" binding:"required"`
	RtnReturnDate   time.Time `json:"rtn_return_date" binding:"required"`
	RtnReturnNumber string    `json:"rtn_return_number" binding:"required"`
	RtnProductCode  string    `json:"rtn_product_code" binding:"required"`
	RtnBatchNumber  string    `json:"rtn_batch_number" binding:"required"`
	RtnUnitCode     string    `json:"rtn_unit_code" binding:"required"`
	RtnQuantity     int       `json:"rtn_quantity" binding:"required"`
}

type StockReturnUpdateDTO struct {
	RtnQuantity *int `json:"rtn_quantity,omitempty" binding:"omitempty"`
}

type StockReturnResponseDTO struct {
	ID               uint64    `json:"id"`
	RtnSystemCode    string    `json:"rtn_system_code"`
	RtnFacilityCode  string    `json:"rtn_facility_code"`
	RtnTimestamp     time.Time `json:"rtn_timestamp"`
	RtnReturnNumber  string    `json:"rtn_return_number"`
	RtnReturnDate    time.Time `json:"rtn_return_date"`
	RtnProductCode   string    `json:"rtn_product_code"`
	RtnBatchNumber   string    `json:"rtn_batch_number"`
	RtnUnitCode      string    `json:"rtn_unit_code"`
	RtnQuantity      int       `json:"rtn_quantity"`
	ValidationStatus int16     `json:"validation_status"`
	SyncStatus       int16     `json:"sync_status"`
}
