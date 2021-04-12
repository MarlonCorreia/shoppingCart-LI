package models

import "shoppingCart-LI/config"

func CreateDiscountCoupon(name string, price float64) error {
	db := config.GetConnection()

	discountCoupon := DiscountCoupon{
		Name:  name,
		Price: price,
	}

	db.Create(&discountCoupon)

	return nil
}

func DeleteDiscountCoupon(discountCouponId uint) error {
	db := config.GetConnection()

	err := db.Delete(&DiscountCoupon{}, discountCouponId).Error
	if err != nil {
		return err
	}

	return nil
}
