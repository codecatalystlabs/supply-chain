package models

import "time"

type BaseModel struct {
	ValidationStatus  int16      `gorm:"column:validation_status;default:0" json:"validation_status"`
	ValidationMessage *string    `gorm:"column:validation_message;size:100" json:"validation_message,omitempty"`
	SyncStatus        int16      `gorm:"column:sync_status;default:0" json:"sync_status"`
	AddDate           *time.Time `gorm:"column:add_date" json:"add_date,omitempty"`
}

type StockOnHand struct {
	ID              uint64    `gorm:"column:id;primaryKey" json:"id"`
	SrcSystemCode   string    `gorm:"column:src_system_code;size:100;not null" json:"src_system_code"`
	SrcFacilityCode string    `gorm:"column:src_facility_code;size:100;not null" json:"src_facility_code"`
	SrcTimestamp    time.Time `gorm:"column:src_timestamp;not null" json:"src_timestamp"`
	SrcProductCode  string    `gorm:"column:src_product_code;size:100;not null" json:"src_product_code"`
	SrcBatchNumber  string    `gorm:"column:src_batch_number;size:100;not null" json:"src_batch_number"`
	SrcQuantity     int       `gorm:"column:src_quantity;not null" json:"src_quantity"`
	SrcExpiryDate   time.Time `gorm:"column:src_expiry_date;not null" json:"src_expiry_date"`
	BaseModel
}

func (StockOnHand) TableName() string {
	return "Stg_Stock_On_Hand"
}
