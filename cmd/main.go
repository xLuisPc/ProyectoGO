package main

import (
	"log"
	"net/http"

	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/handlers"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/personas", handlers.CrearPersona)

	query := `
    CREATE TABLE IF NOT EXISTS personas (
        id SERIAL PRIMARY KEY,
        carrera TEXT,
        genero_accion INTEGER,
        genero_ciencia_ficcion INTEGER,
        genero_comedia INTEGER,
        genero_terror INTEGER,
        genero_documental INTEGER,
        genero_romance INTEGER,
        genero_musicales INTEGER,
        poo REAL,
        calculo_multivariado REAL,
        ctd REAL,
        ingenieria_software REAL,
        bases_datos REAL,
        control_analogo REAL,
        circuitos_digitales REAL,
        promedio REAL
    );
    `
	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatal("Error creando la tabla:", err)
	} else {
		log.Println("Tabla creada exitosamente.")
	}

	log.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
