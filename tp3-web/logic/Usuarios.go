package logic

import (
	"errors"
	"time"
)

type Usuario struct {
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Contraseña string `json:"contraseña"`
}
type Entrenamiento struct {
	IDusuario      int32     `json:"id_usuario"`
	Fecha          time.Time `json:"fecha"`
	Tipo           string    `json:"tipo"`
	Distancia      float64   `json:"distancia"`
	Tiempo         int       `json:"tiempo"`
	Ritmo_promedio float64   `json:"ritmo_promedio"`
	Calorias       int       `json:"calorias"`
	Notas          string    `json:"notas"`
}

func ValidarUsuario(usuario Usuario) error {
	if usuario.Nombre == "" || usuario.Email == "" || usuario.Contraseña == "" {
		return errors.New("Hay algun campo vacio")
	}
	return nil
}
func ValidarEntrenamiento(entrenamiento Entrenamiento) error {
	if entrenamiento.Fecha.IsZero() {
		entrenamiento.Fecha = time.Now()
	}
	if entrenamiento.Tipo == "" {
		return errors.New("Falta el tipo")
	}
	if entrenamiento.Distancia == 0 {
		return errors.New("Falta el distancia")
	}
	if entrenamiento.Tiempo == 0 {
		return errors.New("Falta el tiempo")
	}
	return nil
}
