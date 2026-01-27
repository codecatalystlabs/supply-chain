package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"supply-chain/internals/config"
	"supply-chain/internals/handlers"

	_ "supply-chain/docs" // 👈 REQUIRED for Swagger
)

// @title MoH Emergency Dispatch API
// @version 1.0
// @description Backend API for emergency call & dispatch system
// @termsOfService https://health.go.ug

// @contact.name DHI
// @contact.email ict@moh.go.ug

// @host localhost:8080
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		// try loading .env from the cmd/server folder (when run from repo root)
		if err2 := godotenv.Load("cmd/server/.env"); err2 != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.ConnectDatabase()
	config.SeedDatabase()

	router := gin.Default()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stock := router.Group("/api/v1/stock")
	{
		stock.POST("/on-hand", handlers.CreateStockOnHand)
		stock.GET("/on-hand", handlers.ListStockOnHand)
		stock.GET("/on-hand/:id", handlers.GetStockOnHand)
		stock.PUT("/on-hand/:id", handlers.UpdateStockOnHand)
		stock.DELETE("/on-hand/:id", handlers.DeleteStockOnHand)

		// stock dispensed
		stock.POST("/dispensed", handlers.CreateStockDispensed)
		stock.GET("/dispensed", handlers.ListStockDispensed)
		stock.GET("/dispensed/:id", handlers.GetStockDispensed)
		stock.PUT("/dispensed/:id", handlers.UpdateStockDispensed)
		stock.DELETE("/dispensed/:id", handlers.DeleteStockDispensed)

		// stock returns
		stock.POST("/return", handlers.CreateStockReturn)
		stock.GET("/return", handlers.ListStockReturns)
		stock.GET("/return/:id", handlers.GetStockReturn)
		stock.PUT("/return/:id", handlers.UpdateStockReturn)
		stock.DELETE("/return/:id", handlers.DeleteStockReturn)

		// stock adjustments
		stock.POST("/adjustment", handlers.CreateStockAdjustment)
		stock.GET("/adjustment", handlers.ListStockAdjustments)
		stock.GET("/adjustment/:id", handlers.GetStockAdjustment)
		stock.PUT("/adjustment/:id", handlers.UpdateStockAdjustment)
		stock.DELETE("/adjustment/:id", handlers.DeleteStockAdjustment)
	}
	// Purchase Order routes
	po := router.Group("/api/v1/purchase-order")
	{
		po.POST("/", handlers.CreatePurchaseOrder)
		po.GET("/:id", handlers.GetPurchaseOrder)
		po.PUT("/:id", handlers.UpdatePurchaseOrder)
		po.DELETE("/:id", handlers.DeletePurchaseOrder)
		po.GET("/", handlers.ListPurchaseOrders)
	}

	// Pharmacy stock
	pha := router.Group("/api/v1/pharmacy-stock")
	{
		pha.POST("/", handlers.CreatePharmacyStock)
		pha.GET("/", handlers.ListPharmacyStock)
		pha.GET("/:id", handlers.GetPharmacyStock)
		pha.PUT("/:id", handlers.UpdatePharmacyStock)
		pha.DELETE("/:id", handlers.DeletePharmacyStock)
	}

	// Goods receipt
	grn := router.Group("/api/v1/goods-receipt")
	{
		grn.POST("/", handlers.CreateGoodsReceipt)
		grn.GET("/", handlers.ListGoodsReceipts)
		grn.GET("/:id", handlers.GetGoodsReceipt)
		grn.PUT("/:id", handlers.UpdateGoodsReceipt)
		grn.DELETE("/:id", handlers.DeleteGoodsReceipt)
	}

	// Patient visits
	pv := router.Group("/api/v1/patient-visit")
	{
		pv.POST("/", handlers.CreatePatientVisit)
		pv.GET("/", handlers.ListPatientVisits)
		pv.GET("/:id", handlers.GetPatientVisit)
		pv.PUT("/:id", handlers.UpdatePatientVisit)
		pv.DELETE("/:id", handlers.DeletePatientVisit)
	}

	// Product AMC
	pa := router.Group("/api/v1/product-amc")
	{
		pa.POST("/", handlers.CreateProductAmc)
		pa.GET("/", handlers.ListProductAmc)
		pa.GET("/:id", handlers.GetProductAmc)
		pa.PUT("/:id", handlers.UpdateProductAmc)
		pa.DELETE("/:id", handlers.DeleteProductAmc)
	}

	// Procurement plans
	pp := router.Group("/api/v1/procurement")
	{
		pp.POST("/", handlers.CreateProcurementPlan)
		pp.GET("/", handlers.ListProcurementPlans)
		pp.GET("/:id", handlers.GetProcurementPlan)
		pp.DELETE("/:id", handlers.DeleteProcurementPlan)
	}

	// Warehouse orders and deliveries
	wo := router.Group("/api/v1/warehouse-orders")
	{
		wo.POST("/", handlers.ReceiveWarehouseOrder)
		wo.GET("/", handlers.ListWarehouseOrders)
		wo.GET("/:id", handlers.GetWarehouseOrder)
	}

	// Run server
	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
