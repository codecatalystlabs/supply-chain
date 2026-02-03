package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create EMR integrations
// @Tags EMRIntegration
// @Accept json
// @Produce json
// @Param payload body []dto.EMRIntegrationCreateDTO true "Integrations payload"
// @Success 201 {array} dto.EMRIntegrationResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /emr-integrations [post]
func CreateEMRIntegration(c *gin.Context) {
	var payloads []dto.EMRIntegrationCreateDTO
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var responses []dto.EMRIntegrationResponseDTO
	for _, payload := range payloads {
		// Verify facility exists
		var facility models.Facility
		if err := config.DB.First(&facility, payload.FacilityID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Facility not found"})
			return
		}

		// Check if integration already exists for this facility
		var existing models.EMRIntegration
		if err := config.DB.Where("facility_id = ?", payload.FacilityID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "EMR integration already exists for this facility"})
			return
		}

		syncEnabled := true
		if payload.SyncEnabled != nil {
			syncEnabled = *payload.SyncEnabled
		}

		integration := models.EMRIntegration{
			FacilityID:       payload.FacilityID,
			EMRSystemCode:    payload.EMRSystemCode,
			EMRSystemName:    payload.EMRSystemName,
			EMRSystemVersion: payload.EMRSystemVersion,
			APIEndpoint:      payload.APIEndpoint,
			APIKey:           payload.APIKey,
			APISecret:        payload.APISecret,
			WebhookURL:       payload.WebhookURL,
			SyncEnabled:      syncEnabled,
			SyncFrequency:    payload.SyncFrequency,
			AuthType:         payload.AuthType,
			AuthConfig:       payload.AuthConfig,
			IsActive:         true,
			IsVerified:       false,
			Notes:            payload.Notes,
			CreatedAt:        time.Now(),
		}

		if err := config.DB.Create(&integration).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responses = append(responses, mapToEMRIntegrationResponse(integration))
	}

	c.JSON(http.StatusCreated, responses)
}

