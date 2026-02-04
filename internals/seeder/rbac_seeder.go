package seeder

import (
	"log"
	"strings"
	"supply-chain/internals/models"

	"gorm.io/gorm"
)

// SeedRBAC seeds roles and permissions
func SeedRBAC(db *gorm.DB) {
	// Create Permissions
	permissions := []models.Permission{
		// Facilities
		{Name: "facilities.read", DisplayName: "View Facilities", Resource: "facilities", Action: "read"},
		{Name: "facilities.create", DisplayName: "Create Facilities", Resource: "facilities", Action: "create"},
		{Name: "facilities.update", DisplayName: "Update Facilities", Resource: "facilities", Action: "update"},
		{Name: "facilities.delete", DisplayName: "Delete Facilities", Resource: "facilities", Action: "delete"},

		// Warehouses
		{Name: "warehouses.read", DisplayName: "View Warehouses", Resource: "warehouses", Action: "read"},
		{Name: "warehouses.create", DisplayName: "Create Warehouses", Resource: "warehouses", Action: "create"},
		{Name: "warehouses.update", DisplayName: "Update Warehouses", Resource: "warehouses", Action: "update"},
		{Name: "warehouses.delete", DisplayName: "Delete Warehouses", Resource: "warehouses", Action: "delete"},

		// Pharmacies
		{Name: "pharmacies.read", DisplayName: "View Pharmacies", Resource: "pharmacies", Action: "read"},
		{Name: "pharmacies.create", DisplayName: "Create Pharmacies", Resource: "pharmacies", Action: "create"},
		{Name: "pharmacies.update", DisplayName: "Update Pharmacies", Resource: "pharmacies", Action: "update"},
		{Name: "pharmacies.delete", DisplayName: "Delete Pharmacies", Resource: "pharmacies", Action: "delete"},

		// Procurement Plans
		{Name: "procurement_plans.read", DisplayName: "View Procurement Plans", Resource: "procurement_plans", Action: "read"},
		{Name: "procurement_plans.create", DisplayName: "Create Procurement Plans", Resource: "procurement_plans", Action: "create"},
		{Name: "procurement_plans.update", DisplayName: "Update Procurement Plans", Resource: "procurement_plans", Action: "update"},
		{Name: "procurement_plans.delete", DisplayName: "Delete Procurement Plans", Resource: "procurement_plans", Action: "delete"},

		// Purchase Orders
		{Name: "purchase_orders.read", DisplayName: "View Purchase Orders", Resource: "purchase_orders", Action: "read"},
		{Name: "purchase_orders.create", DisplayName: "Create Purchase Orders", Resource: "purchase_orders", Action: "create"},
		{Name: "purchase_orders.update", DisplayName: "Update Purchase Orders", Resource: "purchase_orders", Action: "update"},
		{Name: "purchase_orders.delete", DisplayName: "Delete Purchase Orders", Resource: "purchase_orders", Action: "delete"},

		// Stock Management
		{Name: "stock.read", DisplayName: "View Stock", Resource: "stock", Action: "read"},
		{Name: "stock.create", DisplayName: "Create Stock", Resource: "stock", Action: "create"},
		{Name: "stock.update", DisplayName: "Update Stock", Resource: "stock", Action: "update"},
		{Name: "stock.delete", DisplayName: "Delete Stock", Resource: "stock", Action: "delete"},
		{Name: "stock.transfer", DisplayName: "Transfer Stock", Resource: "stock", Action: "transfer"},
		{Name: "stock.adjust", DisplayName: "Adjust Stock", Resource: "stock", Action: "adjust"},

		// Patient Visits
		{Name: "patient_visits.read", DisplayName: "View Patient Visits", Resource: "patient_visits", Action: "read"},
		{Name: "patient_visits.create", DisplayName: "Create Patient Visits", Resource: "patient_visits", Action: "create"},

		// Reports
		{Name: "reports.read", DisplayName: "View Reports", Resource: "reports", Action: "read"},

		// Administration
		{Name: "admin.users", DisplayName: "Manage Users", Resource: "admin", Action: "users"},
		{Name: "admin.roles", DisplayName: "Manage Roles", Resource: "admin", Action: "roles"},
	}

	permCount := 0
	for _, perm := range permissions {
		var existing models.Permission
		if err := db.Where("name = ?", perm.Name).First(&existing).Error; err != nil {
			if err := db.Create(&perm).Error; err != nil {
				log.Printf("Failed to create permission %s: %v", perm.Name, err)
			} else {
				permCount++
			}
		}
	}
	log.Printf("✅ Created/verified %d permissions", permCount)

	// Create Roles
	roles := []struct {
		role        models.Role
		permissions []string
	}{
		{
			role: models.Role{
				Name:        "super_admin",
				DisplayName: "Super Administrator",
				Description: ptrString("Full system access"),
			},
			permissions: []string{}, // All permissions
		},
		{
			role: models.Role{
				Name:        "admin",
				DisplayName: "Administrator",
				Description: ptrString("Administrative access"),
			},
			permissions: []string{
				"facilities.read", "facilities.create", "facilities.update", "facilities.delete",
				"warehouses.read", "warehouses.create", "warehouses.update", "warehouses.delete",
				"pharmacies.read", "pharmacies.create", "pharmacies.update", "pharmacies.delete",
				"procurement_plans.read", "procurement_plans.create", "procurement_plans.update",
				"purchase_orders.read", "purchase_orders.create", "purchase_orders.update",
				"stock.read", "stock.create", "stock.update", "stock.transfer", "stock.adjust",
				"patient_visits.read", "reports.read",
				"admin.users", "admin.roles", // Add admin permissions
			},
		},
		{
			role: models.Role{
				Name:        "procurement_officer",
				DisplayName: "Procurement Officer",
				Description: ptrString("Manages procurement and orders"),
			},
			permissions: []string{
				"facilities.read", "warehouses.read",
				"procurement_plans.read", "procurement_plans.create", "procurement_plans.update",
				"purchase_orders.read", "purchase_orders.create", "purchase_orders.update",
				"reports.read",
			},
		},
		{
			role: models.Role{
				Name:        "warehouse_manager",
				DisplayName: "Warehouse Manager",
				Description: ptrString("Manages warehouse operations"),
			},
			permissions: []string{
				"warehouses.read", "warehouses.update",
				"stock.read", "stock.create", "stock.update", "stock.transfer", "stock.adjust",
				"purchase_orders.read",
				"reports.read",
			},
		},
		{
			role: models.Role{
				Name:        "pharmacist",
				DisplayName: "Pharmacist",
				Description: ptrString("Manages pharmacy operations"),
			},
			permissions: []string{
				"pharmacies.read",
				"stock.read", "stock.create", "stock.update",
				"patient_visits.read", "patient_visits.create",
			},
		},
		{
			role: models.Role{
				Name:        "viewer",
				DisplayName: "Viewer",
				Description: ptrString("Read-only access"),
			},
			permissions: []string{
				"facilities.read", "warehouses.read", "pharmacies.read",
				"procurement_plans.read", "purchase_orders.read",
				"stock.read", "patient_visits.read", "reports.read",
			},
		},
	}

	roleCount := 0
	for _, r := range roles {
		var existing models.Role
		if err := db.Where("name = ?", r.role.Name).First(&existing).Error; err != nil {
			if err := db.Create(&r.role).Error; err != nil {
				log.Printf("Failed to create role %s: %v", r.role.Name, err)
				continue
			}
			roleCount++
		} else {
			r.role = existing
		}

		// Assign permissions to role
		if r.role.Name == "super_admin" {
			// Super admin gets all permissions
			var allPerms []models.Permission
			db.Find(&allPerms)
			if err := db.Model(&r.role).Association("Permissions").Replace(allPerms); err != nil {
				log.Printf("⚠️ Failed to assign permissions to super_admin: %v", err)
			} else {
				log.Printf("✅ Assigned %d permissions to super_admin role", len(allPerms))
			}
		} else {
			var perms []models.Permission
			for _, permName := range r.permissions {
				var perm models.Permission
				if err := db.Where("name = ?", permName).First(&perm).Error; err == nil {
					perms = append(perms, perm)
				}
			}
			if err := db.Model(&r.role).Association("Permissions").Replace(perms); err != nil {
				log.Printf("⚠️ Failed to assign permissions to role %s: %v", r.role.Name, err)
			} else {
				log.Printf("✅ Assigned %d permissions to role %s", len(perms), r.role.Name)
			}
		}
	}
	log.Printf("✅ Created/verified %d roles", roleCount)

	// Create default admin user
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		adminUser := models.User{
			Username:  "admin",
			Email:     "admin@moh.go.ug",
			FirstName: "System",
			LastName:  "Administrator",
			IsActive:  true,
		}
		if err := adminUser.SetPassword("admin123"); err != nil {
			log.Printf("Failed to set admin password: %v", err)
		} else {
			if err := db.Create(&adminUser).Error; err != nil {
				log.Printf("Failed to create admin user: %v", err)
			} else {
				// Assign ALL roles to admin user
				var allRoles []models.Role
				if err := db.Find(&allRoles).Error; err == nil {
					db.Model(&adminUser).Association("Roles").Replace(allRoles)
					log.Printf("✅ Assigned all %d roles to admin user", len(allRoles))
				} else {
					log.Printf("Failed to fetch roles: %v", err)
					// Fallback: assign super_admin role
					var superAdminRole models.Role
					if err := db.Where("name = ?", "super_admin").First(&superAdminRole).Error; err == nil {
						db.Model(&adminUser).Association("Roles").Append(&superAdminRole)
					}
				}
			}
		}
	} else {
		// Update existing admin user to have all roles
		var adminUser models.User
		if err := db.Where("username = ?", "admin").First(&adminUser).Error; err == nil {
			var allRoles []models.Role
			if err := db.Find(&allRoles).Error; err == nil {
				db.Model(&adminUser).Association("Roles").Replace(allRoles)
				log.Printf("✅ Updated admin user with all %d roles", len(allRoles))
			}
		}
	}

	log.Println("✅ RBAC data seeded")

	// Verify data exists before seeding relationships
	var roleCheck []models.Role
	var permCheck []models.Permission
	db.Find(&roleCheck)
	db.Find(&permCheck)
	log.Printf("📊 Verification: Found %d roles and %d permissions in database", len(roleCheck), len(permCheck))

	// Seed role permissions, user roles, and user permissions
	SeedRolePermissions(db)
	SeedUserRoles(db)
	SeedUserPermissions(db)
}

