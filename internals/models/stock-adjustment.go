package models

import "time"

type StockAdjustment struct {
	ID                  uint64    `gorm:"primaryKey"`
	AdjSystemCode       string    `gorm:"size:100;not null"`
	AdjFacilityCode     string    `gorm:"size:100;not null"`
	AdjTimestamp        time.Time `gorm:"not null"`
	AdjAdjustmentDate   time.Time `gorm:"not null"`
	AdjAdjustmentType   string    `gorm:"size:100;not null"`
	AdjAdjustmentReason string    `gorm:"size:100;not null"`
	AdjProductCode      string    `gorm:"size:100;not null"`
	AdjBatchNumber      string    `gorm:"size:100;not null"`
	AdjQuantity         int       `gorm:"not null"`
	AdjExpiryDate       time.Time `gorm:"not null"`
	BaseModel
}

func (StockAdjustment) TableName() string {
	return "Stg_Stock_Adjustment"
}
