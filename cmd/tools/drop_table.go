package main

import (
	"log"

	"github.com/xLuisPc/ProyectoGO/internal/db"
)

func main() {
	db.ConnectDB()

	query := `DROP TABLE IF EXISTS personas;`

	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Error eliminando la tabla: %v", err)
	} else {
		log.Println("🗑️ Tabla 'personas' eliminada exitosamente.")
	}
}
