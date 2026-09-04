package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	CorsOrigin  string `env:"CORS_ORIGIN" envDefault:"*"`
	DatabaseURL string `env:"DATABASE_URL"`

	JWTRefreshSecret string `env:"JWT_REFRESH_SECRET" envDefault:"secret"`
	JWTRefreshExpiry int    `env:"JWT_REFRESH_EXPIRY" envDefault:"86400"`
	JWTAccessSecret  string `env:"JWT_ACCESS_SECRET" envDefault:"secret"`
	JWTAccessExpiry  int    `env:"JWT_ACCESS_EXPIRY" envDefault:"600"`
}
