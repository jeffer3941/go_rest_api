package main

import (
	"drink-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	DrinksController := controller.NewDrinksController()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá mundo!",
		})
	})

	server.GET("/drinks", DrinksController.GetDrinksController)

	server.Run(":8080")
}
