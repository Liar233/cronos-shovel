package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type PostgresStorageConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type PostgresStorage struct {
	db     *sql.DB
	config *PostgresStorageConfig
}

func NewPostgresStorage(config *PostgresStorageConfig) *PostgresStorage {
	return &PostgresStorage{
		config: config,
	}
}

func (ps *PostgresStorage) Connect() error {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ps.config.Host,
		ps.config.Port,
		ps.config.User,
		ps.config.Password,
		ps.config.Database,
	)

	var err error

	ps.db, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	if err = ps.db.Ping(); err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStorage) Close() error {

	return ps.db.Close()
}

func (ps *PostgresStorage) GetDB() *sql.DB {

	return ps.db
}
