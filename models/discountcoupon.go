package models

import "shoppingCart-LI/config"

func GetDiscountCoupon(couponId uint) (DiscountCoupon, error) {
	var coupon DiscountCoupon
	db := config.GetConnection()

	err := db.First(&coupon, couponId).Error
	if err != nil {
		return coupon, err
	}

	return coupon, nil
}

func GetallDiscountCoupon() []DiscountCoupon {
	var coupons []DiscountCoupon
	db := config.GetConnection()

	db.Find(&coupons)

	return coupons
}

func CreateDiscountCoupon(coupon DiscountCoupon) {
	db := config.GetConnection()

	db.Create(&coupon)

	return
}

func DeleteDiscountCoupon(discountCoupon *DiscountCoupon) {
	db := config.GetConnection()
	db.Unscoped().Delete(discountCoupon)

	return
}
