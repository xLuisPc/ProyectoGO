package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/xLuisPc/ProyectoGO/internal/models"
	"io/ioutil"
	"log"
	"os"
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

// Agregar100Personas lee un archivo JSON y agrega las personas a la base de datos.
func Agregar100Personas(db *sql.DB, jsonPath string) error {
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("❌ No se pudo abrir el archivo JSON: %v", err)
	} else {
		log.Println("✅ Archivo JSON abierto correctamente")
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error leyendo archivo JSON: %w", err)
	}

	var personas []models.Persona
	if err := json.Unmarshal(bytes, &personas); err != nil {
		return fmt.Errorf("error parseando JSON: %w", err)
	}

	for _, p := range personas {
		_, err := db.Exec(`
			INSERT INTO DBPersonas (
				id, carrera, genero_accion, genero_ciencia_ficcion, genero_comedia,
				genero_terror, genero_documental, genero_romance, genero_musicales,
				poo, ctd, calculo_multivariado, ingenieria_software, bases_datos,
				control_analogo, circuitos_digitales, promedio
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9,
				$10, $11, $12, $13, $14, $15, $16, $17
			)
		`, p.ID, p.Carrera, p.GeneroAccion, p.GeneroCienciaFiccion, p.GeneroComedia,
			p.GeneroTerror, p.GeneroDocumental, p.GeneroRomance, p.GeneroMusicales,
			p.Poo, p.Ctd, p.CalculoMultivariado, p.IngenieriaSoftware, p.BasesDatos,
			p.ControlAnalogo, p.CircuitosDigitales, p.Promedio)

		if err != nil {
			log.Printf("Error insertando ID %d: %v", p.ID, err)
		}
	}
	fmt.Println("Todas las personas fueron insertadas exitosamente.")
	return nil
}
