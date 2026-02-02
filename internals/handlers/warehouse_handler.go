package handlers

import (
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param payload body dto.WarehouseCreateDTO true "Warehouse payload"
// @Success 201 {object} dto.WarehouseResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /warehouses [post]
func CreateWarehouse(c *gin.Context) {
	var payload dto.WarehouseCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if warehouse code already exists
	var existing models.Warehouse
	if err := config.DB.Where("warehouse_code = ?", payload.WarehouseCode).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Warehouse code already exists"})
		return
	}

	isActive := true
	if payload.IsActive != nil {
		isActive = *payload.IsActive
	}

	warehouse := models.Warehouse{
		WarehouseCode: payload.WarehouseCode,
		WarehouseName: payload.WarehouseName,
		WarehouseType: payload.WarehouseType,
		Address:       payload.Address,
		ContactPerson: payload.ContactPerson,
		ContactPhone:  payload.ContactPhone,
		ContactEmail:  payload.ContactEmail,
		IsActive:      isActive,
		CreatedAt:     time.Now(),
	}

	if err := config.DB.Create(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapToWarehouseResponse(warehouse))
}

// @Summary List warehouses
// @Tags Warehouse
// @Produce json
// @Param active query bool false "Filter by active status"
// @Success 200 {array} dto.WarehouseResponseDTO
// @Failure 500 {object} map[string]string
// @Router /warehouses [get]
func ListWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	query := config.DB

	// Apply filters
	if active := c.Query("active"); active != "" {
		if active == "true" {
			query = query.Where("is_active = ?", true)
		} else if active == "false" {
			query = query.Where("is_active = ?", false)
		}
	}

	if err := query.Find(&warehouses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.WarehouseResponseDTO
	for _, w := range warehouses {
		resp = append(resp, mapToWarehouseResponse(w))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get warehouse by ID
// @Tags Warehouse
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} dto.WarehouseResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /warehouses/{id} [get]
func GetWarehouse(c *gin.Context) {
	id := c.Param("id")
	var warehouse models.Warehouse
	if err := config.DB.First(&warehouse, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}
	c.JSON(http.StatusOK, mapToWarehouseResponse(warehouse))
}

// @Summary Update warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param id path int true "Warehouse ID"
// @Param payload body dto.WarehouseUpdateDTO true "Update payload"
// @Success 200 {object} dto.WarehouseResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /warehouses/{id} [put]
func UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var payload dto.WarehouseUpdateDTO
	var warehouse models.Warehouse

	if err := config.DB.First(&warehouse, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.WarehouseName != "" {
		warehouse.WarehouseName = payload.WarehouseName
	}
	if payload.WarehouseType != nil {
		warehouse.WarehouseType = payload.WarehouseType
	}
	if payload.Address != nil {
		warehouse.Address = payload.Address
	}
	if payload.ContactPerson != nil {
		warehouse.ContactPerson = payload.ContactPerson
	}
	if payload.ContactPhone != nil {
		warehouse.ContactPhone = payload.ContactPhone
	}
	if payload.ContactEmail != nil {
		warehouse.ContactEmail = payload.ContactEmail
	}
	if payload.IsActive != nil {
		warehouse.IsActive = *payload.IsActive
	}
	now := time.Now()
	warehouse.UpdatedAt = &now

	if err := config.DB.Save(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapToWarehouseResponse(warehouse))
}

// @Summary Delete warehouse
// @Tags Warehouse
// @Param id path int true "Warehouse ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /warehouses/{id} [delete]
func DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Warehouse{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}

func mapToWarehouseResponse(w models.Warehouse) dto.WarehouseResponseDTO {
	return dto.WarehouseResponseDTO{
		ID:            w.ID,
		WarehouseCode: w.WarehouseCode,
		WarehouseName: w.WarehouseName,
		WarehouseType: w.WarehouseType,
		Address:       w.Address,
		ContactPerson: w.ContactPerson,
		ContactPhone:  w.ContactPhone,
		ContactEmail:  w.ContactEmail,
		IsActive:      w.IsActive,
		CreatedAt:     w.CreatedAt,
		UpdatedAt:     w.UpdatedAt,
	}
}
