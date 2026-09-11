package apperror

type Code string

const (
	CodeInvalidCredentials Code = "INVALID_CREDENTIALS"
	CodeAccountNotFound    Code = "USER_NOT_FOUND"
	CodeAccountLocked      Code = "USER_INACTIVE"
	CodeInvalidToken       Code = "INVALID_TOKEN"
	CodeValidation         Code = "VALIDATION_ERROR"
	CodeInternal           Code = "INTERNAL_ERROR"
)

type AppError struct {
	Code    Code
	Message string
	Err     error
}

func New(code Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(code Code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
