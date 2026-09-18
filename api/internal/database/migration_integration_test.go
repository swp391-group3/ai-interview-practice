//go:build integration

package database_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/swp391-group3/ai-interview-practice/api/internal/testutil"
)

func TestInterviewBlueprintMigration(t *testing.T) {
	// The shared helper applies 000001 up followed by 000002 up.
	db := testutil.NewPostgres(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	hasColumn := func(table, column string, want bool) {
		t.Helper()
		var exists bool
		err := db.Pool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_schema = 'public' AND table_name = $1 AND column_name = $2)", table, column).Scan(&exists)
		if err != nil || exists != want {
			t.Fatalf("%s.%s exists=%v want=%v: %v", table, column, exists, want, err)
		}
	}
	hasColumn("job_descriptions", "blueprint", false)
	hasColumn("interview_sessions", "blueprint_id", false)
	hasColumn("interview_sessions", "blueprint_snapshot", true)
	var nullable string
	if err := db.Pool.QueryRow(ctx, "SELECT is_nullable FROM information_schema.columns WHERE table_name = 'job_descriptions' AND column_name = 'seniority_level'").Scan(&nullable); err != nil || nullable != "NO" {
		t.Fatalf("seniority nullability=%s: %v", nullable, err)
	}
	var user, jd uuid.UUID
	if err := db.Pool.QueryRow(ctx, "INSERT INTO accounts (email, full_name, password_hash) VALUES ('migration@example.com', 'Test', 'unused') RETURNING id").Scan(&user); err != nil {
		t.Fatal(err)
	}
	if err := db.Pool.QueryRow(ctx, "INSERT INTO job_descriptions (user_id, title, seniority_level, raw_text) VALUES ($1, 'Engineer', 'senior', 'Text') RETURNING id", user).Scan(&jd); err != nil {
		t.Fatal(err)
	}
	for range 2 {
		if _, err := db.Pool.Exec(ctx, "INSERT INTO interview_blueprints (job_description_id, difficulty, duration_minutes, question_count, blueprint_data, contract_version) VALUES ($1, 'medium', 30, 5, '{}', 1)", jd); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := db.Pool.Exec(ctx, "DELETE FROM job_descriptions WHERE id = $1", jd); err == nil {
		t.Fatal("blueprint did not restrict JD deletion")
	}
	down, err := os.ReadFile("../../migration/000002_interview_blueprints.down.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Pool.Exec(ctx, string(down)); err != nil {
		t.Fatal(err)
	}
	hasColumn("job_descriptions", "blueprint", true)
	hasColumn("interview_blueprints", "id", false)
	hasColumn("interview_sessions", "blueprint_snapshot", true)
	var count int
	if err := db.Pool.QueryRow(ctx, "SELECT count(*) FROM job_descriptions WHERE id = $1", jd).Scan(&count); err != nil || count != 1 {
		t.Fatalf("rollback changed existing JD: %d, %v", count, err)
	}
	up, err := os.ReadFile("../../migration/000002_interview_blueprints.up.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Pool.Exec(ctx, string(up)); err != nil {
		t.Fatal(err)
	}
}
