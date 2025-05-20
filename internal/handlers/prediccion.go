package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/models"
	"github.com/xLuisPc/ProyectoGO/internal/services"
	"log"
	"net/http"
)

func PredecirCluster(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevo models.Persona
	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Obtener todos los estudiantes con gustos y promedio
	rows, err := db.DB.Query("SELECT genero_accion, genero_ciencia_ficcion, genero_comedia, genero_terror, genero_documental, genero_romance, genero_musicales, promedio FROM dbpersonas")
	if err != nil {
		log.Println("❌ Error al obtener personas:", err)
		http.Error(w, "Error al consultar base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var gustos [][]float64
	var promedios []float64
	for rows.Next() {
		var g1, g2, g3, g4, g5, g6, g7 int
		var promedio float64
		if err := rows.Scan(&g1, &g2, &g3, &g4, &g5, &g6, &g7, &promedio); err != nil {
			continue
		}
		gustos = append(gustos, []float64{
			float64(g1), float64(g2), float64(g3), float64(g4),
			float64(g5), float64(g6), float64(g7),
		})
		promedios = append(promedios, promedio)
	}

	// Vector del nuevo estudiante
	perfilNuevo := []float64{
		float64(nuevo.GeneroAccion),
		float64(nuevo.GeneroCienciaFiccion),
		float64(nuevo.GeneroComedia),
		float64(nuevo.GeneroTerror),
		float64(nuevo.GeneroDocumental),
		float64(nuevo.GeneroRomance),
		float64(nuevo.GeneroMusicales),
	}

	// Predecir cluster
	clusterID, asignaciones := services.PredecirClusterPorGustosConPromedios(gustos, perfilNuevo, 3)

	// Calcular promedio de ese cluster
	var suma float64
	var cantidad int
	for i, asignado := range asignaciones {
		if asignado == clusterID {
			suma += promedios[i]
			cantidad++
		}
	}
	var promedioEstimado float64
	if cantidad > 0 {
		promedioEstimado = suma / float64(cantidad)
	}

	// Devolver JSON con cluster y promedio
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"cluster": %d, "promedio": %.2f}`, clusterID, promedioEstimado)))
}
