package handlers

import (
	"database/sql" // <--- AGREGADO: Necesario para sql.NullString
	"net/http"
	"strconv"
	"time"

	sqlc "tp3-web/db/sqlc"
	"tp3-web/views"
)

// GET /entrenamientos
func (s *Server) EntrenamientosPage(w http.ResponseWriter, r *http.Request) {
	// Por simplicidad uso usuario_id = 1. En producción sacalo de la sesión.
	entrenamientos, err := s.Queries.GetEntrenamientos(r.Context())
	if err != nil {
		http.Error(w, "Error leyendo entrenamientos", http.StatusInternalServerError)
		return
	}

	views.EntrenamientosPage(entrenamientos).Render(r.Context(), w)
}

// POST /entrenamientos (PRG)
func (s *Server) CreateEntrenamiento(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error leyendo formulario", http.StatusBadRequest)
		return
	}

	// si tu formulario incluyera usuario_id, parsealo; por ahora usamos 1
	usuarioID := int32(1)

	fechaStr := r.FormValue("fecha")
	tipo := r.FormValue("tipo")
	distanciaStr := r.FormValue("distancia")
	tiempoStr := r.FormValue("tiempo")
	notas := r.FormValue("notas")

	if fechaStr == "" || tipo == "" || distanciaStr == "" || tiempoStr == "" {
		http.Error(w, "Campos obligatorios faltantes", http.StatusBadRequest)
		return
	}

	fecha, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		http.Error(w, "Fecha inválida", http.StatusBadRequest)
		return
	}

	// Validamos que sea numero, pero para guardarlo usamos el string original
	// ya que sqlc espera un string para tipos DECIMAL/NUMERIC
	_, err = strconv.ParseFloat(distanciaStr, 64)
	if err != nil {
		http.Error(w, "Distancia inválida", http.StatusBadRequest)
		return
	}

	tiempo, err := strconv.Atoi(tiempoStr)
	if err != nil {
		http.Error(w, "Tiempo inválido", http.StatusBadRequest)
		return
	}

	_, err = s.Queries.CreateEntrenamiento(r.Context(), sqlc.CreateEntrenamientoParams{
		Fecha:     fecha,
		Tipo:      tipo,
		Distancia: distanciaStr, // <--- CAMBIO: Usamos el string directo
		Tiempo:    int32(tiempo),
		Notas:     sql.NullString{String: notas, Valid: notas != ""}, // <--- CAMBIO: sql.NullString
		UsuarioID: usuarioID,
	})
	if err != nil {
		http.Error(w, "Error creando entrenamiento: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// PRG: redirigir a la lista
	http.Redirect(w, r, "/entrenamientos", http.StatusSeeOther)
}
