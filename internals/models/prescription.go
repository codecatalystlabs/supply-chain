package models

import "time"

type Prescription struct {
	ID                    uint64    `gorm:"primaryKey" json:"id"`
	PrsSystemCode         string    `gorm:"size:100;not null" json:"prs_system_code"`
	PrsFacilityCode       string    `gorm:"size:100;not null" json:"prs_facility_code"`
	PrsTimestamp          time.Time `gorm:"not null" json:"prs_timestamp"`
	PrsPatientCode        string    `gorm:"size:100;not null" json:"prs_patient_code"`
	PrsProductCode        string    `gorm:"size:100;not null" json:"prs_product_code"`
	PrsPrescribedQuantity float64   `gorm:"not null" json:"prs_prescribed_quantity"`
	PrsPrescriber         *string   `gorm:"size:100" json:"prs_prescriber,omitempty"`
	PrsNotes              *string   `gorm:"size:255" json:"prs_notes,omitempty"`
	BaseModel
}

func (Prescription) TableName() string {
	return "Prescription"
}
