package models

import "time"

type WarehouseOrder struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	WarehouseCode   string    `gorm:"size:100;not null" json:"warehouse_code"`
	OrderNumber     string    `gorm:"size:100;not null" json:"order_number"`
	ReceivedDate    time.Time `json:"received_date"`
	HonoredQuantity int       `json:"honored_quantity"`
	DeliveredCount  int       `json:"delivered_count"`
	Status          string    `gorm:"size:50" json:"status"`
}

func (WarehouseOrder) TableName() string {
	return "Warehouse_Order"
}

type WarehouseDelivery struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	OrderID     uint64    `gorm:"not null" json:"order_id"`
	DeliveryRef string    `gorm:"size:100" json:"delivery_ref"`
	DeliveredAt time.Time `json:"delivered_at"`
	Quantity    int       `json:"quantity"`
	Status      string    `gorm:"size:50" json:"status"`
}

func (WarehouseDelivery) TableName() string {
	return "Warehouse_Delivery"
}
