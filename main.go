package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salmageek/hellogo/controllers"
	"github.com/salmageek/hellogo/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	r.GET("/images", controllers.FindImages)
	r.GET("/images/:id", controllers.GetImage)
	r.Run()
}
