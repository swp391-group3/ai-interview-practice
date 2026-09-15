package config

import "time"

type Config struct {
	LLMProvider   string        `env:"LLM_PROVIDER" envDefault:"gemini"`
	LLMAPIKey     string        `env:"LLM_API_KEY"`
	LLMModel      string        `env:"LLM_MODEL"`
	LLMTimeout    time.Duration `env:"LLM_TIMEOUT" envDefault:"30s"`
	LLMMaxRetries int           `env:"LLM_MAX_RETRIES" envDefault:"1"`

	Port        int    `env:"PORT" envDefault:"3000"`
	CorsOrigin  string `env:"CORS_ORIGIN" envDefault:"*"`
	DatabaseURL string `env:"DATABASE_URL"`

	JWTRefreshSecret string `env:"JWT_REFRESH_SECRET" envDefault:"secret"`
	JWTRefreshExpiry int    `env:"JWT_REFRESH_EXPIRY" envDefault:"86400"`
	JWTAccessSecret  string `env:"JWT_ACCESS_SECRET" envDefault:"secret"`
	JWTAccessExpiry  int    `env:"JWT_ACCESS_EXPIRY" envDefault:"600"`
}
