package model

type Drinks struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float32 `json:"value"`
}
