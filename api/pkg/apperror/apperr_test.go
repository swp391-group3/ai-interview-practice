package apperror_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
)

var _ error = (*apperror.AppError)(nil)

func TestAppError(t *testing.T) {
	cause := errors.New("private database failure")
	for _, tc := range []struct {
		name      string
		err       *apperror.AppError
		wantText  string
		wantCause error
	}{
		{"new", apperror.New(apperror.CodeValidation, "invalid input"), "invalid input", nil},
		{"wrapped cause", apperror.Wrap(apperror.CodeValidation, "invalid input", cause), cause.Error(), cause},
		{"nil cause", apperror.Wrap(apperror.CodeValidation, "invalid input", nil), "invalid input", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.Code != apperror.CodeValidation || tc.err.Message != "invalid input" {
				t.Fatalf("public fields lost: %+v", tc.err)
			}
			if tc.err.Error() != tc.wantText || errors.Unwrap(tc.err) != tc.wantCause || tc.err.Err != tc.wantCause {
				t.Fatalf("error text or cause lost: %#v", tc.err)
			}
			outer := fmt.Errorf("service: %w", tc.err)
			var appErr *apperror.AppError
			if !errors.As(outer, &appErr) || appErr != tc.err {
				t.Fatal("errors.As did not recover the application error")
			}
			if tc.wantCause != nil && !errors.Is(outer, tc.wantCause) {
				t.Fatal("errors.Is did not recover the cause")
			}
		})
	}
}
