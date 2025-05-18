package main

import (
	"log"
	"net/http"

	"github.com/tuusuario/proyecto/internal/db"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API lista"))
	})

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
