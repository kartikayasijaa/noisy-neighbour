package database

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"noisy_neighbour/common/database/config"
	"noisy_neighbour/common/database/helpers"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbInstance := config.NewDBConnection()
	connection, err := dbInstance.Connect()

	db = connection

	return connection, err
}

func Migrate() error {
	dbInstance, err := ConnectDB()
	if err != nil {
		return err
	}

	migrateErr := dbInstance.AutoMigrate(helpers.GetAllModels()...)
	return migrateErr
}