// SeedRolePermissions assigns permissions to roles (RolePermission join table)
func SeedRolePermissions(db *gorm.DB) {
	log.Println("🌱 Seeding role permissions (RolePermission relationships)...")

	// Get all roles and permissions
	var roles []models.Role
	var permissions []models.Permission
	if err := db.Find(&roles).Error; err != nil {
		log.Printf("❌ Error loading roles: %v", err)
		return
	}
	if err := db.Find(&permissions).Error; err != nil {
		log.Printf("❌ Error loading permissions: %v", err)
		return
	}

	if len(roles) == 0 || len(permissions) == 0 {
		log.Println("⚠️ No roles or permissions found, skipping role permission seeding")
		return
	}

	log.Printf("📊 Found %d roles and %d permissions to assign", len(roles), len(permissions))

	// Create permission map for quick lookup
	permMap := make(map[string]models.Permission)
	for _, p := range permissions {
		permMap[p.Name] = p
	}

	// Helper function to safely get permission from map
	getPerm := func(name string) models.Permission {
		if perm, ok := permMap[name]; ok && perm.ID != 0 {
			return perm
		}
		return models.Permission{}
	}

	// Assign permissions to each role based on role name
	for _, role := range roles {
		var permsToAssign []models.Permission

		switch role.Name {
		case "super_admin":
			// Super admin gets ALL permissions
			permsToAssign = permissions
		case "admin":
			// Admin gets most permissions
			permNames := []string{
				"facilities.read", "facilities.create", "facilities.update", "facilities.delete",
				"warehouses.read", "warehouses.create", "warehouses.update", "warehouses.delete",
				"pharmacies.read", "pharmacies.create", "pharmacies.update", "pharmacies.delete",
				"procurement_plans.read", "procurement_plans.create", "procurement_plans.update", "procurement_plans.delete",
				"purchase_orders.read", "purchase_orders.create", "purchase_orders.update", "purchase_orders.delete",
				"stock.read", "stock.create", "stock.update", "stock.delete", "stock.transfer", "stock.adjust",
				"patient_visits.read", "patient_visits.create",
				"reports.read",
				"admin.users", "admin.roles",
			}
			for _, name := range permNames {
				if perm := getPerm(name); perm.ID != 0 {
					permsToAssign = append(permsToAssign, perm)
				}
			}
		case "procurement_officer":
			permNames := []string{
				"facilities.read", "warehouses.read",
				"procurement_plans.read", "procurement_plans.create", "procurement_plans.update",
				"purchase_orders.read", "purchase_orders.create", "purchase_orders.update",
				"reports.read",
			}
			for _, name := range permNames {
				if perm := getPerm(name); perm.ID != 0 {
					permsToAssign = append(permsToAssign, perm)
				}
			}
		case "warehouse_manager":
			permNames := []string{
				"warehouses.read", "warehouses.update",
				"stock.read", "stock.create", "stock.update", "stock.transfer", "stock.adjust",
				"purchase_orders.read",
				"reports.read",
			}
			for _, name := range permNames {
				if perm := getPerm(name); perm.ID != 0 {
					permsToAssign = append(permsToAssign, perm)
				}
			}
		case "pharmacist":
			permNames := []string{
				"pharmacies.read",
				"stock.read", "stock.create", "stock.update",
				"patient_visits.read", "patient_visits.create",
			}
			for _, name := range permNames {
				if perm := getPerm(name); perm.ID != 0 {
					permsToAssign = append(permsToAssign, perm)
				}
			}
		case "viewer":
			permNames := []string{
				"facilities.read", "warehouses.read", "pharmacies.read",
				"procurement_plans.read", "purchase_orders.read",
				"stock.read", "patient_visits.read", "reports.read",
			}
			for _, name := range permNames {
				if perm := getPerm(name); perm.ID != 0 {
					permsToAssign = append(permsToAssign, perm)
				}
			}
		}

		// Replace all permissions for this role using GORM association
		if len(permsToAssign) > 0 {
			if err := db.Model(&role).Association("Permissions").Replace(permsToAssign); err != nil {
				log.Printf("❌ Failed to assign permissions to role %s: %v", role.Name, err)
			} else {
				log.Printf("✅ Assigned %d permissions to role: %s", len(permsToAssign), role.Name)
			}
		} else {
			log.Printf("⚠️ No valid permissions found for role: %s", role.Name)
		}
	}

	log.Println("✅ Role permissions seeded")
}

