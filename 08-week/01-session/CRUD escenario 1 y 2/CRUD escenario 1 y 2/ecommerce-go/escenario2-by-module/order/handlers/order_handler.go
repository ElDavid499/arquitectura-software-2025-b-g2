package handlers

import (
	"escenario2-by-module/order/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	service.CreateOrderDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Crear Order (demo)"})
}

func GetOrders(c *gin.Context) {
	service.GetOrdersDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener Orders (demo)"})
}

func GetOrder(c *gin.Context) {
	service.GetOrderDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener Order por ID (demo)"})
}

func UpdateOrder(c *gin.Context) {
	service.UpdateOrderDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Actualizar Order (demo)"})
}

func DeleteOrder(c *gin.Context) {
	service.DeleteOrderDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Eliminar Order (demo)"})
}
