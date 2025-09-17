package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	CourseID uint   `json:"course_id"`
	Course   Course `gorm:"foreignKey:CourseID"`

	PeriodID uint   `json:"period_id"`
	Period   Period `gorm:"foreignKey:PeriodID"`

	PaymethodID uint      `json:"paymethod_id"`
	Paymethod   Paymethod `gorm:"foreignKey:PaymethodID"`

	EmployeeID uint     `json:"employee_id"`
	Employee   Employee `gorm:"foreignKey:EmployeeID"`

	HealthCareID *uint       `json:"healthcare_id"`
	HealthCare   *HealthCare `gorm:"foreignKey:HealthCareID"`

	Discount     float64 `json:"discount"`
	TotalPrice   float64 `json:"total_price"`
	CustomerName string  `json:"customer_name"`

	StartTime time.Time `json:"start_time"` // เพิ่มเวลาเริ่มนวด
	EndTime   time.Time `json:"end_time"`   // เพิ่มเวลาเลิกนวด
}
