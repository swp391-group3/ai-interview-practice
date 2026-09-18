package config

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds the entire system configuration
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	CORS     CORSConfig     `mapstructure:"cors"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Tracer   TracerConfig   `mapstructure:"tracer"`
	LLM      LLMConfig      `mapstructure:"llm"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"` // development, staging, production
	Version     string `mapstructure:"version"`
	Debug       bool   `mapstructure:"debug"`
}

type ServerConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type DatabaseConfig struct {
	URL             string        `mapstructure:"url"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"name"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	DebugLevel      string        `mapstructure:"debug_level"`
}

// DSN returns the PostgreSQL connection string.
// If URL is set, it is returned directly.
// Otherwise, a postgres:// connection URL is constructed with all
// credentials and parameters properly escaped.
func (d *DatabaseConfig) DSN() string {
	if d.URL != "" {
		return d.URL
	}

	u := &url.URL{
		Scheme: "postgres",
	}

	host := d.Host
	if host == "" {
		host = "localhost"
	}
	port := d.Port
	if port <= 0 {
		port = 5432
	}
	u.Host = net.JoinHostPort(host, strconv.Itoa(port))

	if d.Name != "" {
		u.Path = "/" + strings.TrimPrefix(d.Name, "/")
	}

	if d.User != "" || d.Password != "" {
		u.User = url.UserPassword(d.User, d.Password)
	}

	q := make(url.Values)
	if d.SSLMode != "" {
		q.Set("sslmode", d.SSLMode)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}

	return u.String()
}

type JWTConfig struct {
	AccessSecret       string        `mapstructure:"access_secret"`
	RefreshSecret      string        `mapstructure:"refresh_secret"`
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"`
	ResetTokenExpiry   time.Duration `mapstructure:"reset_token_expiry"`
	VerifyTokenExpiry  time.Duration `mapstructure:"verify_token_expiry"`
	Issuer             string        `mapstructure:"issuer"`
}

