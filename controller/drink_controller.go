package controller

import (
	"drink-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type drinksController struct {
}

func NewDrinksController() drinksController {
	return drinksController{}
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
