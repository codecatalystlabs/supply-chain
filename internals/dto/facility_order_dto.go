package dto

import "time"

// FacilityOrderCreateDTO for creating a new facility order
type FacilityOrderCreateDTO struct {
	OrderRefNumber        *string                    `json:"order_ref_number,omitempty"`
	FacilityID            uint64                     `json:"facility_id" binding:"required"`
	WarehouseID           uint64                     `json:"warehouse_id" binding:"required"`
	OrderDate             time.Time                  `json:"order_date" binding:"required"`
	OrderType             *string                    `json:"order_type,omitempty"`
	Priority              *string                    `json:"priority,omitempty"`
	FinancialYear         *string                    `json:"financial_year,omitempty"`
	OrderCycle            *string                    `json:"order_cycle,omitempty"`
	ProcurementPlanID     *uint64                    `json:"procurement_plan_id,omitempty"`
	ExpectedDeliveryDate  *time.Time                 `json:"expected_delivery_date,omitempty"`
	Notes                 *string                    `json:"notes,omitempty"`
	SourceSystem          *string                    `json:"source_system,omitempty"`
	SourceRecordID        *string                    `json:"source_record_id,omitempty"`
	IdempotencyKey        *string                    `json:"idempotency_key,omitempty"`
	Items                 []FacilityOrderItemCreateDTO `json:"items" binding:"required,min=1"`
}

// FacilityOrderItemCreateDTO for order items
type FacilityOrderItemCreateDTO struct {
	ProductCode        string   `json:"product_code" binding:"required"`
	ProductDescription *string   `json:"product_description,omitempty"`
	UOM                *string   `json:"uom,omitempty"`
	OrderedQuantity    int      `json:"ordered_quantity" binding:"required"`
	UnitPrice           *float64 `json:"unit_price,omitempty"`
	Currency            *string   `json:"currency,omitempty"`
	Notes               *string   `json:"notes,omitempty"`
}

