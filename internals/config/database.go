package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"supply-chain/internals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Warn, // Only log warnings and errors, not every query
				Colorful:      true,
			},
		),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("✅ Database connected")

	// Auto-migrate models
	// RBAC models
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.UserPermission{},
		&models.RolePermission{},
	); err != nil {
		log.Fatal("Failed to migrate RBAC models:", err)
	}

	// Core entities
	if err := DB.AutoMigrate(
		&models.Facility{},
		&models.Pharmacy{},
		&models.Warehouse{},
		&models.EMRIntegration{},
		&models.EMRSyncLog{},
	); err != nil {
		log.Fatal("Failed to migrate core models:", err)
	}
	
	// Procurement and orders
	if err := DB.AutoMigrate(
		&models.ProcurementPlan{},
		&models.ProcurementPlanItem{},
		&models.FacilityOrder{},
		&models.FacilityOrderItem{},
		&models.FacilityDelivery{},
		&models.FacilityDeliveryItem{},
	); err != nil {
		log.Fatal("Failed to migrate order models:", err)
	}
	
	// Stock management
	if err := DB.AutoMigrate(
		&models.StockOnHand{},
		&models.StockDispensed{},
		&models.StockAdjustment{},
		&models.StockTransfer{},
		&models.StockReturn{},
		&models.PharmacyStock{},
		&models.GoodsReceipt{},
	); err != nil {
		log.Fatal("Failed to migrate stock models:", err)
	}
	
	// Legacy and other models
	if err := DB.AutoMigrate(
		&models.PurchaseOrder{},
		&models.HealthCommodityOrder{},
		&models.InterFacilityTransfer{},
		&models.PatientVisit{},
		&models.PatientMetric{},
		&models.ProductAmc{},
		&models.Prescription{},
		&models.WarehouseOrder{},
		&models.WarehouseDelivery{},
	); err != nil {
		log.Fatal("Failed to migrate legacy models:", err)
	}
}
