package service

import (
	"context"

	"github.com/swp391-group3/ai-interview-practice/api/internal/features/jd/domain"
)

func (s *Service) Extract(ctx context.Context, raw string) (domain.StructuredJD, error) {
	jd, err := NormalizeInput(raw)
	if err != nil {
		return domain.StructuredJD{}, err
	}
	return s.extractWithRetry(ctx, jd)
}
