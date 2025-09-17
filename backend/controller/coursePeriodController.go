package controller

import (
	"backend/config"
	"backend/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all CoursePeriods
func GetCoursePeriods(c *gin.Context) {
	var cps []entity.CoursePeriod
	config.DB.Preload("Course").Preload("Period").Find(&cps)
	c.JSON(http.StatusOK, gin.H{"data": cps})
}

// Create a new CoursePeriod
func CreateCoursePeriod(c *gin.Context) {
	var cp entity.CoursePeriod
	if err := c.ShouldBindJSON(&cp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&cp)
	c.JSON(http.StatusOK, gin.H{"data": cp})
}

// Update an existing CoursePeriod by ID
func UpdateCoursePeriod(c *gin.Context) {
	id := c.Param("id")
	var cp entity.CoursePeriod
	if err := config.DB.First(&cp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoursePeriod not found"})
		return
	}

	var input entity.CoursePeriod
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cp.CusPrice = input.CusPrice
	cp.EmpPrice = input.EmpPrice
	cp.CourseID = input.CourseID
	cp.PeriodID = input.PeriodID

	config.DB.Save(&cp)
	c.JSON(http.StatusOK, gin.H{"data": cp})
}

// Delete a CoursePeriod by ID
func DeleteCoursePeriod(c *gin.Context) {
	id := c.Param("id")
	var cp entity.CoursePeriod
	if err := config.DB.First(&cp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CoursePeriod not found"})
		return
	}
	config.DB.Delete(&cp)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Search CoursePeriods by CourseID or PeriodID
func SearchCoursePeriods(c *gin.Context) {
	courseIDStr := c.Query("course_id")
	periodIDStr := c.Query("period_id")

	var cps []entity.CoursePeriod
	db := config.DB.Preload("Course").Preload("Period")

	if courseIDStr != "" {
		courseID, err := strconv.Atoi(courseIDStr)
		if err == nil {
			db = db.Where("course_id = ?", courseID)
		}
	}

	if periodIDStr != "" {
		periodID, err := strconv.Atoi(periodIDStr)
		if err == nil {
			db = db.Where("period_id = ?", periodID)
		}
	}

	db.Find(&cps)
	c.JSON(http.StatusOK, gin.H{"data": cps})
}
