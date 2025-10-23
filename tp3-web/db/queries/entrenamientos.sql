
-- Entrenamientos

-- name: CreateEntrenamiento :one
INSERT INTO entrenamiento (usuario_id, fecha, tipo, distancia, tiempo, ritmo_promedio, calorias, notas)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetEntrenamientosByUsuario :many
SELECT * FROM entrenamiento
WHERE usuario_id = $1
ORDER BY fecha DESC;

-- name: GetEntrenamientoByID :one
SELECT * FROM entrenamiento
WHERE id_entrenamiento = $1;

-- name: UpdateEntrenamiento :one
UPDATE entrenamiento
SET fecha = $2, tipo = $3, distancia = $4, tiempo = $5, ritmo_promedio = $6, calorias = $7, notas = $8
WHERE id_entrenamiento = $1
RETURNING *;

-- name: GetEntrenamientos :many
SELECT * FROM entrenamiento;

-- name: DeleteEntrenamiento :exec
DELETE FROM entrenamiento
WHERE id_entrenamiento = $1;
