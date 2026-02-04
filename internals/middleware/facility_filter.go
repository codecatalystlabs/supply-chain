package middleware

import (
	"strings"
	"supply-chain/internals/config"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FacilityFilter middleware filters data based on user's facility
func FacilityFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only apply to GET requests
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		user, exists := c.Get("user")
		if !exists {
			c.Next()
			return
		}

		u := user.(*models.User)
		
		// Super admins and users without facility can see all data
		if u.FacilityID == nil || u.HasRole("super_admin") {
			c.Next()
			return
		}

		// Set facility filter in context for handlers to use
		c.Set("facility_id", *u.FacilityID)
		c.Set("facility_code", getFacilityCode(*u.FacilityID))
		
		c.Next()
	}
}

// ApplyFacilityFilter applies facility filter to a GORM query
func ApplyFacilityFilter(c *gin.Context, query *gorm.DB, facilityField string) *gorm.DB {
	if facilityID, exists := c.Get("facility_id"); exists {
		return query.Where(facilityField+" = ?", facilityID)
	}
	return query
}

// ApplyFacilityFilterByCode applies facility filter using facility code
func ApplyFacilityFilterByCode(c *gin.Context, query *gorm.DB, facilityCodeField string) *gorm.DB {
	if facilityCode, exists := c.Get("facility_code"); exists {
		return query.Where(facilityCodeField+" = ?", facilityCode)
	}
	return query
}

func getFacilityCode(facilityID uint64) string {
	var facility models.Facility
	if err := config.DB.First(&facility, facilityID).Error; err == nil {
		return facility.FacilityCode
	}
	return ""
}

// FilterFacilityData filters response data based on facility
func FilterFacilityData(c *gin.Context, data interface{}) interface{} {
	user, exists := c.Get("user")
	if !exists {
		return data
	}

	u := user.(*models.User)
	
	// Super admins can see all data
	if u.FacilityID == nil || u.HasRole("super_admin") {
		return data
	}

	// Filter data based on facility
	// This is a generic filter - specific handlers should implement their own filtering
	return data
}

// IsFacilityScoped checks if the current request should be facility-scoped
func IsFacilityScoped(c *gin.Context) bool {
	// Don't apply to API docs or health checks
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/swagger") || path == "/health" {
		return false
	}
	
	user, exists := c.Get("user")
	if !exists {
		return false
	}

	u := user.(*models.User)
	return u.FacilityID != nil && !u.HasRole("super_admin")
}

