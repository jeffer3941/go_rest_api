package usecase

import "drink-api/model"

type DrinksUsecase struct {
	//
}

func NewDrinksController() DrinksUsecase {
	return DrinksUsecase{}
}

func (d DrinksUsecase) GetDrinks() ([]model.Drinks, error) {
	return []model.Drinks{}, nil
}
