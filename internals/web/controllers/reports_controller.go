package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowPatientVisits displays patient visit records
func ShowPatientVisits(c *gin.Context) {
	viewData := &views.ViewData{
		Title:   "Patient Visits",
	}
	services.GetTemplateService().RenderTemplate(c, "patient-visits.tpl", viewData)
}

// ShowProductAMC displays product AMC (Average Monthly Consumption)
func ShowProductAMC(c *gin.Context) {
	viewData := &views.ViewData{
		Title:   "Product AMC",
	}
	services.GetTemplateService().RenderTemplate(c, "product-amc.tpl", viewData)
}
