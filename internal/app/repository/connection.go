package repository

import "database/sql"

type Connection interface {
	GetConnection() (*sql.DB, error)
	ReleaseConnection(conn *sql.DB)
}
