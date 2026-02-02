package seeder

import (
	"log"
	"os"
	"time"

	"supply-chain/internals/config"
	"supply-chain/internals/models"

	"gorm.io/gorm"
)

// Seed inserts sample data for development/testing if tables are empty.
func Seed() {
	db := config.DB
	var cnt int64

	// Support forced reseed for development: set FORCE_SEED=true
	if os.Getenv("FORCE_SEED") == "true" {
		log.Println("force seeding enabled: clearing seeded tables")
		// Allow global deletes for development only
		db = db.Session(&gorm.Session{AllowGlobalUpdate: true})
		_ = db.Delete(&models.StockOnHand{})
		_ = db.Delete(&models.StockDispensed{})
		_ = db.Delete(&models.PurchaseOrder{})
		_ = db.Delete(&models.ProcurementPlanItem{})
		_ = db.Delete(&models.ProcurementPlan{})
		_ = db.Delete(&models.PatientVisit{})
		_ = db.Delete(&models.StockReturn{})
		_ = db.Delete(&models.ProductAmc{})
		_ = db.Delete(&models.StockAdjustment{})
		_ = db.Delete(&models.PharmacyStock{})
		_ = db.Delete(&models.GoodsReceipt{})
		_ = db.Delete(&models.WarehouseDelivery{})
		_ = db.Delete(&models.WarehouseOrder{})
	}

	// =======================
	// StockOnHand
	// =======================
	db.Model(&models.StockOnHand{}).Count(&cnt)
	if cnt == 0 {
		samples := []models.StockOnHand{
			{
				SrcSystemCode:   "ELMO",
				SrcFacilityCode: "HOSP_A",
				SrcTimestamp:    time.Now().Add(-24 * time.Hour),
				SrcProductCode:  "PARA500",
				SrcBatchNumber:  "BATCH-A1",
				SrcQuantity:     150,
				SrcExpiryDate:   time.Now().AddDate(2, 0, 0),
			},
			{
				SrcSystemCode:   "ELMO",
				SrcFacilityCode: "CLINIC_B",
				SrcTimestamp:    time.Now().Add(-72 * time.Hour),
				SrcProductCode:  "AMOX250",
				SrcBatchNumber:  "BATCH-B2",
				SrcQuantity:     60,
				SrcExpiryDate:   time.Now().AddDate(1, 6, 0),
			},
		}

		for _, s := range samples {
			if err := db.Create(&s).Error; err != nil {
				log.Println("seed: failed to create StockOnHand:", err)
			}
		}
	}

	// =======================
	// PurchaseOrders
	// =======================
	db.Model(&models.PurchaseOrder{}).Count(&cnt)
	if cnt == 0 {
		pos := []models.PurchaseOrder{
			{
				OrdSystemCode:      "ORDSYS",
				OrdFacilityCode:    "HOSP_A",
				OrdTimestamp:       time.Now().Add(-48 * time.Hour),
				OrdOrderDate:       time.Now().Add(-48 * time.Hour),
				OrdOrderRefNumber:  "REF-PO-100",
				OrdOrderNumber:     "PO-100",
				OrdProductCode:     "PARA500",
				OrdOrderedQuantity: 200,
			},
			{
				OrdSystemCode:      "ORDSYS",
				OrdFacilityCode:    "CLINIC_B",
				OrdTimestamp:       time.Now().Add(-7 * 24 * time.Hour),
				OrdOrderDate:       time.Now().Add(-7 * 24 * time.Hour),
				OrdOrderRefNumber:  "REF-PO-101",
				OrdOrderNumber:     "PO-101",
				OrdProductCode:     "AMOX250",
				OrdOrderedQuantity: 100,
			},
		}

		for _, po := range pos {
			if err := db.Create(&po).Error; err != nil {
				log.Println("seed: failed to create PurchaseOrder:", err)
			}
		}
	}

	// =======================
	// Procurement Plan
	// =======================
	db.Model(&models.ProcurementPlan{}).Count(&cnt)
	if cnt == 0 {
		plan := models.ProcurementPlan{
			PlanSystemCode: "PROC_SYS",
			StoreCode:      "CENTRAL_STORE",
			CreatedAt:      time.Now().Add(-3 * 24 * time.Hour),
			Notes:          ptrString("Quarterly restock for region north"),
		}

		if err := db.Create(&plan).Error; err != nil {
			log.Println("seed: failed to create ProcurementPlan:", err)
		} else {
			items := []models.ProcurementPlanItem{
				{ProcurementID: plan.ID, ProductCode: "PARA500", Quantity: 1000, NeededBy: time.Now().AddDate(0, 1, 0), Status: "planned"},
				{ProcurementID: plan.ID, ProductCode: "AMOX250", Quantity: 500, NeededBy: time.Now().AddDate(0, 1, 0), Status: "planned"},
			}

			for _, it := range items {
				if err := db.Create(&it).Error; err != nil {
					log.Println("seed: failed to create ProcurementPlanItem:", err)
				}
			}
		}
	}

	// =======================
	// Warehouse Orders + Deliveries
	// =======================
	db.Model(&models.WarehouseOrder{}).Count(&cnt)
	if cnt == 0 {
		// NMS
		wo1 := models.WarehouseOrder{
			WarehouseCode:   "NMS",
			OrderNumber:     "NMS-ORD-900",
			ReceivedDate:    time.Now().Add(-5 * 24 * time.Hour),
			HonoredQuantity: 120,
			DeliveredCount:  2,
			Status:          "partial",
		}

		if err := db.Create(&wo1).Error; err == nil {
			deliveries := []models.WarehouseDelivery{
				{OrderID: wo1.ID, DeliveryRef: "NMS-DEL-900-A", DeliveredAt: time.Now().Add(-4 * 24 * time.Hour), Quantity: 60, Status: "delivered"},
				{OrderID: wo1.ID, DeliveryRef: "NMS-DEL-900-B", DeliveredAt: time.Now().Add(-3 * 24 * time.Hour), Quantity: 60, Status: "delivered"},
			}

			for _, d := range deliveries {
				_ = db.Create(&d)
			}
		}

		// JMS
		wo2 := models.WarehouseOrder{
			WarehouseCode:   "JMS",
			OrderNumber:     "JMS-ORD-701",
			ReceivedDate:    time.Now().Add(-10 * 24 * time.Hour),
			HonoredQuantity: 200,
			DeliveredCount:  3,
			Status:          "fulfilled",
		}

		if err := db.Create(&wo2).Error; err == nil {
			deliveries := []models.WarehouseDelivery{
				{OrderID: wo2.ID, DeliveryRef: "JMS-DEL-701-A", DeliveredAt: time.Now().Add(-9 * 24 * time.Hour), Quantity: 80, Status: "delivered"},
				{OrderID: wo2.ID, DeliveryRef: "JMS-DEL-701-B", DeliveredAt: time.Now().Add(-8 * 24 * time.Hour), Quantity: 60, Status: "delivered"},
				{OrderID: wo2.ID, DeliveryRef: "JMS-DEL-701-C", DeliveredAt: time.Now().Add(-7 * 24 * time.Hour), Quantity: 60, Status: "delivered"},
			}

			for _, d := range deliveries {
				_ = db.Create(&d)
			}
		}
	}

	// =======================
	// Remaining single-record tables
	// =======================
	seedIfEmpty(db, models.GoodsReceipt{}, models.GoodsReceipt{
		GrnSystemCode:            "GRN_SYS",
		GrnFacilityCode:          "HOSP_A",
		GrnTimestamp:             time.Now().Add(-6 * 24 * time.Hour),
		GrnReceiptDate:           time.Now().Add(-6 * 24 * time.Hour),
		GrnFacilityReceiptNumber: "GRN-900",
		GrnWarehouseRefNumber:    "NMS",
		GrnOrderNumber:           "PO-100",
		GrnProductCode:           "PARA500",
		GrnBatchNumber:           "BATCH-A1",
		GrnQuantity:              200,
		GrnExpiryDate:            time.Now().AddDate(2, 0, 0),
		GrnSupplierCode:          "SUPPLY_CO",
	})
}

// =======================
// Helpers
// =======================
func ptrString(s string) *string { return &s }

func seedIfEmpty[T any](db *gorm.DB, model T, data T) {
	var cnt int64
	db.Model(&model).Count(&cnt)
	if cnt == 0 {
		_ = db.Create(&data)
	}
}
