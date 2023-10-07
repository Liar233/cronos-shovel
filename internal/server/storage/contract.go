package storage

import "database/sql"

type ConnectorInterface interface {
	Connect() error
	Close() error
	GetDB() *sql.DB
}
