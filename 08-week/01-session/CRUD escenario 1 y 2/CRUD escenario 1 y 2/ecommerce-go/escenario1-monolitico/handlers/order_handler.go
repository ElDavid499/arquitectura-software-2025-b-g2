package handlers

import (
	"crud-inventory/models"
	"crud-inventory/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	orders, err := services.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := services.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.ID = uint(id)
	if err := services.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
