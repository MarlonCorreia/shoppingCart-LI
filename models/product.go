package models

import "shoppingCart-LI/config"

func CreateProduct(id uint, name string, price float64, status string) error {
	db := config.GetConnection()

	product := Product{
		ID:     id,
		Name:   name,
		Status: status,
		Price:  price,
	}
	db.Create(&product)

	return nil

}

func GetProduct(productId uint) (Product, error) {
	var product Product
	db := config.GetConnection()

	err := db.First(&product, productId).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func GetAllProducts() ([]Product, error) {
	db := config.GetConnection()

	var products []Product

	err := db.Find(&products).Error
	if err != nil {
		return products, nil
	}

	return products, nil
}

func UpdateProduct(productId uint, product Product) error {
	db := config.GetConnection()

	var updatedProduct Product
	err := db.First(&updatedProduct, productId).Error
	if err != nil {
		return err
	}

	if product.Name != "" {
		updatedProduct.Name = product.Name
	}
	if product.Status != "" {
		updatedProduct.Status = product.Status
	}
	if product.Price != 0 {
		updatedProduct.Price = product.Price
	}

	db.Save(&updatedProduct)

	return nil
}

func DeleteProduct(productId uint) error {
	db := config.GetConnection()

	err := db.Delete(&Product{}, productId).Error
	if err != nil {
		return err
	}

	return nil
}

func ProductExists(productId uint) (bool, error) {
	db := config.GetConnection()

	var product Product
	err := db.First(&product, productId).Error
	if err != nil {
		return false, nil
	}

	return true, nil
}
