package controller

import (
	"drink-api/usecase"
	"net/http"
	"strconv"

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
	drinks, err := d.drinksUsecase.GetDrinks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, drinks)
}

func (d *drinksController) GetDrinksByIdController(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	drink, err := d.drinksUsecase.GetDrinksById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, drink)
}
