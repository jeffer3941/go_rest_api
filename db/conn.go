package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	sslmode  = "disable"
)

func ConnectDB() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
