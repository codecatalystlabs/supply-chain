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
	DspPatientHash       string  `gorm:"size:100;not null"`
	DspExpiryDate        time.Time
	BaseModel
}

func (StockDispensed) TableName() string {
	return "Stg_Stock_Dispensed"
}
