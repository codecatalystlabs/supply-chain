package dto

import "time"

// UserCreateDTO for creating a new user
type UserCreateDTO struct {
	Username  string  `json:"username" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=6"`
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	FacilityID *uint64 `json:"facility_id,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
	RoleIDs   []uint64 `json:"role_ids,omitempty"`
}

// UserUpdateDTO for updating a user
type UserUpdateDTO struct {
	Email     *string  `json:"email,omitempty"`
	FirstName *string  `json:"first_name,omitempty"`
	LastName  *string  `json:"last_name,omitempty"`
	FacilityID *uint64 `json:"facility_id,omitempty"`
	IsActive  *bool    `json:"is_active,omitempty"`
	RoleIDs   []uint64 `json:"role_ids,omitempty"`
}

// UserResponseDTO for API responses
type UserResponseDTO struct {
	ID          uint64     `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	FacilityID  *uint64    `json:"facility_id,omitempty"`
	FacilityCode *string   `json:"facility_code,omitempty"`
	FacilityName *string   `json:"facility_name,omitempty"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Roles       []RoleResponseDTO `json:"roles,omitempty"`
}

// RoleResponseDTO for role information in user responses
type RoleResponseDTO struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
}

