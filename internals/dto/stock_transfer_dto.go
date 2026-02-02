package dto

import "time"

// StockTransferCreateDTO for creating a new stock transfer
type StockTransferCreateDTO struct {
	TransferType   string     `json:"transfer_type" binding:"required,oneof=inter_facility intra_facility"`
	FromFacilityID uint64     `json:"from_facility_id" binding:"required"`
	FromPharmacyID *uint64    `json:"from_pharmacy_id,omitempty"` // Required for intra_facility
	ToFacilityID   uint64     `json:"to_facility_id" binding:"required"`
	ToPharmacyID   *uint64    `json:"to_pharmacy_id,omitempty"` // Required for intra_facility
	ProductCode    string     `json:"product_code" binding:"required"`
	BatchNumber    *string    `json:"batch_number,omitempty"`
	Quantity       int        `json:"quantity" binding:"required,gt=0"`
	ExpiryDate     *time.Time `json:"expiry_date,omitempty"`
	TransferDate   time.Time  `json:"transfer_date" binding:"required"`
	RequestedBy    *string    `json:"requested_by,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
}

// StockTransferUpdateDTO for updating a transfer
type StockTransferUpdateDTO struct {
	Status     *string    `json:"status,omitempty" binding:"omitempty,oneof=pending in_transit completed cancelled"`
	ApprovedBy *string    `json:"approved_by,omitempty"`
	ReceivedBy *string    `json:"received_by,omitempty"`
	ReceivedAt *time.Time `json:"received_at,omitempty"`
	Notes      *string    `json:"notes,omitempty"`
}

// StockTransferResponseDTO for API responses
type StockTransferResponseDTO struct {
	ID               uint64     `json:"id"`
	TransferRef      string     `json:"transfer_ref"`
	TransferType     string     `json:"transfer_type"`
	FromFacilityID   uint64     `json:"from_facility_id"`
	FromFacilityCode string     `json:"from_facility_code,omitempty"`
	FromPharmacyID   *uint64    `json:"from_pharmacy_id,omitempty"`
	FromPharmacyCode *string    `json:"from_pharmacy_code,omitempty"`
	ToFacilityID     uint64     `json:"to_facility_id"`
	ToFacilityCode   string     `json:"to_facility_code,omitempty"`
	ToPharmacyID     *uint64    `json:"to_pharmacy_id,omitempty"`
	ToPharmacyCode   *string    `json:"to_pharmacy_code,omitempty"`
	ProductCode      string     `json:"product_code"`
	BatchNumber      *string    `json:"batch_number,omitempty"`
	Quantity         int        `json:"quantity"`
	ExpiryDate       *time.Time `json:"expiry_date,omitempty"`
	TransferDate     time.Time  `json:"transfer_date"`
	Status           string     `json:"status"`
	RequestedBy      *string    `json:"requested_by,omitempty"`
	ApprovedBy       *string    `json:"approved_by,omitempty"`
	ReceivedBy       *string    `json:"received_by,omitempty"`
	ReceivedAt       *time.Time `json:"received_at,omitempty"`
	Notes            *string    `json:"notes,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}
