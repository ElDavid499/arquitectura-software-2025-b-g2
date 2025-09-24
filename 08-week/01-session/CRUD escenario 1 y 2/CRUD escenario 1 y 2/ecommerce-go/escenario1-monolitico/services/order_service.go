package services

import (
	"crud-inventory/models"
	"crud-inventory/repositories"
)

func CreateOrder(order *models.Order) error {
	return repositories.CreateOrder(order)
}

func GetAllOrders() ([]models.Order, error) {
	return repositories.GetAllOrders()
}

func GetOrderByID(id uint) (models.Order, error) {
	return repositories.GetOrderByID(id)
}

func UpdateOrder(order *models.Order) error {
	return repositories.UpdateOrder(order)
}

func DeleteOrder(id uint) error {
	return repositories.DeleteOrder(id)
}
