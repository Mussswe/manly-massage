package controller

import (
	"net/http"

	"backend/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB // ตัวแปร global สำหรับ DB

func SetDB(db *gorm.DB) {
	DB = db
}

// DTO สำหรับ response
type EmployeeRevenue struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Revenue float64 `json:"revenue"`
}

// Handler สำหรับ route /employees/revenue
func GetEmployeesRevenue(c *gin.Context) {
	var revenues []EmployeeRevenue

	err := DB.Model(&entity.Employee{}).
		Select("employees.id, employees.name, IFNULL(SUM(payments.amount),0) as revenue").
		Joins("LEFT JOIN payments ON payments.employee_id = employees.id").
		Group("employees.id").
		Scan(&revenues).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": revenues})
}
