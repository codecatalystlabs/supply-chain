package models

import "time"

type ProductAmc struct {
	ID              uint64 `gorm:"primaryKey"`
	AmcSystemCode   string `gorm:"size:100;not null"`
	AmcFacilityCode string `gorm:"size:100;not null"`
	AmcTimestamp    time.Time
	AmcProductCode  string `gorm:"size:100;not null"`
	AmcProductName  string `gorm:"size:100;not null"`
	AmcDate         time.Time
	AmcMonth        int16   `gorm:"not null"`
	AmcYear         int16   `gorm:"not null"`
	AmcValue        float64 `gorm:"not null"`
	AmcDaysOutStock *float64
	BaseModel
}

func (ProductAmc) TableName() string {
	return "Stg_Product_Amc"
}
