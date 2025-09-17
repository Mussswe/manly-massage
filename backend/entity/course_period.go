package entity

import "gorm.io/gorm"

type CoursePeriod struct {
	gorm.Model

	CusPrice float64 `json:"cus_price"`
	EmpPrice float64 `json:"emp_price"`

	CourseID uint   `json:"course_id"`
	Course   Course `gorm:"foreignKey:CourseID" json:"course"`

	PeriodID uint   `json:"period_id"`
	Period   Period `gorm:"foreignKey:PeriodID" json:"period"`
}
