package server

import (
	"github.com/xLuisPc/ProyectoGO/internal/handlers"
	"net/http"
)

func SetupRouter() *http.Server {
	mux := http.NewServeMux()

	// API REST
	mux.HandleFunc("/api/estudiantes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListarPersonas(w, r)
		case http.MethodPost:
			handlers.CrearPersona(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/estadisticas", handlers.ObtenerClusters)
	mux.HandleFunc("/api/prediccion", handlers.PredecirCluster)

	// Archivos estáticos
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// HTML
	mux.HandleFunc("/", rootHandler("web/templates/index.html"))
	mux.HandleFunc("/add", rootHandler("web/templates/add.html"))
	mux.HandleFunc("/estadisticas", rootHandler("web/templates/estadisticas.html"))

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func rootHandler(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && file == "web/templates/index.html" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, file)
	}
}
