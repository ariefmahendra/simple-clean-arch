package config

import (
	"database/sql"
	"fmt"
	"simple-rest-api-clean-arch/helper"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "12345"
	dbname   = "category"
)

func DatabaseConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	return db
}
