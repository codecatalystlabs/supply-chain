package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"supply-chain/internals/config"
	"supply-chain/internals/dto"
	"supply-chain/internals/models"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// @Summary Create procurement plan
// @Tags ProcurementPlan
// @Accept json
// @Produce json
// @Param payload body dto.ProcurementPlanCreateDTO true "Procurement plan payload"
// @Success 201 {object} dto.ProcurementPlanResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /procurement [post]
func CreateProcurementPlan(c *gin.Context) {
	var payload dto.ProcurementPlanCreateDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan := models.ProcurementPlan{
		PlanSystemCode:  payload.PlanSystemCode,
		StoreCode:       payload.StoreCode,
		FacilityID:      payload.FacilityID,
		FacilityCode:    payload.FacilityCode,
		FacilityName:    valueOrEmpty(payload.FacilityName),
		LevelOfCare:     payload.LevelOfCare,
		District:        payload.District,
		Region:          payload.Region,
		Zone:            payload.Zone,
		FinancialYear:   payload.FinancialYear,
		PlanPeriodType:  payload.PlanPeriodType,
		PlanPeriodStart: payload.PlanPeriodStart,
		PlanPeriodEnd:   payload.PlanPeriodEnd,
		CreatedAt:       time.Now(),
		Notes:           payload.Notes,
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create items
	for _, it := range payload.Items {
		item := models.ProcurementPlanItem{
			ProcurementID: plan.ID,
			ProductCode:   it.ProductCode,
			Quantity:      it.Quantity,
			NeededBy:      it.NeededBy,
			Status:        "planned",
		}
		_ = config.DB.Create(&item)
	}

	// load items
	var items []models.ProcurementPlanItem
	config.DB.Where("procurement_id = ?", plan.ID).Find(&items)

	resp := mapToProcurementPlanResponse(plan, items)
	c.JSON(http.StatusCreated, resp)
}

// @Summary List procurement plans
// @Tags ProcurementPlan
// @Produce json
// @Success 200 {array} dto.ProcurementPlanResponseDTO
// @Failure 500 {object} map[string]string
// @Router /procurement [get]
func ListProcurementPlans(c *gin.Context) {
	var plans []models.ProcurementPlan
	query := config.DB

	// Scope by logged-in user if facility-scoped (plan stores facility_id as code; user has facility id)
	if user, exists := c.Get("user"); exists {
		u := user.(*models.User)
		if u.FacilityID != nil && !u.HasRole("super_admin") {
			var fac models.Facility
			if err := config.DB.Select("facility_code").First(&fac, *u.FacilityID).Error; err == nil {
				query = query.Where("facility_id = ?", fac.FacilityCode)
			} else {
				query = query.Where("facility_id = ?", *u.FacilityID)
			}
		}
	}

	// Optional explicit filters for facility, district, region, year (case-insensitive where applicable)
	if fid := c.Query("facility_id"); fid != "" {
		query = query.Where("facility_id = ?", fid)
	}
	if district := c.Query("district"); district != "" {
		// Match plan's district (case-insensitive) or plans whose facility is in this district
		query = query.Where(
			"(district IS NOT NULL AND LOWER(TRIM(district)) = LOWER(TRIM(?))) OR facility_id IN (SELECT facility_code FROM facilities WHERE district IS NOT NULL AND LOWER(TRIM(district)) = LOWER(TRIM(?)))",
			district, district,
		)
	}
	if region := c.Query("region"); region != "" {
		// Match plan's region (case-insensitive) or plans whose facility is in this region
		query = query.Where(
			"(region IS NOT NULL AND LOWER(TRIM(region)) = LOWER(TRIM(?))) OR facility_id IN (SELECT facility_code FROM facilities WHERE region IS NOT NULL AND LOWER(TRIM(region)) = LOWER(TRIM(?)))",
			region, region,
		)
	}
	if fy := c.Query("financial_year"); fy != "" {
		query = query.Where("financial_year = ?", fy)
	}
	if level := c.Query("level_of_care"); level != "" {
		query = query.Where("(level_of_care IS NOT NULL AND LOWER(TRIM(level_of_care)) = LOWER(TRIM(?))) OR facility_id IN (SELECT facility_code FROM facilities WHERE level_of_care IS NOT NULL AND LOWER(TRIM(level_of_care)) = LOWER(TRIM(?)))", level, level)
	}

	if err := query.Find(&plans).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.ProcurementPlanResponseDTO
	for _, p := range plans {
		var items []models.ProcurementPlanItem
		config.DB.Where("procurement_id = ?", p.ID).Find(&items)
		resp = append(resp, mapToProcurementPlanResponse(p, items))
	}
	c.JSON(200, resp)
}

// @Summary Get procurement plan
// @Tags ProcurementPlan
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} dto.ProcurementPlanResponseDTO
// @Failure 404 {object} map[string]string
// @Router /procurement/{id} [get]
func GetProcurementPlan(c *gin.Context) {
	id := c.Param("id")
	var plan models.ProcurementPlan
	if err := config.DB.First(&plan, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Plan not found"})
		return
	}
	var items []models.ProcurementPlanItem
	config.DB.Where("procurement_id = ?", plan.ID).Find(&items)
	c.JSON(200, mapToProcurementPlanResponse(plan, items))
}

