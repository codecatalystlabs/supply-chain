package models

type PurchaseOrder struct {
	ID                   uint64    `gorm:"primaryKey"`
	OrdSystemCode        string    `gorm:"size:100;not null"`
	OrdFacilityCode      string    `gorm:"size:100;not null"`
	OrdTimestamp         time.Time `gorm:"not null"`
	OrdOrderDate         time.Time `gorm:"not null"`
	OrdOrderRefNumber    string    `gorm:"size:100;not null"`
	OrdOrderNumber       string    `gorm:"size:100;not null"`
	OrdProductCode       string    `gorm:"size:100;not null"`
	OrdOrderedQuantity   int       `gorm:"not null"`
	ValidationStatus     int16     `gorm:"default:0"`
	ValidationMessage    *string   `gorm:"size:100"`
	SyncStatus           int16     `gorm:"default:0"`
	AddDate              *time.Time
}

func (PurchaseOrder) TableName() string {
	return "Stg_Purchase_Order"
}
s