package config

import (
	"log"
	"time"

	"supply-chain/internals/models"
)

// SeedDatabase populates the database with comprehensive sample data
func SeedDatabase() {
	log.Println("🌱 Seeding database with comprehensive data...")

	// Core entities (must be seeded first)
	seedWarehouses()
	seedFacilities()
	seedPharmacies()
	seedEMRIntegrations()

	// Procurement and orders
	seedProcurementPlans()
	seedFacilityOrders()

	// Stock management
	seedStockOnHand()
	seedPharmacyStock()
	seedStockAdjustments()
	seedStockTransfers()
	seedStockDispensed()
	seedStockReturns()
	seedGoodsReceipts()

	// Other data
	seedPurchaseOrders()
	seedProductAmc()
	seedPatientVisits()

	// Legacy data
	seedWarehouseOrders()
	seedWarehouseDeliveries()

	log.Println("✅ Database seeding completed successfully")
}

// Products and Facilities
var facilities = []string{"HC_KAMPALA_001", "HC_JINJA_002", "HC_MBARARA_003", "HC_GULU_004"}
var products = []struct {
	code string
	name string
}{
	{"PROD_PARACETAMOL_500", "Paracetamol 500mg"},
	{"PROD_IBUPROFEN_400", "Ibuprofen 400mg"},
	{"PROD_AMOXICILLIN_250", "Amoxicillin 250mg"},
	{"PROD_COTRIMOXAZOLE_480", "Cotrimoxazole 480mg"},
	{"PROD_METFORMIN_500", "Metformin 500mg"},
}
var warehouses = []struct {
	code string
	name string
}{
	{"JMS", "Joint Medical Stores"},
	{"NMS", "National Medical Stores"},
}

// Seed core entities
func seedWarehouses() {
	warehouses := []models.Warehouse{
		{
			WarehouseCode: "JMS",
			WarehouseName: "Joint Medical Stores",
			WarehouseType: stringPtr("Private"),
			Address:       stringPtr("Kampala, Uganda"),
			ContactPhone:  stringPtr("+256-XXX-XXXX"),
			ContactEmail:  stringPtr("info@jms.co.ug"),
			IsActive:      true,
			CreatedAt:     time.Now(),
		},
		{
			WarehouseCode: "NMS",
			WarehouseName: "National Medical Stores",
			WarehouseType: stringPtr("National"),
			Address:       stringPtr("Entebbe, Uganda"),
			ContactPhone:  stringPtr("+256-XXX-XXXX"),
			ContactEmail:  stringPtr("info@nms.go.ug"),
			IsActive:      true,
			CreatedAt:     time.Now(),
		},
	}

	for _, warehouse := range warehouses {
		if err := DB.Create(&warehouse).Error; err != nil {
			log.Printf("⚠️ Failed to seed warehouse: %v", err)
		}
	}
	log.Println("✅ Warehouses seeded (2 records)")
}

