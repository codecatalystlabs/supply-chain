package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowFacilities displays all facilities
func ShowFacilities(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Facilities",
	}
	services.GetTemplateService().RenderTemplate(c, "facilities.tpl", viewData)
}

// ShowFacilityOrders displays facility orders
func ShowFacilityOrders(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Facility Orders",
	}
	services.GetTemplateService().RenderTemplate(c, "facility-orders.tpl", viewData)
}
