package controller

import (
	"drink-api/model"
	"drink-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type drinksController struct {
	drinksUsecase usecase.DrinksUsecase
}

func NewDrinksController(usecase usecase.DrinksUsecase) drinksController {
	return drinksController{
		drinksUsecase: usecase,
	}
}

func (d *drinksController) GetDrinksController(ctx *gin.Context) {
	drinks := []model.Drinks{
		{
			Id:          1,
			Name:        "Cerveja",
			Description: "Cerveja lata",
			Value:       "5,00",
		},
	}

	ctx.JSON(http.StatusOK, drinks)
}
