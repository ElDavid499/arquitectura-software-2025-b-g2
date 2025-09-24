package main

import (
	"escenario2-by-module/product/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Product Service corriendo en http://localhost:8082")

	r := gin.Default()

	// Rutas CRUD de Product
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.Run(":8082")
}
