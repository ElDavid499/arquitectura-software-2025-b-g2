package main

import (
	"escenario2-by-module/user/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 User Service corriendo en http://localhost:8081")

	r := gin.Default()

	// Rutas CRUD de User
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.Run(":8081")
}
