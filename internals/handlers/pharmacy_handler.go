package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// NOTE: Pharmacy creation via API has been disabled as per requirements.

// @Summary List pharmacies
// @Tags Pharmacy
// @Produce json
// @Param facility_id query int false "Filter by facility ID"
// @Param active query bool false "Filter by active status"
// @Success 200 {array} dto.PharmacyResponseDTO
// @Failure 500 {object} map[string]string
// @Router /pharmacies [get]
func ListPharmacies(c *gin.Context) {
	var pharmacies []models.Pharmacy
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

	if err := query.Find(&pharmacies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.PharmacyResponseDTO
	for _, p := range pharmacies {
		resp = append(resp, mapToPharmacyResponse(p))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get pharmacy by ID
// @Tags Pharmacy
// @Produce json
// @Param id path int true "Pharmacy ID"
// @Success 200 {object} dto.PharmacyResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /pharmacies/{id} [get]
func GetPharmacy(c *gin.Context) {
	id := c.Param("id")
	var pharmacy models.Pharmacy
	if err := config.DB.Preload("Facility").First(&pharmacy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pharmacy not found"})
		return
	}
	c.JSON(http.StatusOK, mapToPharmacyResponse(pharmacy))
}

// @Summary Update pharmacy
// @Tags Pharmacy
// @Accept json
// @Produce json
// @Param id path int true "Pharmacy ID"
// @Param payload body dto.PharmacyUpdateDTO true "Update payload"
// @Success 200 {object} dto.PharmacyResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /pharmacies/{id} [put]
func UpdatePharmacy(c *gin.Context) {
	id := c.Param("id")
	var payload dto.PharmacyUpdateDTO
	var pharmacy models.Pharmacy

	if err := config.DB.First(&pharmacy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pharmacy not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.PharmacyName != nil {
		pharmacy.PharmacyName = *payload.PharmacyName
	}
	if payload.PharmacyType != nil {
		pharmacy.PharmacyType = payload.PharmacyType
	}
	if payload.IsActive != nil {
		pharmacy.IsActive = *payload.IsActive
	}
	now := time.Now()
	pharmacy.UpdatedAt = &now

	if err := config.DB.Save(&pharmacy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("Facility").First(&pharmacy, pharmacy.ID)
	c.JSON(http.StatusOK, mapToPharmacyResponse(pharmacy))
}

// @Summary Delete pharmacy
// @Tags Pharmacy
// @Param id path int true "Pharmacy ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /pharmacies/{id} [delete]
func DeletePharmacy(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Pharmacy{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pharmacy deleted successfully"})
}

func mapToPharmacyResponse(p models.Pharmacy) dto.PharmacyResponseDTO {
	var facility *dto.FacilityResponseDTO
	if p.Facility.ID != 0 {
		// Map facility without pharmacies to avoid circular dependency
		facility = &dto.FacilityResponseDTO{
			ID:            p.Facility.ID,
			FacilityCode:  p.Facility.FacilityCode,
			FacilityName:  p.Facility.FacilityName,
			DHIS2Code:     p.Facility.DHIS2Code,
			LevelOfCare:   p.Facility.LevelOfCare,
			District:      p.Facility.District,
			Region:        p.Facility.Region,
			Zone:          p.Facility.Zone,
			Address:       p.Facility.Address,
			ContactPerson: p.Facility.ContactPerson,
			ContactPhone:  p.Facility.ContactPhone,
			ContactEmail:  p.Facility.ContactEmail,
			IsActive:      p.Facility.IsActive,
			EMRSystemCode: p.Facility.EMRSystemCode,
			EMRSystemName: p.Facility.EMRSystemName,
			CreatedAt:     p.Facility.CreatedAt,
			UpdatedAt:     p.Facility.UpdatedAt,
			Pharmacies:    nil, // Don't include pharmacies to avoid circular reference
		}
	}

	return dto.PharmacyResponseDTO{
		ID:           p.ID,
		FacilityID:   p.FacilityID,
		PharmacyCode: p.PharmacyCode,
		PharmacyName: p.PharmacyName,
		PharmacyType: p.PharmacyType,
		IsActive:     p.IsActive,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		Facility:     facility,
	}
}
