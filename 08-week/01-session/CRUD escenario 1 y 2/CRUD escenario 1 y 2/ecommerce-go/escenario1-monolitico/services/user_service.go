package services

import (
	"crud-inventory/models"
	"crud-inventory/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserByID(id uint) (models.User, error) {
	return repositories.GetUserByID(id)
}

func UpdateUser(user *models.User) error {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
