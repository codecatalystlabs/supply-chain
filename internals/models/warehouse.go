package models

import "time"

// Warehouse represents a medical warehouse (JMS, NMS, etc.)
type Warehouse struct {
	ID            uint64     `gorm:"primaryKey" json:"id"`
	WarehouseCode string     `gorm:"size:100;uniqueIndex;not null" json:"warehouse_code"` // e.g., JMS, NMS
	WarehouseName string     `gorm:"size:255;not null" json:"warehouse_name"`
	WarehouseType *string    `gorm:"size:100" json:"warehouse_type,omitempty"` // e.g., National, Regional, Private
	Address       *string    `gorm:"size:500" json:"address,omitempty"`
	ContactPerson *string    `gorm:"size:255" json:"contact_person,omitempty"`
	ContactPhone  *string    `gorm:"size:50" json:"contact_phone,omitempty"`
	ContactEmail  *string    `gorm:"size:255" json:"contact_email,omitempty"`
	IsActive      bool       `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`

	// Relationships
	ProcurementPlans []ProcurementPlan `gorm:"foreignKey:WarehouseID" json:"procurement_plans,omitempty"`
	Orders           []FacilityOrder   `gorm:"foreignKey:WarehouseID" json:"orders,omitempty"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}
