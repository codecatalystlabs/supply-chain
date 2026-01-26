package models

import "time"

type ProcurementPlan struct {
	ID             uint64    `gorm:"primaryKey" json:"id"`
	PlanSystemCode string    `gorm:"size:100;not null" json:"plan_system_code"`
	StoreCode      string    `gorm:"size:100;not null" json:"store_code"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	Notes          *string   `gorm:"size:255" json:"notes,omitempty"`
}

func (ProcurementPlan) TableName() string {
	return "Procurement_Plan"
}

type ProcurementPlanItem struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	ProcurementID uint64    `gorm:"not null" json:"procurement_id"`
	ProductCode   string    `gorm:"size:100;not null" json:"product_code"`
	Quantity      int       `gorm:"not null" json:"quantity"`
	NeededBy      time.Time `json:"needed_by"`
	Status        string    `gorm:"size:50" json:"status"` // e.g., planned/ordered/received
}

func (ProcurementPlanItem) TableName() string {
	return "Procurement_Plan_Item"
}
