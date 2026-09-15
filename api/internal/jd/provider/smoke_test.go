//go:build llm_smoke

package provider

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env/v11"

	"github.com/swp391-group3/ai-interview-practice/api/internal/jd/service"
	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/ai"
	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
)

// TestLLMSmoke is explicitly opt-in and never compiled in the default test suite.
// Use only a synthetic JD; never print credentials or raw request headers.
func TestLLMSmoke(t *testing.T) {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		t.Fatal("invalid environment configuration; check LLM_TIMEOUT and LLM_MAX_RETRIES")
	}
	chatModel, err := ai.NewChatModel(context.Background(), cfg)
	if err != nil {
		t.Fatal("invalid Gemini configuration; check provider, API key, model and timeout")
	}
	adapter, err := New(chatModel)
	if err != nil {
		t.Fatal("invalid JD adapter configuration")
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
		t.Fatal("Gemini extraction failed; provider details omitted to protect credentials")
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
