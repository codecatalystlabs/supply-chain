package dto

import "time"

type ProcurementPlanItemCreateDTO struct {
	ProductCode string    `json:"product_code" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required"`
	NeededBy    time.Time `json:"needed_by" binding:"required"`
}

type ProcurementPlanCreateDTO struct {
	PlanSystemCode string                         `json:"plan_system_code" binding:"required"`
	StoreCode      string                         `json:"store_code" binding:"required"`
	Notes          *string                        `json:"notes,omitempty"`
	Items          []ProcurementPlanItemCreateDTO `json:"items" binding:"required,dive"`
}

type ProcurementPlanItemResponseDTO struct {
	ID            uint64    `json:"id"`
	ProcurementID uint64    `json:"procurement_id"`
	ProductCode   string    `json:"product_code"`
	Quantity      int       `json:"quantity"`
	NeededBy      time.Time `json:"needed_by"`
	Status        string    `json:"status"`
}

type ProcurementPlanResponseDTO struct {
	ID             uint64                           `json:"id"`
	PlanSystemCode string                           `json:"plan_system_code"`
	StoreCode      string                           `json:"store_code"`
	CreatedAt      time.Time                        `json:"created_at"`
	Notes          *string                          `json:"notes,omitempty"`
	Items          []ProcurementPlanItemResponseDTO `json:"items,omitempty"`
}
