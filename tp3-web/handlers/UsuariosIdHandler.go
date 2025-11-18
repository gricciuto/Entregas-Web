package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) usuarioIdHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "URL invalida", http.StatusBadRequest)
		log.Print("Error, id de usuario invalido")
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID de usuario invalido", http.StatusBadRequest)
		log.Print("Error, id de usuario invalido")
		return
	}
	switch r.Method {
	case http.MethodGet:
		s.getUsuarioByID(w, r, id)
	case http.MethodDelete:
		s.deleteUsuario(w, r, id)
	default:
		http.Error(w, "Metodo no autorizado", http.StatusMethodNotAllowed)
		log.Print("Error, metodo no definido para usuario")
	}
}
func (s *Server) deleteUsuario(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	err := s.queries.DeleteUsuario(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al borrar usuario", http.StatusBadRequest)
		log.Print("Error de la base al borrar el usuario ID %i", id)
	}
	log.Printf("Se borro el usuario %v", id)
}
func (s *Server) getUsuarioByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	usuario, err := s.queries.GetUsuarioByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer usuarios: %v", err), http.StatusInternalServerError)
		log.Printf("Error de la base de datos al mostrar el usuario ID %v", id)
		return
	}
	json.NewEncoder(w).Encode(usuario)
	log.Printf("Se mostro el usuario ID %v", id)
}
