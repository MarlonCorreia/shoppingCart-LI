package models

import (
	"errors"
	"shoppingCart-LI/config"
)

func CreateCart() Cart {
	var cart Cart

	db := config.GetConnection()

	db.Create(&cart)

	return cart
}

func GetCart(cartId uint) (Cart, error) {
	var cart Cart
	db := config.GetConnection()

	err := db.First(&cart, cartId).Error
	if err != nil {
		return cart, err
	}

	db.Preload("Orders.Product").Find(&cart)
	db.Preload("DiscountCoupons").Find(&cart)

	return cart, nil
}

func AddProductToCart(cart *Cart, product *Product, qty int64) {
	db := config.GetConnection()

	db.Preload("Orders.Product").Find(&cart)

	for _, v := range cart.Orders {
		if product.ID == v.Product.ID {
			v.Quantity = v.Quantity + qty
			db.Save(&v)

			return
		}
	}

	order := CreateOrder(product, qty)

	db.Model(&cart).Association("Orders").Append(&order)

	return
}

func RemoveProductFromnCart(cart *Cart, product *Product, qty int64) error {
	db := config.GetConnection()
	db.Preload("Orders.Product").Find(&cart)

	for _, v := range cart.Orders {
		if v.Product.ID == product.ID {
			if qty <= 0 {
				qty = v.Quantity
			}
			v.Quantity = v.Quantity - qty

			if v.Quantity <= 0 {
				DeleteOrder(&v)

			}
			db.Save(&v)
			return nil
		}
	}
	err := errors.New("Product not in cart")
	return err
}

func CleanCart(cart *Cart) {
	db := config.GetConnection()

	db.Preload("Orders.Product").Find(&cart)
	db.Preload("DiscountCoupons").Find(&cart)

	for _, v := range cart.Orders {
		DeleteOrder(&v)
	}

	for _, v := range cart.DiscountCoupons {
		DeleteDiscountCouponFromCart(cart, &v)
	}

	db.Save(&cart)
	return
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
