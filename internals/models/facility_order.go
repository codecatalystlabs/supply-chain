package models

import "time"

// FacilityOrder represents an order placed by a facility to a warehouse
type FacilityOrder struct {
	ID             uint64  `gorm:"primaryKey" json:"id"`
	OrderNumber    string  `gorm:"size:100;uniqueIndex;not null" json:"order_number"`
	OrderRefNumber *string `gorm:"size:100;uniqueIndex" json:"order_ref_number,omitempty"` // External reference

	// Facility information
	FacilityID   uint64 `gorm:"not null;index" json:"facility_id"`
	FacilityCode string `gorm:"size:100;not null;index" json:"facility_code"`

	// Warehouse information
	WarehouseID   uint64 `gorm:"not null;index" json:"warehouse_id"`
	WarehouseCode string `gorm:"size:100;not null;index" json:"warehouse_code"`

	// Order details
	OrderDate   time.Time `gorm:"not null;index" json:"order_date"`
	OrderType   *string   `gorm:"size:50" json:"order_type,omitempty"`                          // e.g., routine, emergency, replenishment
	OrderStatus string    `gorm:"size:50;not null;default:'pending';index" json:"order_status"` // pending, approved, sent, processing, fulfilled, cancelled
	Priority    *string   `gorm:"size:50" json:"priority,omitempty"`                            // low, normal, high, urgent

	// Financial year and cycle
	FinancialYear *string `gorm:"size:50" json:"financial_year,omitempty"`
	OrderCycle    *string `gorm:"size:50" json:"order_cycle,omitempty"`

	// Procurement plan reference
	ProcurementPlanID *uint64 `gorm:"index" json:"procurement_plan_id,omitempty"`

	// Approval workflow
	SubmittedBy *string    `gorm:"size:255" json:"submitted_by,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	ApprovedBy  *string    `gorm:"size:255" json:"approved_by,omitempty"`
	ApprovedAt  *time.Time `json:"approved_at,omitempty"`

	// Delivery information
	ExpectedDeliveryDate *time.Time `json:"expected_delivery_date,omitempty"`
	ActualDeliveryDate   *time.Time `json:"actual_delivery_date,omitempty"`

	// Totals
	TotalItems    int      `gorm:"default:0" json:"total_items"`
	TotalQuantity int      `gorm:"default:0" json:"total_quantity"`
	TotalValue    *float64 `json:"total_value,omitempty"`

	// Notes and metadata
	Notes          *string `gorm:"type:text" json:"notes,omitempty"`
	SourceSystem   *string `gorm:"size:100" json:"source_system,omitempty"` // EMR system identifier
	SourceRecordID *string `gorm:"size:100" json:"source_record_id,omitempty"`
	IdempotencyKey *string `gorm:"size:100;uniqueIndex" json:"idempotency_key,omitempty"`

	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Facility        Facility            `gorm:"foreignKey:FacilityID" json:"facility,omitempty"`
	Warehouse       Warehouse           `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	ProcurementPlan *ProcurementPlan    `gorm:"foreignKey:ProcurementPlanID" json:"procurement_plan,omitempty"`
	OrderItems      []FacilityOrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order_items,omitempty"`
	Deliveries      []FacilityDelivery  `gorm:"foreignKey:OrderID" json:"deliveries,omitempty"`
}

func (FacilityOrder) TableName() string {
	return "facility_orders"
}

