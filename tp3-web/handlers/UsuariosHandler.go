package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	sqlc "tp3-web/db/sqlc"
	"tp3-web/logic"
)

func (s *Server) UsuariosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		s.postUsuarios(w, r)
	case http.MethodGet:
		s.getUsuarios(w, r)
	default:

		http.Error(w, "Metodo no autorizado", http.StatusMethodNotAllowed)
		log.Print("Se accedio los usuarios con un metodo no definido")
	}
}

func postUsuarios(w http.ResponseWriter, r *http.Request) {
	var u logic.Usuario

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		log.Print("Error al decodificar los datos recibidos del POST para crear un usuario")
		return
	}

	// Validar el usuario (por ejemplo, email, contraseña, etc.)
	if err := logic.ValidarUsuario(u); err != nil {
		http.Error(w, fmt.Sprintf("Los datos recibidos para el nuevo usuario no son correctos"), http.StatusBadRequest)
		log.Print("Se intento crear un usuario con datos incorrectos")
		return
	}

	// Crear el usuario en la base de datos
	ctx := context.Background()
	arg := sqlc.CreateUsuarioParams{
		Nombre:     u.Nombre,
		Email:      u.Email,
		Contraseña: u.Contraseña,
	}

	usuarioCreado, err := s.queries.CreateUsuario(ctx, arg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear usuario: %v", err), http.StatusInternalServerError)
		log.Print("Error de la base de datos al crear un nuevo usuario")
		return
	}

	json.NewEncoder(w).Encode(usuarioCreado)
	log.Printf("Se creo un nuevo usuario con ID %v", usuarioCreado.IDUsuario)
}
func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario, err := queries.GetUsuarios(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer usuarios: %v", err), http.StatusInternalServerError)
		log.Print("Error de la base de datos al mostrar todos los usuarios")
		return
	}
	json.NewEncoder(w).Encode(usuario)
	log.Print("Se mostraron todos los usuarios")
}
