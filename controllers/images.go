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
