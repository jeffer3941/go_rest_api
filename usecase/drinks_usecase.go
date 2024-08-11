package usecase

import (
	"drink-api/model"
	"drink-api/repository"
)

type DrinksUsecase struct {
	//
	repository repository.DrinksRepository
}

func NewDrinksController(repo repository.DrinksRepository) DrinksUsecase {
	return DrinksUsecase{
		repository: repo,
	}
}

func (d DrinksUsecase) GetDrinks() ([]model.Drinks, error) {
	return d.repository.GetDrinks()
}

func (d DrinksUsecase) GetDrinksById(id int) (model.Drinks, error) {
	return d.repository.GetDrinksById(id)
}

func (d DrinksUsecase) CreateDrinks(drink model.Drinks) (model.Drinks, error) {
	return d.repository.CreateDrinks(drink)
}

func (d DrinksUsecase) UpdateDrinks(id int, drink model.Drinks) (model.Drinks, error) {
	return d.repository.UpdateDrinks(id, drink)
}

func (d DrinksUsecase) DeleteDrinks(id int) error {
	return d.repository.DeleteDrinks(id)
}
