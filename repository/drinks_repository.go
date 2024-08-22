package repository

import (
	"drink-api/model"
	"fmt"

	"gorm.io/gorm"
)

type DrinksRepository struct {
	connection *gorm.DB
}

func NewDrinksRepository(connection *gorm.DB) DrinksRepository {
	return DrinksRepository{
		connection: connection,
	}
}

func (dr *DrinksRepository) GetDrinks() ([]model.Drinks, error) {
	var drinks []model.Drinks

	if err := dr.connection.Find(&drinks).Error; err != nil {
		fmt.Println(err)
		return []model.Drinks{}, err
	}

	return drinks, nil
}

func (dr *DrinksRepository) GetDrinksById(id int) (model.Drinks, error) {
	var drinks model.Drinks

	if err := dr.connection.First(&drinks, id).Error; err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}

	return drinks, nil
}

func (dr *DrinksRepository) CreateDrinks(drink model.Drinks) (model.Drinks, error) {
	result := dr.connection.Create(&drink)
	if result.Error != nil {
		fmt.Println(result.Error)
		return model.Drinks{}, result.Error
	}
	return drink, nil
}

func (dr *DrinksRepository) UpdateDrinks(id int, updateDrink model.Drinks) (model.Drinks, error) {
	var drink model.Drinks

	if err := dr.connection.First(&drink, id).Error; err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}

	drink.Name = updateDrink.Name
	drink.Description = updateDrink.Description
	drink.Value = updateDrink.Value

	if err := dr.connection.Save(&drink).Error; err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}

	return drink, nil
}

func (dr *DrinksRepository) DeleteDrinks(id int) error {
	if er := dr.connection.Delete(&model.Drinks{}, id); er.Error != nil {
		fmt.Println(er.Error)
		return er.Error
	}
	return nil
}
