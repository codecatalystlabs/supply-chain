package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowStockOnHand displays stock on hand records
func ShowStockOnHand(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock on Hand",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-on-hand.tpl", viewData)
}

// ShowStockDispensed displays stock dispensed records
func ShowStockDispensed(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock Dispensed",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-dispensed.tpl", viewData)
}

// ShowStockAdjustments displays stock adjustment records
func ShowStockAdjustments(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock Adjustments",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-adjustments.tpl", viewData)
}

// ShowStockReturns displays stock return records
func ShowStockReturns(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock Returns",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-returns.tpl", viewData)
}

// ShowStockTransfers displays stock transfer records
func ShowStockTransfers(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock Transfers",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-transfers.tpl", viewData)
}

