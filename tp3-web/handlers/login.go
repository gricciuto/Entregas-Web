package handlers

import (
	"context"
	"crypto/subtle"
	"net/http"
	"strings"

	"tp3-web/views"
)

// GET /login
func (s *Server) LoginGet(w http.ResponseWriter, r *http.Request) {
	views.LoginPage("").Render(r.Context(), w)
}

// POST /login
// Este login es básico: busca por email y compara contraseña literal (NO recomendado para producción).
// Recomendación: guardar hashes bcrypt y comparar con bcrypt.CompareHashAndPassword.
func (s *Server) LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		views.LoginPage("Error leyendo formulario").Render(r.Context(), w)
		return
	}

	email := strings.TrimSpace(r.FormValue("email"))
	pass := r.FormValue("contraseña")

	if email == "" || pass == "" {
		views.LoginPage("Complete todos los campos").Render(r.Context(), w)
		return
	}

	// Buscar usuario por email
	user, err := s.Queries.GetUsuarioByEmail(context.Background(), email)
	if err != nil {
		views.LoginPage("Usuario/Email no encontrado").Render(r.Context(), w)
		return
	}

	// Comparación segura (literal); en producción comparar hashes bcrypt
	if subtle.ConstantTimeCompare([]byte(user.Contraseña), []byte(pass)) != 1 {
		views.LoginPage("Credenciales incorrectas").Render(r.Context(), w)
		return
	}

	// TODO: crear sesión/cookie segura. Por ahora redirigir a entrenamientos
	http.Redirect(w, r, "/entrenamientos", http.StatusSeeOther)
}
