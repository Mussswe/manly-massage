package entity

import "gorm.io/gorm"

type Paymethod struct {
	gorm.Model
	Name    string    `json:"paymethod"`
	Payment []Payment `gorm:"foreignKey:PaymethodID"` // แก้ typo: "forigkey: PaymentID " → "foreignKey:PaymethodID"
}
