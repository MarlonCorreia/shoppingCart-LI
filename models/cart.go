package models

import (
	"shoppingCart-LI/config"
)

func CreateCart() (Cart, error) {
	var cart Cart

	db := config.GetConnection()

	db.Create(&cart)

	return cart, nil
}

func GetCart(cartId uint) (Cart, error) {
	var cart Cart
	db := config.GetConnection()

	err := db.First(&cart, cartId).Error
	if err != nil {
		return cart, err
	}

	db.Preload("Product").Find(&cart.Orders)
	db.Preload("DiscountCoupons").Find(&cart)

	return cart, nil
}

func AddProductToCart(cartId uint, productId uint, qty int64) error {
	db := config.GetConnection()

	exists, err := ProductExists(productId)
	if err != nil {
		return err
	}
	if exists {
		var cart Cart
		db.Find(&cart, cartId)
		db.Preload("Product").Find(&cart.Orders)

		for _, v := range cart.Orders {
			if productId == v.Product.ID {
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

func RemoveProductFromnCart(productId uint, cartId uint, quantitty int64) error {
	db := config.GetConnection()

	var cart Cart

	err := db.First(&cart, cartId).Error
	db.Preload("Product").Find(&cart.Orders)

	if err != nil {
		return err
	}

	for _, v := range cart.Orders {
		if v.Product.ID == productId {
			v.Quantity = v.Quantity - quantitty

			if v.Quantity <= 0 {
				DeleteOrder(&v)

				return nil
			}

			db.Save(&v)
		}
	}

	return nil
}

func CleanCart(cartID uint) error {
	db := config.GetConnection()

	var c Cart
	db.Preload("Product").Find(&c.Orders)
	db.Preload("DiscountCoupons").Find(&c)
	err := db.Find(&c, cartID).Error
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

func AddDiscountCouponToCart(cart *Cart, coupon *DiscountCoupon) {
	db := config.GetConnection()
	db.Preload("DiscountCoupons").Find(&cart)

	db.Model(cart).Association("DiscountCoupons").Append(coupon)

}

func DeleteDiscountCouponFromCart(cart *Cart, coupon *DiscountCoupon) {
	db := config.GetConnection()
	db.Preload("DiscountCoupons").Find(&cart)

	db.Model(cart).Association("DiscountCoupons").Delete(coupon)
}
