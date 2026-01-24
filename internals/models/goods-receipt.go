package models

import "time"

type GoodsReceipt struct {
	ID                       uint64    `gorm:"primaryKey"`
	GrnSystemCode            string    `gorm:"size:100;not null"`
	GrnFacilityCode          string    `gorm:"size:100;not null"`
	GrnTimestamp             time.Time `gorm:"not null"`
	GrnReceiptDate           time.Time `gorm:"not null"`
	GrnFacilityReceiptNumber string    `gorm:"size:100;not null"`
	GrnWarehouseRefNumber    string    `gorm:"size:100;not null"`
	GrnOrderNumber           string    `gorm:"size:100;not null"`
	GrnProductCode           string    `gorm:"size:100;not null"`
	GrnBatchNumber           string    `gorm:"size:100;not null"`
	GrnQuantity              int       `gorm:"not null"`
	GrnExpiryDate            time.Time `gorm:"not null"`
	GrnSupplierCode          string    `gorm:"size:100;not null"`
	ValidationStatus         int16     `gorm:"default:0"`
	ValidationMessage        *string   `gorm:"size:100"`
	SyncStatus               int16     `gorm:"default:0"`
	AddDate                  *time.Time
}

func (GoodsReceipt) TableName() string {
	return "Stg_Goods_Receipt"
}
