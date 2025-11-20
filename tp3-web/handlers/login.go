package handlers

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"strings"
	sqlc "tp3-web/db/sqlc"

	"tp3-web/views"
)

// GET /login
func (s *Server) LoginGet(w http.ResponseWriter, r *http.Request) {
	views.LoginPage("").Render(r.Context(), w)
	log.Println("Se mostro la pagina de login")
}
func (s *Server) RegisterGet(w http.ResponseWriter, r *http.Request) {
	views.RegisterPage("").Render(r.Context(), w)
	log.Println("Se mostro la pagina de register")
}
func (s *Server) RegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		views.RegisterPage("Error leyendo formulario").Render(r.Context(), w)
		return
	}

	nombre := strings.TrimSpace(r.FormValue("nombre")) // Asegúrate de que el formulario tenga un campo 'nombre'
	email := strings.TrimSpace(r.FormValue("email"))
	pass := r.FormValue("contraseña")

	if nombre == "" || email == "" || pass == "" {
		views.RegisterPage("Complete todos los campos").Render(r.Context(), w)
		log.Println("Se intentó registrar con un campo incompleto")
		return
	}

	// Parámetros para la función de SQLC (adaptar según tu struct)
	params := sqlc.CreateUsuarioParams{
		Nombre:     nombre,
		Email:      email,
		Contraseña: pass,
	}

	// Crear usuario en la base de datos
	_, err := s.Queries.CreateUsuario(context.Background(), params)
	if err != nil {
		// Manejo básico de error (ej: el email ya existe)
		if strings.Contains(err.Error(), "duplicate key") {
			views.RegisterPage("El email ya está registrado").Render(r.Context(), w)
			log.Printf("Intento de registro con email duplicado: %s", email)
			return
		}
		views.RegisterPage("Error al crear el usuario").Render(r.Context(), w)
		log.Printf("Error al crear usuario: %v", err)
		return
	}

	log.Printf("Usuario registrado exitosamente: %s", email)

	// Redirigir al login para que el usuario inicie sesión
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// POST /login
// Este login es básico: busca por email y compara contraseña literal (NO recomendado para producción).
// Recomendación: guardar hashes bcrypt y comparar con bcrypt.CompareHashAndPassword.
func (s *Server) LoginPost(w http.ResponseWriter, r *http.Request) {
	log.Println("entro a login")
	if err := r.ParseForm(); err != nil {
		views.LoginPage("Error leyendo formulario").Render(r.Context(), w)
		return
	}

	email := strings.TrimSpace(r.FormValue("email"))
	pass := r.FormValue("contraseña")

	if email == "" || pass == "" {
		views.LoginPage("Complete todos los campos").Render(r.Context(), w)
		log.Println("Se intento ingresar con un campo incompleto")
		return
	}

	// Buscar usuario por email
	user, err := s.Queries.GetUsuarioByEmail(context.Background(), email)
	if err != nil {
		views.LoginPage("Usuario/Email no encontrado").Render(r.Context(), w)
		log.Println("Se intento entrar con credenciales que no se encuentran")
		return
	}

	// Comparación segura (literal); en producción comparar hashes bcrypt
	if subtle.ConstantTimeCompare([]byte(user.Contraseña), []byte(pass)) != 1 {
		views.LoginPage("Credenciales incorrectas").Render(r.Context(), w)
		log.Println("Se intento entrar con credenciales incorrectas")
		return
	}
	url := fmt.Sprintf("/entrenamientos/%d", user.IDUsuario)
	// TODO: crear sesión/cookie segura. Por ahora redirigir a entrenamientos
	http.Redirect(w, r, url, http.StatusSeeOther)
}
