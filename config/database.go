package config

import (
	"fmt"
	"inoutgo/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5132
	user     = "admin"
	password = "admin"
	dbName   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	helpers.ErrorPanic(err)

	return db
}