// SeedUserRoles assigns roles to users (UserRole join table)
func SeedUserRoles(db *gorm.DB) {
	log.Println("🌱 Seeding user roles (UserRole relationships)...")

	var users []models.User
	var roles []models.Role
	if err := db.Find(&users).Error; err != nil {
		log.Printf("❌ Error loading users: %v", err)
		return
	}
	if err := db.Find(&roles).Error; err != nil {
		log.Printf("❌ Error loading roles: %v", err)
		return
	}

	if len(users) == 0 || len(roles) == 0 {
		log.Println("⚠️ No users or roles found, skipping user role seeding")
		return
	}

	log.Printf("📊 Found %d users and %d roles to assign", len(users), len(roles))

	// Create role map
	roleMap := make(map[string]models.Role)
	for _, r := range roles {
		roleMap[r.Name] = r
	}

	// Assign roles to users based on username
	for _, user := range users {
		var rolesToAssign []models.Role

		switch user.Username {
		case "admin":
			// Admin user gets ALL roles
			rolesToAssign = roles
			log.Printf("🔑 Admin user found, assigning all %d roles", len(roles))
		default:
			// Other users get viewer role by default (can be customized)
			if viewerRole, ok := roleMap["viewer"]; ok {
				rolesToAssign = []models.Role{viewerRole}
			}
		}

		// Replace all roles for this user using GORM association
		if len(rolesToAssign) > 0 {
			if err := db.Model(&user).Association("Roles").Replace(rolesToAssign); err != nil {
				log.Printf("❌ Failed to assign roles to user %s: %v", user.Username, err)
			} else {
				roleNames := make([]string, len(rolesToAssign))
				for i, r := range rolesToAssign {
					roleNames[i] = r.Name
				}
				log.Printf("✅ Assigned roles [%s] to user: %s", strings.Join(roleNames, ", "), user.Username)
			}
		}
	}

	log.Println("✅ User roles seeded")
}

