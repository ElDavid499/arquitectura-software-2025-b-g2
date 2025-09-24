package handlers

import (
	"escenario2-by-module/product/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	service.CreateProductDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Crear Product (demo)"})
}

func GetProducts(c *gin.Context) {
	service.GetProductsDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener Products (demo)"})
}

func GetProduct(c *gin.Context) {
	service.GetProductDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Obtener Product por ID (demo)"})
}

func UpdateProduct(c *gin.Context) {
	service.UpdateProductDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Actualizar Product (demo)"})
}

func DeleteProduct(c *gin.Context) {
	service.DeleteProductDemo()
	c.JSON(http.StatusOK, gin.H{"message": "Eliminar Product (demo)"})
}
