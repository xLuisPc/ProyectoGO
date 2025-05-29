package main

import (
	"github.com/xLuisPc/ProyectoGO/internal/db"
	"log"
	"os"
)

func main() {
	// Conectar a la base de datos
	db.ConnectDB()
	log.Println("🚀 Base de datos conectada correctamente")

	// Obtener puerto desde variable de entorno (usado por Railway)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto para pruebas locales
	}

}
