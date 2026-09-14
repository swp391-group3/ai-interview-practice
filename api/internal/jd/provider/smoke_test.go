//go:build llm_smoke

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env/v11"

	"github.com/swp391-group3/ai-interview-practice/api/internal/jd/service"
	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
)

// TestLLMSmoke is explicitly opt-in and never compiled in the default test suite.
// Use only a synthetic JD; never print credentials or raw request headers.
func TestLLMSmoke(t *testing.T) {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		t.Fatal("invalid environment configuration; check LLM_TIMEOUT and LLM_MAX_RETRIES")
	}
	if cfg.LLMProvider != "openai_compatible" {
		t.Fatal("smoke test requires LLM_PROVIDER=openai_compatible")
	}
	adapter, err := New(Config{BaseURL: cfg.LLMBaseURL, APIKey: cfg.LLMAPIKey, Model: cfg.LLMModel, Timeout: cfg.LLMTimeout})
	if err != nil {
		t.Fatal("invalid LLM adapter configuration; check base URL, API key, model and timeout")
	}
	s, err := service.New(adapter, cfg.LLMMaxRetries)
	if err != nil {
		t.Fatal("invalid LLM retry budget")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*cfg.LLMTimeout+time.Second)
	defer cancel()
	jd := `We are hiring a Senior Backend Engineer for a financial services platform.
Required technical requirements include Go, PostgreSQL, Kafka, API design, and
automated testing. Docker is preferred, and experience with Spring Boot is valued.
Knowledge of payment processing and financial transaction systems is required to
build reliable backend services for high-volume financial workflows.`
	got, err := s.Extract(ctx, jd)
	if err != nil {
		var providerErr *HTTPError
		if errors.As(err, &providerErr) {
			// The body is preserved so account/billing/model failures can be
			// reported accurately. Remove any echoed key before displaying it.
			body := strings.ReplaceAll(providerErr.Body, cfg.LLMAPIKey, "[REDACTED]")
			if encoded, marshalErr := json.Marshal(cfg.LLMAPIKey); marshalErr == nil {
				body = strings.ReplaceAll(body, string(encoded[1:len(encoded)-1]), "[REDACTED]")
			}
			t.Fatalf("LLM provider returned HTTP %d; provider response: %s", providerErr.StatusCode, body)
		}
		t.Fatalf("LLM extraction failed: %v", err)
	}
	body, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"LLM authentication, configured model and extraction succeeded:\n%s",
		strings.ReplaceAll(string(body), cfg.LLMAPIKey, "[REDACTED]"),
	)
}
