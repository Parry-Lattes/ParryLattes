package repository

import (
	"database/sql"
)

type Repository struct {
	Connection *sql.DB
}
