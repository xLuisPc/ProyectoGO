package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	const (
		host     = "turntable.proxy.rlwy.net"
		port     = 13930
		user     = "postgres"
		password = "rIjtdMpRegkyAbfcaaQeYHzvqjvwvBjr"
		dbname   = "railway"
	)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error // declarar err antes
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("✅ Conexión exitosa a PostgreSQL en Railway")
}
