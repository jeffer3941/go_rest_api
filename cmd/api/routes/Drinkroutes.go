package route

import (
	"drink-api/controller"
	"drink-api/db"
	"drink-api/repository"
	"drink-api/usecase"

	"github.com/gin-gonic/gin"
)

func DrinksRoutes(server *gin.Engine) {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	DrinksRepository := repository.NewDrinksRepository(dbConnection)

	DrinksUsecase := usecase.NewDrinksController(DrinksRepository)

	DrinksController := controller.NewDrinksController(DrinksUsecase)

	drinkGroup := server.Group("/v1")
	{
		drinkGroup.GET("/drinks", DrinksController.GetDrinksController)

		drinkGroup.GET("/drink/:id", DrinksController.GetDrinksByIdController)

		drinkGroup.POST("/createdrink", DrinksController.CreateDrinksController)

		drinkGroup.PUT("/updatedrink", DrinksController.UpdateDrinksController)

		drinkGroup.DELETE("/deletedrink", DrinksController.DeleteDrinksController)
	}

}
