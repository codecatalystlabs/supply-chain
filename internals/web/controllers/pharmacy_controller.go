package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowPharmacies displays all pharmacies
func ShowPharmacies(c *gin.Context) {
	viewData := &views.ViewData{
		Title:   "Pharmacies",
	}
	services.GetTemplateService().RenderTemplate(c, "pharmacies.tpl", viewData)
}

// ShowPharmacyStock displays pharmacy stock
func ShowPharmacyStock(c *gin.Context) {
	viewData := &views.ViewData{
		Title:   "Pharmacy Stock",
	}
	services.GetTemplateService().RenderTemplate(c, "pharmacy-stock.tpl", viewData)
}
