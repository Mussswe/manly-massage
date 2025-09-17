package controller

import (
	"backend/config"
	"backend/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------------- Course CRUD -------------------

// Get all courses
func GetCourses(c *gin.Context) {
	var courses []entity.Course
	config.DB.Find(&courses)
	c.JSON(http.StatusOK, gin.H{"data": courses})
}

// Get single course by ID
func GetCourse(c *gin.Context) {
	id := c.Param("id")
	var course entity.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}

// Create a new course
func CreateCourse(c *gin.Context) {
	var course entity.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&course)
	c.JSON(http.StatusOK, gin.H{"data": course})
}

// Update a course
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course entity.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	var input entity.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.Name = input.Name
	config.DB.Save(&course)
	c.JSON(http.StatusOK, gin.H{"data": course})
}

// Delete a course
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	var course entity.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	config.DB.Delete(&course)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
