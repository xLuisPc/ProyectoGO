package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/handlers"
)

func main() {
	db.ConnectDB()

	// Admin: ejecutar acciones con argumentos
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "create-table":
			db.CreateTable()
			return
		case "drop-table":
			db.DropTable()
			return
		}
	}

	http.HandleFunc("/personas", handlers.CrearPersona)

	log.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
