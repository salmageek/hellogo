package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salmageek/hellogo/models"
)

// Gets all images
func FindImages(c *gin.Context) {
	var images []models.Image
	models.DB.Find(&images)
	c.JSON(http.StatusOK, gin.H{"data": images})
}

// Get image detail by ID
func GetImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var image []models.Image
	models.DB.First(&image, id)
	c.JSON(http.StatusOK, gin.H{"data": image})
}
