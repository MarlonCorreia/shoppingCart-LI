package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	godotenv.Load()

	dsn := "host=" + os.Getenv("DB_HOST") + " " +
		"user=" + os.Getenv("DB_USER") + " " +
		"password=" + os.Getenv("DB_PASSWORD") + " " +
		"dbname=" + os.Getenv("DB_NAME") + " " +
		"port=" + os.Getenv("DB_PORT") + " " +
		"sslmode=disable"

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
