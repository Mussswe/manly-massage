package entity

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name          string         `json:"name"`
	CoursePeriods []CoursePeriod `gorm:"foreignKey:CourseID"`
}
