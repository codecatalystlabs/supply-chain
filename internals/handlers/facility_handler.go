package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create facility
// @Tags Facility
// @Accept json
// @Produce json
// @Param payload body dto.FacilityCreateDTO true "Facility payload"
// @Success 201 {object} dto.FacilityResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facilities [post]
func CreateFacility(c *gin.Context) {
	var payload dto.FacilityCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if facility code already exists
	var existing models.Facility
	if err := config.DB.Where("facility_code = ?", payload.FacilityCode).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Facility code already exists"})
		return
	}

	isActive := true
	if payload.IsActive != nil {
		isActive = *payload.IsActive
	}

	facility := models.Facility{
		FacilityCode:  payload.FacilityCode,
		FacilityName:  payload.FacilityName,
		DHIS2Code:     payload.DHIS2Code,
		LevelOfCare:   payload.LevelOfCare,
		District:      payload.District,
		Region:        payload.Region,
		Zone:          payload.Zone,
		Address:       payload.Address,
		ContactPerson: payload.ContactPerson,
		ContactPhone:  payload.ContactPhone,
		ContactEmail:  payload.ContactEmail,
		EMRSystemCode: payload.EMRSystemCode,
		EMRSystemName: payload.EMRSystemName,
		IsActive:      isActive,
		CreatedAt:     time.Now(),
	}

	if err := config.DB.Create(&facility).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapToFacilityResponse(facility))
}

// @Summary List distinct regions from facilities
// @Tags Facility
// @Produce json
// @Success 200 {array} string
// @Router /facilities/regions [get]
func ListRegions(c *gin.Context) {
	var regions []string
	config.DB.Model(&models.Facility{}).Where("region IS NOT NULL AND region != ''").Distinct("region").Pluck("region", &regions)
	c.JSON(http.StatusOK, regions)
}

// @Summary List distinct districts from facilities
// @Tags Facility
// @Produce json
// @Param region query string false "Filter by region"
// @Success 200 {array} string
// @Router /facilities/districts [get]
func ListDistricts(c *gin.Context) {
	query := config.DB.Model(&models.Facility{}).Where("district IS NOT NULL AND district != ''").Distinct("district")
	if region := c.Query("region"); region != "" {
		query = query.Where("region = ?", region)
	}
	var districts []string
	query.Pluck("district", &districts)
	c.JSON(http.StatusOK, districts)
}

