package config

import (
	"fmt"
	"golang-crud/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5400
	user     = "postgres"
	password = "1234"
	dbName   = "CRUD"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host = %s port =%d user =%s password=%s dbname=%s sslmode= disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ReturnError(err)
	return db
}
