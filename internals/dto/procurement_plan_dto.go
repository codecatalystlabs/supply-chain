package dto

import "time"

type ProcurementPlanItemCreateDTO struct {
	ProductCode string    `json:"product_code" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required"`
	NeededBy    time.Time `json:"needed_by" binding:"required"`
}

type ProcurementPlanCreateDTO struct {
	PlanSystemCode string                         `json:"plan_system_code" binding:"required"`
	StoreCode      string                         `json:"store_code" binding:"required"`
	FacilityID     string                         `json:"facility_id" binding:"required"`
	FinancialYear  string                         `json:"financial_year" binding:"required"`
	FacilityCode   *string                        `json:"facility_code,omitempty"`
	FacilityName   *string                        `json:"facility_name,omitempty"`
	LevelOfCare    *string                        `json:"level_of_care,omitempty"`
	District       *string                        `json:"district,omitempty"`
	Region         *string                        `json:"region,omitempty"`
	Zone           *string                        `json:"zone,omitempty"`
	PlanPeriodType *string                        `json:"plan_period_type,omitempty"`
	PlanPeriodStart *time.Time                    `json:"plan_period_start,omitempty"`
	PlanPeriodEnd   *time.Time                    `json:"plan_period_end,omitempty"`
	Notes           *string                       `json:"notes,omitempty"`
	Items           []ProcurementPlanItemCreateDTO `json:"items" binding:"required,dive"`
}

type ProcurementPlanItemResponseDTO struct {
	ID            uint64    `json:"id"`
	ProcurementID uint64    `json:"procurement_id"`
	ProductCode   string    `json:"product_code"`
	Quantity      int       `json:"quantity"`
	NeededBy      time.Time `json:"needed_by"`
	Status        string    `json:"status"`
}

type ProcurementPlanResponseDTO struct {
	ID             uint64                           `json:"id"`
	PlanSystemCode string                           `json:"plan_system_code"`
	StoreCode      string                           `json:"store_code"`
	FacilityID     string                           `json:"facility_id"`
	FacilityCode   *string                          `json:"facility_code,omitempty"`
	FacilityName   string                           `json:"facility_name"`
	LevelOfCare    *string                          `json:"level_of_care,omitempty"`
	District       *string                          `json:"district,omitempty"`
	Region         *string                          `json:"region,omitempty"`
	Zone           *string                          `json:"zone,omitempty"`
	ProductCode    string                           `json:"product_code"`
	ProductDescription string                       `json:"product_description"`
	UnitOfMeasure  *string                          `json:"unit_of_measure,omitempty"`
	Section        *string                          `json:"section,omitempty"`
	SubSection     *string                          `json:"sub_section,omitempty"`
	VenClassification *string                       `json:"ven_classification,omitempty"`
	UnitPrice      *float64                         `json:"unit_price,omitempty"`
	Currency       *string                          `json:"currency,omitempty"`
	PreviousBiMonthlyPlannedQty *int                `json:"previous_bi_monthly_planned_qty,omitempty"`
	PastAvgNmsIssuePlanQty     *float64             `json:"past_avg_nms_issue_plan_qty,omitempty"`
	AverageMonthlyConsumption   *float64            `json:"average_monthly_consumption,omitempty"`
	AverageDaysOutOfStock       *int                `json:"average_days_out_of_stock,omitempty"`
	AdjustedAmc                 *float64            `json:"adjusted_amc,omitempty"`
	BiMonthlyPlanQty            *int               `json:"bi_monthly_plan_qty,omitempty"`
	Comment        *string                          `json:"comment,omitempty"`
	FundedQty      *int                             `json:"funded_qty,omitempty"`
	IndicativeAnnualBudget       *float64           `json:"indicative_annual_budget,omitempty"`
	CalculatedAnnualProcurement  *float64            `json:"calculated_annual_procurement,omitempty"`
	IndicativeBiMonthlyBudget   *float64           `json:"indicative_bi_monthly_budget,omitempty"`
	CalculatedBiMonthlyProcurement *float64         `json:"calculated_bi_monthly_procurement,omitempty"`
	RemainingBudget              *float64           `json:"remaining_budget,omitempty"`
	PercentBudgetRemaining       *float64           `json:"percent_budget_remaining,omitempty"`
	FinancialYear  string                           `json:"financial_year"`
	PlanPeriodType *string                          `json:"plan_period_type,omitempty"`
	PlanPeriodStart *time.Time                      `json:"plan_period_start,omitempty"`
	PlanPeriodEnd   *time.Time                      `json:"plan_period_end,omitempty"`
	CreatedAt      time.Time                        `json:"created_at"`
	Notes          *string                          `json:"notes,omitempty"`
	Items          []ProcurementPlanItemResponseDTO `json:"items,omitempty"`
}
