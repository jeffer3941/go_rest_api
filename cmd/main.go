package main

import (
	"drink-api/controller"
	"drink-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	DrinksUsecase := usecase.NewDrinksController()

	DrinksController := controller.NewDrinksController(DrinksUsecase)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá mundo!",
		})
	})

	server.GET("/drinks", DrinksController.GetDrinksController)

	server.Run(":8080")
}
