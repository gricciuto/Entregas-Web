-- Usuarios

-- name: CreateUsuario :one
INSERT INTO usuario (nombre, email, contraseña)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUsuarioByEmail :one
SELECT * FROM usuario
WHERE email = $1;

-- name: GetUsuarioByID :one
SELECT * FROM usuario
WHERE id_usuario = $1;

-- name: GetUsuarios :many
SELECT * FROM usuario;

-- name: DeleteUsuario :exec
DELETE FROM usuario
WHERE id_usuario = $1;