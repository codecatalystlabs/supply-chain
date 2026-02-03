package dto

import "time"

// FacilityCreateDTO for creating a new facility
type FacilityCreateDTO struct {
	FacilityCode  string  `json:"facility_code" binding:"required"`
	FacilityName  string  `json:"facility_name" binding:"required"`
	DHIS2Code     *string `json:"dhis2_code,omitempty"`
	LevelOfCare   *string `json:"level_of_care,omitempty"`
	District      *string `json:"district,omitempty"`
	Region        *string `json:"region,omitempty"`
	Zone          *string `json:"zone,omitempty"`
	Address       *string `json:"address,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	ContactPhone  *string `json:"contact_phone,omitempty"`
	ContactEmail  *string `json:"contact_email,omitempty"`
	EMRSystemCode *string `json:"emr_system_code,omitempty"`
	EMRSystemName *string `json:"emr_system_name,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
}

// FacilityUpdateDTO for updating a facility
type FacilityUpdateDTO struct {
	FacilityName  *string `json:"facility_name,omitempty"`
	DHIS2Code     *string `json:"dhis2_code,omitempty"`
	LevelOfCare   *string `json:"level_of_care,omitempty"`
	District      *string `json:"district,omitempty"`
	Region        *string `json:"region,omitempty"`
	Zone          *string `json:"zone,omitempty"`
	Address       *string `json:"address,omitempty"`
	ContactPerson *string `json:"contact_person,omitempty"`
	ContactPhone  *string `json:"contact_phone,omitempty"`
	ContactEmail  *string `json:"contact_email,omitempty"`
	EMRSystemCode *string `json:"emr_system_code,omitempty"`
	EMRSystemName *string `json:"emr_system_name,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
}

// FacilityResponseDTO for API responses
type FacilityResponseDTO struct {
	ID            uint64     `json:"id"`
	FacilityCode  string     `json:"facility_code"`
	FacilityName  string     `json:"facility_name"`
	DHIS2Code     *string    `json:"dhis2_code,omitempty"`
	LevelOfCare   *string    `json:"level_of_care,omitempty"`
	District      *string    `json:"district,omitempty"`
	Region        *string    `json:"region,omitempty"`
	Zone          *string    `json:"zone,omitempty"`
	Address       *string    `json:"address,omitempty"`
	ContactPerson *string    `json:"contact_person,omitempty"`
	ContactPhone  *string    `json:"contact_phone,omitempty"`
	ContactEmail   *string   `json:"contact_email,omitempty"`
	IsActive      bool       `json:"is_active"`
	EMRSystemCode *string    `json:"emr_system_code,omitempty"`
	EMRSystemName *string    `json:"emr_system_name,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Pharmacies    []PharmacyResponseDTO `json:"pharmacies,omitempty"`
}
