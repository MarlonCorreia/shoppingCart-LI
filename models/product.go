package models

import "shoppingCart-LI/config"

func CreateProduct(id uint, name string, price float64, status string) {
	db := config.GetConnection()

	product := Product{
		ID:     id,
		Name:   name,
		Status: status,
		Price:  price,
	}
	db.Create(&product)

	return
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

func UpdateProduct(product *Product, updateTo Product) {
	db := config.GetConnection()

	if updateTo.Name != "" {
		product.Name = updateTo.Name
	}
	if updateTo.Price != 0 {
		product.Price = updateTo.Price

	}
	if updateTo.Status != "" {
		product.Status = updateTo.Status
	}
	db.Save(&product)

	return
}

func DeleteProduct(product *Product) {
	db := config.GetConnection()

	DeleteOrdersByProductId(product.ID)

	db.Unscoped().Delete(product)

}
