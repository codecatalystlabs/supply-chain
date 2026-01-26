package models

import "time"

type PurchaseOrder struct {
	ID                 uint64    `gorm:"primaryKey" json:"id"`
	OrdSystemCode      string    `gorm:"size:100;not null" json:"ord_system_code"`
	OrdFacilityCode    string    `gorm:"size:100;not null" json:"ord_facility_code"`
	OrdTimestamp       time.Time `gorm:"not null" json:"ord_timestamp"`
	OrdOrderDate       time.Time `gorm:"not null" json:"ord_order_date"`
	OrdOrderRefNumber  string    `gorm:"size:100;not null" json:"ord_order_ref_number"`
	OrdOrderNumber     string    `gorm:"size:100;not null" json:"ord_order_number"`
	OrdProductCode     string    `gorm:"size:100;not null" json:"ord_product_code"`
	OrdOrderedQuantity int       `gorm:"not null" json:"ord_ordered_quantity"`
	BaseModel
}

func (PurchaseOrder) TableName() string {
	return "Stg_Purchase_Order"
}
