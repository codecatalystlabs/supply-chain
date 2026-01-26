package models

import "time"

type PharmacyStock struct {
	ID              uint64 `gorm:"primaryKey"`
	PhaSystemCode   string `gorm:"size:100;not null"`
	PhaFacilityCode string `gorm:"size:100;not null"`
	PhaTimestamp    time.Time
	PhaProductCode  string    `gorm:"size:100;not null"`
	PhaBatchNumber  string    `gorm:"size:100;not null"`
	PhaQuantity     int       `gorm:"not null"`
	PhaExpiryDate   time.Time `gorm:"not null"`
	BaseModel
}

func (PharmacyStock) TableName() string {
	return "Stg_Pharmacy_Stock"
}
