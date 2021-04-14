package models

import (
	"shoppingCart-LI/config"
)

func CreateOrder(product *Product, qty int64) Order {
	var order Order
	db := config.GetConnection()

	order.Quantity = qty
	order.Product = *product
	order.ProductID = product.ID

	db.Create(&order)

	return order

}

func DeleteOrder(order *Order) {
	db := config.GetConnection()

	db.Delete(order)

	return
}

func DeleteOrdersByProductId(productId uint) {
	db := config.GetConnection()

	var orders []Order
	db.Find(&orders)

	for _, v := range orders {
		db.Preload("Product").Find(&v)
		if v.Product.ID == productId {
			db.Delete(&v)
		}
	}

}
