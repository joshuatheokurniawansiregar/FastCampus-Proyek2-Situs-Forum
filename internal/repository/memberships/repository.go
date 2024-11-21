package memberships

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(sql *sql.DB) *repository {
	return &repository{
		db: sql,
	}
}


