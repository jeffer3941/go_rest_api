package main

import (
	"drink-api/controller"
	"drink-api/db"
	"drink-api/repository"
	"drink-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	DrinksRepository := repository.NewDrinksRepository(dbConnection)

	DrinksUsecase := usecase.NewDrinksController(DrinksRepository)

	DrinksController := controller.NewDrinksController(DrinksUsecase)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá mundo!",
		})
	})

	server.GET("/drinks", DrinksController.GetDrinksController)

	server.GET("/drink/:id", DrinksController.GetDrinksByIdController)

	server.POST("/createdrink", DrinksController.CreateDrinksController)

	server.POST("/updatedrink", DrinksController.UpdateDrinksController)

	server.POST("/deletedrink", DrinksController.DeleteDrinksController)

	server.Run(":8080")
}