// @Summary List EMR integrations
// @Tags EMRIntegration
// @Produce json
// @Param facility_id query int false "Filter by facility ID"
// @Param active query bool false "Filter by active status"
// @Success 200 {array} dto.EMRIntegrationResponseDTO
// @Failure 500 {object} map[string]string
// @Router /emr-integrations [get]
func ListEMRIntegrations(c *gin.Context) {
	var integrations []models.EMRIntegration
	query := config.DB.Preload("Facility")

	// Apply filters
	if facilityID := c.Query("facility_id"); facilityID != "" {
		query = query.Where("facility_id = ?", facilityID)
	}
	if active := c.Query("active"); active != "" {
		if active == "true" {
			query = query.Where("is_active = ?", true)
		} else if active == "false" {
			query = query.Where("is_active = ?", false)
		}
	}

	if err := query.Find(&integrations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.EMRIntegrationResponseDTO
	for _, i := range integrations {
		resp = append(resp, mapToEMRIntegrationResponse(i))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get EMR integration by ID
// @Tags EMRIntegration
// @Produce json
// @Param id path int true "Integration ID"
// @Success 200 {object} dto.EMRIntegrationResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /emr-integrations/{id} [get]
func GetEMRIntegration(c *gin.Context) {
	id := c.Param("id")
	var integration models.EMRIntegration
	if err := config.DB.Preload("Facility").First(&integration, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Integration not found"})
		return
	}
	c.JSON(http.StatusOK, mapToEMRIntegrationResponse(integration))
}

// @Summary Update EMR integration
// @Tags EMRIntegration
// @Accept json
// @Produce json
// @Param id path int true "Integration ID"
// @Param payload body dto.EMRIntegrationUpdateDTO true "Update payload"
// @Success 200 {object} dto.EMRIntegrationResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /emr-integrations/{id} [put]
func UpdateEMRIntegration(c *gin.Context) {
	id := c.Param("id")
	var payload dto.EMRIntegrationUpdateDTO
	var integration models.EMRIntegration

	if err := config.DB.First(&integration, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Integration not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.EMRSystemVersion != nil {
		integration.EMRSystemVersion = payload.EMRSystemVersion
	}
	if payload.APIEndpoint != nil {
		integration.APIEndpoint = payload.APIEndpoint
	}
	if payload.APIKey != nil {
		integration.APIKey = payload.APIKey
	}
	if payload.APISecret != nil {
		integration.APISecret = payload.APISecret
	}
	if payload.WebhookURL != nil {
		integration.WebhookURL = payload.WebhookURL
	}
	if payload.SyncEnabled != nil {
		integration.SyncEnabled = *payload.SyncEnabled
	}
	if payload.SyncFrequency != nil {
		integration.SyncFrequency = payload.SyncFrequency
	}
	if payload.AuthType != nil {
		integration.AuthType = payload.AuthType
	}
	if payload.AuthConfig != nil {
		integration.AuthConfig = payload.AuthConfig
	}
	if payload.IsActive != nil {
		integration.IsActive = *payload.IsActive
	}
	if payload.Notes != nil {
		integration.Notes = payload.Notes
	}
	now := time.Now()
	integration.UpdatedAt = &now

	if err := config.DB.Save(&integration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("Facility").First(&integration, integration.ID)
	c.JSON(http.StatusOK, mapToEMRIntegrationResponse(integration))
}

// @Summary Verify EMR integration
// @Tags EMRIntegration
// @Param id path int true "Integration ID"
// @Param payload body map[string]string true "Verify payload" SchemaExample({"verified_by": "System Admin"})
// @Success 200 {object} dto.EMRIntegrationResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /emr-integrations/{id}/verify [post]
func VerifyEMRIntegration(c *gin.Context) {
	id := c.Param("id")
	var integration models.EMRIntegration
	var payload struct {
		VerifiedBy string `json:"verified_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&integration, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Integration not found"})
		return
	}

	now := time.Now()
	integration.IsVerified = true
	integration.VerifiedAt = &now
	integration.VerifiedBy = &payload.VerifiedBy
	integration.UpdatedAt = &now

	if err := config.DB.Save(&integration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("Facility").First(&integration, integration.ID)
	c.JSON(http.StatusOK, mapToEMRIntegrationResponse(integration))
}

// @Summary Get EMR sync logs
// @Tags EMRIntegration
// @Produce json
// @Param id path int true "Integration ID"
// @Param limit query int false "Limit results"
// @Success 200 {array} dto.EMRSyncLogResponseDTO
// @Failure 500 {object} map[string]string
// @Router /emr-integrations/{id}/sync-logs [get]
func GetEMRSyncLogs(c *gin.Context) {
	id := c.Param("id")
	var logs []models.EMRSyncLog
	query := config.DB.Where("integration_id = ?", id).Order("started_at DESC")

	if limit := c.Query("limit"); limit != "" {
		query = query.Limit(100) // Default limit
	}

	if err := query.Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.EMRSyncLogResponseDTO
	for _, log := range logs {
		resp = append(resp, mapToEMRSyncLogResponse(log))
	}
	c.JSON(http.StatusOK, resp)
}

func mapToEMRIntegrationResponse(i models.EMRIntegration) dto.EMRIntegrationResponseDTO {
	var facilityCode string
	if i.Facility.ID != 0 {
		facilityCode = i.Facility.FacilityCode
	}

	return dto.EMRIntegrationResponseDTO{
		ID:               i.ID,
		FacilityID:       i.FacilityID,
		FacilityCode:     facilityCode,
		EMRSystemCode:    i.EMRSystemCode,
		EMRSystemName:    i.EMRSystemName,
		EMRSystemVersion: i.EMRSystemVersion,
		APIEndpoint:      i.APIEndpoint,
		SyncEnabled:      i.SyncEnabled,
		SyncFrequency:    i.SyncFrequency,
		LastSyncAt:       i.LastSyncAt,
		LastSyncStatus:   i.LastSyncStatus,
		LastSyncMessage:  i.LastSyncMessage,
		AuthType:         i.AuthType,
		IsActive:         i.IsActive,
		IsVerified:       i.IsVerified,
		VerifiedAt:       i.VerifiedAt,
		VerifiedBy:       i.VerifiedBy,
		Notes:            i.Notes,
		CreatedAt:        i.CreatedAt,
		UpdatedAt:        i.UpdatedAt,
	}
}

func mapToEMRSyncLogResponse(l models.EMRSyncLog) dto.EMRSyncLogResponseDTO {
	return dto.EMRSyncLogResponseDTO{
		ID:                l.ID,
		IntegrationID:     l.IntegrationID,
		SyncType:          l.SyncType,
		SyncDirection:     l.SyncDirection,
		Status:            l.Status,
		RecordsProcessed:  l.RecordsProcessed,
		RecordsSuccessful: l.RecordsSuccessful,
		RecordsFailed:     l.RecordsFailed,
		ErrorMessage:      l.ErrorMessage,
		StartedAt:         l.StartedAt,
		CompletedAt:       l.CompletedAt,
		DurationSeconds:   l.DurationSeconds,
	}
}
