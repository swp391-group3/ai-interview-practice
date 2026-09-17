//go:build integration

// Package testutil supplies isolated infrastructure for repository integration tests.
package testutil

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// Postgres holds a migrated, disposable database. Pool closes before its container.
type Postgres struct {
	DSN  string
	Pool *pgxpool.Pool
}

// NewPostgres starts PostgreSQL 18.6 and applies the checked-in baseline migration.
// Call once per test (or parent test); all resources are registered with t.Cleanup.
func NewPostgres(t testing.TB) *Postgres {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	container, err := postgres.Run(ctx, "postgres:18.6",
		postgres.WithDatabase("ai_interview_practice"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test-only-password"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("start PostgreSQL: %v", err)
	}
	testcontainers.CleanupContainer(t, container)
	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("PostgreSQL DSN: %v", err)
	}
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("connect PostgreSQL: %v", err)
	}
	t.Cleanup(pool.Close)
	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("ping PostgreSQL: %v", err)
	}
	_, source, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate baseline migration")
	}
	migration, err := os.ReadFile(filepath.Join(filepath.Dir(source), "../../migration/000001_init.up.sql"))
	if err != nil {
		t.Fatalf("read baseline migration: %v", err)
	}
	if _, err := pool.Exec(ctx, string(migration)); err != nil {
		t.Fatalf("apply baseline migration: %v", err)
	}
	return &Postgres{DSN: dsn, Pool: pool}
}
