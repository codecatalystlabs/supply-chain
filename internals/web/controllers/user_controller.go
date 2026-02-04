package controllers

import (
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowUsers displays all users
func ShowUsers(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "User Management",
	}
	services.GetTemplateService().RenderTemplate(c, "users.tpl", viewData)
}

// ShowRoles displays all roles
func ShowRoles(c *gin.Context) {
	viewData := &views.ViewData{
		Title: "Role Management",
	}
	services.GetTemplateService().RenderTemplate(c, "roles.tpl", viewData)
}

