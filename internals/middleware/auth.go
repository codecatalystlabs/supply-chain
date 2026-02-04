package middleware

import (
	"net/http"
	"strings"

	"supply-chain/internals/config"
	"supply-chain/internals/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired middleware checks if user is authenticated
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			// Check if it's an API request
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
				c.Abort()
				return
			}
			// Redirect to login for web requests
			c.Redirect(http.StatusFound, "/cp/login")
			c.Abort()
			return
		}

		// Load user from database with facility
		var user models.User
		if err := config.DB.Preload("Roles.Permissions").Preload("Permissions").Preload("Facility").First(&user, userID).Error; err != nil {
			session.Clear()
			session.Save()
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			} else {
				c.Redirect(http.StatusFound, "/cp/login")
			}
			c.Abort()
			return
		}

		// Check if user is active
		if !user.IsActive {
			session.Clear()
			session.Save()
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusForbidden, gin.H{"error": "Account is inactive"})
			} else {
				c.Redirect(http.StatusFound, "/cp/login?error=Account+is+inactive")
			}
			c.Abort()
			return
		}

		// Store user in context
		c.Set("user", &user)
		c.Set("user_id", user.ID)
		c.Next()
	}
}

// RequireRole middleware checks if user has a specific role
func RequireRole(roleNames ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			} else {
				c.Redirect(http.StatusFound, "/cp/login")
			}
			c.Abort()
			return
		}

		u := user.(*models.User)
		hasRole := false
		for _, roleName := range roleNames {
			if u.HasRole(roleName) {
				hasRole = true
				break
			}
		}

		if !hasRole {
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			} else {
				c.HTML(http.StatusForbidden, "error.tpl", gin.H{
					"title": "Access Denied",
					"error": "You do not have permission to access this resource",
				})
			}
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequirePermission middleware checks if user has a specific permission
func RequirePermission(permissionNames ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			} else {
				c.Redirect(http.StatusFound, "/cp/login")
			}
			c.Abort()
			return
		}

		u := user.(*models.User)
		hasPermission := false
		for _, permName := range permissionNames {
			if u.HasPermission(permName) {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusForbidden, gin.H{
					"error":                "Insufficient permissions",
					"required_permissions": permissionNames,
				})
			} else {
				// Return JSON error for web requests too, since we don't have an error template
				c.JSON(http.StatusForbidden, gin.H{
					"error":                "Access Denied",
					"message":              "You do not have permission to access this resource",
					"required_permissions": permissionNames,
				})
			}
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAnyPermission middleware checks if user has at least one of the specified permissions
func RequireAnyPermission(permissionNames ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			if strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			} else {
				c.Redirect(http.StatusFound, "/cp/login")
			}
			c.Abort()
			return
		}

		u := user.(*models.User)
		for _, permName := range permissionNames {
			if u.HasPermission(permName) {
				c.Next()
				return
			}
		}

		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		} else {
			c.HTML(http.StatusForbidden, "error.tpl", gin.H{
				"title": "Access Denied",
				"error": "You do not have permission to access this resource",
			})
		}
		c.Abort()
	}
}
