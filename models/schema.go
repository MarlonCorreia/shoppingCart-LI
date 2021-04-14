package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model `json:"-"`
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
}

type DiscountCoupon struct {
	gorm.Model `json:"-"`
	ID         uint    `json:"id"`
	Price      float64 `json:"discount"`
	Name       string  `json:"name"`
}

type Order struct {
	gorm.Model `json:"-"`
	ID         uint
	ProductID  uint `json:"-"`
	Quantity   int64
	Product    Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Cart struct {
	gorm.Model      `json:"-"`
	ID              uint
	Orders          []Order          `gorm:"many2many:cart_order;"`
	DiscountCoupons []DiscountCoupon `gorm:"many2many:cart_discountcoupon;"`
}

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Token      string `json:"-"`
	CartID     uint   `json:"cartId"`
	Cart       Cart   `json:"-"`
}
