package controllers

import (
	"log"
	"net/http"
	"time"
	"supply-chain/internals/config"
	"supply-chain/internals/models"
	"supply-chain/internals/web/services"
	"supply-chain/internals/web/views"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ShowLoginPage renders the login page
func ShowLoginPage(c *gin.Context) {
	// Redirect if already logged in
	session := sessions.Default(c)
	if userID := session.Get("user_id"); userID != nil {
		c.Redirect(302, "/cp/home")
		return
	}

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
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Username and password are required"},
		}
		views.RenderLogin(c, data)
		return
	}

	// Find user by username or email
	var user models.User
	if err := config.DB.Where("username = ? OR email = ?", username, username).Preload("Roles.Permissions").Preload("Permissions").First(&user).Error; err != nil {
		log.Printf("Login failed for username: %s - User not found\n", username)
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Invalid username or password"},
		}
		views.RenderLogin(c, data)
		return
	}

	// Check if user is active
	if !user.IsActive {
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Account is inactive. Please contact administrator."},
		}
		views.RenderLogin(c, data)
		return
	}

	// Verify password
	if !user.CheckPassword(password) {
		log.Printf("Login failed for username: %s - Invalid password\n", username)
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Invalid username or password"},
		}
		views.RenderLogin(c, data)
		return
	}

	// Update last login
	now := time.Now()
	user.LastLoginAt = &now
	config.DB.Save(&user)

	// Set session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		log.Printf("Error saving session: %v\n", err)
		data := &views.ViewData{
			Title:  "Login",
			Errors: map[string]string{"general": "Failed to create session. Please try again."},
		}
		views.RenderLogin(c, data)
		return
	}

	log.Printf("User %s logged in successfully\n", username)
	c.Redirect(302, "/cp/home")
}

// ShowHome renders the home page
func ShowHome(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Redirect(302, "/cp/login")
		return
	}
	u := user.(*models.User)
	
	// Ensure facility is loaded
	if u.FacilityID != nil && u.Facility == nil {
		config.DB.Preload("Facility").First(u, u.ID)
	}

	// Fetch dashboard statistics
	var stats struct {
		TotalFacilities      int64
		TotalWarehouses      int64
		TotalPharmacies      int64
		TotalProcurementPlans int64
		TotalPurchaseOrders  int64
		TotalStockOnHand     int64
		TotalStockTransfers  int64
		PendingTransfers     int64
	}

	config.DB.Model(&models.Facility{}).Count(&stats.TotalFacilities)
	config.DB.Model(&models.Warehouse{}).Count(&stats.TotalWarehouses)
	config.DB.Model(&models.Pharmacy{}).Count(&stats.TotalPharmacies)
	config.DB.Model(&models.ProcurementPlan{}).Count(&stats.TotalProcurementPlans)
	config.DB.Model(&models.PurchaseOrder{}).Count(&stats.TotalPurchaseOrders)
	config.DB.Model(&models.StockOnHand{}).Count(&stats.TotalStockOnHand)
	config.DB.Model(&models.StockTransfer{}).Count(&stats.TotalStockTransfers)
	config.DB.Model(&models.StockTransfer{}).Where("status = ?", "pending").Count(&stats.PendingTransfers)

	// Create a user map with facility name for template access
	userData := map[string]interface{}{
		"ID":        u.ID,
		"Username":  u.Username,
		"FirstName": u.FirstName,
		"LastName":  u.LastName,
		"IsActive":  u.IsActive,
	}
	if u.Facility != nil {
		userData["FacilityName"] = u.Facility.FacilityName
		userData["FacilityCode"] = u.Facility.FacilityCode
	}
	
	data := &views.ViewData{
		Title: "Home",
		Data: map[string]interface{}{
			"user":  userData,
			"stats": stats,
		},
		CurrentUser: u,
		Username:    u.FirstName + " " + u.LastName,
	}
	services.GetTemplateService().RenderTemplate(c, "home.tpl", data)
}

// ShowDashboard renders the dashboard page (alias for ShowHome)
func ShowDashboard(c *gin.Context) {
	ShowHome(c)
}

// Logout handles user logout
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	
	// Clear all session values
	session.Clear()
	
	// Save the cleared session
	if err := session.Save(); err != nil {
		log.Printf("Error clearing session: %v\n", err)
	}
	
	// Redirect to login page
	c.Redirect(http.StatusFound, "/cp/login")
}
