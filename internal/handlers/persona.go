package handlers

import (
	"encoding/json"
	"github.com/xLuisPc/ProyectoGO/internal/db"
	"github.com/xLuisPc/ProyectoGO/internal/models"
	"net/http"
)

func CrearPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var persona models.Persona
	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Calcular promedio
	suma := persona.POO + persona.CalculoMultivariado + persona.CTD + persona.IngenieriaSoftware +
		persona.BasesDeDatos + persona.ControlAnalogo + persona.CircuitosDigitales
	persona.Promedio = suma / 7

	query := `
        INSERT INTO personas (
            carrera, genero_accion, genero_ciencia_ficcion, genero_comedia,
            genero_terror, genero_documental, genero_romance, genero_musicales,
            poo, calculo_multivariado, ctd, ingenieria_software, bases_datos,
            control_analogo, circuitos_digitales, promedio
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
    `
	_, err = db.DB.Exec(query,
		persona.Carrera,
		persona.GeneroAccion,
		persona.GeneroCienciaFiccion,
		persona.GeneroComedia,
		persona.GeneroTerror,
		persona.GeneroDocumental,
		persona.GeneroRomance,
		persona.GeneroMusicales,
		persona.POO,
		persona.CalculoMultivariado,
		persona.CTD,
		persona.IngenieriaSoftware,
		persona.BasesDeDatos,
		persona.ControlAnalogo,
		persona.CircuitosDigitales,
		persona.Promedio,
	)
	if err != nil {
		http.Error(w, "Error al insertar en la base de datos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Persona agregada correctamente"))
}
