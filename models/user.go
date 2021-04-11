package models

import (
	"shoppingCart-LI/config"
	"shoppingCart-LI/utils"

	"gorm.io/gorm"
)

func CreateUser(name string, password string) error {
	db, err := config.GetConnection()
	if err != nil {
		return err
	}
	cart, err := CreateCart()
	if err != nil {
		return err
	}

	u := User{
		Name:     name,
		Password: password,
		Token:    utils.GenerateToken(name, password),
		CartID:   cart.ID,
		Cart:     cart,
	}
	db.Create(&u)

	return nil
}

func GetUserToken(name string, password string) (string, error) {
	db, err := config.GetConnection()
	if err != nil {
		return "", err
	}

	var u User
	err = db.Where("name = ? AND password = ?", name, password).First(&u).Error
	if err != nil {
		return "", nil
	}

	return u.Token, nil

}

func CheckUserExists(name string, password string) (bool, error) {
	db, err := config.GetConnection()
	if err != nil {
		return false, err
	}

	var u User
	err = db.Where("name = ? AND password = ?", name, password).First(&u).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func UserCart(db *gorm.DB, userId uint) Cart {
	var user User
	db.Preload("Cart").Find(&user, userId)

	return user.Cart
}
