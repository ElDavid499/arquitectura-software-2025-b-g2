package repositories

import (
	"crud-inventory/models"
	"crud-inventory/utils"
)

func CreateOrder(order *models.Order) error {
	return utils.DB.Create(order).Error
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := utils.DB.Find(&orders).Error
	return orders, err
}

func GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	err := utils.DB.First(&order, id).Error
	return order, err
}

func UpdateOrder(order *models.Order) error {
	return utils.DB.Save(order).Error
}

func DeleteOrder(id uint) error {
	return utils.DB.Delete(&models.Order{}, id).Error
}
