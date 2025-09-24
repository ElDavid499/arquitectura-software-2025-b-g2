package service

import (
	"escenario2-by-module/user/repository"
	"fmt"
)

func CreateUserDemo() {
	fmt.Println("⚙️ [Service] Crear User (demo)")
	repository.SaveUserDemo()
}

func GetUsersDemo() {
	fmt.Println("⚙️ [Service] Listar Users (demo)")
	repository.GetAllUsersDemo()
}

func GetUserDemo() {
	fmt.Println("⚙️ [Service] Obtener User por ID (demo)")
	repository.GetUserByIdDemo()
}

func UpdateUserDemo() {
	fmt.Println("⚙️ [Service] Actualizar User (demo)")
	repository.UpdateUserDemo()
}

func DeleteUserDemo() {
	fmt.Println("⚙️ [Service] Eliminar User (demo)")
	repository.DeleteUserDemo()
}
