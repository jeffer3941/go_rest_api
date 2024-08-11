package repository

import (
	"database/sql"
	"drink-api/model"
	"fmt"
)

type DrinksRepository struct {
	connection *sql.DB
}

func NewDrinksRepository(connection *sql.DB) DrinksRepository {
	return DrinksRepository{
		connection: connection,
	}
}

func (dr *DrinksRepository) GetDrinks() ([]model.Drinks, error) {
	query := "SELECT id, name, description, value FROM drinks"
	rows, err := dr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Drinks{}, err
	}

	var drinks []model.Drinks
	var DrinkObj model.Drinks

	for rows.Next() {
		err := rows.Scan(&DrinkObj.Id, &DrinkObj.Name, &DrinkObj.Description, &DrinkObj.Value)
		if err != nil {
			fmt.Println(err)
			return []model.Drinks{}, err
		}
		drinks = append(drinks, DrinkObj)
	}

	return drinks, nil
}

func (dr *DrinksRepository) GetDrinksById(id int) (model.Drinks, error) {
	query := "SELECT id, name, description, value FROM drinks WHERE id = $1"
	row := dr.connection.QueryRow(query, id)
	var drinks model.Drinks
	err := row.Scan(&drinks.Id, &drinks.Name, &drinks.Description, &drinks.Value)
	if err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}
	return drinks, nil
}
