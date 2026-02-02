package dto

import "time"

// WarehouseCreateDTO for creating a new warehouse
type WarehouseCreateDTO struct {
	WarehouseCode string  `json:"warehouse_code" binding:"required"`
	WarehouseName string  `json:"warehouse_name" binding:"required"`
	WarehouseType *string `json:"warehouse_type,omitempty"`
	Address       *string `json:"address,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	ContactPhone  *string `json:"contact_phone,omitempty"`
	ContactEmail  *string `json:"contact_email,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
}

// WarehouseUpdateDTO for updating a warehouse
type WarehouseUpdateDTO struct {
	WarehouseName string  `json:"warehouse_name,omitempty"`
	WarehouseType *string `json:"warehouse_type,omitempty"`
	Address       *string `json:"address,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	ContactPhone  *string `json:"contact_phone,omitempty"`
	ContactEmail  *string `json:"contact_email,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
}

// WarehouseResponseDTO for API responses
type WarehouseResponseDTO struct {
	ID            uint64     `json:"id"`
	WarehouseCode string     `json:"warehouse_code"`
	WarehouseName string     `json:"warehouse_name"`
	WarehouseType *string    `json:"warehouse_type,omitempty"`
	Address       *string    `json:"address,omitempty"`
	ContactPerson *string    `json:"contact_person,omitempty"`
	ContactPhone  *string    `json:"contact_phone,omitempty"`
	ContactEmail  *string    `json:"contact_email,omitempty"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
