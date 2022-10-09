package config

import (
	"assignment2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type DBConfig map[string]string

func Dc() string {
	config := Config()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		config["host"],
		config["user"],
		config["password"],
		config["dbName"],
		config["sslMode"],
	)
}

func StartDB() {
	var err error
	dc := Dc()

	db, err = gorm.Open(postgres.Open(dc))
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
