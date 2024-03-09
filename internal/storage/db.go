package storage

import (
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) *Postgres {
	return &Postgres{
		conn: conn,
	}
}