// SeedUserPermissions assigns permissions directly to users (UserPermission join table)
func SeedUserPermissions(db *gorm.DB) {
	log.Println("🌱 Seeding user permissions (UserPermission relationships)...")

	var users []models.User
	var permissions []models.Permission
	if err := db.Find(&users).Error; err != nil {
		log.Printf("❌ Error loading users: %v", err)
		return
	}
	if err := db.Find(&permissions).Error; err != nil {
		log.Printf("❌ Error loading permissions: %v", err)
		return
	}

	if len(users) == 0 || len(permissions) == 0 {
		log.Println("⚠️ No users or permissions found, skipping user permission seeding")
		return
	}

	log.Printf("📊 Found %d users and %d permissions", len(users), len(permissions))

	// Assign direct permissions to users (optional - usually permissions come from roles)
	// This is useful for granting specific permissions that override role-based permissions
	for _, user := range users {
		var permsToAssign []models.Permission

		switch user.Username {
		case "admin":
			// Admin user gets ALL permissions directly (in addition to role-based)
			permsToAssign = permissions
			log.Printf("🔑 Admin user found, assigning all %d direct permissions", len(permissions))
		default:
			// Other users don't get direct permissions (they get them from roles)
			permsToAssign = []models.Permission{}
		}

		// Replace all direct permissions for this user using GORM association
		if len(permsToAssign) > 0 {
			// Reload user to ensure we have the latest version
			var currentUser models.User
			if err := db.First(&currentUser, user.ID).Error; err != nil {
				log.Printf("❌ Failed to reload user %s: %v", user.Username, err)
				continue
			}

			if err := db.Model(&currentUser).Association("Permissions").Replace(permsToAssign); err != nil {
				log.Printf("❌ Failed to assign permissions to user %s: %v", user.Username, err)
			} else {
				log.Printf("✅ Assigned %d direct permissions to user: %s", len(permsToAssign), user.Username)

				// Verify the assignment
				var assignedPerms []models.Permission
				db.Model(&currentUser).Association("Permissions").Find(&assignedPerms)
				log.Printf("   Verified: User %s now has %d direct permissions", user.Username, len(assignedPerms))
			}
		} else {
			// Clear any existing direct permissions for non-admin users
			var currentUser models.User
			if err := db.First(&currentUser, user.ID).Error; err == nil {
				if err := db.Model(&currentUser).Association("Permissions").Clear(); err != nil {
					log.Printf("⚠️ Failed to clear permissions for user %s: %v", user.Username, err)
				}
			}
		}
	}

	log.Println("✅ User permissions seeded")
}

func ptrString(s string) *string { return &s }
