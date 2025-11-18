package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	sqlc "tp3-web/db/sqlc"
	"tp3-web/logic"
)

func (s *Server) EntrenamientosIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "URL invalida", http.StatusBadRequest)
		log.Print("Se intento acceder a entrenamientos con una ID invalida")
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID de usuario invalido", http.StatusBadRequest)
		log.Print("Se intento acceder a entrenamientos con una ID invalida")
		return
	}
	switch r.Method {
	case http.MethodGet:
		s.getEntrenamientosId(w, r, id)
	case http.MethodPut:
		s.actualizarEntrenamiento(w, r, id)
	case http.MethodDelete:
		s.eliminarEntrenamiento(w, r, id)
	}
}
func (s *Server) eliminarEntrenamiento(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	err := s.Queries.DeleteEntrenamiento(r.Context(), int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No se pudo elimiar el entrenamiento"), http.StatusInternalServerError)
		log.Print("No se pudo elimiar el entrenamiento")
		return
	}
	log.Printf("Se borro el entrenamiento %v", id)
}
func (s *Server) actualizarEntrenamiento(w http.ResponseWriter, r *http.Request, id int) {
	// Lo unico que no se va a poder actualizar es la id de los entrenamientos, ya que eso lo maneja la base de datos
	var entrenamientoActualizado logic.Entrenamiento

	err := json.NewDecoder(r.Body).Decode(&entrenamientoActualizado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error al leer la modificacion de un entrenamiento")
		return
	}

	entrenamientoActualizadoSQL := sqlc.UpdateEntrenamientoParams{
		IDEntrenamiento: int32(id), //Este campo no se actualiza, se usa para que la base de datos busque que entrenamiento tiene que modificar
		Fecha:           entrenamientoActualizado.Fecha,
		Tipo:            entrenamientoActualizado.Tipo,
		Tiempo:          int32(entrenamientoActualizado.Tiempo),
		RitmoPromedio: sql.NullString{
			String: fmt.Sprintf("%.2f", entrenamientoActualizado.Ritmo_promedio),
			Valid:  entrenamientoActualizado.Ritmo_promedio != 0,
		},
		Distancia: fmt.Sprintf("%.2f", entrenamientoActualizado.Distancia),
		Calorias: sql.NullInt32{
			Int32: int32(entrenamientoActualizado.Calorias),
			Valid: entrenamientoActualizado.Calorias != 0,
		},
		Notas: sql.NullString{
			String: entrenamientoActualizado.Notas,
			Valid:  entrenamientoActualizado.Notas != "",
		},
	}
	entrenamiento, err := s.Queries.UpdateEntrenamiento(r.Context(), entrenamientoActualizadoSQL)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se pudo actualizar el entrenamiento"), http.StatusInternalServerError)
		log.Print("Error al actualizar un entrenamiento en la base de datos")
		return
	}
	json.NewEncoder(w).Encode(entrenamiento)
	log.Printf("Se actualizo el entrenamiento ID %v en la base de datos", entrenamiento.IDEntrenamiento)
}
func (s *Server) getEntrenamientosId(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	entrenamiento, err := s.Queries.GetEntrenamientoByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No se encontro un entrenamiento con ese ID de usuario"), http.StatusBadRequest)
		log.Print("Error, no se encontro un entrenamiento con el ID %i", id)
	}

	json.NewEncoder(w).Encode(entrenamiento)
	log.Printf("Se mostro el entrenamiento ID %v", id)
}
