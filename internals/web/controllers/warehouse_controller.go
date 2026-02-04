package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowWarehouses displays all warehouses
func ShowWarehouses(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Warehouses",
	}
	services.GetTemplateService().RenderTemplate(c, "warehouses.tpl", viewData)
}

// ShowWarehouseOrders displays warehouse orders
func ShowWarehouseOrders(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Warehouse Orders",
	}
	services.GetTemplateService().RenderTemplate(c, "warehouse-orders.tpl", viewData)
}

// ShowGoodsReceipt displays goods receipt records
func ShowGoodsReceipt(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Goods Receipt",
	}
	services.GetTemplateService().RenderTemplate(c, "goods-receipt.tpl", viewData)
}
