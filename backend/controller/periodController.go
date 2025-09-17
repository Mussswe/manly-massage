package controller

import (
	"backend/config"
	"backend/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPeriods(c *gin.Context) {
	var periods []entity.Period
	config.DB.Find(&periods)
	c.JSON(http.StatusOK, gin.H{"data": periods})
}

func CreatePeriod(c *gin.Context) {
	var period entity.Period
	if err := c.ShouldBindJSON(&period); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&period)
	c.JSON(http.StatusOK, gin.H{"data": period})
}
