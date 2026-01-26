package models

import "time"

type StockReturn struct {
	ID              uint64 `gorm:"primaryKey"`
	RtnSystemCode   string `gorm:"size:100;not null"`
	RtnFacilityCode string `gorm:"size:100;not null"`
	RtnTimestamp    time.Time
	RtnReturnNumber string    `gorm:"size:100;not null"`
	RtnReturnDate   time.Time `gorm:"not null"`
	RtnProductCode  string    `gorm:"size:100;not null"`
	RtnBatchNumber  string    `gorm:"size:100;not null"`
	RtnUnitCode     string    `gorm:"size:100;not null"`
	RtnQuantity     int       `gorm:"not null"`
	BaseModel
}

func (StockReturn) TableName() string {
	return "Stg_Stock_Return"
}
