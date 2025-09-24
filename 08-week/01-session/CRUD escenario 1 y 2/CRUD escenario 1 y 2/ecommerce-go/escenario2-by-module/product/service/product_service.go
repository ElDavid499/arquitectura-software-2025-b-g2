package service

import (
	"escenario2-by-module/product/repository"
	"fmt"
)

func CreateProductDemo() {
	fmt.Println("⚙️ [Service] Crear Product (demo)")
	repository.SaveProductDemo()
}

func GetProductsDemo() {
	fmt.Println("⚙️ [Service] Listar Products (demo)")
	repository.GetAllProductsDemo()
}

func GetProductDemo() {
	fmt.Println("⚙️ [Service] Obtener Product por ID (demo)")
	repository.GetProductByIdDemo()
}

func UpdateProductDemo() {
	fmt.Println("⚙️ [Service] Actualizar Product (demo)")
	repository.UpdateProductDemo()
}

func DeleteProductDemo() {
	fmt.Println("⚙️ [Service] Eliminar Product (demo)")
	repository.DeleteProductDemo()
}