// @Summary Delete procurement plan
// @Tags ProcurementPlan
// @Param id path int true "ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /procurement/{id} [delete]
func DeleteProcurementPlan(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.ProcurementPlan{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	config.DB.Where("procurement_id = ?", id).Delete(&models.ProcurementPlanItem{})
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}

func mapToProcurementPlanResponse(p models.ProcurementPlan, items []models.ProcurementPlanItem) dto.ProcurementPlanResponseDTO {
	var itResp []dto.ProcurementPlanItemResponseDTO
	for _, it := range items {
		itResp = append(itResp, dto.ProcurementPlanItemResponseDTO{
			ID:            it.ID,
			ProcurementID: it.ProcurementID,
			ProductCode:   it.ProductCode,
			Quantity:      it.Quantity,
			NeededBy:      it.NeededBy,
			Status:        it.Status,
		})
	}
	return dto.ProcurementPlanResponseDTO{
		ID:              p.ID,
		PlanSystemCode:  p.PlanSystemCode,
		StoreCode:       p.StoreCode,
		FacilityID:      p.FacilityID,
		FacilityCode:    p.FacilityCode,
		FacilityName:    p.FacilityName,
		LevelOfCare:     p.LevelOfCare,
		District:        p.District,
		Region:          p.Region,
		Zone:            p.Zone,
		ProductCode:     p.ProductCode,
		ProductDescription: p.ProductDescription,
		UnitOfMeasure:   p.UnitOfMeasure,
		Section:         p.Section,
		SubSection:      p.SubSection,
		VenClassification: p.VenClassification,
		UnitPrice:       p.UnitPrice,
		Currency:        p.Currency,
		PreviousBiMonthlyPlannedQty: p.PreviousBiMonthlyPlannedQty,
		PastAvgNmsIssuePlanQty:      p.PastAvgNmsIssuePlanQty,
		AverageMonthlyConsumption:   p.AverageMonthlyConsumption,
		AverageDaysOutOfStock:       p.AverageDaysOutOfStock,
		AdjustedAmc:                 p.AdjustedAmc,
		BiMonthlyPlanQty:            p.BiMonthlyPlanQty,
		Comment:         p.Comment,
		FundedQty:       p.FundedQty,
		IndicativeAnnualBudget:      p.IndicativeAnnualBudget,
		CalculatedAnnualProcurement: p.CalculatedAnnualProcurement,
		IndicativeBiMonthlyBudget:   p.IndicativeBiMonthlyBudget,
		CalculatedBiMonthlyProcurement: p.CalculatedBiMonthlyProcurement,
		RemainingBudget:             p.RemainingBudget,
		PercentBudgetRemaining:      p.PercentBudgetRemaining,
		FinancialYear:   p.FinancialYear,
		PlanPeriodType:  p.PlanPeriodType,
		PlanPeriodStart: p.PlanPeriodStart,
		PlanPeriodEnd:   p.PlanPeriodEnd,
		CreatedAt:       p.CreatedAt,
		Notes:           p.Notes,
		Items:           itResp,
	}
}

// valueOrEmpty safely dereferences an optional string pointer, defaulting to empty.
func valueOrEmpty(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// @Summary Upload procurement plan Excel
// @Tags ProcurementPlan
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Procurement plan XLS/XLSX file"
// @Param facility_id formData string true "Facility ID"
// @Param facility_name formData string false "Facility name"
// @Param financial_year formData string true "Financial year, e.g. 2025/26"
// @Success 201 {array} dto.ProcurementPlanResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /procurement-plans/upload-xls [post]
func UploadProcurementPlanXLS(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	facilityID := c.PostForm("facility_id")
	if facilityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "facility_id is required"})
		return
	}
	facilityName := c.PostForm("facility_name")
	financialYear := c.PostForm("financial_year")
	if financialYear == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "financial_year is required"})
		return
	}
	region := c.PostForm("region")
	zone := c.PostForm("zone")
	district := c.PostForm("district")
	levelOfCare := c.PostForm("level_of_care")
	ownership := strings.ToUpper(strings.TrimSpace(c.PostForm("ownership")))
	storeCode := ""
	switch ownership {
	case "GOV":
		storeCode = "NMS"
	case "PNFP":
		storeCode = "JMS"
	}

	src, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	xls, err := excelize.OpenReader(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Excel file"})
		return
	}
	defer func() { _ = xls.Close() }()

	// Use the first sheet as the main procurement plan
	sheetName := xls.GetSheetName(0)
	if sheetName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Excel file has no sheets"})
		return
	}

	plans, err := importProcurementSheet(xls, sheetName, facilityID, facilityName, financialYear, region, zone, district, levelOfCare, storeCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Replace existing plan for this facility + financial year (re-upload = no duplicates)
	if res := config.DB.Where("facility_id = ? AND financial_year = ?", facilityID, financialYear).Delete(&models.ProcurementPlan{}); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	var responses []dto.ProcurementPlanResponseDTO
	for _, p := range plans {
		if err := config.DB.Create(&p).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responses = append(responses, mapToProcurementPlanResponse(p, nil))
	}

	c.JSON(http.StatusCreated, responses)
}

