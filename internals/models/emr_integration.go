package models

import "time"

// EMRIntegration represents an EMR system integration at a facility
type EMRIntegration struct {
	ID               uint64  `gorm:"primaryKey" json:"id"`
	FacilityID       uint64  `gorm:"not null;uniqueIndex" json:"facility_id"`
	EMRSystemCode    string  `gorm:"size:100;not null;index" json:"emr_system_code"` // e.g., OPENMRS, DHIS2, etc.
	EMRSystemName    string  `gorm:"size:255;not null" json:"emr_system_name"`
	EMRSystemVersion *string `gorm:"size:50" json:"emr_system_version,omitempty"`

	// API Configuration
	APIEndpoint *string `gorm:"size:500" json:"api_endpoint,omitempty"`
	APIKey      *string `gorm:"size:255" json:"api_key,omitempty"`     // Encrypted in production
	APISecret   *string `gorm:"size:255" json:"api_secret,omitempty"`  // Encrypted in production
	WebhookURL  *string `gorm:"size:500" json:"webhook_url,omitempty"` // Our webhook to receive EMR events

	// Integration settings
	SyncEnabled     bool       `gorm:"default:true" json:"sync_enabled"`
	SyncFrequency   *string    `gorm:"size:50" json:"sync_frequency,omitempty"` // realtime, hourly, daily
	LastSyncAt      *time.Time `json:"last_sync_at,omitempty"`
	LastSyncStatus  *string    `gorm:"size:50" json:"last_sync_status,omitempty"` // success, failed, partial
	LastSyncMessage *string    `gorm:"type:text" json:"last_sync_message,omitempty"`

	// Authentication
	AuthType   *string `gorm:"size:50" json:"auth_type,omitempty"`     // api_key, oauth2, basic
	AuthConfig *string `gorm:"type:text" json:"auth_config,omitempty"` // JSON config

	// Status
	IsActive   bool       `gorm:"default:true" json:"is_active"`
	IsVerified bool       `gorm:"default:false" json:"is_verified"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	VerifiedBy *string    `gorm:"size:255" json:"verified_by,omitempty"`

	// Metadata
	Notes     *string    `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Facility Facility     `gorm:"foreignKey:FacilityID" json:"facility,omitempty"`
	SyncLogs []EMRSyncLog `gorm:"foreignKey:IntegrationID" json:"sync_logs,omitempty"`
}

func (EMRIntegration) TableName() string {
	return "emr_integrations"
}

// EMRSyncLog tracks synchronization events between EMR and central system
type EMRSyncLog struct {
	ID                uint64     `gorm:"primaryKey" json:"id"`
	IntegrationID     uint64     `gorm:"not null;index" json:"integration_id"`
	SyncType          string     `gorm:"size:50;not null;index" json:"sync_type"` // stock, order, patient, etc.
	SyncDirection     string     `gorm:"size:50;not null" json:"sync_direction"`  // emr_to_central, central_to_emr
	Status            string     `gorm:"size:50;not null;index" json:"status"`    // success, failed, partial
	RecordsProcessed  int        `gorm:"default:0" json:"records_processed"`
	RecordsSuccessful int        `gorm:"default:0" json:"records_successful"`
	RecordsFailed     int        `gorm:"default:0" json:"records_failed"`
	ErrorMessage      *string    `gorm:"type:text" json:"error_message,omitempty"`
	StartedAt         time.Time  `gorm:"not null" json:"started_at"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
	DurationSeconds   *int       `json:"duration_seconds,omitempty"`

	// Relationships
	Integration EMRIntegration `gorm:"foreignKey:IntegrationID" json:"integration,omitempty"`
}

func (EMRSyncLog) TableName() string {
	return "emr_sync_logs"
}
