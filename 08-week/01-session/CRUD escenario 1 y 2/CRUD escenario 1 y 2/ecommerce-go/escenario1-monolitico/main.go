package main

import (
	"crud-inventory/handlers"
	"crud-inventory/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Iniciando Escenario 1 - Monolítico")

	// Conectar a la base de datos
	utils.ConnectDB()

	// Iniciar servidor con Gin
	r := gin.Default()

	// User CRUD
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// Product CRUD
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	// Order CRUD
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)
	r.GET("/orders/:id", handlers.GetOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)

	// Run server
	r.Run(":8080")
}
