package provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/database"
	authRepo "github.com/swp391-group3/ai-interview-practice/api/internal/features/auth/repository"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/tracer"
)

func ProvideLogger(cfg *config.Config) (*logger.Logger, error) {
	return logger.New(logger.Config{
		Level:      cfg.Logger.Level,
		Format:     cfg.Logger.Format,
		Output:     cfg.Logger.Output,
		TimeFormat: cfg.Logger.TimeFormat,
	})
}

func ProvideTracer(cfg *config.Config) (*tracer.Tracer, error) {
	return tracer.New(tracer.Config{
		Enabled:     cfg.Tracer.Enabled,
		ServiceName: cfg.Tracer.ServiceName,
		Endpoint:    cfg.Tracer.Endpoint,
		Insecure:    cfg.Tracer.Insecure,
		SampleRate:  cfg.Tracer.SampleRate,
		Environment: cfg.App.Environment,
		Version:     cfg.App.Version,
	})
}

func ProvideDatabasePool(cfg *config.Config, log *logger.Logger) (*pgxpool.Pool, error) {
	pool, err := database.NewDatabasePool(cfg.Database.DSN())
	if err != nil {
		log.Error("Failed to connect to database", logger.Error(err))
		return nil, err
	}
	log.Info("Database connection pool established successfully")
	return pool, nil
}

func ProvideAuthQueries(pool *pgxpool.Pool) *authRepo.Queries {
	return authRepo.New(pool)
}
