package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowInventory displays inventory management overview
func ShowInventory(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Inventory Management",
	}
	services.GetTemplateService().RenderTemplate(c, "inventory.tpl", viewData)
}

// ShowStockList displays all stock records
func ShowStockList(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Stock Records",
	}
	services.GetTemplateService().RenderTemplate(c, "stock-list.tpl", viewData)
}
