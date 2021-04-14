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

func GetUser(userId uint) (User, error) {
	db := config.GetConnection()

	var user User
	err := db.First(&user, userId).Error
	db.Preload("Cart").Find(&user)

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

func CheckTokenExists(token string, cartId uint) bool {
	db := config.GetConnection()
	var user User

	err := db.Where("token = ?", token).First(&user).Error
	if err != nil || user.CartID != cartId {
		return false
	}

	return true
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
