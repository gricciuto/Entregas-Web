package handlers

import (
	"database/sql"

	sqlc "tp3-web/db/sqlc"
)

type Server struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}