// FacilityOrderItem represents a line item in a facility order
type FacilityOrderItem struct {
	ID                 uint64  `gorm:"primaryKey" json:"id"`
	OrderID            uint64  `gorm:"not null;index" json:"order_id"`
	ProductCode        string  `gorm:"size:100;not null;index" json:"product_code"`
	ProductDescription *string `gorm:"size:255" json:"product_description,omitempty"`
	UOM                *string `gorm:"size:50" json:"uom,omitempty"`

	// Quantities
	OrderedQuantity   int  `gorm:"not null" json:"ordered_quantity"`
	HonoredQuantity   *int `json:"honored_quantity,omitempty"`   // Quantity actually provided by warehouse
	DeliveredQuantity *int `json:"delivered_quantity,omitempty"` // Quantity actually delivered

	// Pricing
	UnitPrice  *float64 `json:"unit_price,omitempty"`
	TotalPrice *float64 `json:"total_price,omitempty"`
	Currency   *string  `gorm:"size:10" json:"currency,omitempty"`

	// Status
	Status string `gorm:"size:50;default:'pending'" json:"status"` // pending, partially_fulfilled, fulfilled, cancelled

	// Notes
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Order FacilityOrder `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}

func (FacilityOrderItem) TableName() string {
	return "facility_order_items"
}

// FacilityDelivery represents a delivery from warehouse to facility
type FacilityDelivery struct {
	ID             uint64  `gorm:"primaryKey" json:"id"`
	OrderID        uint64  `gorm:"not null;index" json:"order_id"`
	DeliveryRef    string  `gorm:"size:100;uniqueIndex;not null" json:"delivery_ref"`
	DeliveryNumber *string `gorm:"size:100;uniqueIndex" json:"delivery_number,omitempty"`

	// Delivery details
	DeliveredAt  time.Time `gorm:"not null;index" json:"delivered_at"`
	DeliveryDate time.Time `gorm:"not null" json:"delivery_date"`
	Status       string    `gorm:"size:50;not null;default:'pending';index" json:"status"` // pending, in_transit, delivered, received, cancelled

	// Delivery personnel
	DeliveredBy *string    `gorm:"size:255" json:"delivered_by,omitempty"`
	ReceivedBy  *string    `gorm:"size:255" json:"received_by,omitempty"`
	ReceivedAt  *time.Time `json:"received_at,omitempty"`

	// Vehicle/Transport
	VehicleNumber *string `gorm:"size:100" json:"vehicle_number,omitempty"`
	DriverName    *string `gorm:"size:255" json:"driver_name,omitempty"`

	// Totals
	TotalItems    int `gorm:"default:0" json:"total_items"`
	TotalQuantity int `gorm:"default:0" json:"total_quantity"`

	// Notes
	Notes          *string `gorm:"type:text" json:"notes,omitempty"`
	ConditionNotes *string `gorm:"type:text" json:"condition_notes,omitempty"` // Notes on condition of goods

	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Order         FacilityOrder          `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	DeliveryItems []FacilityDeliveryItem `gorm:"foreignKey:DeliveryID;constraint:OnDelete:CASCADE" json:"delivery_items,omitempty"`
}

func (FacilityDelivery) TableName() string {
	return "facility_deliveries"
}

// FacilityDeliveryItem represents a line item in a delivery
type FacilityDeliveryItem struct {
	ID          uint64  `gorm:"primaryKey" json:"id"`
	DeliveryID  uint64  `gorm:"not null;index" json:"delivery_id"`
	OrderItemID *uint64 `gorm:"index" json:"order_item_id,omitempty"` // Link to original order item

	ProductCode        string     `gorm:"size:100;not null;index" json:"product_code"`
	ProductDescription *string    `gorm:"size:255" json:"product_description,omitempty"`
	BatchNumber        *string    `gorm:"size:100" json:"batch_number,omitempty"`
	ExpiryDate         *time.Time `json:"expiry_date,omitempty"`

	Quantity         int  `gorm:"not null" json:"quantity"`
	ReceivedQuantity *int `json:"received_quantity,omitempty"` // Actual quantity received (may differ due to damage, etc.)

	UnitPrice  *float64 `json:"unit_price,omitempty"`
	TotalPrice *float64 `json:"total_price,omitempty"`

	// Condition
	Condition      *string `gorm:"size:50" json:"condition,omitempty"` // good, damaged, expired
	ConditionNotes *string `gorm:"type:text" json:"condition_notes,omitempty"`

	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Relationships
	Delivery  FacilityDelivery   `gorm:"foreignKey:DeliveryID" json:"delivery,omitempty"`
	OrderItem *FacilityOrderItem `gorm:"foreignKey:OrderItemID" json:"order_item,omitempty"`
}

func (FacilityDeliveryItem) TableName() string {
	return "facility_delivery_items"
}
