package model

type Drinks struct {
	Id          int    `json:"id_drink"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"price"`
}
