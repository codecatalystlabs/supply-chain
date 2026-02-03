package handlers

import (
	"fmt"
	"net/http"
	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create facility order
// @Tags FacilityOrder
// @Accept json
// @Produce json
// @Param payload body dto.FacilityOrderCreateDTO true "Order payload"
// @Success 201 {object} dto.FacilityOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders [post]
func CreateFacilityOrder(c *gin.Context) {
	var payload dto.FacilityOrderCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify facility exists
	var facility models.Facility
	if err := config.DB.First(&facility, payload.FacilityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Facility not found"})
		return
	}

	// Verify warehouse exists
	var warehouse models.Warehouse
	if err := config.DB.First(&warehouse, payload.WarehouseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	// Generate order number
	orderNumber := fmt.Sprintf("ORD-%s-%s-%d", facility.FacilityCode, warehouse.WarehouseCode, time.Now().Unix())

	// Calculate totals
	totalItems := len(payload.Items)
	totalQuantity := 0
	var totalValue *float64
	for _, item := range payload.Items {
		totalQuantity += item.OrderedQuantity
		if item.UnitPrice != nil {
			itemTotal := float64(item.OrderedQuantity) * *item.UnitPrice
			if totalValue == nil {
				totalValue = new(float64)
			}
			*totalValue += itemTotal
		}
	}

	order := models.FacilityOrder{
		OrderNumber:          orderNumber,
		OrderRefNumber:       payload.OrderRefNumber,
		FacilityID:           payload.FacilityID,
		FacilityCode:         facility.FacilityCode,
		WarehouseID:          payload.WarehouseID,
		WarehouseCode:        warehouse.WarehouseCode,
		OrderDate:            payload.OrderDate,
		OrderType:            payload.OrderType,
		OrderStatus:          "pending",
		Priority:             payload.Priority,
		FinancialYear:        payload.FinancialYear,
		OrderCycle:           payload.OrderCycle,
		ProcurementPlanID:    payload.ProcurementPlanID,
		ExpectedDeliveryDate: payload.ExpectedDeliveryDate,
		TotalItems:           totalItems,
		TotalQuantity:        totalQuantity,
		TotalValue:           totalValue,
		Notes:                payload.Notes,
		SourceSystem:         payload.SourceSystem,
		SourceRecordID:       payload.SourceRecordID,
		IdempotencyKey:       payload.IdempotencyKey,
		CreatedAt:            time.Now(),
	}

	// Check idempotency if provided
	if payload.IdempotencyKey != nil {
		var existing models.FacilityOrder
		if err := config.DB.Where("idempotency_key = ?", *payload.IdempotencyKey).First(&existing).Error; err == nil {
			config.DB.Preload("OrderItems").Preload("Deliveries").First(&existing, existing.ID)
			c.JSON(http.StatusOK, mapToFacilityOrderResponse(existing))
			return
		}
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create order items
	for _, itemDTO := range payload.Items {
		var itemTotalPrice *float64
		if itemDTO.UnitPrice != nil {
			total := float64(itemDTO.OrderedQuantity) * *itemDTO.UnitPrice
			itemTotalPrice = &total
		}

		item := models.FacilityOrderItem{
			OrderID:            order.ID,
			ProductCode:        itemDTO.ProductCode,
			ProductDescription: itemDTO.ProductDescription,
			UOM:                itemDTO.UOM,
			OrderedQuantity:    itemDTO.OrderedQuantity,
			UnitPrice:          itemDTO.UnitPrice,
			TotalPrice:         itemTotalPrice,
			Currency:           itemDTO.Currency,
			Status:             "pending",
			Notes:              itemDTO.Notes,
			CreatedAt:          time.Now(),
		}

		if err := config.DB.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create order item: %v", err)})
			return
		}
	}

	// Reload with relationships
	config.DB.Preload("OrderItems").Preload("Deliveries").First(&order, order.ID)
	c.JSON(http.StatusCreated, mapToFacilityOrderResponse(order))
}

// @Summary List facility orders
// @Tags FacilityOrder
// @Produce json
// @Param facility_id query int false "Filter by facility ID"
// @Param warehouse_id query int false "Filter by warehouse ID"
// @Param status query string false "Filter by order status"
// @Success 200 {array} dto.FacilityOrderResponseDTO
// @Failure 500 {object} map[string]string
// @Router /facility-orders [get]
func ListFacilityOrders(c *gin.Context) {
	var orders []models.FacilityOrder
	query := config.DB.Preload("OrderItems").Preload("Deliveries")

	// Apply filters
	if facilityID := c.Query("facility_id"); facilityID != "" {
		query = query.Where("facility_id = ?", facilityID)
	}
	if warehouseID := c.Query("warehouse_id"); warehouseID != "" {
		query = query.Where("warehouse_id = ?", warehouseID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("order_status = ?", status)
	}

	if err := query.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.FacilityOrderResponseDTO
	for _, o := range orders {
		resp = append(resp, mapToFacilityOrderResponse(o))
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get facility order by ID
// @Tags FacilityOrder
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.FacilityOrderResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders/{id} [get]
func GetFacilityOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.FacilityOrder
	if err := config.DB.Preload("OrderItems").Preload("Deliveries").Preload("Deliveries.DeliveryItems").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, mapToFacilityOrderResponse(order))
}

// @Summary Update facility order
// @Tags FacilityOrder
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param payload body dto.FacilityOrderUpdateDTO true "Update payload"
// @Success 200 {object} dto.FacilityOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders/{id} [put]
func UpdateFacilityOrder(c *gin.Context) {
	id := c.Param("id")
	var payload dto.FacilityOrderUpdateDTO
	var order models.FacilityOrder

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if payload.OrderStatus != nil {
		order.OrderStatus = *payload.OrderStatus
	}
	if payload.Priority != nil {
		order.Priority = payload.Priority
	}
	if payload.ExpectedDeliveryDate != nil {
		order.ExpectedDeliveryDate = payload.ExpectedDeliveryDate
	}
	if payload.Notes != nil {
		order.Notes = payload.Notes
	}
	now := time.Now()
	order.UpdatedAt = &now

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("OrderItems").Preload("Deliveries").First(&order, order.ID)
	c.JSON(http.StatusOK, mapToFacilityOrderResponse(order))
}

// @Summary Submit facility order for approval
// @Tags FacilityOrder
// @Param id path int true "Order ID"
// @Param payload body map[string]string true "Submit payload" SchemaExample({"submitted_by": "Dr. John Doe"})
// @Success 200 {object} dto.FacilityOrderResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders/{id}/submit [post]
func SubmitFacilityOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.FacilityOrder
	var payload struct {
		SubmittedBy string `json:"submitted_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	now := time.Now()
	order.OrderStatus = "submitted"
	order.SubmittedBy = &payload.SubmittedBy
	order.SubmittedAt = &now
	order.UpdatedAt = &now

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("OrderItems").Preload("Deliveries").First(&order, order.ID)
	c.JSON(http.StatusOK, mapToFacilityOrderResponse(order))
}

// @Summary Approve facility order
// @Tags FacilityOrder
// @Param id path int true "Order ID"
// @Param payload body map[string]string true "Approve payload" SchemaExample({"approved_by": "Pharmacy Manager"})
// @Success 200 {object} dto.FacilityOrderResponseDTO
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders/{id}/approve [post]
func ApproveFacilityOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.FacilityOrder
	var payload struct {
		ApprovedBy string `json:"approved_by" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	now := time.Now()
	order.OrderStatus = "approved"
	order.ApprovedBy = &payload.ApprovedBy
	order.ApprovedAt = &now
	order.UpdatedAt = &now

	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("OrderItems").Preload("Deliveries").First(&order, order.ID)
	c.JSON(http.StatusOK, mapToFacilityOrderResponse(order))
}

// @Summary Create delivery for facility order
// @Tags FacilityOrder
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param payload body dto.FacilityDeliveryCreateDTO true "Delivery payload"
// @Success 201 {object} dto.FacilityDeliveryResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /facility-orders/{id}/deliveries [post]
func CreateFacilityDelivery(c *gin.Context) {
	orderID := c.Param("id")
	var payload dto.FacilityDeliveryCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify order exists
	var order models.FacilityOrder
	if err := config.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Generate delivery reference
	deliveryRef := fmt.Sprintf("DEL-%s-%d", order.OrderNumber, time.Now().Unix())

	// Calculate totals
	totalItems := len(payload.Items)
	totalQuantity := 0
	for _, item := range payload.Items {
		totalQuantity += item.Quantity
	}

	delivery := models.FacilityDelivery{
		OrderID:        order.ID,
		DeliveryRef:    deliveryRef,
		DeliveryNumber: payload.DeliveryNumber,
		DeliveredAt:    payload.DeliveredAt,
		DeliveryDate:   payload.DeliveryDate,
		Status:         "pending",
		DeliveredBy:    payload.DeliveredBy,
		VehicleNumber:  payload.VehicleNumber,
		DriverName:     payload.DriverName,
		TotalItems:     totalItems,
		TotalQuantity:  totalQuantity,
		Notes:          payload.Notes,
		ConditionNotes: payload.ConditionNotes,
		CreatedAt:      time.Now(),
	}

	if err := config.DB.Create(&delivery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create delivery items
	for _, itemDTO := range payload.Items {
		var itemTotalPrice *float64
		if itemDTO.UnitPrice != nil {
			receivedQty := itemDTO.Quantity
			if itemDTO.ReceivedQuantity != nil {
				receivedQty = *itemDTO.ReceivedQuantity
			}
			total := float64(receivedQty) * *itemDTO.UnitPrice
			itemTotalPrice = &total
		}

		item := models.FacilityDeliveryItem{
			DeliveryID:         delivery.ID,
			OrderItemID:        itemDTO.OrderItemID,
			ProductCode:        itemDTO.ProductCode,
			ProductDescription: itemDTO.ProductDescription,
			BatchNumber:        itemDTO.BatchNumber,
			ExpiryDate:         itemDTO.ExpiryDate,
			Quantity:           itemDTO.Quantity,
			ReceivedQuantity:   itemDTO.ReceivedQuantity,
			UnitPrice:          itemDTO.UnitPrice,
			TotalPrice:         itemTotalPrice,
			Condition:          itemDTO.Condition,
			ConditionNotes:     itemDTO.ConditionNotes,
			CreatedAt:          time.Now(),
		}

		if err := config.DB.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create delivery item: %v", err)})
			return
		}
	}

	// Update order status
	order.OrderStatus = "processing"
	now := time.Now()
	order.UpdatedAt = &now
	config.DB.Save(&order)

	// Reload with relationships
	config.DB.Preload("DeliveryItems").First(&delivery, delivery.ID)
	c.JSON(http.StatusCreated, mapToFacilityDeliveryResponse(delivery))
}

// Helper functions
func mapToFacilityOrderResponse(o models.FacilityOrder) dto.FacilityOrderResponseDTO {
	var items []dto.FacilityOrderItemResponseDTO
	for _, item := range o.OrderItems {
		items = append(items, mapToFacilityOrderItemResponse(item))
	}

	var deliveries []dto.FacilityDeliveryResponseDTO
	for _, del := range o.Deliveries {
		deliveries = append(deliveries, mapToFacilityDeliveryResponse(del))
	}

	return dto.FacilityOrderResponseDTO{
		ID:                   o.ID,
		OrderNumber:          o.OrderNumber,
		OrderRefNumber:       o.OrderRefNumber,
		FacilityID:           o.FacilityID,
		FacilityCode:         o.FacilityCode,
		WarehouseID:          o.WarehouseID,
		WarehouseCode:        o.WarehouseCode,
		OrderDate:            o.OrderDate,
		OrderType:            o.OrderType,
		OrderStatus:          o.OrderStatus,
		Priority:             o.Priority,
		FinancialYear:        o.FinancialYear,
		OrderCycle:           o.OrderCycle,
		ProcurementPlanID:    o.ProcurementPlanID,
		SubmittedBy:          o.SubmittedBy,
		SubmittedAt:          o.SubmittedAt,
		ApprovedBy:           o.ApprovedBy,
		ApprovedAt:           o.ApprovedAt,
		ExpectedDeliveryDate: o.ExpectedDeliveryDate,
		ActualDeliveryDate:   o.ActualDeliveryDate,
		TotalItems:           o.TotalItems,
		TotalQuantity:        o.TotalQuantity,
		TotalValue:           o.TotalValue,
		Notes:                o.Notes,
		SourceSystem:         o.SourceSystem,
		CreatedAt:            o.CreatedAt,
		UpdatedAt:            o.UpdatedAt,
		Items:                items,
		Deliveries:           deliveries,
	}
}

func mapToFacilityOrderItemResponse(item models.FacilityOrderItem) dto.FacilityOrderItemResponseDTO {
	return dto.FacilityOrderItemResponseDTO{
		ID:                 item.ID,
		OrderID:            item.OrderID,
		ProductCode:        item.ProductCode,
		ProductDescription: item.ProductDescription,
		UOM:                item.UOM,
		OrderedQuantity:    item.OrderedQuantity,
		HonoredQuantity:    item.HonoredQuantity,
		DeliveredQuantity:  item.DeliveredQuantity,
		UnitPrice:          item.UnitPrice,
		TotalPrice:         item.TotalPrice,
		Currency:           item.Currency,
		Status:             item.Status,
		Notes:              item.Notes,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}
}

func mapToFacilityDeliveryResponse(d models.FacilityDelivery) dto.FacilityDeliveryResponseDTO {
	var items []dto.FacilityDeliveryItemResponseDTO
	for _, item := range d.DeliveryItems {
		items = append(items, mapToFacilityDeliveryItemResponse(item))
	}

	return dto.FacilityDeliveryResponseDTO{
		ID:             d.ID,
		OrderID:        d.OrderID,
		DeliveryRef:    d.DeliveryRef,
		DeliveryNumber: d.DeliveryNumber,
		DeliveredAt:    d.DeliveredAt,
		DeliveryDate:   d.DeliveryDate,
		Status:         d.Status,
		DeliveredBy:    d.DeliveredBy,
		ReceivedBy:     d.ReceivedBy,
		ReceivedAt:     d.ReceivedAt,
		VehicleNumber:  d.VehicleNumber,
		DriverName:     d.DriverName,
		TotalItems:     d.TotalItems,
		TotalQuantity:  d.TotalQuantity,
		Notes:          d.Notes,
		ConditionNotes: d.ConditionNotes,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
		Items:          items,
	}
}

func mapToFacilityDeliveryItemResponse(item models.FacilityDeliveryItem) dto.FacilityDeliveryItemResponseDTO {
	return dto.FacilityDeliveryItemResponseDTO{
		ID:                 item.ID,
		DeliveryID:         item.DeliveryID,
		OrderItemID:        item.OrderItemID,
		ProductCode:        item.ProductCode,
		ProductDescription: item.ProductDescription,
		BatchNumber:        item.BatchNumber,
		ExpiryDate:         item.ExpiryDate,
		Quantity:           item.Quantity,
		ReceivedQuantity:   item.ReceivedQuantity,
		UnitPrice:          item.UnitPrice,
		TotalPrice:         item.TotalPrice,
		Condition:          item.Condition,
		ConditionNotes:     item.ConditionNotes,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}
}
