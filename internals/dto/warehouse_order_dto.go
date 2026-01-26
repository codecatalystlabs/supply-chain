package dto

import "time"

type WarehouseDeliveryCreateDTO struct {
	DeliveryRef string    `json:"delivery_ref" binding:"required"`
	DeliveredAt time.Time `json:"delivered_at"`
	Quantity    int       `json:"quantity" binding:"required"`
	Status      string    `json:"status"`
}

type WarehouseOrderCreateDTO struct {
	WarehouseCode   string                       `json:"warehouse_code" binding:"required"`
	OrderNumber     string                       `json:"order_number" binding:"required"`
	ReceivedDate    time.Time                    `json:"received_date"`
	HonoredQuantity int                          `json:"honored_quantity"`
	Status          string                       `json:"status"`
	Deliveries      []WarehouseDeliveryCreateDTO `json:"deliveries"`
}

type WarehouseDeliveryResponseDTO struct {
	ID          uint64    `json:"id"`
	OrderID     uint64    `json:"order_id"`
	DeliveryRef string    `json:"delivery_ref"`
	DeliveredAt time.Time `json:"delivered_at"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
}

type WarehouseOrderResponseDTO struct {
	ID              uint64                         `json:"id"`
	WarehouseCode   string                         `json:"warehouse_code"`
	OrderNumber     string                         `json:"order_number"`
	ReceivedDate    time.Time                      `json:"received_date"`
	HonoredQuantity int                            `json:"honored_quantity"`
	DeliveredCount  int                            `json:"delivered_count"`
	Status          string                         `json:"status"`
	Deliveries      []WarehouseDeliveryResponseDTO `json:"deliveries,omitempty"`
}
