package models

import "time"

type PatientVisit struct {
	ID                 uint64 `gorm:"primaryKey"`
	VstSystemCode      string `gorm:"size:100;not null"`
	VstFacilityCode    string `gorm:"size:100;not null"`
	VstTimestamp       time.Time
	VstPatientCode     string `gorm:"size:100;not null"`
	VstSex             string `gorm:"size:10;not null"`
	VstAge             int16  `gorm:"not null"`
	VstVisitDate       time.Time
	VstProductCode     string  `gorm:"size:100;not null"`
	VstBatchNumber     string  `gorm:"size:100;not null"`
	VstQuantity        float64 `gorm:"not null"`
	VstRegimenCode     string  `gorm:"size:100;not null"`
	VstPatientCategory string  `gorm:"size:100;not null"`
	ValidationStatus   int16   `gorm:"default:0"`
	ValidationMessage  *string `gorm:"size:100"`
	SyncStatus         int16   `gorm:"default:0"`
	AddDate            *time.Time
}

func (PatientVisit) TableName() string {
	return "Stg_Patient_Visit"
}
