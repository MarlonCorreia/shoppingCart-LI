package mock

import (
	"shoppingCart-LI/config"
	"shoppingCart-LI/models"
	"shoppingCart-LI/utils"
)

func InputMockedInfo() {
	db := config.GetConnection()

	diaPais := models.DiscountCoupon{
		Name:  "#dia-dos-pais",
		Price: 350,
	}
	db.Create(&diaPais)

	blackFriday := models.DiscountCoupon{
		Name:  "#black-friday",
		Price: 200,
	}
	db.Create(&blackFriday)

	natal := models.DiscountCoupon{
		Name:  "#feliz-natal",
		Price: 435,
	}
	db.Create(&natal)

	product1 := models.Product{
		ID:     13345,
		Name:   "Iphone 11 32GB",
		Price:  4500,
		Status: "Available",
	}
	db.Create(&product1)

	product2 := models.Product{
		ID:     44567,
		Name:   "Motorola Moto 65s",
		Price:  2100,
		Status: "Available",
	}
	db.Create(&product2)

	product3 := models.Product{
		ID:     55678,
		Name:   "Samsung Galaxy S11",
		Price:  3000,
		Status: "Available",
	}
	db.Create(&product3)

	product4 := models.Product{
		ID:     6678,
		Name:   "Xiaomi Mi Mix",
		Price:  2600,
		Status: "Available",
	}
	db.Create(&product4)

	product5 := models.Product{
		ID:     99234,
		Name:   "LG 65 64GB",
		Price:  1800,
		Status: "Available",
	}
	db.Create(&product5)

	orderLuke := models.Order{
		ID:        4938,
		ProductID: product1.ID,
		Product:   product1,
		Quantity:  1,
	}
	db.Create(&orderLuke)

	cartLuke := models.Cart{
		ID:     345,
		Orders: []models.Order{orderLuke},
	}

	luke := models.User{
		Name:     "Luke Skywalker",
		Password: "iamyourfather",
		Token:    utils.GenerateToken("Luke Skywalker", "iamyourfather"),
		CartID:   cartLuke.ID,
		Cart:     cartLuke,
	}
	db.Create(&luke)

	orderLeia1 := models.Order{
		ID:        234,
		ProductID: product2.ID,
		Product:   product2,
		Quantity:  1,
	}
	db.Create(&orderLeia1)

	orderLeia2 := models.Order{
		ID:        9101,
		ProductID: product3.ID,
		Product:   product3,
		Quantity:  2,
	}
	db.Create(&orderLeia2)

	cartLeia := models.Cart{
		ID:              444,
		Orders:          []models.Order{orderLeia1, orderLeia2},
		DiscountCoupons: []models.DiscountCoupon{blackFriday},
	}

	leia := models.User{
		Name:     "Leia Skywalker",
		Password: "iloveyou",
		Token:    utils.GenerateToken("Leia Skywalker", "iloveyou"),
		CartID:   cartLeia.ID,
		Cart:     cartLeia,
	}
	db.Create(&leia)

	orderYoda1 := models.Order{
		ID:        10394,
		ProductID: product4.ID,
		Product:   product4,
		Quantity:  1,
	}
	db.Create(&orderYoda1)

	orderYoda2 := models.Order{
		ID:        0144,
		ProductID: product5.ID,
		Product:   product5,
		Quantity:  2,
	}
	db.Create(&orderYoda2)

	orderYoda3 := models.Order{
		ID:        7392,
		ProductID: product1.ID,
		Product:   product1,
		Quantity:  1,
	}
	db.Create(&orderYoda3)

	cartYoda := models.Cart{
		ID:              5555,
		Orders:          []models.Order{orderYoda1, orderYoda2, orderYoda3},
		DiscountCoupons: []models.DiscountCoupon{diaPais, natal},
	}

	yoda := models.User{
		Name:     "Yoda",
		Password: "drowssap",
		Token:    utils.GenerateToken("Yoda", "drowssap"),
		CartID:   cartYoda.ID,
		Cart:     cartYoda,
	}
	db.Create(&yoda)
}
