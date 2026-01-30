package models

import "time"

// Facility represents a health facility in the system
type Facility struct {
	ID            uint64     `gorm:"primaryKey" json:"id"`
	FacilityCode  string     `gorm:"size:100;uniqueIndex;not null" json:"facility_code"`
	FacilityName  string     `gorm:"size:255;not null" json:"facility_name"`
	DHIS2Code     *string    `gorm:"size:100;uniqueIndex" json:"dhis2_code,omitempty"`
	LevelOfCare   *string    `gorm:"size:100" json:"level_of_care,omitempty"` // e.g., HCII, HCIII, HCIV, Hospital
	District      *string    `gorm:"size:100" json:"district,omitempty"`
	Region        *string    `gorm:"size:100" json:"region,omitempty"`
	Zone          *string    `gorm:"size:100" json:"zone,omitempty"`
	Address       *string    `gorm:"size:500" json:"address,omitempty"`
	ContactPerson *string    `gorm:"size:255" json:"contact_person,omitempty"`
	ContactPhone  *string    `gorm:"size:50" json:"contact_phone,omitempty"`
	ContactEmail  *string    `gorm:"size:255" json:"contact_email,omitempty"`
	IsActive      bool       `gorm:"default:true" json:"is_active"`
	EMRSystemCode *string    `gorm:"size:100" json:"emr_system_code,omitempty"` // EMR system identifier
	EMRSystemName *string    `gorm:"size:255" json:"emr_system_name,omitempty"` // e.g., OpenMRS, DHIS2, etc.
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Pharmacies []Pharmacy      `gorm:"foreignKey:FacilityID;constraint:OnDelete:CASCADE" json:"pharmacies,omitempty"`
	Orders     []FacilityOrder `gorm:"foreignKey:FacilityID" json:"orders,omitempty"`
}

func (Facility) TableName() string {
	return "facilities"
}

// Pharmacy represents a pharmacy within a facility
// Facilities can have multiple pharmacies (e.g., Main Pharmacy, Pediatric Pharmacy, etc.)
type Pharmacy struct {
	ID           uint64     `gorm:"primaryKey" json:"id"`
	FacilityID   uint64     `gorm:"not null;index" json:"facility_id"`
	PharmacyCode string     `gorm:"size:100;not null" json:"pharmacy_code"` // Unique within facility
	PharmacyName string     `gorm:"size:255;not null" json:"pharmacy_name"`
	PharmacyType *string    `gorm:"size:100" json:"pharmacy_type,omitempty"` // e.g., Main, Pediatric, ART, etc.
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Facility Facility `gorm:"foreignKey:FacilityID" json:"facility,omitempty"`
}

func (Pharmacy) TableName() string {
	return "pharmacies"
}
