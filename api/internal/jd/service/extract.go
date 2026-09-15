package service

import "context"

func (s *Service) Extract(ctx context.Context, raw string) (StructuredJD, error) {
	jd, err := NormalizeInput(raw)
	if err != nil {
		return StructuredJD{}, err
	}
	return s.extractWithRetry(ctx, jd)
}
