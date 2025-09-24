package main

import (
	"escenario2-by-module/order/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Order Service corriendo en http://localhost:8083")

	r := gin.Default()

	// Rutas CRUD de Order
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)
	r.GET("/orders/:id", handlers.GetOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)

	r.Run(":8083")
}
