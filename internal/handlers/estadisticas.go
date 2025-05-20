package handlers

import (
	"encoding/json"
	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/models"
	"github.com/xLuisPc/ProyectoGO/internal/services"
	"log"
	"net/http"
	"strconv"
)

func ObtenerClusters(w http.ResponseWriter, r *http.Request) {
	genero := r.URL.Query().Get("genero")
	if genero == "" {
		http.Error(w, "Parámetro 'genero' requerido", http.StatusBadRequest)
		return
	}

	// Leer K desde la URL
	k := 3 // valor por defecto
	if kStr := r.URL.Query().Get("k"); kStr != "" {
		if parsed, err := strconv.Atoi(kStr); err == nil && parsed >= 2 && parsed <= 10 {
			k = parsed
		} else {
			log.Println("⚠️ Valor inválido de k, usando 3")
		}
	}

	log.Println("✅ Generando clusters con género:", genero, "y K =", k)

	rows, err := db.DB.Query("SELECT * FROM dbpersonas")
	if err != nil {
		log.Println("❌ Error al obtener personas:", err)
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var personas []models.Persona
	for rows.Next() {
		var p models.Persona
		err := rows.Scan(&p.ID, &p.Carrera, &p.GeneroAccion, &p.GeneroCienciaFiccion,
			&p.GeneroComedia, &p.GeneroTerror, &p.GeneroDocumental, &p.GeneroRomance,
			&p.GeneroMusicales, &p.Poo, &p.CalculoMultivariado, &p.Ctd, &p.IngenieriaSoftware,
			&p.BasesDatos, &p.ControlAnalogo, &p.CircuitosDigitales, &p.Promedio)
		if err != nil {
			log.Println("❌ Error al escanear persona:", err)
			http.Error(w, "Error al procesar resultados", http.StatusInternalServerError)
			return
		}
		personas = append(personas, p)
	}

	// Llamar a K-means
	clusters := services.KMeansPorGenero(personas, genero, k)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}
