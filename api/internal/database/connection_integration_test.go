//go:build integration

package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/swp391-group3/ai-interview-practice/api/internal/testutil"
)

func TestPostgresBaseline(t *testing.T) {
	db := testutil.NewPostgres(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, db.DSN)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = conn.Close(context.Background()) }()
	var database, version string
	if err := conn.QueryRow(ctx, "SELECT current_database(), current_setting('server_version_num')").Scan(&database, &version); err != nil {
		t.Fatal(err)
	}
	if database != "ai_interview_practice" || version != "180006" {
		t.Fatalf("database=%s version=%s", database, version)
	}
	for _, table := range []string{"accounts", "technical_domains", "skills", "job_descriptions", "job_description_skills", "avatar_profiles", "interview_sessions", "session_turns", "performance_reports"} {
		var exists bool
		if err := conn.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1)", table).Scan(&exists); err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Errorf("missing baseline table %s", table)
		}
	}
}
