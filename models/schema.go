package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID     uint
	Name   string
	Price  float64
	Status string
}

type DiscountCoupon struct {
	gorm.Model
	ID    uint
	Price float64
	Name  string
}

type Order struct {
	gorm.Model
	ID        uint
	ProductID uint
	Quantity  int64
	Product   Product
}

type Cart struct {
	gorm.Model
	ID              uint
	Orders          []Order          `gorm:"many2many:cart_order;"`
	DiscountCoupons []DiscountCoupon `gorm:"many2many:cart_discountcoupon;"`
}

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Password string
	Token    string
	CartID   uint
	Cart     Cart
}
