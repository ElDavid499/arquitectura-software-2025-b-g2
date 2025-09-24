package handlers

import (
	"escenario2-by-module/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	service.CreateUserDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Crear User (demo)"})
}

func GetUsers(c *gin.Context) {
	service.GetUsersDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener Users (demo)"})
}

func GetUser(c *gin.Context) {
	service.GetUserDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener User por ID (demo)"})
}

func UpdateUser(c *gin.Context) {
	service.UpdateUserDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Actualizar User (demo)"})
}

func DeleteUser(c *gin.Context) {
	service.DeleteUserDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Eliminar User (demo)"})
}
