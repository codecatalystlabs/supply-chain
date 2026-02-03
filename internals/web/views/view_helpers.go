package views

import (
	"supply-chain/internals/web/services"

	"github.com/gin-gonic/gin"
)

// ViewData represents the common data structure for all views
type ViewData struct {
	Title             string            // Page title
	Message           string            // Display message
	Data              interface{}       // Page-specific data
	Errors            map[string]string // Form errors
	Success           bool              // Success flag
	CurrentUser       interface{}       // Current user info
	Username          string            // User display name for header
	IsAdmin           bool              // User is admin
	CanManageServices bool              // User can manage services
	CanExploreData    bool              // User can explore data
	Content           string            // Content template name for dynamic layout (deprecated)
	ContentBlock      string            // Content block name to render
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
