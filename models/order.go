package models

import (
	"shoppingCart-LI/config"
)

func CreateOrder(productId uint) (Order, error) {
	var order Order
	db := config.GetConnection()

	prod, err := GetProduct(productId)
	if err != nil {
		return order, err
	}
	order.Quantity = 1
	order.Product = prod
	order.ProductID = prod.ID
	db.Create(&order)

	return order, nil

}

func DeleteOrder(order *Order) error {
	db := config.GetConnection()

	db.Delete(order)

	return nil
}

func DeleteOrdersByProductId(productId uint) error {
	db := config.GetConnection()

	var orders []Order
	db.Find(&orders)

	for _, v := range orders {
		db.Preload("Product").Find(&v)
		if v.Product.ID == productId {
			db.Delete(&v)
		}
	}
	return nil

}
