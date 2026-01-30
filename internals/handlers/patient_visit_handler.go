package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create patient visit
// @Tags PatientVisit
// @Accept json
// @Produce json
// @Param payload body dto.PatientVisitCreateDTO true "Patient visit payload"
// @Success 201 {object} dto.PatientVisitResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patient-visit [post]
func CreatePatientVisit(c *gin.Context) {
	var payload dto.PatientVisitCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.PatientVisit{
		VstSystemCode:      payload.VstSystemCode,
		VstFacilityCode:    payload.VstFacilityCode,
		VstTimestamp:       time.Now(),
		VstVisitDate:       payload.VstVisitDate,
		VstPatientCode:     payload.VstPatientCode,
		VstSex:             payload.VstSex,
		VstAge:             payload.VstAge,
		VstProductCode:     payload.VstProductCode,
		VstBatchNumber:     payload.VstBatchNumber,
		VstQuantity:        payload.VstQuantity,
		VstRegimenCode:     payload.VstRegimenCode,
		VstPatientCategory: payload.VstPatientCategory,
	}

	if err := config.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapToPatientVisitResponse(record))
}

func ListPatientVisits(c *gin.Context) {
	var records []models.PatientVisit
	if err := config.DB.Find(&records).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.PatientVisitResponseDTO
	for _, r := range records {
		resp = append(resp, mapToPatientVisitResponse(r))
	}
	c.JSON(200, resp)
}

func GetPatientVisit(c *gin.Context) {
	id := c.Param("id")
	var record models.PatientVisit
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(200, mapToPatientVisitResponse(record))
}

func UpdatePatientVisit(c *gin.Context) {
	id := c.Param("id")
	var payload dto.PatientVisitUpdateDTO
	var record models.PatientVisit
	if err := config.DB.First(&record, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if payload.VstQuantity != nil {
		record.VstQuantity = *payload.VstQuantity
	}
	if err := config.DB.Save(&record).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mapToPatientVisitResponse(record))
}

func DeletePatientVisit(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.PatientVisit{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToPatientVisitResponse(m models.PatientVisit) dto.PatientVisitResponseDTO {
	return dto.PatientVisitResponseDTO{
		ID:                 m.ID,
		VstSystemCode:      m.VstSystemCode,
		VstFacilityCode:    m.VstFacilityCode,
		VstTimestamp:       m.VstTimestamp,
		VstPatientCode:     m.VstPatientCode,
		VstSex:             m.VstSex,
		VstAge:             m.VstAge,
		VstVisitDate:       m.VstVisitDate,
		VstProductCode:     m.VstProductCode,
		VstBatchNumber:     m.VstBatchNumber,
		VstQuantity:        m.VstQuantity,
		VstRegimenCode:     m.VstRegimenCode,
		VstPatientCategory: m.VstPatientCategory,
		ValidationStatus:   m.ValidationStatus,
		SyncStatus:         m.SyncStatus,
	}
}
