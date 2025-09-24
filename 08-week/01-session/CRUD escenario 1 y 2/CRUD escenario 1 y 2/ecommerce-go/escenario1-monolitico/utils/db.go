package utils

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"crud-inventory/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Conexión a SQLite
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	// Migraciones automáticas
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("Error en migraciones: ", err)
	}

	log.Println("✅ Base de datos conectada y migraciones ejecutadas")
}
