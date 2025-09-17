package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name    string    `json:"name"`
	Revenue float64   `json:"revenue"`
	Payment []Payment `gorm:"foreignKey:EmployeeID"` // ถูกต้อง
}
