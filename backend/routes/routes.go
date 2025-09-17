package routes

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Course routes
	api.GET("/courses", controller.GetCourses)
	api.GET("/courses/:id", controller.GetCourse)
	api.POST("/courses", controller.CreateCourse)
	api.PUT("/courses/:id", controller.UpdateCourse)
	api.DELETE("/courses/:id", controller.DeleteCourse)

	// Period routes
	api.GET("/periods", controller.GetPeriods)
	api.POST("/periods", controller.CreatePeriod)

	// CoursePeriod routes
	api.GET("/courseperiods", controller.GetCoursePeriods)
	api.GET("/courseperiods/search", controller.SearchCoursePeriods)
	api.POST("/courseperiods", controller.CreateCoursePeriod)
	api.PUT("/courseperiods/:id", controller.UpdateCoursePeriod)
	api.DELETE("/courseperiods/:id", controller.DeleteCoursePeriod)

	api.POST("/payments", controller.CreatePayment)
	api.GET("/payments", controller.GetPayments)
	api.PUT("/payments/:id", controller.UpdatePayment)

	api.GET("/employees/revenue", controller.GetEmployeesRevenue)

	api.GET("/employees", controller.GetEmployees)
	api.GET("/paymethods", controller.GetPaymethods)
	api.GET("/healthcares", controller.GetHealthcares)
}
