package ai

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/gemini"
	"github.com/cloudwego/eino/components/model"
	"google.golang.org/genai"

	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
)

func NewChatModel(ctx context.Context, cfg config.Config) (model.BaseChatModel, error) {
	if cfg.LLMProvider != "gemini" {
		return nil, fmt.Errorf("unsupported LLM provider %q: expected gemini", cfg.LLMProvider)
	}
	if strings.TrimSpace(cfg.LLMAPIKey) == "" || strings.ContainsAny(cfg.LLMAPIKey, "\r\n") {
		return nil, fmt.Errorf("LLM API key is required and must not contain newlines")
	}
	if strings.TrimSpace(cfg.LLMModel) == "" {
		return nil, fmt.Errorf("LLM model is required")
	}
	if cfg.LLMTimeout <= 0 {
		return nil, fmt.Errorf("LLM timeout must be positive")
	}
	attempts := int32(1)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend:    genai.BackendGeminiAPI,
		APIKey:     cfg.LLMAPIKey,
		HTTPClient: &http.Client{Timeout: cfg.LLMTimeout, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }},
		HTTPOptions: genai.HTTPOptions{
			Timeout: &cfg.LLMTimeout,
			// The application owns retries; one Generate must make one API attempt.
			RetryOptions: &genai.HTTPRetryOptions{Attempts: &attempts},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create Gemini client: %w", err)
	}
	return gemini.NewChatModel(ctx, &gemini.Config{Client: client, Model: cfg.LLMModel})
}
