package models

import "time"

// PatientMetric represents aggregated patient data metrics
type PatientMetric struct {
	ID                 uint64     `gorm:"primaryKey" json:"id"`
	Section            string     `gorm:"size:100;not null" json:"section"`
	MetricCode         string     `gorm:"size:100;not null" json:"metric_code"`
	MetricName         *string    `gorm:"size:255" json:"metric_name,omitempty"`
	MetricValue        int        `json:"metric_value"`
	PeriodStart        time.Time  `json:"period_start"`
	PeriodEnd          time.Time  `json:"period_end"`
	AgeGroup           *string    `gorm:"size:50" json:"age_group,omitempty"`
	Sex                *string    `gorm:"size:10" json:"sex,omitempty"`
	RegimenLine        *string    `gorm:"size:100" json:"regimen_line,omitempty"`
	ConditionOrDrug    *string    `gorm:"size:255" json:"condition_or_drug,omitempty"`
	DisaggregationJSON *string    `gorm:"type:json" json:"disaggregation_json,omitempty"`
	AnonymizationFlag  *bool      `json:"anonymization_flag,omitempty"`
	ReportingLevel     *string    `gorm:"size:100" json:"reporting_level,omitempty"`
	SourceSystem       *string    `gorm:"size:100" json:"source_system,omitempty"`
	SourceRecordID     *string    `gorm:"size:100" json:"source_record_id,omitempty"`
	SubmittedBy        *string    `gorm:"size:100" json:"submitted_by,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

func (PatientMetric) TableName() string {
	return "Patient_Metric"
}
