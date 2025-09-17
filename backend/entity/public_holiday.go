package entity

import (
	"time"

	"gorm.io/gorm"
)

type PublicHoliday struct {
	gorm.Model
	HolidayDate time.Time `json:"holiday_date" gorm:"index"`     // เพิ่ม index สำหรับค้นหาเร็ว
	Name        string    `json:"name" gorm:"size:100;not null"` // กำหนดขนาดและ not null
	State       string    `json:"state" gorm:"size:50"`          // จำกัดความยาว state
}
