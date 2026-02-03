package views

import (
	"supply-chain/internals/web/services"

	"github.com/gin-gonic/gin"
)

// ViewData represents the common data structure for all views
type ViewData struct {
	Title       string
	Message     string
	Data        interface{}
	Errors      map[string]string
	Success     bool
	CurrentUser interface{}
}

// RenderLogin renders the login page
func RenderLogin(c *gin.Context, data *ViewData) {
	if data == nil {
		data = &ViewData{
			Title:   "Login",
			Errors:  make(map[string]string),
			Success: false,
		}
	}
	if data.Errors == nil {
		data.Errors = make(map[string]string)
	}

	templateService := services.GetTemplateService()
	templateService.RenderTemplate(c, "login.tpl", data)
}

// RenderDashboard renders the dashboard page
func RenderDashboard(c *gin.Context, data *ViewData) {
	if data == nil {
		data = &ViewData{
			Title: "Dashboard",
		}
	}

	templateService := services.GetTemplateService()
	templateService.RenderTemplate(c, "dashboard.tpl", data)
}

// RenderIndex renders the index/home page
func RenderIndex(c *gin.Context, data *ViewData) {
	if data == nil {
		data = &ViewData{
			Title: "Home",
		}
	}

	templateService := services.GetTemplateService()
	templateService.RenderTemplate(c, "index.tpl", data)
}

// RenderError renders error page
func RenderError(c *gin.Context, statusCode int, message string) {
	data := &ViewData{
		Title:   "Error",
		Message: message,
	}

	templateService := services.GetTemplateService()
	c.Status(statusCode)
	templateService.RenderTemplate(c, "error.tpl", data)
}
