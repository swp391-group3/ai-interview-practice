package provider

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	jdprovider "github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/provider"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/repository"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/service"
	"github.com/swp391-group3/ai-interview-practice/api/internal/handler"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/ai"
)

func ProvideJDRepository(pool *pgxpool.Pool) *repository.Repository {
	return repository.NewRepository(pool)
}
func ProvideJDService(cfg *config.Config, repo *repository.Repository) (*service.Application, error) {
	model, err := ai.NewChatModel(context.Background(), cfg.LLM)
	if err != nil {
		return nil, err
	}
	extractor, err := jdprovider.New(model)
	if err != nil {
		return nil, err
	}
	extraction, err := service.New(extractor, cfg.LLM.MaxRetries)
	if err != nil {
		return nil, err
	}
	return service.NewApplication(extraction, repo)
}
func ProvideJDHandler(app *service.Application) *handler.JDHandler { return handler.NewJDHandler(app) }
