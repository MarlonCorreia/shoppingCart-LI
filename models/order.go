package models

import "shoppingCart-LI/config"

func CreateOrder(productId uint) (Order, error) {
	var order Order
	db := config.GetConnection()

	prod, err := GetProduct(productId)
	if err != nil {
		return order, err
	}

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
