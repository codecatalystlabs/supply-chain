package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

// User represents a system user
type User struct {
	ID           uint64    `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:100;not null;uniqueIndex" json:"username"`
	Email        string    `gorm:"size:255;not null;uniqueIndex" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	FirstName    string    `gorm:"size:100" json:"first_name"`
	LastName     string    `gorm:"size:100" json:"last_name"`
	FacilityID   *uint64   `gorm:"index" json:"facility_id,omitempty"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Facility     *Facility  `gorm:"foreignKey:FacilityID" json:"facility,omitempty"`
	Roles        []Role     `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	Permissions  []Permission `gorm:"many2many:user_permissions;" json:"permissions,omitempty"`
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword verifies if the provided password matches the user's password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// HasRole checks if user has a specific role
func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

// HasPermission checks if user has a specific permission
func (u *User) HasPermission(permissionName string) bool {
	// Check direct permissions
	for _, perm := range u.Permissions {
		if perm.Name == permissionName {
			return true
		}
	}
	// Check permissions through roles
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if perm.Name == permissionName {
				return true
			}
		}
	}
	return false
}

// Role represents a user role
type Role struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	DisplayName string    `gorm:"size:255" json:"display_name"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Users       []User       `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// Permission represents a system permission
type Permission struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null;uniqueIndex" json:"name"`
	DisplayName string    `gorm:"size:255" json:"display_name"`
	Resource    string    `gorm:"size:100;not null;index" json:"resource"` // e.g., "facilities", "warehouses"
	Action      string    `gorm:"size:100;not null;index" json:"action"`   // e.g., "create", "read", "update", "delete"
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Users []User `gorm:"many2many:user_permissions;" json:"users,omitempty"`
	Roles []Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	UserID    uint64    `gorm:"primaryKey" json:"user_id"`
	RoleID    uint64    `gorm:"primaryKey" json:"role_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

// UserPermission represents the many-to-many relationship between users and permissions
type UserPermission struct {
	UserID       uint64    `gorm:"primaryKey" json:"user_id"`
	PermissionID uint64    `gorm:"primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
}

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	RoleID       uint64    `gorm:"primaryKey" json:"role_id"`
	PermissionID uint64    `gorm:"primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
}

