package models

import "shoppingCart-LI/config"

func CreateOrder(productId uint) (Order, error) {
	var order Order
	db, err := config.GetConnection()
	if err != nil {
		return order, nil
	}
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
	db, err := config.GetConnection()
	if err != nil {
		return err
	}
	db.Delete(order)

	return nil
}
