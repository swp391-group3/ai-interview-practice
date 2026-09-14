package config

import (
	"testing"
	"time"

	"github.com/caarlos0/env/v11"
)

func TestLLMEnvironment(t *testing.T) {
	cfg, err := env.ParseAsWithOptions[Config](env.Options{Environment: map[string]string{}})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.LLMProvider != "opencode_zen" || cfg.LLMBaseURL != "https://opencode.ai/zen/v1" || cfg.LLMModel != "" || cfg.LLMAPIKey != "" || cfg.LLMTimeout != 30*time.Second || cfg.LLMMaxRetries != 1 {
		t.Fatal("unexpected LLM defaults")
	}
	cfg, err = env.ParseAsWithOptions[Config](env.Options{Environment: map[string]string{
		"LLM_PROVIDER": "openai_compatible", "LLM_BASE_URL": "http://localhost:9999/v1",
		"LLM_API_KEY": "test-key", "LLM_MODEL": "custom-model", "LLM_TIMEOUT": "2s", "LLM_MAX_RETRIES": "0",
		"DATABASE_URL": "postgres://local/example", "DB_DATABASE": "compose-only",
	}})
	if err != nil {
		t.Fatal(err)
	}
	if cfg.LLMProvider != "openai_compatible" || cfg.LLMBaseURL != "http://localhost:9999/v1" || cfg.LLMAPIKey != "test-key" || cfg.LLMModel != "custom-model" || cfg.LLMTimeout != 2*time.Second || cfg.LLMMaxRetries != 0 || cfg.DatabaseURL != "postgres://local/example" {
		t.Fatal("environment overrides were not parsed correctly")
	}
	for _, values := range []map[string]string{{"LLM_TIMEOUT": "invalid"}, {"LLM_MAX_RETRIES": "invalid"}} {
		if _, err := env.ParseAsWithOptions[Config](env.Options{Environment: values}); err == nil {
			t.Fatal("invalid environment accepted")
		}
	}
}
