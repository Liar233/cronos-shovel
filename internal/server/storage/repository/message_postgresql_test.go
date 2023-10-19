package repository

import (
	"database/sql"
	"testing"
)

func TestMessagePostgresqlRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

func TestMessagePostgresqlRepository_FindOne(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

func TestMessagePostgresqlRepository_GetList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

func TestMessagePostgresqlRepository_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

func TestMessagePostgresqlRepository_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

type storageMock struct {
	db *sql.DB
}

func newStorageMock(db *sql.DB) *storageMock {
	return &storageMock{
		db: db,
	}
}

func (s *storageMock) Connect() error {
	return nil
}

func (s *storageMock) Close() error {
	return nil
}

func (s *storageMock) GetDB() *sql.DB {
	return s.db
}
