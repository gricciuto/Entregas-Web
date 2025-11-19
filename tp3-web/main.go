package main

import (
	"context"
	"database/sql"

	"fmt"
	"log"
	"net/http"
	"os"

	"time"
	handlers "tp3-web/handlers"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	sqlc "tp3-web/db/sqlc"
)

// Estructura global para compartir la conexión y queries

func main() {
	// Cargar variables de entorno
	_ = godotenv.Load(".env")
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

	server := &handlers.Server{
		DB:      db,
		Queries: sqlc.New(db),
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handler para la página principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.LoginGet(w, r)
		case http.MethodPost:
			server.LoginPost(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/entrenamientos/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			server.DeleteEntrenamiento(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/entrenamientos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.EntrenamientosPage(w, r)
		case http.MethodPost:
			server.CreateEntrenamiento(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
