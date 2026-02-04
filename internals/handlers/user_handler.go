package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary List users
// @Tags User
// @Produce json
// @Success 200 {array} dto.UserResponseDTO
// @Failure 500 {object} map[string]string
// @Router /users [get]
func ListUsers(c *gin.Context) {
	var users []models.User
	query := config.DB.Preload("Roles").Preload("Facility")

	// Apply facility filter if user is facility-scoped
	if user, exists := c.Get("user"); exists {
		u := user.(*models.User)
		if u.FacilityID != nil {
			query = query.Where("facility_id = ?", *u.FacilityID)
		}
	}

	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.UserResponseDTO
	for _, u := range users {
		resp = append(resp, mapToUserResponse(u))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get user by ID
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Preload("Roles").Preload("Facility").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check facility access
	if user, exists := c.Get("user"); exists {
		u := user.(*models.User)
		if u.FacilityID != nil && (user.(*models.User).FacilityID == nil || *u.FacilityID != *user.(*models.User).FacilityID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}
	}

	c.JSON(http.StatusOK, mapToUserResponse(user))
}

// @Summary Create user
// @Tags User
// @Accept json
// @Produce json
// @Param payload body dto.UserCreateDTO true "User payload"
// @Success 201 {object} dto.UserResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var payload dto.UserCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username or email already exists
	var existing models.User
	if err := config.DB.Where("username = ? OR email = ?", payload.Username, payload.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exists"})
		return
	}

	user := models.User{
		Username:  payload.Username,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		FacilityID: payload.FacilityID,
		IsActive:  true,
	}

	if payload.IsActive != nil {
		user.IsActive = *payload.IsActive
	}

	if err := user.SetPassword(payload.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set password"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assign roles
	if len(payload.RoleIDs) > 0 {
		var roles []models.Role
		config.DB.Where("id IN ?", payload.RoleIDs).Find(&roles)
		config.DB.Model(&user).Association("Roles").Replace(roles)
	}

	// Reload with relationships
	config.DB.Preload("Roles").Preload("Facility").First(&user, user.ID)

	c.JSON(http.StatusCreated, mapToUserResponse(user))
}

// @Summary Update user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param payload body dto.UserUpdateDTO true "Update payload"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var payload dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.Email != nil {
		user.Email = *payload.Email
	}
	if payload.FirstName != nil {
		user.FirstName = *payload.FirstName
	}
	if payload.LastName != nil {
		user.LastName = *payload.LastName
	}
	if payload.FacilityID != nil {
		user.FacilityID = payload.FacilityID
	}
	if payload.IsActive != nil {
		user.IsActive = *payload.IsActive
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update roles
	if payload.RoleIDs != nil {
		var roles []models.Role
		config.DB.Where("id IN ?", payload.RoleIDs).Find(&roles)
		config.DB.Model(&user).Association("Roles").Replace(roles)
	}

	// Reload with relationships
	config.DB.Preload("Roles").Preload("Facility").First(&user, user.ID)

	c.JSON(http.StatusOK, mapToUserResponse(user))
}

// @Summary Delete user
// @Tags User
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func mapToUserResponse(u models.User) dto.UserResponseDTO {
	var roles []dto.RoleResponseDTO
	for _, r := range u.Roles {
		roles = append(roles, dto.RoleResponseDTO{
			ID:          r.ID,
			Name:        r.Name,
			DisplayName: r.DisplayName,
		})
	}

	resp := dto.UserResponseDTO{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FacilityID: u.FacilityID,
		IsActive:  u.IsActive,
		LastLoginAt: u.LastLoginAt,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Roles:     roles,
	}

	if u.Facility != nil {
		facilityCode := u.Facility.FacilityCode
		facilityName := u.Facility.FacilityName
		resp.FacilityCode = &facilityCode
		resp.FacilityName = &facilityName
	}

	return resp
}

