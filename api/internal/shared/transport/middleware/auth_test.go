package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/token"
)

func TestRequireAuth(t *testing.T) {
	cfg := &config.Config{JWTAccessSecret: "test-access-secret", JWTRefreshSecret: "test-refresh-secret"}
	id := uuid.New()
	generate := func(secret string, expiry int) string {
		t.Helper()
		value, err := token.GenerateToken(id, secret, expiry)
		if err != nil {
			t.Fatalf("generate token: %v", err)
		}
		return value
	}
	access := generate(cfg.JWTAccessSecret, 600)
	for _, tc := range []struct {
		name   string
		header string
		valid  bool
	}{
		{name: "missing header"},
		{name: "wrong scheme", header: "Basic " + access},
		{name: "missing token", header: "Bearer"},
		{name: "blank token", header: "Bearer   "},
		{name: "malformed token", header: "Bearer not-a-jwt"},
		{name: "wrong secret", header: "Bearer " + generate("wrong-secret", 600)},
		{name: "expired token", header: "Bearer " + generate(cfg.JWTAccessSecret, -60)},
		{name: "refresh token", header: "Bearer " + generate(cfg.JWTRefreshSecret, 86400)},
		{name: "extra credentials", header: "Bearer " + access + " extra"},
		{name: "valid access token", header: "Bearer " + access, valid: true},
		{name: "case insensitive scheme", header: "bEaReR " + access, valid: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.New()
			called := false
			var ctx *gin.Context
			r.Use(func(c *gin.Context) { ctx = c; c.Next() })
			group := r.Group("/protected")
			group.Use(RequireAuth(cfg))
			group.POST("/:user_id", func(c *gin.Context) {
				called = true
				got, ok := CurrentUserID(c)
				if !ok || got != id {
					t.Errorf("CurrentUserID = (%v, %t), want (%v, true)", got, ok, id)
				}
				c.Status(http.StatusNoContent)
			})
			// Untrusted identity inputs must neither authenticate nor override the subject.
			spoofed := uuid.New().String()
			req := httptest.NewRequest(http.MethodPost, "/protected/"+spoofed+"?user_id="+spoofed,
				strings.NewReader(`{"user_id":"`+spoofed+`"}`))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("user_id", spoofed)
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)
			if called != tc.valid {
				t.Fatalf("protected handler called = %t, want %t", called, tc.valid)
			}
			if tc.valid {
				if recorder.Code != http.StatusNoContent || ctx.IsAborted() {
					t.Fatalf("valid request: status = %d, aborted = %t", recorder.Code, ctx.IsAborted())
				}
				return
			}
			if recorder.Code != http.StatusUnauthorized || !ctx.IsAborted() {
				t.Fatalf("invalid request: status = %d, aborted = %t", recorder.Code, ctx.IsAborted())
			}
			var body response.Envelope
			if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode error envelope: %v", err)
			}
			if body.Success || body.Data != nil || body.Error == nil || body.Error.Code != apperror.CodeInvalidToken || body.Error.Message == "" {
				t.Fatalf("unexpected error envelope: %s", recorder.Body.String())
			}
			if got, ok := CurrentUserID(ctx); ok || got != uuid.Nil {
				t.Fatalf("failed authentication stored identity: (%v, %t)", got, ok)
			}
		})
	}
}

func TestCurrentUserID(t *testing.T) {
	id := uuid.New()
	for _, tc := range []struct {
		name  string
		value any
		set   bool
		valid bool
	}{
		{name: "missing"},
		{name: "nil value", set: true},
		{name: "string UUID", value: id.String(), set: true},
		{name: "wrong type", value: 123, set: true},
		{name: "UUID pointer", value: &id, set: true},
		{name: "UUID", value: id, set: true, valid: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			if tc.set {
				c.Set(currentUserIDKey, tc.value)
			}
			want := uuid.Nil
			if tc.valid {
				want = id
			}
			if got, ok := CurrentUserID(c); got != want || ok != tc.valid {
				t.Fatalf("CurrentUserID = (%v, %t), want (%v, %t)", got, ok, want, tc.valid)
			}
		})
	}
	if got, ok := CurrentUserID(nil); got != uuid.Nil || ok {
		t.Fatalf("CurrentUserID(nil) = (%v, %t), want (uuid.Nil, false)", got, ok)
	}
}
