package main

import (
	"backend/config"
	"backend/entity"
	"backend/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&entity.Course{},
		&entity.Period{},
		&entity.CoursePeriod{},
		&entity.Employee{},
		&entity.Paymethod{},
		&entity.HealthCare{},
		&entity.Payment{},
		&entity.PublicHoliday{}, // <--- ต้องมีตัวนี้
	)
	config.SetHoliday()
	//config.SetupData()
	//config.SeedData()
	fmt.Println("✅ Database setup completed")

	r := gin.Default()

	// ⚡ เพิ่ม cors middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	r.Run(":8080")
}
