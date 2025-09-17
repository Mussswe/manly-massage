package entity

import "gorm.io/gorm"

type HealthCare struct {
	gorm.Model
	Name    string    `json:"name"`
	Hicaps  float64   `json:"hicaps"`
	Payment []Payment `gorm:"foreignKey:HealthCareID"` // แก้ typo: HealtCareID → HealthCareID
}
