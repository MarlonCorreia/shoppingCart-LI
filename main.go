package main

import (
	"fmt"
	"shoppingCart-LI/config"
	"shoppingCart-LI/models"
	"shoppingCart-LI/server"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println("status: ", err)
	}

	schema := []interface{}{
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.DiscountCoupon{},
		&models.Cart{},
	}

	config.ApllyMigrations(schema)

	server.RunServer()

}
