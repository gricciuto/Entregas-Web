package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	sqlc "tp3-web/db/sqlc"
	"tp3-web/logic"
)

// Estructura global para compartir la conexión y queries
type Server struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func main() {
	// Cargar variables de entorno
	_ = godotenv.Load("credentials.env")
	user := os.Getenv("POSTGRES_ADMIN")
	pass := os.Getenv("POSTGRES_ADMIN_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", user, pass, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error conectando a la DB: %v", err)
	}
	defer db.Close()

	// Verificar conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("No se pudo conectar a Postgres: %v", err)
	}

	server := &Server{
		db:      db,
		queries: sqlc.New(db),
	}

	// Registrar handlers
	http.HandleFunc("/usuarios", server.usuariosHandler)
	http.HandleFunc("/usuario/", server.usuarioIdHandler)
	http.HandleFunc("/entrenamientos", server.entrenamientosHandler)
	http.HandleFunc("/entrenamiento/", server.entrenamientosIdHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler para entrenamientos con ID
func (s *Server) entrenamientosIdHandler(w http.ResponseWriter, r *http.Request) {
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
	err := s.queries.DeleteEntrenamiento(r.Context(), int32(id))
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
	entrenamiento, err := s.queries.UpdateEntrenamiento(r.Context(), entrenamientoActualizadoSQL)
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
	entrenamiento, err := s.queries.GetEntrenamientoByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No se encontro un entrenamiento con ese ID de usuario"), http.StatusBadRequest)
		log.Print("Error, no se encontro un entrenamiento con el ID %i", id)
	}
	json.NewEncoder(w).Encode(entrenamiento)
	log.Printf("Se mostro el entrenamiento ID %v", id)
}

// Handler para entrenamientos
func (s *Server) entrenamientosHandler(w http.ResponseWriter, r *http.Request) {
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
	entrenamiento, err := s.queries.GetEntrenamientos(r.Context())
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
			String: fmt.Sprintf("%.2f", entrenamientoNuevo.Ritmo_promedio),
			Valid:  entrenamientoNuevo.Ritmo_promedio != 0,
		},
		Distancia: fmt.Sprintf("%.2f", entrenamientoNuevo.Distancia),
		Calorias: sql.NullInt32{
			Int32: int32(entrenamientoNuevo.Calorias),
			Valid: entrenamientoNuevo.Calorias != 0,
		},
		Notas: sql.NullString{
			String: entrenamientoNuevo.Notas,
			Valid:  entrenamientoNuevo.Notas != "",
		},
	}
	entrenamiento, err := s.queries.CreateEntrenamiento(r.Context(), entrenamientoNuevoSQL)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se pudo crear el entrenamiento"), http.StatusInternalServerError)
		log.Print("Error al crear el entrenamiento en la base de datos")
		return
	}
	json.NewEncoder(w).Encode(entrenamiento)
	log.Printf("Se creo el entrenamiento con ID %v", entrenamiento.IDEntrenamiento)
}

// Handler para Usuarios
func (s *Server) usuariosHandler(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) postUsuarios(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuario, err := s.queries.GetUsuarios(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer usuarios: %v", err), http.StatusInternalServerError)
		log.Print("Error de la base de datos al mostrar todos los usuarios")
		return
	}
	json.NewEncoder(w).Encode(usuario)
	log.Print("Se mostraron todos los usuarios")
}

// Handler de usuario por id
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
