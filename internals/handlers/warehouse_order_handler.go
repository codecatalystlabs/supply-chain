package handlers

import (
	"net/http"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
)

// @Summary Receive warehouse order
// @Tags WarehouseOrder
// @Accept json
// @Produce json
// @Param payload body dto.WarehouseOrderCreateDTO true "Warehouse order payload"
// @Success 201 {object} dto.WarehouseOrderResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/warehouse-orders [post]
func ReceiveWarehouseOrder(c *gin.Context) {
	var payload dto.WarehouseOrderCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.WarehouseOrder{
		WarehouseCode:   payload.WarehouseCode,
		OrderNumber:     payload.OrderNumber,
		ReceivedDate:    payload.ReceivedDate,
		HonoredQuantity: payload.HonoredQuantity,
		DeliveredCount:  len(payload.Deliveries),
		Status:          payload.Status,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create deliveries
	for _, d := range payload.Deliveries {
		dl := models.WarehouseDelivery{
			OrderID:     order.ID,
			DeliveryRef: d.DeliveryRef,
			DeliveredAt: d.DeliveredAt,
			Quantity:    d.Quantity,
			Status:      d.Status,
		}
		_ = config.DB.Create(&dl)
	}

	var deliveries []models.WarehouseDelivery
	config.DB.Where("order_id = ?", order.ID).Find(&deliveries)

	resp := mapToWarehouseOrderResponse(order, deliveries)
	c.JSON(http.StatusCreated, resp)
}

// @Summary List warehouse orders
// @Tags WarehouseOrder
// @Produce json
// @Success 200 {array} dto.WarehouseOrderResponseDTO
// @Failure 500 {object} map[string]string
// @Router /api/v1/warehouse-orders [get]
func ListWarehouseOrders(c *gin.Context) {
	var orders []models.WarehouseOrder
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.WarehouseOrderResponseDTO
	for _, o := range orders {
		var dels []models.WarehouseDelivery
		config.DB.Where("order_id = ?", o.ID).Find(&dels)
		resp = append(resp, mapToWarehouseOrderResponse(o, dels))
	}
	c.JSON(200, resp)
}

// @Summary Get warehouse order
// @Tags WarehouseOrder
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.WarehouseOrderResponseDTO
// @Failure 404 {object} map[string]string
// @Router /api/v1/warehouse-orders/{id} [get]
func GetWarehouseOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.WarehouseOrder
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}
	var dels []models.WarehouseDelivery
	config.DB.Where("order_id = ?", order.ID).Find(&dels)
	c.JSON(200, mapToWarehouseOrderResponse(order, dels))
}

func mapToWarehouseOrderResponse(o models.WarehouseOrder, dels []models.WarehouseDelivery) dto.WarehouseOrderResponseDTO {
	var dresp []dto.WarehouseDeliveryResponseDTO
	for _, d := range dels {
		dresp = append(dresp, dto.WarehouseDeliveryResponseDTO{
			ID:          d.ID,
			OrderID:     d.OrderID,
			DeliveryRef: d.DeliveryRef,
			DeliveredAt: d.DeliveredAt,
			Quantity:    d.Quantity,
			Status:      d.Status,
		})
	}
	return dto.WarehouseOrderResponseDTO{
		ID:              o.ID,
		WarehouseCode:   o.WarehouseCode,
		OrderNumber:     o.OrderNumber,
		ReceivedDate:    o.ReceivedDate,
		HonoredQuantity: o.HonoredQuantity,
		DeliveredCount:  o.DeliveredCount,
		Status:          o.Status,
		Deliveries:      dresp,
	}
}
