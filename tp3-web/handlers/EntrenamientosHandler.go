package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	sqlc "tp3-web/db/sqlc"
	"tp3-web/logic"
)

const peso int = 70

func (s *Server) EntrenamientosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		s.getEntrenamientos(w, r)
	case http.MethodPost:
		s.crearEntrenamiento(w, r)
	default:
		http.Error(w, fmt.Sprintf("Metodo no permitido"), http.StatusMethodNotAllowed)
		log.Print("Error, se intento acceder a un metodo no definido de entrenamientos")
	}
}
func (s *Server) getEntrenamientos(w http.ResponseWriter, r *http.Request) {
	entrenamiento, err := s.Queries.GetEntrenamientos(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener los entrenamientos"), http.StatusInternalServerError)
		log.Print("Error al mostrar los entrenamientos")
		return
	}
	json.NewEncoder(w).Encode(entrenamiento)
	log.Print("Se mostraron todos los entrenamientos")
}
func (s *Server) crearEntrenamiento(w http.ResponseWriter, r *http.Request) {
	var entrenamientoNuevo logic.Entrenamiento

	err := json.NewDecoder(r.Body).Decode(&entrenamientoNuevo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer los datos recibidos"), http.StatusBadRequest)
		log.Print("Error al leer los datos recibidos al crear un entrenamiento")
		return
	}

	entrenamientoNuevoSQL := sqlc.CreateEntrenamientoParams{
		UsuarioID: entrenamientoNuevo.IDusuario,
		Fecha:     entrenamientoNuevo.Fecha,
		Tipo:      entrenamientoNuevo.Tipo,
		Tiempo:    int32(entrenamientoNuevo.Tiempo),
		RitmoPromedio: sql.NullString{
			String: fmt.Sprintf("%.2f", entrenamientoNuevo.Distancia/float64(entrenamientoNuevo.Tiempo)),
			Valid:  entrenamientoNuevo.Ritmo_promedio != 0,
		},
		Distancia: fmt.Sprintf("%.2f", entrenamientoNuevo.Distancia),
		Calorias: sql.NullInt32{
			Int32: int32(float64(entrenamientoNuevo.Tiempo) * float64(peso) * entrenamientoNuevo.Ritmo_promedio),
			Valid: entrenamientoNuevo.Calorias != 0,
		},
		Notas: sql.NullString{
			String: entrenamientoNuevo.Notas,
			Valid:  entrenamientoNuevo.Notas != "",
		},
	}
	entrenamiento, err := s.Queries.CreateEntrenamiento(r.Context(), entrenamientoNuevoSQL)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se pudo crear el entrenamiento"), http.StatusInternalServerError)
		log.Print("Error al crear el entrenamiento en la base de datos")
		return
	}
	json.NewEncoder(w).Encode(entrenamiento)
	log.Printf("Se creo el entrenamiento con ID %v", entrenamiento.IDEntrenamiento)
}
