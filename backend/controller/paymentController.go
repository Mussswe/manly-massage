package controller

import (
	"backend/config"
	"backend/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ===== Employees =====
func GetEmployees(c *gin.Context) {
	var employees []entity.Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// ===== Paymethods =====
func GetPaymethods(c *gin.Context) {
	var paymethods []entity.Paymethod
	if err := config.DB.Find(&paymethods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paymethods})
}

// ===== Healthcares =====
func GetHealthcares(c *gin.Context) {
	var healthcares []entity.HealthCare
	if err := config.DB.Find(&healthcares).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": healthcares})
}

// ===== CreatePayment =====
func CreatePayment(c *gin.Context) {
	var input struct {
		CourseID     uint    `json:"course_id"`
		PeriodID     uint    `json:"period_id"`
		PaymethodID  uint    `json:"paymethod_id"`
		EmployeeID   uint    `json:"employee_id"`
		HealthCareID *uint   `json:"healthcare_id"` // optional
		Discount     float64 `json:"discount"`
		CustomerName string  `json:"customer_name"`
		StartTime    string  `json:"start_time"` // ISO string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ ใช้ timezone Australia/Sydney
	loc, _ := time.LoadLocation("Australia/Sydney")
	start, err := time.ParseInLocation(time.RFC3339, input.StartTime, loc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_time"})
		return
	}

	// ดึง CoursePeriod
	var cp entity.CoursePeriod
	if err := config.DB.Where("course_id = ? AND period_id = ?", input.CourseID, input.PeriodID).
		Preload("Course").Preload("Period").
		First(&cp).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoursePeriod not found"})
		return
	}

	// คำนวณ total price
	totalPrice := cp.CusPrice - input.Discount
	if totalPrice < 0 {
		totalPrice = 0
	}

	// healthcare
	if input.HealthCareID != nil {
		var hc entity.HealthCare
		if err := config.DB.First(&hc, *input.HealthCareID).Error; err == nil {
			totalPrice -= hc.Hicaps
			if totalPrice < 0 {
				totalPrice = 0
			}
		}
	}

	// ดึง duration ของ period และคำนวณ end_time
	var period entity.Period
	if err := config.DB.First(&period, input.PeriodID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Period not found"})
		return
	}
	end := start.Add(time.Duration(period.Duration) * time.Minute)

	// สร้าง Payment
	payment := entity.Payment{
		CourseID:     cp.CourseID,
		PeriodID:     cp.PeriodID,
		PaymethodID:  input.PaymethodID,
		EmployeeID:   input.EmployeeID,
		HealthCareID: input.HealthCareID,
		Discount:     input.Discount,
		TotalPrice:   totalPrice,
		CustomerName: input.CustomerName,
		StartTime:    start,
		EndTime:      end,
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// ===== GetPayments =====
func GetPayments(c *gin.Context) {
	var payments []entity.Payment
	if err := config.DB.Preload("Course").
		Preload("Period").
		Preload("Employee").
		Preload("Paymethod").
		Preload("HealthCare").
		Order("created_at desc").
		Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	loc, _ := time.LoadLocation("Australia/Sydney")

	type PaymentResponse struct {
		ID             uint    `json:"id"`
		CustomerName   string  `json:"customer_name"`
		CourseID       uint    `json:"course_id"`
		CourseName     string  `json:"course_name"`
		PeriodID       uint    `json:"period_id"`
		Duration       int     `json:"duration"`
		EmployeeID     uint    `json:"employee_id"`
		EmployeeName   string  `json:"employee_name"`
		PaymethodID    uint    `json:"paymethod_id"`
		PaymethodName  string  `json:"paymethod_name"`
		HealthCareID   *uint   `json:"healthcare_id,omitempty"`
		HealthCareName *string `json:"healthcare_name,omitempty"`
		Hicaps         float64 `json:"hicaps"`
		Discount       float64 `json:"discount"`
		TotalPrice     float64 `json:"total_price"`
		StartTime      string  `json:"start_time"`
		EndTime        string  `json:"end_time"`
	}

	var resp []PaymentResponse
	for _, p := range payments {
		var hcName *string
		var hcAmount float64
		if p.HealthCare != nil {
			hcName = &p.HealthCare.Name
			hcAmount = p.HealthCare.Hicaps
		}

		resp = append(resp, PaymentResponse{
			ID:             p.ID,
			CustomerName:   p.CustomerName,
			CourseID:       p.CourseID,
			CourseName:     p.Course.Name,
			PeriodID:       p.PeriodID,
			Duration:       p.Period.Duration,
			EmployeeID:     p.EmployeeID,
			EmployeeName:   p.Employee.Name,
			PaymethodID:    p.PaymethodID,
			PaymethodName:  p.Paymethod.Name,
			HealthCareID:   p.HealthCareID,
			HealthCareName: hcName,
			Hicaps:         hcAmount,
			Discount:       p.Discount,
			TotalPrice:     p.TotalPrice,
			StartTime:      p.StartTime.In(loc).Format(time.RFC3339), // ✅ Sydney time
			EndTime:        p.EndTime.In(loc).Format(time.RFC3339),   // ✅ Sydney time
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// ===== UpdatePayment =====
func UpdatePayment(c *gin.Context) {
	var input struct {
		CourseID     uint    `json:"course_id"`
		PeriodID     uint    `json:"period_id"`
		PaymethodID  uint    `json:"paymethod_id"`
		EmployeeID   uint    `json:"employee_id"`
		HealthCareID *uint   `json:"healthcare_id"`
		Discount     float64 `json:"discount"`
		CustomerName string  `json:"customer_name"`
		StartTime    string  `json:"start_time"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ดึง payment
	id := c.Param("id")
	var payment entity.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// ✅ ใช้ timezone Australia/Sydney
	loc, _ := time.LoadLocation("Australia/Sydney")
	start, err := time.ParseInLocation(time.RFC3339, input.StartTime, loc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_time"})
		return
	}

	// คำนวณ total price
	totalPrice := input.Discount
	var cp entity.CoursePeriod
	if err := config.DB.Where("course_id = ? AND period_id = ?", input.CourseID, input.PeriodID).First(&cp).Error; err == nil {
		totalPrice = cp.CusPrice - input.Discount
		if totalPrice < 0 {
			totalPrice = 0
		}

		if input.HealthCareID != nil {
			var hc entity.HealthCare
			if err := config.DB.First(&hc, *input.HealthCareID).Error; err == nil {
				totalPrice -= hc.Hicaps
				if totalPrice < 0 {
					totalPrice = 0
				}
			}
		}
	}

	// ดึง duration ของ period และคำนวณ end_time
	var period entity.Period
	if err := config.DB.First(&period, input.PeriodID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Period not found"})
		return
	}
	end := start.Add(time.Duration(period.Duration) * time.Minute)

	// อัปเดต payment
	payment.CourseID = input.CourseID
	payment.PeriodID = input.PeriodID
	payment.PaymethodID = input.PaymethodID
	payment.EmployeeID = input.EmployeeID
	payment.HealthCareID = input.HealthCareID
	payment.Discount = input.Discount
	payment.CustomerName = input.CustomerName
	payment.TotalPrice = totalPrice
	payment.StartTime = start
	payment.EndTime = end

	if err := config.DB.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}
func GetPublicHolidays(c *gin.Context) {
	var holidays []entity.PublicHoliday
	if err := config.DB.Find(&holidays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": holidays})
}
