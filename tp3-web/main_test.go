package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"
	sqlc "tp3-web/db/sqlc"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func setupDB(t *testing.T) (*sql.DB, *sqlc.Queries) {
	t.Helper()
	_ = godotenv.Load("credentials.env")
	user := os.Getenv("POSTGRES_ADMIN")
	pass := os.Getenv("POSTGRES_ADMIN_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")
	connStr := "postgresql://" + user + ":" + pass + "@localhost:5432/" + dbname + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Error conectando a la DB: %v", err)
	}
	return db, sqlc.New(db)
}

func cleanupUsuario(t *testing.T, db *sql.DB, email string) {
	t.Helper()
	_, _ = db.Exec("DELETE FROM entrenamiento WHERE usuario_id IN (SELECT id_usuario FROM usuario WHERE email = $1)", email)
	_, _ = db.Exec("DELETE FROM usuario WHERE email = $1", email)
}

func TestCrearYBuscarUsuario(t *testing.T) {
	ctx := context.Background()
	db, queries := setupDB(t)
	defer db.Close()

	testEmail := "hla@example.com"
	cleanupUsuario(t, db, testEmail)

	usuario, err := queries.CreateUsuario(ctx, sqlc.CreateUsuarioParams{
		Nombre:     "Carlos",
		Email:      testEmail,
		Contraseña: "upaaaa",
	})
	if err != nil {
		t.Fatalf("Error creando usuario: %v", err)
	}

	u, err := queries.GetUsuarioByEmail(ctx, testEmail)
	if err != nil {
		t.Fatalf("Error buscando usuario por email: %v", err)
	}
	if u.Email != usuario.Email {
		t.Fatalf("El email del usuario no coincide")
	}

	// Limpieza final
	cleanupUsuario(t, db, testEmail)
}

func TestCrearYListarEntrenamiento(t *testing.T) {
	ctx := context.Background()
	db, queries := setupDB(t)
	defer db.Close()

	testEmail := "hashdhs@example.com"
	cleanupUsuario(t, db, testEmail)

	usuario, err := queries.CreateUsuario(ctx, sqlc.CreateUsuarioParams{
		Nombre:     "carlos",
		Email:      testEmail,
		Contraseña: "1234",
	})
	if err != nil {
		t.Fatalf("Error creando usuario: %v", err)
	}

	entreno, err := queries.CreateEntrenamiento(ctx, sqlc.CreateEntrenamientoParams{
		UsuarioID:     usuario.IDUsuario,
		Fecha:         time.Now(),
		Tipo:          "Testing Run",
		Distancia:     fmt.Sprintf("%.2f", 10.0), // string
		Tiempo:        60,
		RitmoPromedio: sql.NullString{String: fmt.Sprintf("%.2f", 6.0), Valid: true}, // NullString
		Calorias:      sql.NullInt32{Int32: 300, Valid: true},
		Notas:         sql.NullString{String: "Test training", Valid: true},
	})
	if err != nil {
		t.Fatalf("Error creando entrenamiento: %v", err)
	}

	entrenamientos, err := queries.GetEntrenamientosByUsuario(ctx, usuario.IDUsuario)
	if err != nil {
		t.Fatalf("Error listando entrenamientos: %v", err)
	}

	if len(entrenamientos) == 0 {
		t.Fatalf("No se encontraron entrenamientos para el usuario")
	}

	found := false
	for _, e := range entrenamientos {
		if e.IDEntrenamiento == entreno.IDEntrenamiento {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Entrenamiento creado no encontrado en la lista")
	}

	// Limpieza final
	cleanupUsuario(t, db, testEmail)
}