// @Summary List facilities
// @Tags Facility
// @Produce json
// @Param active query bool false "Filter by active status"
// @Param region query string false "Filter by region"
// @Param district query string false "Filter by district"
// @Param search query string false "Search by facility code or name"
// @Success 200 {array} dto.FacilityResponseDTO
// @Failure 500 {object} map[string]string
// @Router /facilities [get]
func ListFacilities(c *gin.Context) {
	var facilities []models.Facility
	query := config.DB

	// Apply facility filter if user is facility-scoped
	if user, exists := c.Get("user"); exists {
		u := user.(*models.User)
		if u.FacilityID != nil && !u.HasRole("super_admin") {
			query = query.Where("id = ?", *u.FacilityID)
		}
	}

	// Apply filters
	if active := c.Query("active"); active != "" {
		if active == "true" {
			query = query.Where("is_active = ?", true)
		} else if active == "false" {
			query = query.Where("is_active = ?", false)
		}
	}
	if region := c.Query("region"); region != "" {
		query = query.Where("region IS NOT NULL AND LOWER(TRIM(region)) = LOWER(TRIM(?))", region)
	}
	if district := c.Query("district"); district != "" {
		query = query.Where("district IS NOT NULL AND LOWER(TRIM(district)) = LOWER(TRIM(?))", district)
	}
	if search := c.Query("search"); search != "" {
		searchLike := "%" + search + "%"
		query = query.Where("facility_code ILIKE ? OR facility_name ILIKE ?", searchLike, searchLike)
	}

	if err := query.Preload("Pharmacies").Find(&facilities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.FacilityResponseDTO
	for _, f := range facilities {
		resp = append(resp, mapToFacilityResponse(f))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get facility by ID
// @Tags Facility
// @Produce json
// @Param id path int true "Facility ID"
// @Success 200 {object} dto.FacilityResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facilities/{id} [get]
func GetFacility(c *gin.Context) {
	id := c.Param("id")
	var facility models.Facility
	if err := config.DB.Preload("Pharmacies").First(&facility, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Facility not found"})
		return
	}
	c.JSON(http.StatusOK, mapToFacilityResponse(facility))
}

// @Summary Update facility
// @Tags Facility
// @Accept json
// @Produce json
// @Param id path int true "Facility ID"
// @Param payload body dto.FacilityUpdateDTO true "Update payload"
// @Success 200 {object} dto.FacilityResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facilities/{id} [put]
func UpdateFacility(c *gin.Context) {
	id := c.Param("id")
	var payload dto.FacilityUpdateDTO
	var facility models.Facility

	if err := config.DB.First(&facility, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Facility not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.FacilityName != nil {
		facility.FacilityName = *payload.FacilityName
	}
	if payload.DHIS2Code != nil {
		facility.DHIS2Code = payload.DHIS2Code
	}
	if payload.LevelOfCare != nil {
		facility.LevelOfCare = payload.LevelOfCare
	}
	if payload.District != nil {
		facility.District = payload.District
	}
	if payload.Region != nil {
		facility.Region = payload.Region
	}
	if payload.Zone != nil {
		facility.Zone = payload.Zone
	}
	if payload.Address != nil {
		facility.Address = payload.Address
	}
	if payload.ContactPerson != nil {
		facility.ContactPerson = payload.ContactPerson
	}
	if payload.ContactPhone != nil {
		facility.ContactPhone = payload.ContactPhone
	}
	if payload.ContactEmail != nil {
		facility.ContactEmail = payload.ContactEmail
	}
	if payload.EMRSystemCode != nil {
		facility.EMRSystemCode = payload.EMRSystemCode
	}
	if payload.EMRSystemName != nil {
		facility.EMRSystemName = payload.EMRSystemName
	}
	if payload.IsActive != nil {
		facility.IsActive = *payload.IsActive
	}
	now := time.Now()
	facility.UpdatedAt = &now

	if err := config.DB.Save(&facility).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapToFacilityResponse(facility))
}

// @Summary Delete facility
// @Tags Facility
// @Param id path int true "Facility ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facilities/{id} [delete]
func DeleteFacility(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Facility{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Facility deleted successfully"})
}

func mapToFacilityResponse(f models.Facility) dto.FacilityResponseDTO {
	var pharmacies []dto.PharmacyResponseDTO
	for _, p := range f.Pharmacies {
		// Map pharmacy without facility to avoid circular dependency
		pharmacies = append(pharmacies, dto.PharmacyResponseDTO{
			ID:           p.ID,
			FacilityID:   p.FacilityID,
			PharmacyCode: p.PharmacyCode,
			PharmacyName: p.PharmacyName,
			PharmacyType: p.PharmacyType,
			IsActive:     p.IsActive,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
			Facility:     nil, // Don't include facility to avoid circular reference
		})
	}

	return dto.FacilityResponseDTO{
		ID:            f.ID,
		FacilityCode:  f.FacilityCode,
		FacilityName:  f.FacilityName,
		DHIS2Code:     f.DHIS2Code,
		LevelOfCare:   f.LevelOfCare,
		District:      f.District,
		Region:        f.Region,
		Zone:          f.Zone,
		Address:       f.Address,
		ContactPerson: f.ContactPerson,
		ContactPhone:  f.ContactPhone,
		ContactEmail:  f.ContactEmail,
		IsActive:      f.IsActive,
		EMRSystemCode: f.EMRSystemCode,
		EMRSystemName: f.EMRSystemName,
		CreatedAt:     f.CreatedAt,
		UpdatedAt:     f.UpdatedAt,
		Pharmacies:    pharmacies,
	}
}