type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	ExposeHeaders    []string `mapstructure:"expose_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

type LoggerConfig struct {
	Level      string `mapstructure:"level"`  // debug, info, warn, error
	Format     string `mapstructure:"format"` // json or console
	Output     string `mapstructure:"output"` // stdout or file path
	TimeFormat string `mapstructure:"time_format"`
}

type TracerConfig struct {
	Enabled     bool    `mapstructure:"enabled"`
	ServiceName string  `mapstructure:"service_name"`
	Endpoint    string  `mapstructure:"endpoint"`
	Insecure    bool    `mapstructure:"insecure"`
	SampleRate  float64 `mapstructure:"sample_rate"`
}

type LLMConfig struct {
	Provider   string        `mapstructure:"provider"`
	APIKey     string        `mapstructure:"api_key"`
	Model      string        `mapstructure:"model"`
	Timeout    time.Duration `mapstructure:"timeout"`
	MaxRetries int           `mapstructure:"max_retries"`
}

// Load loads configuration from file and overrides with environment variables
func Load(configPath string) (*Config, error) {
	v := viper.New()

	setDefaults(v)

	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./configs")
		v.AddConfigPath(".")
		v.AddConfigPath("/etc/ai-interview/")
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)

	// Explicitly bind environment variables
	for key, env := range map[string]string{
		"database.url":             "DATABASE_URL",
		"database.name":            "DB_DATABASE",
		"database.user":            "DB_USERNAME",
		"database.password":        "DB_PASSWORD",
		"database.port":            "DB_PORT",
		"jwt.access_secret":        "JWT_ACCESS_SECRET",
		"jwt.refresh_secret":       "JWT_REFRESH_SECRET",
		"jwt.access_token_expiry":  "JWT_ACCESS_EXPIRY",
		"jwt.refresh_token_expiry": "JWT_REFRESH_EXPIRY",
		"llm.provider":             "LLM_PROVIDER",
		"llm.api_key":              "LLM_API_KEY",
		"llm.model":                "LLM_MODEL",
		"llm.timeout":              "LLM_TIMEOUT",
		"llm.max_retries":          "LLM_MAX_RETRIES",
	} {
		if err := v.BindEnv(key, env); err != nil {
			return nil, fmt.Errorf("bind %s: %w", key, err)
		}
	}

	// The original JWT env contract uses integer seconds; also accept Go durations.
	for _, key := range []string{"jwt.access_token_expiry", "jwt.refresh_token_expiry"} {
		value := v.GetString(key)
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			v.Set(key, value+"s")
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	// App defaults
	v.SetDefault("app.name", "ai-interview-practice")
	v.SetDefault("app.environment", "development")
	v.SetDefault("app.version", "1.0.0")
	v.SetDefault("app.debug", true)

	// Server defaults
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 3000)
	v.SetDefault("server.read_timeout", "10s")
	v.SetDefault("server.write_timeout", "30s")
	v.SetDefault("server.shutdown_timeout", "5s")

	// Database defaults
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "postgres")
	v.SetDefault("database.name", "ai_interview_practice")
	v.SetDefault("database.ssl_mode", "disable")
	v.SetDefault("database.max_open_conns", 25)
	v.SetDefault("database.max_idle_conns", 10)
	v.SetDefault("database.conn_max_lifetime", "5m")

	// JWT defaults (no default secrets in base viper config)
	v.SetDefault("jwt.access_token_expiry", "600s")
	v.SetDefault("jwt.refresh_token_expiry", "86400s")
	v.SetDefault("jwt.reset_token_expiry", "1h")
	v.SetDefault("jwt.verify_token_expiry", "24h")
	v.SetDefault("jwt.issuer", "ai-interview-practice")

	// CORS defaults
	v.SetDefault("cors.allow_origins", []string{"http://localhost:3000"})
	v.SetDefault("cors.allow_methods", []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	v.SetDefault("cors.allow_headers", []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"})
	v.SetDefault("cors.expose_headers", []string{"X-Request-ID"})
	v.SetDefault("cors.allow_credentials", true)
	v.SetDefault("cors.max_age", 86400)

	// Logger defaults
	v.SetDefault("logger.level", "debug")
	v.SetDefault("logger.format", "console")
	v.SetDefault("logger.output", "stdout")
	v.SetDefault("logger.time_format", "2006-01-02T15:04:05.000Z07:00")

	// Tracer defaults
	v.SetDefault("tracer.enabled", false)
	v.SetDefault("tracer.service_name", "ai-interview-practice-api")
	v.SetDefault("tracer.endpoint", "localhost:4317")
	v.SetDefault("tracer.insecure", true)
	v.SetDefault("tracer.sample_rate", 1.0)

	// LLM defaults
	v.SetDefault("llm.provider", "gemini")
	v.SetDefault("llm.model", "gemini-3.1-flash-lite")
	v.SetDefault("llm.timeout", "30s")
	v.SetDefault("llm.max_retries", 1)
}

// Validate checks configuration integrity and security constraints.
func (c *Config) Validate() error {
	if c.CORS.AllowCredentials {
		if len(c.CORS.AllowOrigins) == 0 {
			return errors.New("credentialed CORS requires explicit allowed origins")
		}
		for _, origin := range c.CORS.AllowOrigins {
			if origin == "*" || strings.TrimSpace(origin) == "" {
				return errors.New("credentialed CORS requires explicit allowed origins, without wildcard")
			}
		}
	}
	if c.JWT.AccessTokenExpiry <= 0 || c.JWT.RefreshTokenExpiry <= 0 {
		return errors.New("JWT expiry must be positive")
	}
	if c.IsProduction() {
		if strings.TrimSpace(c.JWT.AccessSecret) == "" {
			return errors.New("jwt.access_secret is required in production")
		}
		if strings.TrimSpace(c.JWT.RefreshSecret) == "" {
			return errors.New("jwt.refresh_secret is required in production")
		}
		if c.JWT.AccessSecret == c.JWT.RefreshSecret {
			return errors.New("jwt.access_secret and jwt.refresh_secret must be distinct")
		}

		knownInsecure := map[string]bool{
			"access-secret":      true,
			"refresh-secret":     true,
			"secret":             true,
			"dev-access-secret":  true,
			"dev-refresh-secret": true,
			"change-me":          true,
		}
		if knownInsecure[c.JWT.AccessSecret] || knownInsecure[c.JWT.RefreshSecret] {
			return errors.New("jwt secrets must not use default or placeholder values in production")
		}
	} else {
		// Keep isolated defaults for development and testing environments only
		if c.JWT.AccessSecret == "" {
			c.JWT.AccessSecret = "dev-access-secret"
		}
		if c.JWT.RefreshSecret == "" {
			c.JWT.RefreshSecret = "dev-refresh-secret"
		}
		if c.JWT.AccessSecret == c.JWT.RefreshSecret {
			return errors.New("jwt.access_secret and jwt.refresh_secret must be distinct")
		}
	}

	if c.Tracer.Enabled {
		if c.Tracer.SampleRate < 0.0 || c.Tracer.SampleRate > 1.0 {
			return fmt.Errorf("tracer.sample_rate must be between 0.0 and 1.0, got %f", c.Tracer.SampleRate)
		}
	}

	return nil
}

func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}
