package controllers

import (
	"log"
	"supply-chain/internals/web/views"

	"github.com/gin-gonic/gin"
)

// ShowLoginPage renders the login page
func ShowLoginPage(c *gin.Context) {
	data := &views.ViewData{
		Title:   "Login",
		Errors:  make(map[string]string),
		Success: false,
	}

	// Check if there are any flash messages
	if flashMessage, exists := c.GetQuery("message"); exists {
		data.Message = flashMessage
	}

	if flashError, exists := c.GetQuery("error"); exists {
		data.Errors["general"] = flashError
	}

	views.RenderLogin(c, data)
}

// HandleLogin processes login form submission
func HandleLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// TODO: Validate credentials against database
	log.Printf("Login attempt for email: %s\n", email)

	// For now, redirect to dashboard on any login attempt
	// In production, validate against actual user database
	if email == "" || password == "" {
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Email and password are required"},
		}
		views.RenderLogin(c, data)
		return
	}

	// TODO: Set session/JWT token
	c.Redirect(302, "/cp/dashboard")
}

// ShowDashboard renders the dashboard page
func ShowDashboard(c *gin.Context) {
	// TODO: Check if user is authenticated
	data := &views.ViewData{
		Title: "Dashboard",
	}
	views.RenderDashboard(c, data)
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// TODO: Clear session/token
	c.Redirect(302, "/cp/login?message=Logged+out+successfully")
}
