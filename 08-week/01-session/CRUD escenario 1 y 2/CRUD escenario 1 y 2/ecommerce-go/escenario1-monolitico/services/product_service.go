package services

import (
	"crud-inventory/models"
	"crud-inventory/repositories"
)

func CreateProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}

func GetAllProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func GetProductByID(id uint) (models.Product, error) {
	return repositories.GetProductByID(id)
}

func UpdateProduct(product *models.Product) error {
	return repositories.UpdateProduct(product)
}

func DeleteProduct(id uint) error {
	return repositories.DeleteProduct(id)
}
