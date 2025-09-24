package service

import (
	"escenario2-by-module/order/repository"
	"fmt"
)

func CreateOrderDemo() {
	fmt.Println("⚙️ [Service] Crear Order (demo)")
	repository.SaveOrderDemo()
}

func GetOrdersDemo() {
	fmt.Println("⚙️ [Service] Listar Orders (demo)")
	repository.GetAllOrdersDemo()
}

func GetOrderDemo() {
	fmt.Println("⚙️ [Service] Obtener Order por ID (demo)")
	repository.GetOrderByIdDemo()
}

func UpdateOrderDemo() {
	fmt.Println("⚙️ [Service] Actualizar Order (demo)")
	repository.UpdateOrderDemo()
}

func DeleteOrderDemo() {
	fmt.Println("⚙️ [Service] Eliminar Order (demo)")
	repository.DeleteOrderDemo()
}
