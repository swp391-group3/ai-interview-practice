//go:build integration

package repository_test

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/domain"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/repository"
	"github.com/swp391-group3/ai-interview-practice/api/internal/testutil"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
)

func TestJDRepository(t *testing.T) {
	db := testutil.NewPostgres(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	owner, other := uuid.New(), uuid.New()
	for _, id := range []uuid.UUID{owner, other} {
		if _, err := db.Pool.Exec(ctx, "INSERT INTO accounts (id, email, full_name, password_hash) VALUES ($1, $2, 'Test', 'unused')", id, id.String()+"@example.com"); err != nil {
			t.Fatal(err)
		}
	}
	repo := repository.NewRepository(db.Pool)
	seniority := domain.Senior
	input := domain.CreateInput{RawText: "Original JD text", StructuredJD: domain.StructuredJD{
		Title: "Backend Engineer", SeniorityLevel: &seniority,
		Skills:       []domain.ExtractedSkill{{Name: "Go", Category: domain.ProgrammingLanguage, Requirement: domain.Required}},
		Technologies: []string{"PostgreSQL"}, DomainKnowledge: []string{"Payments"},
	}}
	invalid := input
	invalid.StructuredJD.SeniorityLevel = nil
	if _, err := repo.Create(ctx, owner, invalid); err == nil {
		t.Fatal("accepted missing seniority")
	}
	invalidSeniority := domain.Seniority("unknown")
	invalid.StructuredJD.SeniorityLevel = &invalidSeniority
	if _, err := repo.Create(ctx, owner, invalid); err == nil {
		t.Fatal("accepted invalid seniority")
	}
	for _, title := range []string{"", " \t"} {
		invalidTitle := input
		invalidTitle.StructuredJD.Title = title
		if _, err := repo.Create(ctx, owner, invalidTitle); err == nil {
			t.Fatalf("accepted blank title %q", title)
		} else {
			var appErr *apperror.AppError
			if !errors.As(err, &appErr) || appErr.Code != apperror.CodeValidation {
				t.Fatalf("want validation error for blank title, got %v", err)
			}
		}
	}

	created, err := repo.Create(ctx, owner, input)
	if err != nil {
		t.Fatal(err)
	}
	if created.ID == uuid.Nil || created.UserID != owner || created.RawText != input.RawText || created.Status != domain.StatusCustomized ||
		!reflect.DeepEqual(created.StructuredJD, input.StructuredJD) || created.CreatedAt.IsZero() || created.UpdatedAt.IsZero() {
		t.Fatalf("create did not preserve reviewed JD: %+v", created)
	}
	checkJSON := func(id uuid.UUID, want domain.StructuredJD) {
		t.Helper()
		var data []byte
		if err := db.Pool.QueryRow(ctx, "SELECT parsed_data FROM job_descriptions WHERE id = $1", id).Scan(&data); err != nil {
			t.Fatal(err)
		}
		var got map[string]json.RawMessage
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if len(got) != 3 {
			t.Fatalf("unexpected parsed_data keys: %s", data)
		}
		expected := map[string]any{"skills": want.Skills, "technologies": want.Technologies, "domainKnowledge": want.DomainKnowledge}
		for key, value := range expected {
			encoded, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var gotValue, wantValue any
			if err := json.Unmarshal(got[key], &gotValue); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(encoded, &wantValue); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(gotValue, wantValue) {
				t.Fatalf("parsed_data %s mismatch", key)
			}
		}
	}
	checkJSON(created.ID, input.StructuredJD)
	got, err := repo.Get(ctx, owner, created.ID)
	if err != nil || !reflect.DeepEqual(got, created) {
		t.Fatalf("owner get: %+v, %v", got, err)
	}
	assertNotFound := func(err error) {
		t.Helper()
		var appErr *apperror.AppError
		if !errors.As(err, &appErr) || appErr.Code != apperror.CodeJDNotFound {
			t.Fatalf("want JD not found, got %v", err)
		}
	}
	_, err = repo.Get(ctx, other, created.ID)
	assertNotFound(err)
	otherJD, err := repo.Create(ctx, other, input)
	if err != nil {
		t.Fatal(err)
	}
	for user, id := range map[uuid.UUID]uuid.UUID{owner: created.ID, other: otherJD.ID} {
		listed, err := repo.List(ctx, user)
		if err != nil || len(listed) != 1 || listed[0].ID != id || listed[0].UserID != user {
			t.Fatalf("list ownership: %+v, %v", listed, err)
		}
	}
	updatedInput := domain.UpdateInput{StructuredJD: input.StructuredJD}
	level := domain.Lead
	updatedInput.StructuredJD.Title = "Lead Engineer"
	updatedInput.StructuredJD.SeniorityLevel = &level
	updatedInput.StructuredJD.Skills = []domain.ExtractedSkill{{Name: "SQL", Category: domain.Database, Requirement: domain.Preferred}}
	updatedInput.StructuredJD.Technologies = []string{}
	updatedInput.StructuredJD.DomainKnowledge = []string{"Banking"}
	_, err = repo.Update(ctx, other, created.ID, updatedInput)
	assertNotFound(err)
	assertNotFound(repo.Delete(ctx, other, created.ID))
	got, err = repo.Get(ctx, owner, created.ID)
	if err != nil || !reflect.DeepEqual(got, created) {
		t.Fatalf("unauthorized mutation changed JD: %+v, %v", got, err)
	}
	updated, err := repo.Update(ctx, owner, created.ID, updatedInput)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(updated.StructuredJD, updatedInput.StructuredJD) || updated.RawText != input.RawText ||
		!updated.CreatedAt.Equal(created.CreatedAt) || !updated.UpdatedAt.After(created.UpdatedAt) {
		t.Fatalf("owner update mismatch: %+v", updated)
	}
	checkJSON(created.ID, updatedInput.StructuredJD)
	got, err = repo.Get(ctx, owner, created.ID)
	if err != nil || !reflect.DeepEqual(got, updated) {
		t.Fatalf("update not persisted: %+v, %v", got, err)
	}
	if _, err := db.Pool.Exec(ctx, "INSERT INTO interview_blueprints (job_description_id, difficulty, duration_minutes, question_count, blueprint_data, contract_version) VALUES ($1, 'medium', 30, 5, '{}', 1)", created.ID); err != nil {
		t.Fatal(err)
	}
	assertNotFound(repo.Delete(ctx, other, created.ID))
	var inUse *apperror.AppError
	if err := repo.Delete(ctx, owner, created.ID); !errors.As(err, &inUse) || inUse.Code != apperror.CodeJDInUse {
		t.Fatalf("want JD_IN_USE, got %v", err)
	}
	if _, err := db.Pool.Exec(ctx, "DELETE FROM interview_blueprints WHERE job_description_id = $1", created.ID); err != nil {
		t.Fatal(err)
	}
	if err := repo.Delete(ctx, owner, created.ID); err != nil {
		t.Fatal(err)
	}
	_, err = repo.Get(ctx, owner, created.ID)
	assertNotFound(err)
	assertNotFound(repo.Delete(ctx, owner, created.ID))
	listed, err := repo.List(ctx, owner)
	if err != nil || len(listed) != 0 {
		t.Fatalf("list after delete: %+v, %v", listed, err)
	}
	if _, err := repo.Get(ctx, other, otherJD.ID); err != nil {
		t.Fatal(err)
	}
}
