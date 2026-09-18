package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/domain"
	"github.com/swp391-group3/ai-interview-practice/api/internal/middleware"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/token"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeJDService struct {
	JDService
	calls         int
	user          uuid.UUID
	err           error
	limit, offset int32
	items         []domain.ListItem
}

func (f *fakeJDService) List(_ context.Context, u uuid.UUID, limit, offset int32) ([]domain.ListItem, int64, error) {
	f.calls++
	f.user, f.limit, f.offset = u, limit, offset
	return f.items, 123, f.err
}

func TestJDListPagination(t *testing.T) {
	user := uuid.New()
	cfg := &config.Config{JWT: config.JWTConfig{AccessSecret: "access-test", RefreshSecret: "refresh-test"}}
	jwt, err := token.GenerateToken(user, cfg.JWT.AccessSecret, 600)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		query         string
		limit, offset int32
		invalid       bool
	}{
		{"", 20, 0, false}, {"?limit=100&offset=12", 100, 12, false},
		{"?limit=1&offset=2147483647", 1, 2147483647, false},
		{"?limit=0", 0, 0, true}, {"?limit=101", 0, 0, true},
		{"?limit=-1", 0, 0, true}, {"?offset=-1", 0, 0, true},
		{"?limit=abc", 0, 0, true}, {"?offset=1.5", 0, 0, true},
		{"?limit=", 0, 0, true}, {"?offset=", 0, 0, true},
		{"?offset=2147483648", 0, 0, true}, {"?limit=1&limit=2", 0, 0, true},
	} {
		t.Run(tc.query, func(t *testing.T) {
			f := &fakeJDService{items: []domain.ListItem{{ID: uuid.New(), Title: "Engineer", SeniorityLevel: domain.Senior, Status: domain.StatusCustomized}}}
			r := gin.New()
			r.Use(middleware.RequireAuth(cfg))
			r.GET("/jds", NewJDHandler(f).List)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/jds"+tc.query, nil)
			req.Header.Set("Authorization", "Bearer "+jwt)
			r.ServeHTTP(w, req)
			if tc.invalid {
				var body response.Envelope
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
					t.Fatal(err)
				}
				if w.Code != 400 || f.calls != 0 || body.Error == nil || body.Error.Code != apperror.CodeValidation {
					t.Fatalf("invalid response: %s", w.Body.String())
				}
				return
			}
			var body struct {
				Success bool                          `json:"success"`
				Data    response.Page[map[string]any] `json:"data"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if w.Code != 200 || !body.Success || f.user != user || f.limit != tc.limit || f.offset != tc.offset || body.Data.Pagination != (response.Pagination{Limit: tc.limit, Offset: tc.offset, Total: 123}) {
				t.Fatalf("unexpected response: %s", w.Body.String())
			}
			if len(body.Data.Items) != 1 || len(body.Data.Items[0]) != 6 {
				t.Fatalf("unexpected fields: %+v", body.Data.Items)
			}
			for _, key := range []string{"id", "title", "seniorityLevel", "status", "createdAt", "updatedAt"} {
				if _, ok := body.Data.Items[0][key]; !ok {
					t.Fatalf("missing %s", key)
				}
			}
		})
	}
}

func (f *fakeJDService) Get(_ context.Context, u, _ uuid.UUID) (domain.JD, error) {
	f.calls++
	f.user = u
	return domain.JD{}, f.err
}
func (f *fakeJDService) Update(_ context.Context, u, _ uuid.UUID, in domain.UpdateInput) (domain.JD, error) {
	f.calls++
	f.user = u
	return domain.JD{RawText: "original", StructuredJD: in.StructuredJD}, f.err
}
func TestJDTransport(t *testing.T) {
	cfg := &config.Config{JWT: config.JWTConfig{AccessSecret: "access-test", RefreshSecret: "refresh-test"}}
	user := uuid.New()
	jwt, err := token.GenerateToken(user, cfg.JWT.AccessSecret, 600)
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New().String()
	for _, tc := range []struct {
		name, method, path, body string
		status                   int
		code                     apperror.Code
		calls                    int
		noAuth                   bool
		err                      error
	}{
		{"uuid", "GET", "/jds/bad", "", 400, apperror.CodeValidation, 0, false, nil},
		{"json", "PUT", "/jds/" + id, "{", 400, apperror.CodeValidation, 0, false, nil},
		{"missing shape", "PUT", "/jds/" + id, "{}", 400, apperror.CodeValidation, 0, false, nil},
		{"immutable raw", "PUT", "/jds/" + id, `{"structuredJD":{},"rawText":"replace"}`, 400, apperror.CodeValidation, 0, false, nil},
		{"nested raw", "PUT", "/jds/" + id, `{"structuredJD":{"rawText":"replace"}}`, 400, apperror.CodeValidation, 0, false, nil},
		{"client ownership", "PUT", "/jds/" + id, `{"structuredJD":{},"user_id":"fake"}`, 400, apperror.CodeValidation, 0, false, nil},
		{"missing", "GET", "/jds/" + id, "", 404, apperror.CodeJDNotFound, 1, false, apperror.New(apperror.CodeJDNotFound, "job description not found")},
		{"update user", "PUT", "/jds/" + id, `{"structuredJD":{"title":"reviewed"}}`, 200, "", 1, false, nil},
		{"auth", "GET", "/jds/" + id, "", 401, apperror.CodeInvalidToken, 0, true, nil},
		{"empty list", "GET", "/jds", "", 200, "", 1, false, nil},
		{"list error", "GET", "/jds", "", 500, apperror.CodeInternal, 1, false, apperror.New(apperror.CodeInternal, "persistence failed")},
		{"list auth", "GET", "/jds", "", 401, apperror.CodeInvalidToken, 0, true, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := &fakeJDService{err: tc.err}
			h := NewJDHandler(f)
			r := gin.New()
			r.Use(middleware.RequireAuth(cfg))
			r.GET("/jds/:id", h.Get)
			r.GET("/jds", h.List)
			r.PUT("/jds/:id", h.Update)
			req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			if !tc.noAuth {
				req.Header.Set("Authorization", "Bearer "+jwt)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.status || f.calls != tc.calls {
				t.Fatalf("status=%d calls=%d body=%s", w.Code, f.calls, w.Body.String())
			}
			if f.calls > 0 && f.user != user {
				t.Fatal("authenticated identity not forwarded")
			}
			var body response.Envelope
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if tc.code != "" && (body.Error == nil || body.Error.Code != tc.code) {
				t.Fatalf("wrong envelope: %+v", body)
			}
			if strings.Contains(w.Body.String(), `"userId"`) {
				t.Fatal("public response exposes userId")
			}
			if tc.status == 200 && tc.method == "PUT" && !strings.Contains(w.Body.String(), `"rawText":"original"`) {
				t.Fatal("raw text changed")
			}
			if tc.name == "empty list" && !strings.Contains(w.Body.String(), `"items":[]`) {
				t.Fatal("empty list must be an array")
			}
		})
	}
}
