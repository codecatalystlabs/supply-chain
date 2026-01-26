package dto

import "time"

type ProductAmcCreateDTO struct {
	AmcSystemCode   string    `json:"amc_system_code" binding:"required"`
	AmcFacilityCode string    `json:"amc_facility_code" binding:"required"`
	AmcDate         time.Time `json:"amc_date" binding:"required"`
	AmcProductCode  string    `json:"amc_product_code" binding:"required"`
	AmcProductName  string    `json:"amc_product_name" binding:"required"`
	AmcMonth        int16     `json:"amc_month" binding:"required"`
	AmcYear         int16     `json:"amc_year" binding:"required"`
	AmcValue        float64   `json:"amc_value" binding:"required"`
}

type ProductAmcUpdateDTO struct {
	AmcValue *float64 `json:"amc_value,omitempty" binding:"omitempty"`
}

type ProductAmcResponseDTO struct {
	ID               uint64    `json:"id"`
	AmcSystemCode    string    `json:"amc_system_code"`
	AmcFacilityCode  string    `json:"amc_facility_code"`
	AmcTimestamp     time.Time `json:"amc_timestamp"`
	AmcProductCode   string    `json:"amc_product_code"`
	AmcProductName   string    `json:"amc_product_name"`
	AmcDate          time.Time `json:"amc_date"`
	AmcMonth         int16     `json:"amc_month"`
	AmcYear          int16     `json:"amc_year"`
	AmcValue         float64   `json:"amc_value"`
	AmcDaysOutStock  *float64  `json:"amc_days_out_stock,omitempty"`
	ValidationStatus int16     `json:"validation_status"`
	SyncStatus       int16     `json:"sync_status"`
}
