package repository

import "testing"

func TestDelayPostgresqlRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}

func TestDelayPostgresqlRepository_Delete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}
