package models

import "time"

type PurchaseOrder struct {
	ID                 uint64    `gorm:"primaryKey" json:"id"`
	OrdSystemCode      string    `gorm:"size:100;not null" json:"ord_system_code"`
	OrdFacilityCode    string    `gorm:"size:100;not null" json:"ord_facility_code"`
	OrdTimestamp       time.Time `gorm:"not null" json:"ord_timestamp"`
	OrdOrderDate       time.Time `gorm:"not null" json:"ord_order_date"`
	OrdOrderRefNumber  string    `gorm:"size:100;not null" json:"ord_order_ref_number"`
	OrdOrderNumber     string    `gorm:"size:100;not null" json:"ord_order_number"`
	OrdProductCode     string    `gorm:"size:100;not null" json:"ord_product_code"`
	OrdOrderedQuantity int       `gorm:"not null" json:"ord_ordered_quantity"`
	BaseModel
}

// Extended fields for health commodity orders
type HealthCommodityOrder struct {
	ID                 uint64     `gorm:"primaryKey" json:"id"`
	OrderID            string     `gorm:"size:100;not null" json:"order_id"`
	OrderStatus        *string    `gorm:"size:50" json:"order_status,omitempty"`
	FacilityID         string     `gorm:"size:100;not null" json:"facility_id"`
	FacilityCode       *string    `gorm:"size:100" json:"facility_code,omitempty"`
	FacilityName       *string    `gorm:"size:255" json:"facility_name,omitempty"`
	LevelOfCare        *string    `gorm:"size:100" json:"level_of_care,omitempty"`
	District           *string    `gorm:"size:100" json:"district,omitempty"`
	Region             *string    `gorm:"size:100" json:"region,omitempty"`
	Zone               *string    `gorm:"size:100" json:"zone,omitempty"`
	FinancialYear      *string    `gorm:"size:50" json:"financial_year,omitempty"`
	ProductCode        string     `gorm:"size:100;not null" json:"product_code"`
	ProductDescription *string    `gorm:"size:255" json:"product_description,omitempty"`
	UOM                *string    `gorm:"size:50" json:"uom,omitempty"`
	OrderType          *string    `gorm:"size:50" json:"order_type,omitempty"`
	DateCreated        time.Time  `json:"date_created"`
	OrderCycle         *string    `gorm:"size:50" json:"order_cycle,omitempty"`
	PeriodOutOfStock   *int       `json:"period_out_of_stock,omitempty"`
	OpeningBalance     *int       `json:"opening_balance,omitempty"`
	QuantityReceived   *int       `json:"quantity_received,omitempty"`
	QuantityConsumed   *int       `json:"quantity_consumed,omitempty"`
	AdjustedAmc        *float64   `json:"adjusted_amc,omitempty"`
	Adjustments        *int       `json:"adjustments,omitempty"`
	ClosingBalance     *int       `json:"closing_balance,omitempty"`
	BatchNumber        *string    `gorm:"size:100" json:"batch_number,omitempty"`
	ExpiryDate         *time.Time `json:"expiry_date,omitempty"`
	InvoiceNumber      *string    `gorm:"size:100" json:"invoice_number,omitempty"`
	GrnNumber          *string    `gorm:"size:100" json:"grn_number,omitempty"`
	SupplierID         *string    `gorm:"size:100" json:"supplier_id,omitempty"`
	DeliveryETA        *time.Time `json:"delivery_eta,omitempty"`
	ActualDeliveryDate *time.Time `json:"actual_delivery_date,omitempty"`
	IdempotencyKey     *string    `gorm:"size:100" json:"idempotency_key,omitempty"`
	CorrectionOf       *string    `gorm:"size:100" json:"correction_of,omitempty"`
	SourceSystem       *string    `gorm:"size:100" json:"source_system,omitempty"`
	SourceRecordID     *string    `gorm:"size:100" json:"source_record_id,omitempty"`
	SubmittedBy        *string    `gorm:"size:100" json:"submitted_by,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

func (PurchaseOrder) TableName() string {
	return "Stg_Purchase_Order"
}
