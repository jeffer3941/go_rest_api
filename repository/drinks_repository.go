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

func (dr *DrinksRepository) CreateDrinks(drink model.Drinks) (model.Drinks, error) {
	query := "INSERT INTO drinks (name, description, value) VALUES ($1, $2, $3) RETURNING id"
	err := dr.connection.QueryRow(query, drink.Name, drink.Description, drink.Value).Scan(&drink.Id)
	if err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}
	return drink, nil
}

func (dr *DrinksRepository) UpdateDrinks(id int, drink model.Drinks) (model.Drinks, error) {
	query := "UPDATE drinks SET name = $1, description = $2, value = $3 WHERE id = $4"
	_, err := dr.connection.Exec(query, drink.Name, drink.Description, drink.Value, id)
	if err != nil {
		fmt.Println(err)
		return model.Drinks{}, err
	}
	return drink, nil
}

func (dr *DrinksRepository) DeleteDrinks(id int) error {
	query := "DELETE FROM drinks WHERE id = $1"
	_, err := dr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