// importProcurementSheet parses a hospital-level procurement plan sheet into plan rows.
func importProcurementSheet(f *excelize.File, sheetName, facilityID, facilityName, financialYear, region, zone, district, levelOfCare, storeCode string) ([]models.ProcurementPlan, error) {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var plans []models.ProcurementPlan
	var currentSection string

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		code := strings.TrimSpace(row[0])
		desc := ""
		if len(row) > 1 {
			desc = strings.TrimSpace(row[1])
		}

		// Header / section rows: first cell not a product code, treat as section
		if code == "" || !looksLikeProductCode(code) {
			if code != "" {
				currentSection = code
			}
			continue
		}

		plan := models.ProcurementPlan{
			PlanSystemCode: "XLS_IMPORT",
			StoreCode:      storeCode,
			FacilityID:     facilityID,
			FacilityName:   facilityName,
			ProductCode:    code,
			ProductDescription: desc,
			Section:        stringPtr(currentSection),
			FinancialYear:  financialYear,
			CreatedAt:      time.Now(),
		}
		if region != "" {
			plan.Region = &region
		}
		if zone != "" {
			plan.Zone = &zone
		}
		if district != "" {
			plan.District = &district
		}
		if levelOfCare != "" {
			plan.LevelOfCare = &levelOfCare
		}

		// Column order: 0=CODE, 1=DESCRIPTION, 2=UNIT, 3=VEN, 4=PRICE, 5=FY2024/25 BIMONTHLY PLANNED QTY,
		// 6=PAST AVG. NMS ISSUE PLAN Q1, 7=2024 ADJUSTED AMC, 8=FY2025/26 BIMONTHLY PLAN QTY, 9=TOTAL, 10=COMMENT, 11=FUNDED QTY
		if len(row) > 2 {
			if u := strings.TrimSpace(row[2]); u != "" {
				plan.UnitOfMeasure = &u
			}
		}
		if len(row) > 3 {
			if ven := strings.TrimSpace(row[3]); ven != "" {
				plan.VenClassification = &ven
			}
		}
		if len(row) > 4 {
			if price, err := strconv.ParseFloat(strings.TrimSpace(row[4]), 64); err == nil {
				plan.UnitPrice = &price
			}
		}
		if len(row) > 5 {
			if qty, err := strconv.Atoi(strings.TrimSpace(row[5])); err == nil {
				plan.PreviousBiMonthlyPlannedQty = &qty
			}
		}
		if len(row) > 6 {
			if v, err := strconv.ParseFloat(strings.TrimSpace(row[6]), 64); err == nil {
				plan.PastAvgNmsIssuePlanQty = &v
			}
		}
		if len(row) > 7 {
			if v, err := strconv.ParseFloat(strings.TrimSpace(row[7]), 64); err == nil {
				plan.AdjustedAmc = &v
			}
		}
		if len(row) > 8 {
			if qty, err := strconv.Atoi(strings.TrimSpace(row[8])); err == nil {
				plan.BiMonthlyPlanQty = &qty
			}
		}
		if len(row) > 10 {
			if c := strings.TrimSpace(row[10]); c != "" {
				plan.Comment = &c
			}
		}
		if len(row) > 11 {
			if qty, err := strconv.Atoi(strings.TrimSpace(row[11])); err == nil {
				plan.FundedQty = &qty
			}
		}

		plans = append(plans, plan)
	}

	return plans, nil
}

func looksLikeProductCode(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
