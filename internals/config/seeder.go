package config

import (
	"log"
	"time"

	"supply-chain/internals/models"
)

// SeedDatabase populates the database with comprehensive sample data
func SeedDatabase() {
	log.Println("🌱 Seeding database with comprehensive data...")

	seedStockOnHand()
	seedPurchaseOrders()
	seedProductAmc()
	seedGoodsReceipts()
	seedPatientVisits()
	seedPharmacyStock()
	seedProcurementPlans()
	seedStockAdjustments()
	seedStockDispensed()
	seedStockReturns()
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

func seedStockOnHand() {
	stocks := []models.StockOnHand{
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: "HC_KAMPALA_001",
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_PARACETAMOL_500",
			SrcBatchNumber:  "BATCH_2025_001",
			SrcQuantity:     500,
			SrcExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: "HC_KAMPALA_001",
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_IBUPROFEN_400",
			SrcBatchNumber:  "BATCH_2025_002",
			SrcQuantity:     300,
			SrcExpiryDate:   time.Now().AddDate(1, 6, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: "HC_JINJA_002",
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_AMOXICILLIN_250",
			SrcBatchNumber:  "BATCH_2025_003",
			SrcQuantity:     1000,
			SrcExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: "HC_MBARARA_003",
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_COTRIMOXAZOLE_480",
			SrcBatchNumber:  "BATCH_2025_004",
			SrcQuantity:     750,
			SrcExpiryDate:   time.Now().AddDate(1, 3, 0),
		},
		{
			SrcSystemCode:   "HEALTHSYS-01",
			SrcFacilityCode: "HC_GULU_004",
			SrcTimestamp:    time.Now(),
			SrcProductCode:  "PROD_METFORMIN_500",
			SrcBatchNumber:  "BATCH_2025_005",
			SrcQuantity:     2000,
			SrcExpiryDate:   time.Now().AddDate(1, 9, 0),
		},
	}

	for _, stock := range stocks {
		if err := DB.Create(&stock).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock on hand: %v", err)
		}
	}
	log.Println("✅ Stock on hand data seeded (5 records)")
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
	stocks := []models.PharmacyStock{
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: "HC_KAMPALA_001",
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_PARACETAMOL_500",
			PhaBatchNumber:  "BATCH_2025_001",
			PhaQuantity:     450,
			PhaExpiryDate:   time.Now().AddDate(1, 0, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: "HC_JINJA_002",
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_AMOXICILLIN_250",
			PhaBatchNumber:  "BATCH_2025_003",
			PhaQuantity:     4500,
			PhaExpiryDate:   time.Now().AddDate(2, 0, 0),
		},
		{
			PhaSystemCode:   "HEALTHSYS-01",
			PhaFacilityCode: "HC_MBARARA_003",
			PhaTimestamp:    time.Now(),
			PhaProductCode:  "PROD_COTRIMOXAZOLE_480",
			PhaBatchNumber:  "BATCH_2025_004",
			PhaQuantity:     2800,
			PhaExpiryDate:   time.Now().AddDate(1, 8, 0),
		},
	}

	for _, stock := range stocks {
		if err := DB.Create(&stock).Error; err != nil {
			log.Printf("⚠️ Failed to seed pharmacy stock: %v", err)
		}
	}
	log.Println("✅ Pharmacy stock data seeded (3 records)")
}

func seedProcurementPlans() {
	plans := []models.ProcurementPlan{
		{
			PlanSystemCode: "HEALTHSYS-01",
			StoreCode:      "JMS",
			CreatedAt:      time.Now(),
			Notes:          stringPtr("Q1 2025 procurement plan for JMS"),
		},
		{
			PlanSystemCode: "HEALTHSYS-01",
			StoreCode:      "NMS",
			CreatedAt:      time.Now(),
			Notes:          stringPtr("Q1 2025 procurement plan for NMS"),
		},
	}

	for _, plan := range plans {
		if err := DB.Create(&plan).Error; err != nil {
			log.Printf("⚠️ Failed to seed procurement plan: %v", err)
			continue
		}

		// Add plan items
		items := []models.ProcurementPlanItem{
			{
				ProcurementID: plan.ID,
				ProductCode:   "PROD_PARACETAMOL_500",
				Quantity:      5000,
				NeededBy:      time.Now().AddDate(0, 1, 0),
				Status:        "planned",
			},
			{
				ProcurementID: plan.ID,
				ProductCode:   "PROD_AMOXICILLIN_250",
				Quantity:      8000,
				NeededBy:      time.Now().AddDate(0, 1, 0),
				Status:        "ordered",
			},
		}

		for _, item := range items {
			if err := DB.Create(&item).Error; err != nil {
				log.Printf("⚠️ Failed to seed procurement plan item: %v", err)
			}
		}
	}
	log.Println("✅ Procurement plan data seeded (2 plans with 4 items)")
}

func seedStockAdjustments() {
	adjustments := []models.StockAdjustment{
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     "HC_KAMPALA_001",
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -2),
			AdjAdjustmentType:   "inventory_count",
			AdjAdjustmentReason: "Physical count variance",
			AdjProductCode:      "PROD_PARACETAMOL_500",
			AdjBatchNumber:      "BATCH_2025_001",
			AdjQuantity:         -50,
			AdjExpiryDate:       time.Now().AddDate(1, 0, 0),
		},
		{
			AdjSystemCode:       "HEALTHSYS-01",
			AdjFacilityCode:     "HC_JINJA_002",
			AdjTimestamp:        time.Now(),
			AdjAdjustmentDate:   time.Now().AddDate(0, 0, -1),
			AdjAdjustmentType:   "damage",
			AdjAdjustmentReason: "Damaged during storage",
			AdjProductCode:      "PROD_AMOXICILLIN_250",
			AdjBatchNumber:      "BATCH_2025_003",
			AdjQuantity:         -100,
			AdjExpiryDate:       time.Now().AddDate(2, 0, 0),
		},
	}

	for _, adj := range adjustments {
		if err := DB.Create(&adj).Error; err != nil {
			log.Printf("⚠️ Failed to seed stock adjustment: %v", err)
		}
	}
	log.Println("✅ Stock adjustment data seeded (2 records)")
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

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
