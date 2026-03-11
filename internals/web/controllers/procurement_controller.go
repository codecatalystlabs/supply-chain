package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowProcurement displays procurement overview
func ShowProcurement(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Procurement",
	}
	services.GetTemplateService().RenderTemplate(c, "procurement.tpl", viewData)
}

// ShowPurchaseOrders displays all purchase orders
func ShowPurchaseOrders(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Purchase Orders",
	}
	services.GetTemplateService().RenderTemplate(c, "purchase-orders.tpl", viewData)
}

// ShowProcurementPlans displays procurement plans
func ShowProcurementPlans(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Procurement Plans",
	}
	services.GetTemplateService().RenderTemplate(c, "procurement-plans.tpl", viewData)
}

// ShowProcurementPlanImport displays the import procurement plan (XLS) page
func ShowProcurementPlanImport(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Import Procurement Plan",
	}
	services.GetTemplateService().RenderTemplate(c, "procurement-plan-import.tpl", viewData)
}
