package dto

import "time"

type StockDispensedCreateDTO struct {
	DspSystemCode        string    `json:"dsp_system_code" binding:"required"`
	DspFacilityCode      string    `json:"dsp_facility_code" binding:"required"`
	DspDispenseDate      time.Time `json:"dsp_dispense_date" binding:"required"`
	DspProductCode       string    `json:"dsp_product_code" binding:"required"`
	DspBatchNumber       string    `json:"dsp_batch_number" binding:"required"`
	DspDispensedQuantity float64   `json:"dsp_dispensed_quantity" binding:"required"`
	DspPatientHash       string    `json:"dsp_patient_hash" binding:"required"`
	DspExpiryDate        time.Time `json:"dsp_expiry_date" binding:"required"`
}

type StockDispensedUpdateDTO struct {
	DspDispensedQuantity *float64 `json:"dsp_dispensed_quantity,omitempty" binding:"omitempty"`
}

type StockDispensedResponseDTO struct {
	ID                   uint64    `json:"id"`
	DspSystemCode        string    `json:"dsp_system_code"`
	DspFacilityCode      string    `json:"dsp_facility_code"`
	DspTimestamp         time.Time `json:"dsp_timestamp"`
	DspDispenseDate      time.Time `json:"dsp_dispense_date"`
	DspProductCode       string    `json:"dsp_product_code"`
	DspBatchNumber       string    `json:"dsp_batch_number"`
	DspDispensedQuantity float64   `json:"dsp_dispensed_quantity"`
	DspPatientHash       string    `json:"dsp_patient_hash"`
	DspExpiryDate        time.Time `json:"dsp_expiry_date"`
	ValidationStatus     int16     `json:"validation_status"`
	SyncStatus           int16     `json:"sync_status"`
}
