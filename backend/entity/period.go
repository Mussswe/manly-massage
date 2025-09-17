package entity

import (
	"gorm.io/gorm"
)

type Period struct {
	gorm.Model
	Duration      int            `json:"duration"` // แทนชื่อ Period
	CoursePeriods []CoursePeriod `gorm:"foreignKey:PeriodID"`
}
