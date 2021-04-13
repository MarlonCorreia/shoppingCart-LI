package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dsn := "host=localhost user=marlon password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	return nil

}

func GetConnection() *gorm.DB {
	return DB
}

func ApllyMigrations(models []interface{}) error {
	for _, v := range models {
		DB.AutoMigrate(v)
	}

	return nil
}
