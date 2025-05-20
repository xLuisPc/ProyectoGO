package services

import (
	"github.com/xLuisPc/ProyectoGO/internal/models"
	"math"
	"math/rand"
	"time"
)

type Cluster struct {
	ID       int
	Personas []models.Persona
}

// KMeansPorGenero aplica clustering sobre afinidad por un género + notas
func KMeansPorGenero(personas []models.Persona, genero string, k int) []Cluster {
	rand.Seed(time.Now().UnixNano())

	var dataset [][]float64
	for _, p := range personas {
		var afinidad float64
		switch genero {
		case "genero_accion":
			afinidad = float64(p.GeneroAccion)
		case "genero_ciencia_ficcion":
			afinidad = float64(p.GeneroCienciaFiccion)
		case "genero_comedia":
			afinidad = float64(p.GeneroComedia)
		case "genero_terror":
			afinidad = float64(p.GeneroTerror)
		case "genero_documental":
			afinidad = float64(p.GeneroDocumental)
		case "genero_romance":
			afinidad = float64(p.GeneroRomance)
		case "genero_musicales":
			afinidad = float64(p.GeneroMusicales)
		default:
			afinidad = 0
		}

		entry := []float64{
			afinidad,
			p.Poo, p.CalculoMultivariado, p.Ctd, p.IngenieriaSoftware,
			p.BasesDatos, p.ControlAnalogo, p.CircuitosDigitales, p.Promedio,
		}
		dataset = append(dataset, entry)
	}

	// Inicializar centroides aleatorios
	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		centroids[i] = dataset[rand.Intn(len(dataset))]
	}

	assignments := make([]int, len(dataset))
	maxIter := 100

	for iter := 0; iter < maxIter; iter++ {
		// Asignar puntos al cluster más cercano
		for i, point := range dataset {
			minDist := math.MaxFloat64
			for j, centroid := range centroids {
				dist := euclidean(point, centroid)
				if dist < minDist {
					minDist = dist
					assignments[i] = j
				}
			}
		}

		// Recalcular centroides
		newCentroids := make([][]float64, k)
		counts := make([]int, k)
		for i := 0; i < k; i++ {
			newCentroids[i] = make([]float64, len(dataset[0]))
		}
		for i, cluster := range assignments {
			for j := range dataset[i] {
				newCentroids[cluster][j] += dataset[i][j]
			}
			counts[cluster]++
		}
		for i := 0; i < k; i++ {
			for j := range newCentroids[i] {
				if counts[i] > 0 {
					newCentroids[i][j] /= float64(counts[i])
				}
			}
		}
		centroids = newCentroids
	}

	// Agrupar estudiantes según cluster asignado
	clusters := make([]Cluster, k)
	for i := range clusters {
		clusters[i].ID = i
	}
	for i, idx := range assignments {
		clusters[idx].Personas = append(clusters[idx].Personas, personas[i])
	}

	return clusters
}

// Calcula la distancia euclidiana entre dos vectores
func euclidean(a, b []float64) float64 {
	var sum float64
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func PredecirClusterPorGustos(dataset [][]float64, nuevo []float64, k int) int {
	// Inicializar centroides
	rand.Seed(time.Now().UnixNano())
	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		centroids[i] = dataset[rand.Intn(len(dataset))]
	}

	assignments := make([]int, len(dataset))

	// Entrenamiento K-means (solo sobre gustos)
	for iter := 0; iter < 100; iter++ {
		// Asignar a cluster más cercano
		for i, punto := range dataset {
			minDist := math.MaxFloat64
			for j, centro := range centroids {
				if d := euclidean(punto, centro); d < minDist {
					minDist = d
					assignments[i] = j
				}
			}
		}

		// Recalcular centroides
		newCentroids := make([][]float64, k)
		counts := make([]int, k)
		for i := 0; i < k; i++ {
			newCentroids[i] = make([]float64, len(dataset[0]))
		}
		for i, c := range assignments {
			for j := range dataset[i] {
				newCentroids[c][j] += dataset[i][j]
			}
			counts[c]++
		}
		for i := 0; i < k; i++ {
			if counts[i] > 0 {
				for j := range newCentroids[i] {
					newCentroids[i][j] /= float64(counts[i])
				}
			}
		}
		centroids = newCentroids
	}

	// Predecir el cluster más cercano para el nuevo perfil
	minDist := math.MaxFloat64
	clusterAsignado := 0
	for j, centro := range centroids {
		if d := euclidean(nuevo, centro); d < minDist {
			minDist = d
			clusterAsignado = j
		}
	}
	return clusterAsignado
}

func PredecirClusterPorGustosConPromedios(dataset [][]float64, nuevo []float64, k int) (int, []int) {
	rand.Seed(time.Now().UnixNano())

	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		centroids[i] = dataset[rand.Intn(len(dataset))]
	}

	assignments := make([]int, len(dataset))

	for iter := 0; iter < 100; iter++ {
		for i, punto := range dataset {
			minDist := math.MaxFloat64
			for j, centro := range centroids {
				if d := euclidean(punto, centro); d < minDist {
					minDist = d
					assignments[i] = j
				}
			}
		}

		newCentroids := make([][]float64, k)
		counts := make([]int, k)
		for i := 0; i < k; i++ {
			newCentroids[i] = make([]float64, len(dataset[0]))
		}
		for i, c := range assignments {
			for j := range dataset[i] {
				newCentroids[c][j] += dataset[i][j]
			}
			counts[c]++
		}
		for i := 0; i < k; i++ {
			if counts[i] > 0 {
				for j := range newCentroids[i] {
					newCentroids[i][j] /= float64(counts[i])
				}
			}
		}
		centroids = newCentroids
	}

	// Predecir cluster del nuevo
	minDist := math.MaxFloat64
	clusterAsignado := 0
	for j, centro := range centroids {
		if d := euclidean(nuevo, centro); d < minDist {
			minDist = d
			clusterAsignado = j
		}
	}
	return clusterAsignado, assignments
}
