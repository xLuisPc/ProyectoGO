package main

import (
	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/handlers"
	"log"
	"net/http"
	"os"
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
		case "agg-100":
			db.Agregar100Personas(db.DB, "C:\\Users\\luigi\\Desktop\\GOProyect\\estudiantes_aleatorios.json")
			return
		}
	}

	http.HandleFunc("/personas", handlers.CrearPersona)

	http.HandleFunc("/api/estudiantes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ListarPersonas(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CrearPersona(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Archivos estáticos (JS, CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Ruta para añadir estudiantes
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/add.html")
	})

	// Ruta principal (solo raíz "/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "web/templates/index.html")
	})

	log.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
