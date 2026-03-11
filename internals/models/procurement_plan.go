package models

import "time"

type ProcurementPlan struct {
	ID             uint64    `gorm:"primaryKey" json:"id"`
	PlanSystemCode string    `gorm:"size:100;not null" json:"plan_system_code"`
	WarehouseID    *uint64   `gorm:"index" json:"warehouse_id,omitempty"` // Link to warehouse
	StoreCode      string    `gorm:"size:100;not null;index" json:"store_code"` // Legacy field, maps to warehouse code
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	Notes          *string   `gorm:"size:255" json:"notes,omitempty"`
	
	// Relationships
	Warehouse      *Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	Items          []ProcurementPlanItem `gorm:"foreignKey:ProcurementID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
	// Extended fields for annual procurement
	FacilityID                  string     `gorm:"size:100" json:"facility_id,omitempty"`
	FacilityCode                *string    `gorm:"size:100" json:"facility_code,omitempty"`
	FacilityName                string     `gorm:"size:255" json:"facility_name"`
	LevelOfCare                 *string    `gorm:"size:100" json:"level_of_care,omitempty"`
	District                    *string    `gorm:"size:100" json:"district,omitempty"`
	Region                      *string    `gorm:"size:100" json:"region,omitempty"`
	Zone                        *string    `gorm:"size:100" json:"zone,omitempty"`
	ProductCode                 string     `gorm:"size:100" json:"product_code"`
	ProductDescription          string     `gorm:"size:255" json:"product_description"`
	UnitOfMeasure               *string    `gorm:"size:50" json:"unit_of_measure,omitempty"`
	Section                     *string    `gorm:"size:100" json:"section,omitempty"`
	SubSection                  *string    `gorm:"size:100" json:"sub_section,omitempty"`
	VenClassification           *string    `gorm:"size:100" json:"ven_classification,omitempty"`
	UnitPrice                   *float64   `json:"unit_price,omitempty"`
	Currency                    *string    `gorm:"size:10" json:"currency,omitempty"`
	PreviousBiMonthlyPlannedQty *int       `json:"previous_bi_monthly_planned_qty,omitempty"`
	PastAvgNmsIssuePlanQty     *float64   `json:"past_avg_nms_issue_plan_qty,omitempty"`
	AverageMonthlyConsumption   *float64   `json:"average_monthly_consumption,omitempty"`
	AverageDaysOutOfStock       *int       `json:"average_days_out_of_stock,omitempty"`
	AdjustedAmc                 *float64   `json:"adjusted_amc,omitempty"`
	BiMonthlyPlanQty            *int       `json:"bi_monthly_plan_qty,omitempty"`
	Comment                     *string    `gorm:"size:500" json:"comment,omitempty"`
	FundedQty                   *int       `json:"funded_qty,omitempty"`
	// Budget summary (often same for all rows in a facility plan)
	IndicativeAnnualBudget       *float64   `json:"indicative_annual_budget,omitempty"`
	CalculatedAnnualProcurement  *float64   `json:"calculated_annual_procurement,omitempty"`
	IndicativeBiMonthlyBudget    *float64   `json:"indicative_bi_monthly_budget,omitempty"`
	CalculatedBiMonthlyProcurement *float64  `json:"calculated_bi_monthly_procurement,omitempty"`
	RemainingBudget              *float64   `json:"remaining_budget,omitempty"`
	PercentBudgetRemaining       *float64   `json:"percent_budget_remaining,omitempty"`
	FinancialYear               string     `gorm:"size:50" json:"financial_year"`
	PlanPeriodType              *string    `gorm:"size:50" json:"plan_period_type,omitempty"`
	PlanPeriodStart             *time.Time `json:"plan_period_start,omitempty"`
	PlanPeriodEnd               *time.Time `json:"plan_period_end,omitempty"`
	PlanVersion                 *string    `gorm:"size:50" json:"plan_version,omitempty"`
	ApprovalStatus              *string    `gorm:"size:50" json:"approval_status,omitempty"`
	SourceSystem                *string    `gorm:"size:100" json:"source_system,omitempty"`
	SourceRecordID              *string    `gorm:"size:100" json:"source_record_id,omitempty"`
	SubmittedBy                 *string    `gorm:"size:100" json:"submitted_by,omitempty"`
	UpdatedAt                   *time.Time `json:"updated_at,omitempty"`
}

func (ProcurementPlan) TableName() string {
	return "Procurement_Plan"
}

type ProcurementPlanItem struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	ProcurementID uint64    `gorm:"not null" json:"procurement_id"`
	ProductCode   string    `gorm:"size:100;not null" json:"product_code"`
	Quantity      int       `gorm:"not null" json:"quantity"`
	NeededBy      time.Time `json:"needed_by"`
	Status        string    `gorm:"size:50" json:"status"` // e.g., planned/ordered/received
	// Optional descriptive fields
	ProductDescription *string `gorm:"size:255" json:"product_description,omitempty"`
	UOM                *string `gorm:"size:50" json:"uom,omitempty"`
}

func (ProcurementPlanItem) TableName() string {
	return "Procurement_Plan_Item"
}
