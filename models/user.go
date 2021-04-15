package models

import (
	"shoppingCart-LI/config"
	"shoppingCart-LI/utils"

	"gorm.io/gorm"
)

func CreateUser(name string, password string) error {
	db := config.GetConnection()

	cart := CreateCart()

	u := User{
		Name:     name,
		Password: password,
		Token:    utils.GenerateToken(name, password),
		CartID:   cart.ID,
		Cart:     cart,
	}
	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserByToken(token string) (User, error) {
	db := config.GetConnection()

	var user User
	err := db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserToken(name string, password string) string {
	db := config.GetConnection()

	var u User
	db.Where("name = ? AND password = ?", name, password).First(&u)

	return u.Token

}

func UserCartByToken(token string) (*Cart, error) {
	db := config.GetConnection()
	var user User

	err := db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return &user.Cart, err
	}
	db.Preload("Cart.Orders.Product").Find(&user)
	db.Preload("DiscountCoupons").Find(&user.Cart)

	return &user.Cart, nil
}

func CheckUserExists(name string, password string) bool {
	db := config.GetConnection()

	var u User
	err := db.Where("name = ? AND password = ?", name, password).First(&u).Error
	if err != nil {
		return false
	}

	return true
}

func UserCart(db *gorm.DB, userId uint) Cart {
	var user User
	db.Preload("Cart").Find(&user, userId)

	return user.Cart
}
