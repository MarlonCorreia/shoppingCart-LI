package models

import (
	"shoppingCart-LI/config"
)

func CreateCart() (Cart, error) {
	var cart Cart

	db, err := config.GetConnection()
	if err != nil {
		return cart, err
	}
	db.Create(&cart)

	return cart, nil
}

func GetCart(cartId uint) (Cart, error) {
	var cart Cart
	db, err := config.GetConnection()
	if err != nil {
		return cart, err
	}

	err = db.First(&cart, cartId).Error
	if err != nil {
		return cart, err
	}

	db.Preload("Product").Find(&cart.Orders)
	db.Preload("DiscountCoupons").Find(&cart)

	return cart, nil
}

func AddProductToCart(cartId uint, productId uint, qty int64) error {
	db, err := config.GetConnection()
	if err != nil {
		return err
	}

	exists, err := ProductExists(productId)
	if err != nil {
		return err
	}
	if exists {
		var cart Cart
		db.Find(&cart, cartId)
		db.Preload("Product").Find(&cart.Orders)

		for _, v := range cart.Orders {
			if productId == v.ProductID {
				v.Quantity = v.Quantity + qty
				db.Save(&v)

				return nil
			}
		}

		order, err := CreateOrder(productId)
		if err != nil {
			return err
		}

		db.Model(&cart).Association("Orders").Append(&order)

	}

	return nil
}

func CleanCart(cartID uint) error {
	db, err := config.GetConnection()
	if err != nil {
		return err
	}

	var c Cart
	db.Preload("Product").Find(&c.Orders)
	db.Preload("DiscountCoupons").Find(&c)
	err = db.Find(&c, cartID).Error
	if err != nil {
		return err
	}

	for _, v := range c.Orders {
		err = DeleteOrder(&v)
		if err != nil {
			return err
		}
	}
	db.Delete(&c.DiscountCoupons)

	db.Save(&c)
	return nil
}
