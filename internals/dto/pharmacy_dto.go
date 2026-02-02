package dto

import "time"

// PharmacyCreateDTO for creating a new pharmacy
type PharmacyCreateDTO struct {
	FacilityID   uint64  `json:"facility_id" binding:"required"`
	PharmacyCode string  `json:"pharmacy_code" binding:"required"`
	PharmacyName string  `json:"pharmacy_name" binding:"required"`
	PharmacyType *string `json:"pharmacy_type,omitempty"`
	IsActive     *bool   `json:"is_active,omitempty"`
}

// PharmacyUpdateDTO for updating a pharmacy
type PharmacyUpdateDTO struct {
	PharmacyName *string `json:"pharmacy_name,omitempty"`
	PharmacyType *string `json:"pharmacy_type,omitempty"`
	IsActive     *bool   `json:"is_active,omitempty"`
}

// PharmacyResponseDTO for API responses
type PharmacyResponseDTO struct {
	ID           uint64               `json:"id"`
	FacilityID   uint64               `json:"facility_id"`
	PharmacyCode string               `json:"pharmacy_code"`
	PharmacyName string               `json:"pharmacy_name"`
	PharmacyType *string              `json:"pharmacy_type,omitempty"`
	IsActive     bool                 `json:"is_active"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    *time.Time           `json:"updated_at,omitempty"`
	Facility     *FacilityResponseDTO `json:"facility,omitempty"`
}