// FacilityOrderUpdateDTO for updating an order
type FacilityOrderUpdateDTO struct {
	OrderStatus         *string `json:"order_status,omitempty"`
	Priority            *string `json:"priority,omitempty"`
	ExpectedDeliveryDate *time.Time `json:"expected_delivery_date,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// FacilityOrderResponseDTO for API responses
type FacilityOrderResponseDTO struct {
	ID                  uint64                      `json:"id"`
	OrderNumber         string                      `json:"order_number"`
	OrderRefNumber      *string                     `json:"order_ref_number,omitempty"`
	FacilityID          uint64                      `json:"facility_id"`
	FacilityCode        string                      `json:"facility_code"`
	WarehouseID         uint64                      `json:"warehouse_id"`
	WarehouseCode       string                      `json:"warehouse_code"`
	OrderDate           time.Time                   `json:"order_date"`
	OrderType           *string                     `json:"order_type,omitempty"`
	OrderStatus         string                      `json:"order_status"`
	Priority            *string                     `json:"priority,omitempty"`
	FinancialYear       *string                    `json:"financial_year,omitempty"`
	OrderCycle          *string                    `json:"order_cycle,omitempty"`
	ProcurementPlanID   *uint64                    `json:"procurement_plan_id,omitempty"`
	SubmittedBy         *string                    `json:"submitted_by,omitempty"`
	SubmittedAt         *time.Time                 `json:"submitted_at,omitempty"`
	ApprovedBy          *string                    `json:"approved_by,omitempty"`
	ApprovedAt          *time.Time                 `json:"approved_at,omitempty"`
	ExpectedDeliveryDate *time.Time                `json:"expected_delivery_date,omitempty"`
	ActualDeliveryDate  *time.Time                 `json:"actual_delivery_date,omitempty"`
	TotalItems          int                        `json:"total_items"`
	TotalQuantity       int                        `json:"total_quantity"`
	TotalValue          *float64                   `json:"total_value,omitempty"`
	Notes               *string                    `json:"notes,omitempty"`
	SourceSystem        *string                    `json:"source_system,omitempty"`
	CreatedAt           time.Time                  `json:"created_at"`
	UpdatedAt           *time.Time                `json:"updated_at,omitempty"`
	Items               []FacilityOrderItemResponseDTO `json:"items,omitempty"`
	Deliveries          []FacilityDeliveryResponseDTO `json:"deliveries,omitempty"`
}

// FacilityOrderItemResponseDTO for order items
type FacilityOrderItemResponseDTO struct {
	ID                uint64     `json:"id"`
	OrderID           uint64     `json:"order_id"`
	ProductCode       string     `json:"product_code"`
	ProductDescription *string    `json:"product_description,omitempty"`
	UOM               *string     `json:"uom,omitempty"`
	OrderedQuantity   int        `json:"ordered_quantity"`
	HonoredQuantity   *int       `json:"honored_quantity,omitempty"`
	DeliveredQuantity *int       `json:"delivered_quantity,omitempty"`
	UnitPrice         *float64   `json:"unit_price,omitempty"`
	TotalPrice        *float64   `json:"total_price,omitempty"`
	Currency          *string    `json:"currency,omitempty"`
	Status            string     `json:"status"`
	Notes             *string    `json:"notes,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

// FacilityDeliveryCreateDTO for creating a delivery
type FacilityDeliveryCreateDTO struct {
	OrderID         uint64                        `json:"order_id" binding:"required"`
	DeliveryNumber  *string                       `json:"delivery_number,omitempty"`
	DeliveredAt     time.Time                     `json:"delivered_at" binding:"required"`
	DeliveryDate    time.Time                     `json:"delivery_date" binding:"required"`
	DeliveredBy     *string                       `json:"delivered_by,omitempty"`
	VehicleNumber   *string                       `json:"vehicle_number,omitempty"`
	DriverName      *string                       `json:"driver_name,omitempty"`
	Notes           *string                       `json:"notes,omitempty"`
	ConditionNotes  *string                       `json:"condition_notes,omitempty"`
	Items           []FacilityDeliveryItemCreateDTO `json:"items" binding:"required,min=1"`
}

// FacilityDeliveryItemCreateDTO for delivery items
type FacilityDeliveryItemCreateDTO struct {
	OrderItemID      *uint64   `json:"order_item_id,omitempty"`
	ProductCode      string    `json:"product_code" binding:"required"`
	ProductDescription *string  `json:"product_description,omitempty"`
	BatchNumber      *string   `json:"batch_number,omitempty"`
	ExpiryDate       *time.Time `json:"expiry_date,omitempty"`
	Quantity         int       `json:"quantity" binding:"required"`
	ReceivedQuantity *int      `json:"received_quantity,omitempty"`
	UnitPrice        *float64  `json:"unit_price,omitempty"`
	Condition        *string   `json:"condition,omitempty"`
	ConditionNotes   *string   `json:"condition_notes,omitempty"`
}

// FacilityDeliveryResponseDTO for API responses
type FacilityDeliveryResponseDTO struct {
	ID              uint64                          `json:"id"`
	OrderID         uint64                          `json:"order_id"`
	DeliveryRef     string                          `json:"delivery_ref"`
	DeliveryNumber  *string                         `json:"delivery_number,omitempty"`
	DeliveredAt     time.Time                       `json:"delivered_at"`
	DeliveryDate    time.Time                       `json:"delivery_date"`
	Status          string                          `json:"status"`
	DeliveredBy     *string                         `json:"delivered_by,omitempty"`
	ReceivedBy      *string                         `json:"received_by,omitempty"`
	ReceivedAt      *time.Time                      `json:"received_at,omitempty"`
	VehicleNumber   *string                         `json:"vehicle_number,omitempty"`
	DriverName      *string                         `json:"driver_name,omitempty"`
	TotalItems      int                             `json:"total_items"`
	TotalQuantity   int                             `json:"total_quantity"`
	Notes           *string                         `json:"notes,omitempty"`
	ConditionNotes  *string                         `json:"condition_notes,omitempty"`
	CreatedAt       time.Time                       `json:"created_at"`
	UpdatedAt       *time.Time                      `json:"updated_at,omitempty"`
	Items           []FacilityDeliveryItemResponseDTO `json:"items,omitempty"`
}

// FacilityDeliveryItemResponseDTO for delivery items
type FacilityDeliveryItemResponseDTO struct {
	ID                uint64     `json:"id"`
	DeliveryID        uint64     `json:"delivery_id"`
	OrderItemID       *uint64    `json:"order_item_id,omitempty"`
	ProductCode       string     `json:"product_code"`
	ProductDescription *string    `json:"product_description,omitempty"`
	BatchNumber       *string    `json:"batch_number,omitempty"`
	ExpiryDate        *time.Time `json:"expiry_date,omitempty"`
	Quantity          int        `json:"quantity"`
	ReceivedQuantity  *int       `json:"received_quantity,omitempty"`
	UnitPrice         *float64   `json:"unit_price,omitempty"`
	TotalPrice        *float64   `json:"total_price,omitempty"`
	Condition         *string    `json:"condition,omitempty"`
	ConditionNotes    *string    `json:"condition_notes,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}
