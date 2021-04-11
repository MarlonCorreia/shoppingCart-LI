package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn string = "host=localhost user=marlon password=postgres dbname=postgres port=5432 sslmode=disable"
	db  *gorm.DB
)

func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ApllyMigrations(models []interface{}) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	for _, v := range models {
		db.AutoMigrate(&v)
	}

	return nil
}
