package db

import (
	"log"
)

func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS DBPersonas (
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
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Error creando la tabla: %v", err)
	} else {
		log.Println("✅ Tabla creada exitosamente.")
	}
}

func DropTable() {
	_, err := DB.Exec("DROP TABLE IF EXISTS personas;")
	if err != nil {
		log.Fatalf("❌ Error eliminando la tabla: %v", err)
	} else {
		log.Println("✅ Tabla  eliminada exitosamente.")
	}
}
