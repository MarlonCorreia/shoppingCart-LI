package db

import (
	"shoppingCart-LI/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=marlon password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}

func Migrations() error {
	db, err := getConnection()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Product{})
	if err != nil {
		return err
	}
	return nil
}
