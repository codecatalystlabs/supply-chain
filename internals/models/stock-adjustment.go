package models

import "time"

// StockAdjustment represents inventory adjustments (damage, loss, found, etc.)
// Can be at facility level or pharmacy level
type StockAdjustment struct {
	ID                  uint64     `gorm:"primaryKey" json:"id"`
	AdjSystemCode       string     `gorm:"size:100;not null" json:"adj_system_code"`
	AdjFacilityCode     string     `gorm:"size:100;not null;index" json:"adj_facility_code"`
	AdjPharmacyID       *uint64    `gorm:"index" json:"adj_pharmacy_id,omitempty"` // If adjustment is at pharmacy level
	AdjTimestamp        time.Time  `gorm:"not null" json:"adj_timestamp"`
	AdjAdjustmentDate   time.Time  `gorm:"not null" json:"adj_adjustment_date"`
	AdjAdjustmentType   string     `gorm:"size:100;not null" json:"adj_adjustment_type"` // e.g., damage, loss, found, expiry, theft
	AdjAdjustmentReason string     `gorm:"size:255;not null" json:"adj_adjustment_reason"`
	AdjProductCode      string     `gorm:"size:100;not null;index" json:"adj_product_code"`
	AdjBatchNumber       string     `gorm:"size:100;not null" json:"adj_batch_number"`
	AdjQuantity          int        `gorm:"not null" json:"adj_quantity"` // Can be negative (loss) or positive (found)
	AdjExpiryDate        time.Time  `gorm:"not null" json:"adj_expiry_date"`
	AdjReferenceNumber   *string    `gorm:"size:100" json:"adj_reference_number,omitempty"` // Reference to physical count document
	AdjApprovedBy        *string    `gorm:"size:255" json:"adj_approved_by,omitempty"`
	AdjNotes             *string    `gorm:"type:text" json:"adj_notes,omitempty"`
	BaseModel
	
	// Relationships
	Pharmacy             *Pharmacy  `gorm:"foreignKey:AdjPharmacyID" json:"pharmacy,omitempty"`
}

func (StockAdjustment) TableName() string {
	return "stock_adjustments"
}

// StockTransfer represents stock movements between facilities or between pharmacies within the same facility
type StockTransfer struct {
	ID                  uint64     `gorm:"primaryKey" json:"id"`
	TransferRef         string     `gorm:"size:100;uniqueIndex;not null" json:"transfer_ref"`
	TransferType        string     `gorm:"size:50;not null;index" json:"transfer_type"` // "inter_facility" or "intra_facility"
	
	// Source location
	FromFacilityID      uint64     `gorm:"not null;index" json:"from_facility_id"`
	FromPharmacyID      *uint64    `gorm:"index" json:"from_pharmacy_id,omitempty"` // Required for intra-facility transfers
	
	// Destination location
	ToFacilityID        uint64     `gorm:"not null;index" json:"to_facility_id"`
	ToPharmacyID        *uint64    `gorm:"index" json:"to_pharmacy_id,omitempty"` // Required for intra-facility transfers
	
	// Product details
	ProductCode         string     `gorm:"size:100;not null;index" json:"product_code"`
	BatchNumber         *string    `gorm:"size:100" json:"batch_number,omitempty"`
	Quantity            int        `gorm:"not null" json:"quantity"`
	ExpiryDate          *time.Time `json:"expiry_date,omitempty"`
	
	// Transfer details
	TransferDate        time.Time  `gorm:"not null" json:"transfer_date"`
	Status              string     `gorm:"size:50;not null;default:'pending';index" json:"status"` // pending, in_transit, completed, cancelled
	RequestedBy         *string    `gorm:"size:255" json:"requested_by,omitempty"`
	ApprovedBy          *string    `gorm:"size:255" json:"approved_by,omitempty"`
	ReceivedBy          *string    `gorm:"size:255" json:"received_by,omitempty"`
	ReceivedAt          *time.Time `json:"received_at,omitempty"`
	Notes               *string    `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt           time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	
	// Relationships
	FromFacility        Facility   `gorm:"foreignKey:FromFacilityID" json:"from_facility,omitempty"`
	ToFacility          Facility   `gorm:"foreignKey:ToFacilityID" json:"to_facility,omitempty"`
	FromPharmacy        *Pharmacy  `gorm:"foreignKey:FromPharmacyID" json:"from_pharmacy,omitempty"`
	ToPharmacy          *Pharmacy  `gorm:"foreignKey:ToPharmacyID" json:"to_pharmacy,omitempty"`
}

func (StockTransfer) TableName() string {
	return "stock_transfers"
}

// Legacy model for backward compatibility - maps to StockTransfer
type InterFacilityTransfer struct {
	ID               uint64    `gorm:"primaryKey" json:"id"`
	TransferRef      string    `gorm:"size:100;not null" json:"transfer_ref"`
	FromFacilityCode string    `gorm:"size:100;not null" json:"from_facility_code"`
	ToFacilityCode   string    `gorm:"size:100;not null" json:"to_facility_code"`
	ProductCode      string    `gorm:"size:100;not null" json:"product_code"`
	BatchNumber      *string   `gorm:"size:100" json:"batch_number,omitempty"`
	Quantity         int       `gorm:"not null" json:"quantity"`
	TransferDate     time.Time `gorm:"not null" json:"transfer_date"`
	Status           string    `gorm:"size:50" json:"status"` // e.g., pending, completed
	Notes            *string   `gorm:"size:255" json:"notes,omitempty"`
	BaseModel
}

func (InterFacilityTransfer) TableName() string {
	return "inter_facility_transfers"
}
