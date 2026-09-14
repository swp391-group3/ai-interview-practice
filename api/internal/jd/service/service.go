package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
)

const extractionFailedMessage = "Job description extraction failed."

// Extractor performs one attempt. Implementations must not retry internally.
// normalizedJD is untrusted data, never a source of executable instructions.
type Extractor interface {
	Extract(ctx context.Context, normalizedJD string) (ExtractionCandidate, error)
}

// RetryableFailure lets adapters identify clearly transient failures without
// leaking provider-specific response types into the service.
type RetryableFailure interface {
	error
	Retryable() bool
}

type Service struct {
	extractor  Extractor
	maxRetries int
}

func New(extractor Extractor, maxRetries int) (*Service, error) {
	if extractor == nil {
		return nil, fmt.Errorf("JD extractor is required")
	}
	if maxRetries < 0 || maxRetries > 1 {
		return nil, fmt.Errorf("JD max retries must be 0 or 1")
	}
	return &Service{extractor: extractor, maxRetries: maxRetries}, nil
}

func (s *Service) Extract(ctx context.Context, raw string) (StructuredJD, error) {
	jd, err := NormalizeInput(raw)
	if err != nil {
		return StructuredJD{}, err
	}
	for attempt := 0; ; attempt++ {
		if err := ctx.Err(); err != nil {
			return StructuredJD{}, apperror.Wrap(apperror.CodeExtractionFailed, extractionFailedMessage, err)
		}
		candidate, err := s.extractor.Extract(ctx, jd)
		if ctx.Err() != nil {
			return StructuredJD{}, apperror.Wrap(apperror.CodeExtractionFailed, extractionFailedMessage, ctx.Err())
		}
		if err == nil {
			var result StructuredJD
			result, err = ValidateCandidate(candidate)
			if err == nil {
				return result, nil
			}
		}
		var appErr *apperror.AppError
		if !errors.As(err, &appErr) ||
			appErr == nil ||
			appErr.Code != apperror.CodeInvalidExtractionOutput {
			err = apperror.Wrap(apperror.CodeExtractionFailed, extractionFailedMessage, err)
		}
		if attempt >= s.maxRetries || !shouldRetry(err) {
			return StructuredJD{}, err
		}
		// One shared budget covers transient failures and malformed/invalid
		// output regeneration. Never replay the invalid output as instructions.
		timer := time.NewTimer(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return StructuredJD{}, apperror.Wrap(apperror.CodeExtractionFailed, extractionFailedMessage, ctx.Err())
		case <-timer.C:
		}
	}
}

func shouldRetry(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var appErr *apperror.AppError
	if errors.As(err, &appErr) &&
		appErr != nil &&
		appErr.Code == apperror.CodeInvalidExtractionOutput {
		return true
	}
	var transient RetryableFailure
	return errors.As(err, &transient) && transient.Retryable()
}
