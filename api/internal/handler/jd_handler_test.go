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
	calls int
	user  uuid.UUID
	err   error
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
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := &fakeJDService{err: tc.err}
			h := NewJDHandler(f)
			r := gin.New()
			r.Use(middleware.RequireAuth(cfg))
			r.GET("/jds/:id", h.Get)
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
			if tc.status == 200 && !strings.Contains(w.Body.String(), `"rawText":"original"`) {
				t.Fatal("raw text changed")
			}
		})
	}
}
