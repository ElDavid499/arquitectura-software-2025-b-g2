package repositories

import (
	"crud-inventory/models"
	"crud-inventory/utils"
)

func CreateProduct(product *models.Product) error {
	return utils.DB.Create(product).Error
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := utils.DB.Find(&products).Error
	return products, err
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := utils.DB.First(&product, id).Error
	return product, err
}

func UpdateProduct(product *models.Product) error {
	return utils.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return utils.DB.Delete(&models.Product{}, id).Error
}
