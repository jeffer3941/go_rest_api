package controller

import (
	"drink-api/model"
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

func (d *drinksController) CreateDrinksController(ctx *gin.Context) {
	var drink model.Drinks
	err := ctx.ShouldBindJSON(&drink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	drink, err = d.drinksUsecase.CreateDrinks(drink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, drink)
}

func (d *drinksController) UpdateDrinksController(ctx *gin.Context) {
	var drink model.Drinks
	err := ctx.ShouldBindJSON(&drink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	id := drink.Id

	drink, err = d.drinksUsecase.UpdateDrinks(id, drink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, drink)
}

func (d *drinksController) DeleteDrinksController(ctx *gin.Context) {
	var drink model.Drinks
	err := ctx.ShouldBindJSON(&drink)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	id := drink.Id

	err = d.drinksUsecase.DeleteDrinks(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "deleted",
	})
}
