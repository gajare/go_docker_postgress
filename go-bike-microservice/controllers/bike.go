package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-bike-microservice/config"  // Import the config package
	"go-bike-microservice/models"
)

type CreateBikeInput struct {
	Model       string  `json:"model" binding:"required"`
	Manufacturer string `json:"manufacturer" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type UpdateBikeInput struct {
	Model       string  `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Price       float64 `json:"price"`
}

func FindBikes(c *gin.Context) {
	var bikes []models.Bike
	config.DB.Find(&bikes)  // Use config.DB

	c.JSON(http.StatusOK, gin.H{"data": bikes})
}

func FindBike(c *gin.Context) {
	var bike models.Bike

	if err := config.DB.Where("id = ?", c.Param("id")).First(&bike).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bike})
}

func CreateBike(c *gin.Context) {
	var input CreateBikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bike := models.Bike{
		Model:        input.Model,
		Manufacturer: input.Manufacturer,
		Price:        input.Price,
	}
	config.DB.Create(&bike)  // Use config.DB

	c.JSON(http.StatusOK, gin.H{"data": bike})
}

func UpdateBike(c *gin.Context) {
	var bike models.Bike
	if err := config.DB.Where("id = ?", c.Param("id")).First(&bike).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input UpdateBikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&bike).Updates(input)  // Use config.DB

	c.JSON(http.StatusOK, gin.H{"data": bike})
}

func DeleteBike(c *gin.Context) {
	var bike models.Bike
	if err := config.DB.Where("id = ?", c.Param("id")).First(&bike).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	config.DB.Delete(&bike)  // Use config.DB

	c.JSON(http.StatusOK, gin.H{"data": true})
}
