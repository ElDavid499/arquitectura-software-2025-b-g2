package repositories

import (
	"crud-inventory/models"
	"crud-inventory/utils"
)

func CreateUser(user *models.User) error {
	return utils.DB.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := utils.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := utils.DB.First(&user, id)
	return user, result.Error
}

func UpdateUser(user *models.User) error {
	return utils.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return utils.DB.Delete(&models.User{}, id).Error
}
