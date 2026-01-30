package handlers

import (
	"net/http"
	"time"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create procurement plan
// @Tags ProcurementPlan
// @Accept json
// @Produce json
// @Param payload body dto.ProcurementPlanCreateDTO true "Procurement plan payload"
// @Success 201 {object} dto.ProcurementPlanResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /procurement [post]
func CreateProcurementPlan(c *gin.Context) {
	var payload dto.ProcurementPlanCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan := models.ProcurementPlan{
		PlanSystemCode: payload.PlanSystemCode,
		StoreCode:      payload.StoreCode,
		CreatedAt:      time.Now(),
		Notes:          payload.Notes,
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create items
	for _, it := range payload.Items {
		item := models.ProcurementPlanItem{
			ProcurementID: plan.ID,
			ProductCode:   it.ProductCode,
			Quantity:      it.Quantity,
			NeededBy:      it.NeededBy,
			Status:        "planned",
		}
		_ = config.DB.Create(&item)
	}

	// load items
	var items []models.ProcurementPlanItem
	config.DB.Where("procurement_id = ?", plan.ID).Find(&items)

	resp := mapToProcurementPlanResponse(plan, items)
	c.JSON(http.StatusCreated, resp)
}

// @Summary List procurement plans
// @Tags ProcurementPlan
// @Produce json
// @Success 200 {array} dto.ProcurementPlanResponseDTO
// @Failure 500 {object} map[string]string
// @Router /procurement [get]
func ListProcurementPlans(c *gin.Context) {
	var plans []models.ProcurementPlan
	if err := config.DB.Find(&plans).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ProcurementPlanResponseDTO
	for _, p := range plans {
		var items []models.ProcurementPlanItem
		config.DB.Where("procurement_id = ?", p.ID).Find(&items)
		resp = append(resp, mapToProcurementPlanResponse(p, items))
	}
	c.JSON(200, resp)
}

// @Summary Get procurement plan
// @Tags ProcurementPlan
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.ProcurementPlanResponseDTO
// @Failure 404 {object} map[string]string
// @Router /procurement/{id} [get]
func GetProcurementPlan(c *gin.Context) {
	id := c.Param("id")
	var plan models.ProcurementPlan
	if err := config.DB.First(&plan, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Plan not found"})
		return
	}
	var items []models.ProcurementPlanItem
	config.DB.Where("procurement_id = ?", plan.ID).Find(&items)
	c.JSON(200, mapToProcurementPlanResponse(plan, items))
}

// @Summary Delete procurement plan
// @Tags ProcurementPlan
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /procurement/{id} [delete]
func DeleteProcurementPlan(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.ProcurementPlan{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	config.DB.Where("procurement_id = ?", id).Delete(&models.ProcurementPlanItem{})
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToProcurementPlanResponse(p models.ProcurementPlan, items []models.ProcurementPlanItem) dto.ProcurementPlanResponseDTO {
	var itResp []dto.ProcurementPlanItemResponseDTO
	for _, it := range items {
		itResp = append(itResp, dto.ProcurementPlanItemResponseDTO{
			ID:            it.ID,
			ProcurementID: it.ProcurementID,
			ProductCode:   it.ProductCode,
			Quantity:      it.Quantity,
			NeededBy:      it.NeededBy,
			Status:        it.Status,
		})
	}
	return dto.ProcurementPlanResponseDTO{
		ID:             p.ID,
		PlanSystemCode: p.PlanSystemCode,
		StoreCode:      p.StoreCode,
		CreatedAt:      p.CreatedAt,
		Notes:          p.Notes,
		Items:          itResp,
	}
}
