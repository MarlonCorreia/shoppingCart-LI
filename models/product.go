package models

import "shoppingCart-LI/config"

func CreateProduct(name string, price float64, status string) error {
	db, err := config.GetConnection()
	if err != nil {
		return err
	}

	product := Product{
		Name:   name,
		Status: status,
		Price:  price,
	}
	db.Create(product)

	return nil

}

func GetProduct(productId uint) (Product, error) {
	var product Product
	db, err := config.GetConnection()

	if err != nil {
		return product, err
	}

	err = db.First(&product, productId).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func UpdateProduct(productId uint, product Product) error {
	db, err := config.GetConnection()
	if err != nil {
		return err
	}

	var updatedProduct Product
	err = db.First(&updatedProduct, productId).Error
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
	db, err := config.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(&Product{}, productId).Error
	if err != nil {
		return err
	}

	return nil
}
