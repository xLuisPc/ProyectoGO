package main

import (
	"log"

	"github.com/xLuisPc/ProyectoGO/internal/db"
)

func main() {
	db.ConnectDB()

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
		log.Fatalf("❌ Error creando la tabla: %v", err)
	} else {
		log.Println("✅ Tabla 'personas' creada exitosamente.")
	}
}