func seedFacilities() {
	facilities := []models.Facility{
		{
			FacilityCode:  "HC_KAMPALA_001",
			FacilityName:  "Kampala Health Center",
			DHIS2Code:     stringPtr("DHIS2_KLA_001"),
			LevelOfCare:   stringPtr("HCIV"),
			District:      stringPtr("Kampala"),
			Region:        stringPtr("Central"),
			Zone:          stringPtr("Central"),
			Address:       stringPtr("Kampala City"),
			ContactPerson: stringPtr("Dr. John Doe"),
			ContactPhone:  stringPtr("+256-700-000001"),
			ContactEmail:  stringPtr("kampala.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("OPENMRS"),
			EMRSystemName: stringPtr("OpenMRS"),
			CreatedAt:     time.Now(),
		},
		{
			FacilityCode:  "HC_JINJA_002",
			FacilityName:  "Jinja Health Center",
			DHIS2Code:     stringPtr("DHIS2_JNJ_002"),
			LevelOfCare:   stringPtr("HCIII"),
			District:      stringPtr("Jinja"),
			Region:        stringPtr("Eastern"),
			Zone:          stringPtr("Eastern"),
			Address:       stringPtr("Jinja Town"),
			ContactPerson: stringPtr("Dr. Jane Smith"),
			ContactPhone:  stringPtr("+256-700-000002"),
			ContactEmail:  stringPtr("jinja.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("DHIS2"),
			EMRSystemName: stringPtr("DHIS2"),
			CreatedAt:     time.Now(),
		},
		{
			FacilityCode:  "HC_MBARARA_003",
			FacilityName:  "Mbarara Health Center",
			DHIS2Code:     stringPtr("DHIS2_MBR_003"),
			LevelOfCare:   stringPtr("HCIV"),
			District:      stringPtr("Mbarara"),
			Region:        stringPtr("Western"),
			Zone:          stringPtr("Western"),
			Address:       stringPtr("Mbarara Town"),
			ContactPerson: stringPtr("Dr. Peter Okello"),
			ContactPhone:  stringPtr("+256-700-000003"),
			ContactEmail:  stringPtr("mbarara.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("OPENMRS"),
			EMRSystemName: stringPtr("OpenMRS"),
			CreatedAt:     time.Now(),
		},
		{
			FacilityCode:  "HC_GULU_004",
			FacilityName:  "Gulu Health Center",
			DHIS2Code:     stringPtr("DHIS2_GUL_004"),
			LevelOfCare:   stringPtr("HCIII"),
			District:      stringPtr("Gulu"),
			Region:        stringPtr("Northern"),
			Zone:          stringPtr("Northern"),
			Address:       stringPtr("Gulu Town"),
			ContactPerson: stringPtr("Dr. Mary Aceng"),
			ContactPhone:  stringPtr("+256-700-000004"),
			ContactEmail:  stringPtr("gulu.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("OPENMRS"),
			EMRSystemName: stringPtr("OpenMRS"),
			CreatedAt:     time.Now(),
		},
		{
			FacilityCode:  "HC_MASINDI_005",
			FacilityName:  "Masindi Health Center",
			DHIS2Code:     stringPtr("DHIS2_MSD_005"),
			LevelOfCare:   stringPtr("HCIII"),
			District:      stringPtr("Masindi"),
			Region:        stringPtr("Western"),
			Zone:          stringPtr("Western"),
			Address:       stringPtr("Masindi Town"),
			ContactPerson: stringPtr("Dr. James Kigozi"),
			ContactPhone:  stringPtr("+256-700-000005"),
			ContactEmail:  stringPtr("masindi.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("OPENMRS"),
			EMRSystemName: stringPtr("OpenMRS"),
			CreatedAt:     time.Now(),
		},
		{
			FacilityCode:  "HC_LIRA_006",
			FacilityName:  "Lira Health Center",
			DHIS2Code:     stringPtr("DHIS2_LIR_006"),
			LevelOfCare:   stringPtr("HCIV"),
			District:      stringPtr("Lira"),
			Region:        stringPtr("Northern"),
			Zone:          stringPtr("Northern"),
			Address:       stringPtr("Lira Town"),
			ContactPerson: stringPtr("Dr. Sarah Nakato"),
			ContactPhone:  stringPtr("+256-700-000006"),
			ContactEmail:  stringPtr("lira.hc@moh.go.ug"),
			IsActive:      true,
			EMRSystemCode: stringPtr("DHIS2"),
			EMRSystemName: stringPtr("DHIS2"),
			CreatedAt:     time.Now(),
		},
	}

	for _, facility := range facilities {
		if err := DB.Create(&facility).Error; err != nil {
			log.Printf("⚠️ Failed to seed facility: %v", err)
		}
	}
	log.Println("✅ Facilities seeded (6 records)")
}

func seedPharmacies() {
	// Get facilities first
	var facilities []models.Facility
	DB.Find(&facilities)

	if len(facilities) == 0 {
		log.Println("⚠️ No facilities found, skipping pharmacy seeding")
		return
	}

	pharmacies := []models.Pharmacy{
		// Kampala - Multiple pharmacies
		{
			FacilityID:   facilities[0].ID, // Kampala
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[0].ID, // Kampala
			PharmacyCode: "PED",
			PharmacyName: "Pediatric Pharmacy",
			PharmacyType: stringPtr("Pediatric"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[0].ID, // Kampala
			PharmacyCode: "ART",
			PharmacyName: "ART Pharmacy",
			PharmacyType: stringPtr("ART"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[0].ID, // Kampala
			PharmacyCode: "EMERG",
			PharmacyName: "Emergency Pharmacy",
			PharmacyType: stringPtr("Emergency"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		// Jinja
		{
			FacilityID:   facilities[1].ID, // Jinja
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[1].ID, // Jinja
			PharmacyCode: "PED",
			PharmacyName: "Pediatric Pharmacy",
			PharmacyType: stringPtr("Pediatric"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		// Mbarara
		{
			FacilityID:   facilities[2].ID, // Mbarara
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[2].ID, // Mbarara
			PharmacyCode: "ART",
			PharmacyName: "ART Pharmacy",
			PharmacyType: stringPtr("ART"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		// Gulu
		{
			FacilityID:   facilities[3].ID, // Gulu
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		// Masindi
		{
			FacilityID:   facilities[4].ID, // Masindi
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		// Lira
		{
			FacilityID:   facilities[5].ID, // Lira
			PharmacyCode: "MAIN",
			PharmacyName: "Main Pharmacy",
			PharmacyType: stringPtr("Main"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
		{
			FacilityID:   facilities[5].ID, // Lira
			PharmacyCode: "PED",
			PharmacyName: "Pediatric Pharmacy",
			PharmacyType: stringPtr("Pediatric"),
			IsActive:     true,
			CreatedAt:    time.Now(),
		},
	}

	for _, pharmacy := range pharmacies {
		if err := DB.Create(&pharmacy).Error; err != nil {
			log.Printf("⚠️ Failed to seed pharmacy: %v", err)
		}
	}
	log.Println("✅ Pharmacies seeded (12 records)")
}

func seedEMRIntegrations() {
	var facilities []models.Facility
	DB.Find(&facilities)

	if len(facilities) == 0 {
		log.Println("⚠️ No facilities found, skipping EMR integration seeding")
		return
	}

	integrations := []models.EMRIntegration{
		{
			FacilityID:       facilities[0].ID,
			EMRSystemCode:    "OPENMRS",
			EMRSystemName:    "OpenMRS",
			EMRSystemVersion: stringPtr("3.0.0"),
			APIEndpoint:      stringPtr("https://emr-kampala.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("realtime"),
			IsActive:         true,
			IsVerified:       true,
			VerifiedAt:       timePtr(time.Now().AddDate(0, 0, -10)),
			VerifiedBy:       stringPtr("System Admin"),
			LastSyncAt:       timePtr(time.Now().AddDate(0, 0, -1)),
			LastSyncStatus:   stringPtr("success"),
			CreatedAt:        time.Now().AddDate(0, 0, -10),
		},
		{
			FacilityID:       facilities[1].ID,
			EMRSystemCode:    "DHIS2",
			EMRSystemName:    "DHIS2",
			EMRSystemVersion: stringPtr("2.40"),
			APIEndpoint:      stringPtr("https://dhis2-jinja.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("hourly"),
			IsActive:         true,
			IsVerified:       true,
			VerifiedAt:       timePtr(time.Now().AddDate(0, 0, -8)),
			VerifiedBy:       stringPtr("System Admin"),
			LastSyncAt:       timePtr(time.Now().Add(-2 * time.Hour)),
			LastSyncStatus:   stringPtr("success"),
			CreatedAt:        time.Now().AddDate(0, 0, -8),
		},
		{
			FacilityID:       facilities[2].ID,
			EMRSystemCode:    "OPENMRS",
			EMRSystemName:    "OpenMRS",
			EMRSystemVersion: stringPtr("3.0.0"),
			APIEndpoint:      stringPtr("https://emr-mbarara.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("realtime"),
			IsActive:         true,
			IsVerified:       true,
			VerifiedAt:       timePtr(time.Now().AddDate(0, 0, -5)),
			VerifiedBy:       stringPtr("System Admin"),
			LastSyncAt:       timePtr(time.Now().Add(-1 * time.Hour)),
			LastSyncStatus:   stringPtr("success"),
			CreatedAt:        time.Now().AddDate(0, 0, -5),
		},
		{
			FacilityID:       facilities[3].ID,
			EMRSystemCode:    "OPENMRS",
			EMRSystemName:    "OpenMRS",
			EMRSystemVersion: stringPtr("3.0.0"),
			APIEndpoint:      stringPtr("https://emr-gulu.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("daily"),
			IsActive:         true,
			IsVerified:       true,
			VerifiedAt:       timePtr(time.Now().AddDate(0, 0, -3)),
			VerifiedBy:       stringPtr("System Admin"),
			LastSyncAt:       timePtr(time.Now().AddDate(0, 0, -1)),
			LastSyncStatus:   stringPtr("success"),
			CreatedAt:        time.Now().AddDate(0, 0, -3),
		},
		{
			FacilityID:       facilities[4].ID,
			EMRSystemCode:    "OPENMRS",
			EMRSystemName:    "OpenMRS",
			EMRSystemVersion: stringPtr("3.0.0"),
			APIEndpoint:      stringPtr("https://emr-masindi.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("hourly"),
			IsActive:         true,
			IsVerified:       false,
			CreatedAt:        time.Now().AddDate(0, 0, -2),
		},
		{
			FacilityID:       facilities[5].ID,
			EMRSystemCode:    "DHIS2",
			EMRSystemName:    "DHIS2",
			EMRSystemVersion: stringPtr("2.40"),
			APIEndpoint:      stringPtr("https://dhis2-lira.moh.go.ug/api"),
			SyncEnabled:      true,
			SyncFrequency:    stringPtr("realtime"),
			IsActive:         true,
			IsVerified:       true,
			VerifiedAt:       timePtr(time.Now().AddDate(0, 0, -1)),
			VerifiedBy:       stringPtr("System Admin"),
			LastSyncAt:       timePtr(time.Now().Add(-30 * time.Minute)),
			LastSyncStatus:   stringPtr("partial"),
			LastSyncMessage:  stringPtr("Some records failed to sync"),
			CreatedAt:        time.Now().AddDate(0, 0, -1),
		},
	}

	for _, integration := range integrations {
		if err := DB.Create(&integration).Error; err != nil {
			log.Printf("⚠️ Failed to seed EMR integration: %v", err)
			continue
		}

		// Create sync logs for verified integrations
		if integration.IsVerified && integration.LastSyncAt != nil {
			syncLog := models.EMRSyncLog{
				IntegrationID:     integration.ID,
				SyncType:          "stock",
				SyncDirection:     "emr_to_central",
				Status:            *integration.LastSyncStatus,
				RecordsProcessed:  100,
				RecordsSuccessful: 95,
				RecordsFailed:     5,
				StartedAt:         *integration.LastSyncAt,
				CompletedAt:       timePtr(integration.LastSyncAt.Add(time.Minute * 5)),
				DurationSeconds:   intPtr(300),
			}
			if *integration.LastSyncStatus == "partial" {
				syncLog.ErrorMessage = stringPtr("Some records failed validation")
			}
			DB.Create(&syncLog)
		}
	}
	log.Println("✅ EMR integrations seeded (6 records with sync logs)")
}

func seedFacilityOrders() {
	var facilities []models.Facility
	var warehouses []models.Warehouse
	DB.Find(&facilities)
	DB.Find(&warehouses)

	if len(facilities) == 0 || len(warehouses) == 0 {
		log.Println("⚠️ No facilities or warehouses found, skipping facility order seeding")
		return
	}

	orders := []models.FacilityOrder{
		// Kampala - JMS - Fulfilled
		{
			OrderNumber:        "ORD-KLA-JMS-001",
			OrderRefNumber:     stringPtr("PO_2025_001"),
			FacilityID:         facilities[0].ID,
			FacilityCode:       facilities[0].FacilityCode,
			WarehouseID:        warehouses[0].ID, // JMS
			WarehouseCode:      warehouses[0].WarehouseCode,
			OrderDate:          time.Now().AddDate(0, 0, -10),
			OrderType:          stringPtr("routine"),
			OrderStatus:        "fulfilled",
			Priority:           stringPtr("normal"),
			FinancialYear:      stringPtr("2024/2025"),
			OrderCycle:         stringPtr("Q1"),
			SubmittedBy:        stringPtr("Dr. John Doe"),
			SubmittedAt:        timePtr(time.Now().AddDate(0, 0, -10)),
			ApprovedBy:         stringPtr("Pharmacy Manager"),
			ApprovedAt:         timePtr(time.Now().AddDate(0, 0, -9)),
			ActualDeliveryDate: timePtr(time.Now().AddDate(0, 0, -7)),
			TotalItems:         3,
			TotalQuantity:      5000,
			SourceSystem:       stringPtr("OPENMRS"),
			CreatedAt:          time.Now().AddDate(0, 0, -10),
		},
		// Jinja - NMS - Processing
		{
			OrderNumber:    "ORD-JNJ-NMS-001",
			OrderRefNumber: stringPtr("PO_2025_002"),
			FacilityID:     facilities[1].ID,
			FacilityCode:   facilities[1].FacilityCode,
			WarehouseID:    warehouses[1].ID, // NMS
			WarehouseCode:  warehouses[1].WarehouseCode,
			OrderDate:      time.Now().AddDate(0, 0, -5),
			OrderType:      stringPtr("routine"),
			OrderStatus:    "processing",
			Priority:       stringPtr("normal"),
			FinancialYear:  stringPtr("2024/2025"),
			OrderCycle:     stringPtr("Q1"),
			SubmittedBy:    stringPtr("Dr. Jane Smith"),
			SubmittedAt:    timePtr(time.Now().AddDate(0, 0, -5)),
			ApprovedBy:     stringPtr("Pharmacy Manager"),
			ApprovedAt:     timePtr(time.Now().AddDate(0, 0, -4)),
			TotalItems:     2,
			TotalQuantity:  3000,
			SourceSystem:   stringPtr("DHIS2"),
			CreatedAt:      time.Now().AddDate(0, 0, -5),
		},
		// Mbarara - JMS - Approved
		{
			OrderNumber:          "ORD-MBR-JMS-001",
			OrderRefNumber:       stringPtr("PO_2025_003"),
			FacilityID:           facilities[2].ID,
			FacilityCode:         facilities[2].FacilityCode,
			WarehouseID:          warehouses[0].ID, // JMS
			WarehouseCode:        warehouses[0].WarehouseCode,
			OrderDate:            time.Now().AddDate(0, 0, -3),
			OrderType:            stringPtr("routine"),
			OrderStatus:          "approved",
			Priority:             stringPtr("normal"),
			FinancialYear:        stringPtr("2024/2025"),
			OrderCycle:           stringPtr("Q1"),
			SubmittedBy:          stringPtr("Dr. Peter Okello"),
			SubmittedAt:          timePtr(time.Now().AddDate(0, 0, -3)),
			ApprovedBy:           stringPtr("Pharmacy Manager"),
			ApprovedAt:           timePtr(time.Now().AddDate(0, 0, -2)),
			ExpectedDeliveryDate: timePtr(time.Now().AddDate(0, 0, 7)),
			TotalItems:           4,
			TotalQuantity:        6000,
			SourceSystem:         stringPtr("OPENMRS"),
			CreatedAt:            time.Now().AddDate(0, 0, -3),
		},
		// Gulu - NMS - Pending
		{
			OrderNumber:    "ORD-GUL-NMS-001",
			OrderRefNumber: stringPtr("PO_2025_004"),
			FacilityID:     facilities[3].ID,
			FacilityCode:   facilities[3].FacilityCode,
			WarehouseID:    warehouses[1].ID, // NMS
			WarehouseCode:  warehouses[1].WarehouseCode,
			OrderDate:      time.Now().AddDate(0, 0, -1),
			OrderType:      stringPtr("emergency"),
			OrderStatus:    "pending",
			Priority:       stringPtr("urgent"),
			FinancialYear:  stringPtr("2024/2025"),
			OrderCycle:     stringPtr("Q1"),
			TotalItems:     2,
			TotalQuantity:  1500,
			SourceSystem:   stringPtr("OPENMRS"),
			CreatedAt:      time.Now().AddDate(0, 0, -1),
		},
		// Lira - JMS - Submitted
		{
			OrderNumber:    "ORD-LIR-JMS-001",
			OrderRefNumber: stringPtr("PO_2025_005"),
			FacilityID:     facilities[5].ID,
			FacilityCode:   facilities[5].FacilityCode,
			WarehouseID:    warehouses[0].ID, // JMS
			WarehouseCode:  warehouses[0].WarehouseCode,
			OrderDate:      time.Now().AddDate(0, 0, -2),
			OrderType:      stringPtr("routine"),
			OrderStatus:    "submitted",
			Priority:       stringPtr("normal"),
			FinancialYear:  stringPtr("2024/2025"),
			OrderCycle:     stringPtr("Q1"),
			SubmittedBy:    stringPtr("Dr. Sarah Nakato"),
			SubmittedAt:    timePtr(time.Now().AddDate(0, 0, -2)),
			TotalItems:     3,
			TotalQuantity:  4000,
			SourceSystem:   stringPtr("DHIS2"),
			CreatedAt:      time.Now().AddDate(0, 0, -2),
		},
	}

	// Product codes for order items
	productItems := map[int][]struct {
		ProductCode string
		Description string
		Quantity    int
		UOM         string
	}{
		0: { // Order 1
			{"PROD_PARACETAMOL_500", "Paracetamol 500mg", 2000, "Tablets"},
			{"PROD_AMOXICILLIN_250", "Amoxicillin 250mg", 2000, "Capsules"},
			{"PROD_IBUPROFEN_400", "Ibuprofen 400mg", 1000, "Tablets"},
		},
		1: { // Order 2
			{"PROD_AMOXICILLIN_250", "Amoxicillin 250mg", 2000, "Capsules"},
			{"PROD_COTRIMOXAZOLE_480", "Cotrimoxazole 480mg", 1000, "Tablets"},
		},
		2: { // Order 3
			{"PROD_PARACETAMOL_500", "Paracetamol 500mg", 2000, "Tablets"},
			{"PROD_METFORMIN_500", "Metformin 500mg", 1500, "Tablets"},
			{"PROD_IBUPROFEN_400", "Ibuprofen 400mg", 1500, "Tablets"},
			{"PROD_COTRIMOXAZOLE_480", "Cotrimoxazole 480mg", 1000, "Tablets"},
		},
		3: { // Order 4
			{"PROD_PARACETAMOL_500", "Paracetamol 500mg", 1000, "Tablets"},
			{"PROD_AMOXICILLIN_250", "Amoxicillin 250mg", 500, "Capsules"},
		},
		4: { // Order 5
			{"PROD_PARACETAMOL_500", "Paracetamol 500mg", 1500, "Tablets"},
			{"PROD_IBUPROFEN_400", "Ibuprofen 400mg", 1500, "Tablets"},
			{"PROD_METFORMIN_500", "Metformin 500mg", 1000, "Tablets"},
		},
	}

	for idx, order := range orders {
		if err := DB.Create(&order).Error; err != nil {
			log.Printf("⚠️ Failed to seed facility order: %v", err)
			continue
		}

		// Add order items
		if items, ok := productItems[idx]; ok {
			for _, itemData := range items {
				item := models.FacilityOrderItem{
					OrderID:            order.ID,
					ProductCode:        itemData.ProductCode,
					ProductDescription: stringPtr(itemData.Description),
					UOM:                stringPtr(itemData.UOM),
					OrderedQuantity:    itemData.Quantity,
					Status:             "pending",
					CreatedAt:          time.Now(),
				}

				// Set honored quantity for fulfilled orders
				if order.OrderStatus == "fulfilled" {
					item.HonoredQuantity = intPtr(itemData.Quantity)
					item.DeliveredQuantity = intPtr(itemData.Quantity)
					item.Status = "fulfilled"
				} else if order.OrderStatus == "processing" {
					item.HonoredQuantity = intPtr(itemData.Quantity)
					item.Status = "partially_fulfilled"
				}

				if err := DB.Create(&item).Error; err != nil {
					log.Printf("⚠️ Failed to seed facility order item: %v", err)
				}
			}
		}
	}
	log.Println("✅ Facility orders seeded (5 orders with items)")
}

func seedStockTransfers() {
	var facilities []models.Facility
	var pharmacies []models.Pharmacy
	DB.Find(&facilities)
	DB.Find(&pharmacies)

	if len(facilities) < 2 {
		log.Println("⚠️ Insufficient facilities found, skipping stock transfer seeding")
		return
	}

	transfers := []models.StockTransfer{
		// Inter-facility transfer - Completed
		{
			TransferRef:    "TRF-INT-001",
			TransferType:   "inter_facility",
			FromFacilityID: facilities[0].ID, // Kampala
			ToFacilityID:   facilities[1].ID, // Jinja
			ProductCode:    "PROD_PARACETAMOL_500",
			BatchNumber:    stringPtr("BATCH_2025_001"),
			Quantity:       500,
			ExpiryDate:     timePtr(time.Now().AddDate(1, 0, 0)),
			TransferDate:   time.Now().AddDate(0, 0, -3),
			Status:         "completed",
			RequestedBy:    stringPtr("Dr. John Doe"),
			ApprovedBy:     stringPtr("Pharmacy Manager"),
			ReceivedBy:     stringPtr("Dr. Jane Smith"),
			ReceivedAt:     timePtr(time.Now().AddDate(0, 0, -2)),
			Notes:          stringPtr("Emergency stock transfer"),
			CreatedAt:      time.Now().AddDate(0, 0, -3),
		},
		// Intra-facility transfer - Completed
		{
			TransferRef:    "TRF-INTRA-001",
			TransferType:   "intra_facility",
			FromFacilityID: facilities[0].ID,  // Kampala
			FromPharmacyID: &pharmacies[0].ID, // Main Pharmacy
			ToFacilityID:   facilities[0].ID,  // Kampala (same facility)
			ToPharmacyID:   &pharmacies[1].ID, // Pediatric Pharmacy
			ProductCode:    "PROD_IBUPROFEN_400",
			BatchNumber:    stringPtr("BATCH_2025_002"),
			Quantity:       200,
			ExpiryDate:     timePtr(time.Now().AddDate(1, 6, 0)),
			TransferDate:   time.Now().AddDate(0, 0, -1),
			Status:         "completed",
			RequestedBy:    stringPtr("Pediatric Pharmacy Manager"),
			ApprovedBy:     stringPtr("Main Pharmacy Manager"),
			ReceivedBy:     stringPtr("Pediatric Pharmacy Staff"),
			ReceivedAt:     timePtr(time.Now()),
			Notes:          stringPtr("Transfer to pediatric pharmacy for pediatric patients"),
			CreatedAt:      time.Now().AddDate(0, 0, -1),
		},
		// Inter-facility transfer - In Transit
		{
			TransferRef:    "TRF-INT-002",
			TransferType:   "inter_facility",
			FromFacilityID: facilities[2].ID, // Mbarara
			ToFacilityID:   facilities[3].ID, // Gulu
			ProductCode:    "PROD_AMOXICILLIN_250",
			BatchNumber:    stringPtr("BATCH_2025_003"),
			Quantity:       750,
			ExpiryDate:     timePtr(time.Now().AddDate(2, 0, 0)),
			TransferDate:   time.Now().AddDate(0, 0, -2),
			Status:         "in_transit",
			RequestedBy:    stringPtr("Dr. Peter Okello"),
			ApprovedBy:     stringPtr("Pharmacy Manager"),
			Notes:          stringPtr("Routine stock redistribution"),
			CreatedAt:      time.Now().AddDate(0, 0, -2),
		},
		// Intra-facility transfer - Pending
		{
			TransferRef:    "TRF-INTRA-002",
			TransferType:   "intra_facility",
			FromFacilityID: facilities[0].ID,  // Kampala
			FromPharmacyID: &pharmacies[0].ID, // Main Pharmacy
			ToFacilityID:   facilities[0].ID,  // Kampala
			ToPharmacyID:   &pharmacies[2].ID, // ART Pharmacy
			ProductCode:    "PROD_COTRIMOXAZOLE_480",
			BatchNumber:    stringPtr("BATCH_2025_004"),
			Quantity:       300,
			ExpiryDate:     timePtr(time.Now().AddDate(1, 3, 0)),
			TransferDate:   time.Now(),
			Status:         "pending",
			RequestedBy:    stringPtr("ART Pharmacy Manager"),
			Notes:          stringPtr("Request for ART pharmacy stock replenishment"),
			CreatedAt:      time.Now(),
		},
		// Inter-facility transfer - Pending
		{
			TransferRef:    "TRF-INT-003",
			TransferType:   "inter_facility",
			FromFacilityID: facilities[5].ID, // Lira
			ToFacilityID:   facilities[4].ID, // Masindi
			ProductCode:    "PROD_METFORMIN_500",
			BatchNumber:    stringPtr("BATCH_2025_005"),
			Quantity:       400,
			ExpiryDate:     timePtr(time.Now().AddDate(1, 9, 0)),
			TransferDate:   time.Now(),
			Status:         "pending",
			RequestedBy:    stringPtr("Dr. Sarah Nakato"),
			Notes:          stringPtr("Request for emergency stock"),
			CreatedAt:      time.Now(),
		},
		// Intra-facility transfer - In Transit
		{
			TransferRef:    "TRF-INTRA-003",
			TransferType:   "intra_facility",
			FromFacilityID: facilities[1].ID,  // Jinja
			FromPharmacyID: &pharmacies[4].ID, // Main Pharmacy
			ToFacilityID:   facilities[1].ID,  // Jinja
			ToPharmacyID:   &pharmacies[5].ID, // Pediatric Pharmacy
			ProductCode:    "PROD_PARACETAMOL_500",
			BatchNumber:    stringPtr("BATCH_2025_001"),
			Quantity:       150,
			ExpiryDate:     timePtr(time.Now().AddDate(1, 0, 0)),
			TransferDate:   time.Now().AddDate(0, 0, -1),
			Status:         "in_transit",
			RequestedBy:    stringPtr("Pediatric Pharmacy Staff"),
			ApprovedBy:     stringPtr("Main Pharmacy Manager"),
			Notes:          stringPtr("Transfer for pediatric patients"),
			CreatedAt:      time.Now().AddDate(0, 0, -1),
		},
	}

	for _, transfer := range transfers {
		if err := DB.Create(&transfer).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock transfer: %v", err)
		}
	}
	log.Println("✅ Stock transfers seeded (6 records: 3 inter-facility, 3 intra-facility)")
}

func seedStockOnHand() {
	var facilities []models.Facility
	DB.Find(&facilities)

	if len(facilities) == 0 {
		log.Println("⚠️ No facilities found, skipping stock on hand seeding")
		return
	}

	stocks := []models.StockOnHand{
		// Kampala
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[0].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_PARACETAMOL_500",
			SrcBatchNumber:  "BATCH_2025_001",
			SrcQuantity:     500,
			SrcExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[0].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_IBUPROFEN_400",
			SrcBatchNumber:  "BATCH_2025_002",
			SrcQuantity:     300,
			SrcExpiryDate:   time.Now().AddDate(1, 6, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[0].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_AMOXICILLIN_250",
			SrcBatchNumber:  "BATCH_2025_003",
			SrcQuantity:     1200,
			SrcExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		// Jinja
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[1].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_AMOXICILLIN_250",
			SrcBatchNumber:  "BATCH_2025_003",
			SrcQuantity:     1000,
			SrcExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[1].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_PARACETAMOL_500",
			SrcBatchNumber:  "BATCH_2025_001",
			SrcQuantity:     800,
			SrcExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[1].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_COTRIMOXAZOLE_480",
			SrcBatchNumber:  "BATCH_2025_004",
			SrcQuantity:     600,
			SrcExpiryDate:   time.Now().AddDate(1, 3, 0),
		},
		// Mbarara
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[2].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_COTRIMOXAZOLE_480",
			SrcBatchNumber:  "BATCH_2025_004",
			SrcQuantity:     750,
			SrcExpiryDate:   time.Now().AddDate(1, 3, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[2].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_METFORMIN_500",
			SrcBatchNumber:  "BATCH_2025_005",
			SrcQuantity:     1500,
			SrcExpiryDate:   time.Now().AddDate(1, 9, 0),
		},
		// Gulu
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[3].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_METFORMIN_500",
			SrcBatchNumber:  "BATCH_2025_005",
			SrcQuantity:     2000,
			SrcExpiryDate:   time.Now().AddDate(1, 9, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[3].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_PARACETAMOL_500",
			SrcBatchNumber:  "BATCH_2025_001",
			SrcQuantity:     600,
			SrcExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		// Masindi
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[4].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_IBUPROFEN_400",
			SrcBatchNumber:  "BATCH_2025_002",
			SrcQuantity:     400,
			SrcExpiryDate:   time.Now().AddDate(1, 6, 0),
		},
		// Lira
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[5].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_AMOXICILLIN_250",
			SrcBatchNumber:  "BATCH_2025_003",
			SrcQuantity:     1800,
			SrcExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: facilities[5].FacilityCode,
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_COTRIMOXAZOLE_480",
			SrcBatchNumber:  "BATCH_2025_004",
			SrcQuantity:     900,
			SrcExpiryDate:   time.Now().AddDate(1, 8, 0),
		},
	}

	for _, stock := range stocks {
		if err := DB.Create(&stock).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock on hand: %v", err)
		}
	}
	log.Println("✅ Stock on hand data seeded (12 records)")
}

func seedPurchaseOrders() {
	orders := []models.PurchaseOrder{
		{
			OrdSystemCode:      "HEALTHSYS-01",
			OrdFacilityCode:    "HC_KAMPALA_001",
			OrdTimestamp:       time.Now(),
			OrdOrderDate:       time.Now(),
			OrdOrderRefNumber:  "PO_REF_2025_001",
			OrdOrderNumber:     "PO_2025_001",
			OrdProductCode:     "PROD_PARACETAMOL_500",
			OrdOrderedQuantity: 2000,
		},
		{
			OrdSystemCode:      "HEALTHSYS-01",
			OrdFacilityCode:    "HC_JINJA_002",
			OrdTimestamp:       time.Now(),
			OrdOrderDate:       time.Now().AddDate(0, 0, -5),
			OrdOrderRefNumber:  "PO_REF_2025_002",
			OrdOrderNumber:     "PO_2025_002",
			OrdProductCode:     "PROD_AMOXICILLIN_250",
			OrdOrderedQuantity: 5000,
		},
		{
			OrdSystemCode:      "HEALTHSYS-01",
			OrdFacilityCode:    "HC_KAMPALA_001",
			OrdTimestamp:       time.Now(),
			OrdOrderDate:       time.Now().AddDate(0, 0, -10),
			OrdOrderRefNumber:  "PO_REF_2025_003",
			OrdOrderNumber:     "PO_2025_003",
			OrdProductCode:     "PROD_IBUPROFEN_400",
			OrdOrderedQuantity: 1500,
		},
		{
			OrdSystemCode:      "HEALTHSYS-01",
			OrdFacilityCode:    "HC_MBARARA_003",
			OrdTimestamp:       time.Now(),
			OrdOrderDate:       time.Now().AddDate(0, 0, -15),
			OrdOrderRefNumber:  "PO_REF_2025_004",
			OrdOrderNumber:     "PO_2025_004",
			OrdProductCode:     "PROD_COTRIMOXAZOLE_480",
			OrdOrderedQuantity: 3000,
		},
		{
			OrdSystemCode:      "HEALTHSYS-01",
			OrdFacilityCode:    "HC_GULU_004",
			OrdTimestamp:       time.Now(),
			OrdOrderDate:       time.Now().AddDate(0, 0, -3),
			OrdOrderRefNumber:  "PO_REF_2025_005",
			OrdOrderNumber:     "PO_2025_005",
			OrdProductCode:     "PROD_METFORMIN_500",
			OrdOrderedQuantity: 4000,
		},
	}

	for _, order := range orders {
		if err := DB.Create(&order).Error; err != nil {
			log.Printf("⚠️ Failed to seed purchase order: %v", err)
		}
	}
	log.Println("✅ Purchase order data seeded (5 records)")
}

func seedProductAmc() {
	amcData := []models.ProductAmc{
		{
			AmcSystemCode:   "HEALTHSYS-01",
			AmcFacilityCode: "HC_KAMPALA_001",
			AmcTimestamp:    time.Now(),
			AmcProductCode:  "PROD_PARACETAMOL_500",
			AmcProductName:  "Paracetamol 500mg",
			AmcDate:         time.Now(),
			AmcMonth:        1,
			AmcYear:         2025,
			AmcValue:        450.5,
		},
		{
			AmcSystemCode:   "HEALTHSYS-01",
			AmcFacilityCode: "HC_JINJA_002",
			AmcTimestamp:    time.Now(),
			AmcProductCode:  "PROD_AMOXICILLIN_250",
			AmcProductName:  "Amoxicillin 250mg",
			AmcDate:         time.Now(),
			AmcMonth:        1,
			AmcYear:         2025,
			AmcValue:        890.25,
		},
		{
			AmcSystemCode:   "HEALTHSYS-01",
			AmcFacilityCode: "HC_KAMPALA_001",
			AmcTimestamp:    time.Now(),
			AmcProductCode:  "PROD_IBUPROFEN_400",
			AmcProductName:  "Ibuprofen 400mg",
			AmcDate:         time.Now(),
			AmcMonth:        1,
			AmcYear:         2025,
			AmcValue:        320.75,
		},
		{
			AmcSystemCode:   "HEALTHSYS-01",
			AmcFacilityCode: "HC_MBARARA_003",
			AmcTimestamp:    time.Now(),
			AmcProductCode:  "PROD_COTRIMOXAZOLE_480",
			AmcProductName:  "Cotrimoxazole 480mg",
			AmcDate:         time.Now(),
			AmcMonth:        1,
			AmcYear:         2025,
			AmcValue:        610.0,
		},
		{
			AmcSystemCode:   "HEALTHSYS-01",
			AmcFacilityCode: "HC_GULU_004",
			AmcTimestamp:    time.Now(),
			AmcProductCode:  "PROD_METFORMIN_500",
			AmcProductName:  "Metformin 500mg",
			AmcDate:         time.Now(),
			AmcMonth:        1,
			AmcYear:         2025,
			AmcValue:        725.5,
		},
	}

	for _, amc := range amcData {
		if err := DB.Create(&amc).Error; err != nil {
			log.Printf("⚠️ Failed to seed product AMC: %v", err)
		}
	}
	log.Println("✅ Product AMC data seeded (5 records)")
}

func seedGoodsReceipts() {
	receipts := []models.GoodsReceipt{
		{
			GrnSystemCode:            "HEALTHSYS-01",
			GrnFacilityCode:          "HC_KAMPALA_001",
			GrnTimestamp:             time.Now(),
			GrnReceiptDate:           time.Now().AddDate(0, 0, -3),
			GrnFacilityReceiptNumber: "GRN_KLA_001",
			GrnWarehouseRefNumber:    "WH_JMS_001",
			GrnOrderNumber:           "PO_2025_001",
			GrnProductCode:           "PROD_PARACETAMOL_500",
			GrnBatchNumber:           "BATCH_2025_001",
			GrnQuantity:              2000,
			GrnExpiryDate:            time.Now().AddDate(1, 6, 0),
			GrnSupplierCode:          "SUPP_001",
		},
		{
			GrnSystemCode:            "HEALTHSYS-01",
			GrnFacilityCode:          "HC_JINJA_002",
			GrnTimestamp:             time.Now(),
			GrnReceiptDate:           time.Now().AddDate(0, 0, -2),
			GrnFacilityReceiptNumber: "GRN_JNJ_002",
			GrnWarehouseRefNumber:    "WH_NMS_001",
			GrnOrderNumber:           "PO_2025_002",
			GrnProductCode:           "PROD_AMOXICILLIN_250",
			GrnBatchNumber:           "BATCH_2025_003",
			GrnQuantity:              5000,
			GrnExpiryDate:            time.Now().AddDate(2, 0, 0),
			GrnSupplierCode:          "SUPP_002",
		},
		{
			GrnSystemCode:            "HEALTHSYS-01",
			GrnFacilityCode:          "HC_MBARARA_003",
			GrnTimestamp:             time.Now(),
			GrnReceiptDate:           time.Now().AddDate(0, 0, -1),
			GrnFacilityReceiptNumber: "GRN_MBR_003",
			GrnWarehouseRefNumber:    "WH_JMS_002",
			GrnOrderNumber:           "PO_2025_004",
			GrnProductCode:           "PROD_COTRIMOXAZOLE_480",
			GrnBatchNumber:           "BATCH_2025_004",
			GrnQuantity:              3000,
			GrnExpiryDate:            time.Now().AddDate(1, 8, 0),
			GrnSupplierCode:          "SUPP_003",
		},
	}

	for _, receipt := range receipts {
		if err := DB.Create(&receipt).Error; err != nil {
			log.Printf("⚠️ Failed to seed goods receipt: %v", err)
		}
	}
	log.Println("✅ Goods receipt data seeded (3 records)")
}

func seedPatientVisits() {
	visits := []models.PatientVisit{
		{
			VstSystemCode:      "HEALTHSYS-01",
			VstFacilityCode:    "HC_KAMPALA_001",
			VstTimestamp:       time.Now(),
			VstPatientCode:     "PAT_KLA_001",
			VstSex:             "M",
			VstAge:             45,
			VstVisitDate:       time.Now(),
			VstProductCode:     "PROD_PARACETAMOL_500",
			VstBatchNumber:     "BATCH_2025_001",
			VstQuantity:        2.0,
			VstRegimenCode:     "REG_FEVER_01",
			VstPatientCategory: "Adult",
		},
		{
			VstSystemCode:      "HEALTHSYS-01",
			VstFacilityCode:    "HC_JINJA_002",
			VstTimestamp:       time.Now(),
			VstPatientCode:     "PAT_JNJ_002",
			VstSex:             "F",
			VstAge:             32,
			VstVisitDate:       time.Now().AddDate(0, 0, -1),
			VstProductCode:     "PROD_AMOXICILLIN_250",
			VstBatchNumber:     "BATCH_2025_003",
			VstQuantity:        3.5,
			VstRegimenCode:     "REG_INFECT_02",
			VstPatientCategory: "Adult",
		},
		{
			VstSystemCode:      "HEALTHSYS-01",
			VstFacilityCode:    "HC_KAMPALA_001",
			VstTimestamp:       time.Now(),
			VstPatientCode:     "PAT_KLA_003",
			VstSex:             "M",
			VstAge:             8,
			VstVisitDate:       time.Now().AddDate(0, 0, -2),
			VstProductCode:     "PROD_IBUPROFEN_400",
			VstBatchNumber:     "BATCH_2025_002",
			VstQuantity:        1.0,
			VstRegimenCode:     "REG_PAIN_03",
			VstPatientCategory: "Pediatric",
		},
	}

	for _, visit := range visits {
		if err := DB.Create(&visit).Error; err != nil {
			log.Printf("⚠️ Failed to seed patient visit: %v", err)
		}
	}
	log.Println("✅ Patient visit data seeded (3 records)")
}

func seedPharmacyStock() {
	var facilities []models.Facility
	DB.Find(&facilities)

	if len(facilities) == 0 {
		log.Println("⚠️ No facilities found, skipping pharmacy stock seeding")
		return
	}

	stocks := []models.PharmacyStock{
		// Kampala
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[0].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_PARACETAMOL_500",
			PhaBatchNumber:  "BATCH_2025_001",
			PhaQuantity:     450,
			PhaExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[0].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_IBUPROFEN_400",
			PhaBatchNumber:  "BATCH_2025_002",
			PhaQuantity:     280,
			PhaExpiryDate:   time.Now().AddDate(1, 6, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[0].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_AMOXICILLIN_250",
			PhaBatchNumber:  "BATCH_2025_003",
			PhaQuantity:     1200,
			PhaExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		// Jinja
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[1].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_AMOXICILLIN_250",
			PhaBatchNumber:  "BATCH_2025_003",
			PhaQuantity:     4500,
			PhaExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[1].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_PARACETAMOL_500",
			PhaBatchNumber:  "BATCH_2025_001",
			PhaQuantity:     800,
			PhaExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		// Mbarara
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[2].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_COTRIMOXAZOLE_480",
			PhaBatchNumber:  "BATCH_2025_004",
			PhaQuantity:     2800,
			PhaExpiryDate:   time.Now().AddDate(1, 8, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[2].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_METFORMIN_500",
			PhaBatchNumber:  "BATCH_2025_005",
			PhaQuantity:     1500,
			PhaExpiryDate:   time.Now().AddDate(1, 9, 0),
		},
		// Gulu
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[3].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_METFORMIN_500",
			PhaBatchNumber:  "BATCH_2025_005",
			PhaQuantity:     2000,
			PhaExpiryDate:   time.Now().AddDate(1, 9, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[3].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_PARACETAMOL_500",
			PhaBatchNumber:  "BATCH_2025_001",
			PhaQuantity:     600,
			PhaExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		// Masindi
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[4].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_IBUPROFEN_400",
			PhaBatchNumber:  "BATCH_2025_002",
			PhaQuantity:     400,
			PhaExpiryDate:   time.Now().AddDate(1, 6, 0),
		},
		// Lira
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[5].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_AMOXICILLIN_250",
			PhaBatchNumber:  "BATCH_2025_003",
			PhaQuantity:     1800,
			PhaExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: facilities[5].FacilityCode,
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_COTRIMOXAZOLE_480",
			PhaBatchNumber:  "BATCH_2025_004",
			PhaQuantity:     900,
			PhaExpiryDate:   time.Now().AddDate(1, 8, 0),
		},
	}

	for _, stock := range stocks {
		if err := DB.Create(&stock).Error; err != nil {
			log.Printf("⚠️ Failed to seed pharmacy stock: %v", err)
		}
	}
	log.Println("✅ Pharmacy stock data seeded (12 records)")
}

func seedProcurementPlans() {
	var warehouses []models.Warehouse
	DB.Find(&warehouses)

	if len(warehouses) == 0 {
		log.Println("⚠️ No warehouses found, skipping procurement plan seeding")
		return
	}

	plans := []models.ProcurementPlan{
		{
			PlanSystemCode:  "HEALTHSYS-01",
			WarehouseID:     &warehouses[0].ID, // JMS
			StoreCode:       warehouses[0].WarehouseCode,
			CreatedAt:       time.Now(),
			Notes:           stringPtr("Q1 2025 procurement plan for JMS"),
			FinancialYear:   "2024/2025",
			PlanPeriodType:  stringPtr("Quarterly"),
			PlanPeriodStart: timePtr(time.Now()),
			PlanPeriodEnd:   timePtr(time.Now().AddDate(0, 3, 0)),
			ApprovalStatus:  stringPtr("approved"),
		},
		{
			PlanSystemCode:  "HEALTHSYS-01",
			WarehouseID:     &warehouses[1].ID, // NMS
			StoreCode:       warehouses[1].WarehouseCode,
			CreatedAt:       time.Now(),
			Notes:           stringPtr("Q1 2025 procurement plan for NMS"),
			FinancialYear:   "2024/2025",
			PlanPeriodType:  stringPtr("Quarterly"),
			PlanPeriodStart: timePtr(time.Now()),
			PlanPeriodEnd:   timePtr(time.Now().AddDate(0, 3, 0)),
			ApprovalStatus:  stringPtr("approved"),
		},
		{
			PlanSystemCode:  "HEALTHSYS-01",
			WarehouseID:     &warehouses[0].ID, // JMS - Q2
			StoreCode:       warehouses[0].WarehouseCode,
			CreatedAt:       time.Now().AddDate(0, -1, 0),
			Notes:           stringPtr("Q2 2025 procurement plan for JMS"),
			FinancialYear:   "2024/2025",
			PlanPeriodType:  stringPtr("Quarterly"),
			PlanPeriodStart: timePtr(time.Now().AddDate(0, 3, 0)),
			PlanPeriodEnd:   timePtr(time.Now().AddDate(0, 6, 0)),
			ApprovalStatus:  stringPtr("draft"),
		},
	}

	planItems := map[int][]models.ProcurementPlanItem{
		0: { // JMS Q1
			{
				ProductCode:        "PROD_PARACETAMOL_500",
				ProductDescription: stringPtr("Paracetamol 500mg"),
				Quantity:           5000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "planned",
				UOM:                stringPtr("Tablets"),
			},
			{
				ProductCode:        "PROD_AMOXICILLIN_250",
				ProductDescription: stringPtr("Amoxicillin 250mg"),
				Quantity:           8000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "ordered",
				UOM:                stringPtr("Capsules"),
			},
			{
				ProductCode:        "PROD_IBUPROFEN_400",
				ProductDescription: stringPtr("Ibuprofen 400mg"),
				Quantity:           3000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "planned",
				UOM:                stringPtr("Tablets"),
			},
		},
		1: { // NMS Q1
			{
				ProductCode:        "PROD_PARACETAMOL_500",
				ProductDescription: stringPtr("Paracetamol 500mg"),
				Quantity:           6000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "ordered",
				UOM:                stringPtr("Tablets"),
			},
			{
				ProductCode:        "PROD_AMOXICILLIN_250",
				ProductDescription: stringPtr("Amoxicillin 250mg"),
				Quantity:           10000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "ordered",
				UOM:                stringPtr("Capsules"),
			},
			{
				ProductCode:        "PROD_COTRIMOXAZOLE_480",
				ProductDescription: stringPtr("Cotrimoxazole 480mg"),
				Quantity:           4000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "planned",
				UOM:                stringPtr("Tablets"),
			},
			{
				ProductCode:        "PROD_METFORMIN_500",
				ProductDescription: stringPtr("Metformin 500mg"),
				Quantity:           5000,
				NeededBy:           time.Now().AddDate(0, 1, 0),
				Status:             "planned",
				UOM:                stringPtr("Tablets"),
			},
		},
		2: { // JMS Q2
			{
				ProductCode:        "PROD_PARACETAMOL_500",
				ProductDescription: stringPtr("Paracetamol 500mg"),
				Quantity:           4500,
				NeededBy:           time.Now().AddDate(0, 4, 0),
				Status:             "planned",
				UOM:                stringPtr("Tablets"),
			},
			{
				ProductCode:        "PROD_AMOXICILLIN_250",
				ProductDescription: stringPtr("Amoxicillin 250mg"),
				Quantity:           7000,
				NeededBy:           time.Now().AddDate(0, 4, 0),
				Status:             "planned",
				UOM:                stringPtr("Capsules"),
			},
		},
	}

	for idx, plan := range plans {
		if err := DB.Create(&plan).Error; err != nil {
			log.Printf("⚠️ Failed to seed procurement plan: %v", err)
			continue
		}

		// Add plan items
		if items, ok := planItems[idx]; ok {
			for _, item := range items {
				item.ProcurementID = plan.ID
				if err := DB.Create(&item).Error; err != nil {
					log.Printf("⚠️ Failed to seed procurement plan item: %v", err)
				}
			}
		}
	}
	log.Println("✅ Procurement plan data seeded (3 plans with 9 items)")
}

func seedStockAdjustments() {
	var facilities []models.Facility
	var pharmacies []models.Pharmacy
	DB.Find(&facilities)
	DB.Find(&pharmacies)

	if len(facilities) == 0 {
		log.Println("⚠️ No facilities found, skipping stock adjustment seeding")
		return
	}

	adjustments := []models.StockAdjustment{
		// Facility-level adjustments
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[0].FacilityCode,
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -5),
			AdjAdjustmentType:   "inventory_count",
			AdjAdjustmentReason: "Physical count variance",
			AdjProductCode:      "PROD_PARACETAMOL_500",
			AdjBatchNumber:      "BATCH_2025_001",
			AdjQuantity:         -50,
			AdjExpiryDate:       time.Now().AddDate(1, 0, 0),
			AdjReferenceNumber:  stringPtr("PHY-COUNT-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[1].FacilityCode,
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -4),
			AdjAdjustmentType:   "damage",
			AdjAdjustmentReason: "Damaged during storage",
			AdjProductCode:      "PROD_AMOXICILLIN_250",
			AdjBatchNumber:      "BATCH_2025_003",
			AdjQuantity:         -100,
			AdjExpiryDate:       time.Now().AddDate(2, 0, 0),
			AdjReferenceNumber:  stringPtr("DAMAGE-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[2].FacilityCode,
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -3),
			AdjAdjustmentType:   "found",
			AdjAdjustmentReason: "Found during stock take",
			AdjProductCode:      "PROD_COTRIMOXAZOLE_480",
			AdjBatchNumber:      "BATCH_2025_004",
			AdjQuantity:         25,
			AdjExpiryDate:       time.Now().AddDate(1, 3, 0),
			AdjReferenceNumber:  stringPtr("FOUND-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[3].FacilityCode,
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -2),
			AdjAdjustmentType:   "theft",
			AdjAdjustmentReason: "Theft reported",
			AdjProductCode:      "PROD_METFORMIN_500",
			AdjBatchNumber:      "BATCH_2025_005",
			AdjQuantity:         -30,
			AdjExpiryDate:       time.Now().AddDate(1, 9, 0),
			AdjReferenceNumber:  stringPtr("THEFT-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
		// Pharmacy-level adjustments
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[0].FacilityCode,
			AdjPharmacyID:       &pharmacies[0].ID, // Main Pharmacy
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -3),
			AdjAdjustmentType:   "expiry",
			AdjAdjustmentReason: "Expired products removed",
			AdjProductCode:      "PROD_IBUPROFEN_400",
			AdjBatchNumber:      "BATCH_2024_001",
			AdjQuantity:         -25,
			AdjExpiryDate:       time.Now().AddDate(0, 0, -10), // Already expired
			AdjReferenceNumber:  stringPtr("EXPIRY-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
			AdjNotes:            stringPtr("Expired products disposed according to protocol"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[0].FacilityCode,
			AdjPharmacyID:       &pharmacies[1].ID, // Pediatric Pharmacy
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -2),
			AdjAdjustmentType:   "damage",
			AdjAdjustmentReason: "Damaged packaging",
			AdjProductCode:      "PROD_PARACETAMOL_500",
			AdjBatchNumber:      "BATCH_2025_001",
			AdjQuantity:         -10,
			AdjExpiryDate:       time.Now().AddDate(1, 0, 0),
			AdjReferenceNumber:  stringPtr("DAMAGE-PED-2025-001"),
			AdjApprovedBy:       stringPtr("Pediatric Pharmacy Manager"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[1].FacilityCode,
			AdjPharmacyID:       &pharmacies[4].ID, // Main Pharmacy Jinja
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -1),
			AdjAdjustmentType:   "inventory_count",
			AdjAdjustmentReason: "Physical count variance",
			AdjProductCode:      "PROD_AMOXICILLIN_250",
			AdjBatchNumber:      "BATCH_2025_003",
			AdjQuantity:         -15,
			AdjExpiryDate:       time.Now().AddDate(2, 0, 0),
			AdjReferenceNumber:  stringPtr("PHY-COUNT-JNJ-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     facilities[2].FacilityCode,
			AdjPharmacyID:       &pharmacies[6].ID, // Main Pharmacy Mbarara
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -1),
			AdjAdjustmentType:   "found",
			AdjAdjustmentReason: "Found during pharmacy stock take",
			AdjProductCode:      "PROD_COTRIMOXAZOLE_480",
			AdjBatchNumber:      "BATCH_2025_004",
			AdjQuantity:         20,
			AdjExpiryDate:       time.Now().AddDate(1, 3, 0),
			AdjReferenceNumber:  stringPtr("FOUND-MBR-2025-001"),
			AdjApprovedBy:       stringPtr("Pharmacy Manager"),
		},
	}

	for _, adj := range adjustments {
		if err := DB.Create(&adj).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock adjustment: %v", err)
		}
	}
	log.Println("✅ Stock adjustment data seeded (8 records: 4 facility-level, 4 pharmacy-level)")
}

func seedStockDispensed() {
	dispensed := []models.StockDispensed{
		{
			DspSystemCode:        "HEALTHSYS-01",
			DspFacilityCode:      "HC_KAMPALA_001",
			DspTimestamp:         time.Now(),
			DspDispenseDate:      time.Now(),
			DspProductCode:       "PROD_PARACETAMOL_500",
			DspBatchNumber:       "BATCH_2025_001",
			DspDispensedQuantity: 25.0,
			DspPatientHash:       "PAT_HASH_001",
			DspExpiryDate:        time.Now().AddDate(1, 0, 0),
		},
		{
			DspSystemCode:        "HEALTHSYS-01",
			DspFacilityCode:      "HC_JINJA_002",
			DspTimestamp:         time.Now(),
			DspDispenseDate:      time.Now().AddDate(0, 0, -1),
			DspProductCode:       "PROD_AMOXICILLIN_250",
			DspBatchNumber:       "BATCH_2025_003",
			DspDispensedQuantity: 35.5,
			DspPatientHash:       "PAT_HASH_002",
			DspExpiryDate:        time.Now().AddDate(2, 0, 0),
		},
		{
			DspSystemCode:        "HEALTHSYS-01",
			DspFacilityCode:      "HC_KAMPALA_001",
			DspTimestamp:         time.Now(),
			DspDispenseDate:      time.Now().AddDate(0, 0, -2),
			DspProductCode:       "PROD_IBUPROFEN_400",
			DspBatchNumber:       "BATCH_2025_002",
			DspDispensedQuantity: 15.0,
			DspPatientHash:       "PAT_HASH_003",
			DspExpiryDate:        time.Now().AddDate(1, 6, 0),
		},
	}

	for _, disp := range dispensed {
		if err := DB.Create(&disp).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock dispensed: %v", err)
		}
	}
	log.Println("✅ Stock dispensed data seeded (3 records)")
}

func seedStockReturns() {
	returns := []models.StockReturn{
		{
			RtnSystemCode:   "HEALTHSYS-01",
			RtnFacilityCode: "HC_KAMPALA_001",
			RtnTimestamp:    time.Now(),
			RtnReturnNumber: "RTN_KLA_001",
			RtnReturnDate:   time.Now().AddDate(0, 0, -1),
			RtnProductCode:  "PROD_PARACETAMOL_500",
			RtnBatchNumber:  "BATCH_2025_001",
			RtnUnitCode:     "UNIT_001",
			RtnQuantity:     50,
		},
		{
			RtnSystemCode:   "HEALTHSYS-01",
			RtnFacilityCode: "HC_JINJA_002",
			RtnTimestamp:    time.Now(),
			RtnReturnNumber: "RTN_JNJ_002",
			RtnReturnDate:   time.Now(),
			RtnProductCode:  "PROD_IBUPROFEN_400",
			RtnBatchNumber:  "BATCH_2025_002",
			RtnUnitCode:     "UNIT_002",
			RtnQuantity:     25,
		},
	}

	for _, ret := range returns {
		if err := DB.Create(&ret).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock return: %v", err)
		}
	}
	log.Println("✅ Stock return data seeded (2 records)")
}

func seedWarehouseOrders() {
	orders := []models.WarehouseOrder{
		{
			WarehouseCode:   "JMS",
			OrderNumber:     "ORD_JMS_001",
			ReceivedDate:    time.Now().AddDate(0, 0, -5),
			HonoredQuantity: 2000,
			DeliveredCount:  1,
			Status:          "delivered",
		},
		{
			WarehouseCode:   "NMS",
			OrderNumber:     "ORD_NMS_001",
			ReceivedDate:    time.Now().AddDate(0, 0, -3),
			HonoredQuantity: 5000,
			DeliveredCount:  2,
			Status:          "delivered",
		},
		{
			WarehouseCode:   "JMS",
			OrderNumber:     "ORD_JMS_002",
			ReceivedDate:    time.Now().AddDate(0, 0, -1),
			HonoredQuantity: 3000,
			DeliveredCount:  0,
			Status:          "pending",
		},
	}

	for _, order := range orders {
		if err := DB.Create(&order).Error; err != nil {
			log.Printf("⚠️ Failed to seed warehouse order: %v", err)
			continue
		}
	}
	log.Println("✅ Warehouse order data seeded (3 records)")
}

func seedWarehouseDeliveries() {
	deliveries := []models.WarehouseDelivery{
		{
			OrderID:     1,
			DeliveryRef: "DEL_JMS_001",
			DeliveredAt: time.Now().AddDate(0, 0, -5),
			Quantity:    2000,
			Status:      "completed",
		},
		{
			OrderID:     2,
			DeliveryRef: "DEL_NMS_001",
			DeliveredAt: time.Now().AddDate(0, 0, -3),
			Quantity:    3000,
			Status:      "completed",
		},
		{
			OrderID:     2,
			DeliveryRef: "DEL_NMS_002",
			DeliveredAt: time.Now().AddDate(0, 0, -2),
			Quantity:    2000,
			Status:      "completed",
		},
	}

	for _, delivery := range deliveries {
		if err := DB.Create(&delivery).Error; err != nil {
			log.Printf("⚠️ Failed to seed warehouse delivery: %v", err)
		}
	}
	log.Println("✅ Warehouse delivery data seeded (3 records)")
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func intPtr(i int) *int {
	return &i
}
