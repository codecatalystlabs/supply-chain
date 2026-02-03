package models

import "time"

type StockDispensed struct {
	ID                   uint64 `gorm:"primaryKey"`
	DspSystemCode        string `gorm:"size:100;not null"`
	DspFacilityCode      string `gorm:"size:100;not null"`
	DspTimestamp         time.Time
	DspDispenseDate      time.Time
	DspProductCode       string  `gorm:"size:100;not null"`
	DspBatchNumber       string  `gorm:"size:100;not null"`
	DspDispensedQuantity float64 `gorm:"not null"`
	// Optional link to a Prescription and the prescribed quantity
	DspPrescriptionID     *uint64  `gorm:"column:dsp_prescription_id" json:"dsp_prescription_id,omitempty"`
	DspPrescribedQuantity *float64 `gorm:"column:dsp_prescribed_quantity" json:"dsp_prescribed_quantity,omitempty"`
	// Comment for cases where dispensed < prescribed
	DspComment     *string `gorm:"column:dsp_comment;size:255" json:"dsp_comment,omitempty"`
	DspPatientHash string  `gorm:"size:100;not null"`
	DspExpiryDate  time.Time
	BaseModel
}

func (StockDispensed) TableName() string {
	return "Stg_Stock_Dispensed"
}
