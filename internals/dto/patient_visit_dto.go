package dto

import "time"

type PatientVisitCreateDTO struct {
	VstSystemCode      string    `json:"vst_system_code" binding:"required"`
	VstFacilityCode    string    `json:"vst_facility_code" binding:"required"`
	VstVisitDate       time.Time `json:"vst_visit_date" binding:"required"`
	VstPatientCode     string    `json:"vst_patient_code" binding:"required"`
	VstSex             string    `json:"vst_sex" binding:"required"`
	VstAge             int16     `json:"vst_age" binding:"required"`
	VstProductCode     string    `json:"vst_product_code" binding:"required"`
	VstBatchNumber     string    `json:"vst_batch_number" binding:"required"`
	VstQuantity        float64   `json:"vst_quantity" binding:"required"`
	VstRegimenCode     string    `json:"vst_regimen_code" binding:"required"`
	VstPatientCategory string    `json:"vst_patient_category" binding:"required"`
}

type PatientVisitUpdateDTO struct {
	VstQuantity *float64 `json:"vst_quantity,omitempty" binding:"omitempty"`
}

type PatientVisitResponseDTO struct {
	ID                 uint64    `json:"id"`
	VstSystemCode      string    `json:"vst_system_code"`
	VstFacilityCode    string    `json:"vst_facility_code"`
	VstTimestamp       time.Time `json:"vst_timestamp"`
	VstPatientCode     string    `json:"vst_patient_code"`
	VstSex             string    `json:"vst_sex"`
	VstAge             int16     `json:"vst_age"`
	VstVisitDate       time.Time `json:"vst_visit_date"`
	VstProductCode     string    `json:"vst_product_code"`
	VstBatchNumber     string    `json:"vst_batch_number"`
	VstQuantity        float64   `json:"vst_quantity"`
	VstRegimenCode     string    `json:"vst_regimen_code"`
	VstPatientCategory string    `json:"vst_patient_category"`
	ValidationStatus   int16     `json:"validation_status"`
	SyncStatus         int16     `json:"sync_status"`
}
