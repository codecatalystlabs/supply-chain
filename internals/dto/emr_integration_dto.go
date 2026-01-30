package dto

import "time"

// EMRIntegrationCreateDTO for creating a new EMR integration
type EMRIntegrationCreateDTO struct {
	FacilityID      uint64  `json:"facility_id" binding:"required"`
	EMRSystemCode   string  `json:"emr_system_code" binding:"required"`
	EMRSystemName   string  `json:"emr_system_name" binding:"required"`
	EMRSystemVersion *string `json:"emr_system_version,omitempty"`
	APIEndpoint     *string  `json:"api_endpoint,omitempty"`
	APIKey          *string  `json:"api_key,omitempty"`
	APISecret       *string  `json:"api_secret,omitempty"`
	WebhookURL      *string  `json:"webhook_url,omitempty"`
	SyncEnabled     *bool    `json:"sync_enabled,omitempty"`
	SyncFrequency   *string  `json:"sync_frequency,omitempty"`
	AuthType        *string  `json:"auth_type,omitempty"`
	AuthConfig      *string  `json:"auth_config,omitempty"`
	Notes           *string  `json:"notes,omitempty"`
}

// EMRIntegrationUpdateDTO for updating an integration
type EMRIntegrationUpdateDTO struct {
	EMRSystemVersion *string `json:"emr_system_version,omitempty"`
	APIEndpoint      *string `json:"api_endpoint,omitempty"`
	APIKey           *string `json:"api_key,omitempty"`
	APISecret        *string `json:"api_secret,omitempty"`
	WebhookURL       *string `json:"webhook_url,omitempty"`
	SyncEnabled      *bool   `json:"sync_enabled,omitempty"`
	SyncFrequency    *string `json:"sync_frequency,omitempty"`
	AuthType         *string `json:"auth_type,omitempty"`
	AuthConfig       *string `json:"auth_config,omitempty"`
	IsActive         *bool   `json:"is_active,omitempty"`
	Notes            *string `json:"notes,omitempty"`
}

// EMRIntegrationResponseDTO for API responses
type EMRIntegrationResponseDTO struct {
	ID                uint64     `json:"id"`
	FacilityID        uint64     `json:"facility_id"`
	FacilityCode      string     `json:"facility_code,omitempty"`
	EMRSystemCode     string     `json:"emr_system_code"`
	EMRSystemName     string     `json:"emr_system_name"`
	EMRSystemVersion  *string    `json:"emr_system_version,omitempty"`
	APIEndpoint       *string    `json:"api_endpoint,omitempty"`
	SyncEnabled       bool       `json:"sync_enabled"`
	SyncFrequency     *string    `json:"sync_frequency,omitempty"`
	LastSyncAt        *time.Time `json:"last_sync_at,omitempty"`
	LastSyncStatus    *string    `json:"last_sync_status,omitempty"`
	LastSyncMessage   *string    `json:"last_sync_message,omitempty"`
	AuthType          *string    `json:"auth_type,omitempty"`
	IsActive          bool       `json:"is_active"`
	IsVerified        bool       `json:"is_verified"`
	VerifiedAt        *time.Time `json:"verified_at,omitempty"`
	VerifiedBy        *string    `json:"verified_by,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

// EMRSyncLogResponseDTO for sync log responses
type EMRSyncLogResponseDTO struct {
	ID                uint64     `json:"id"`
	IntegrationID     uint64     `json:"integration_id"`
	SyncType          string     `json:"sync_type"`
	SyncDirection     string     `json:"sync_direction"`
	Status            string     `json:"status"`
	RecordsProcessed  int        `json:"records_processed"`
	RecordsSuccessful int        `json:"records_successful"`
	RecordsFailed     int        `json:"records_failed"`
	ErrorMessage      *string    `json:"error_message,omitempty"`
	StartedAt         time.Time  `json:"started_at"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
	DurationSeconds   *int       `json:"duration_seconds,omitempty"`
}
