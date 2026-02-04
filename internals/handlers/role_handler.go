package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary List roles
// @Tags Role
// @Produce json
// @Success 200 {array} models.Role
// @Failure 500 {object} map[string]string
// @Router /roles [get]
func ListRoles(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// @Summary Get role by ID
// @Tags Role
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.Role
// @Failure 404 {object} map[string]string
// @Router /roles/{id} [get]
func GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.Preload("Permissions").First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

