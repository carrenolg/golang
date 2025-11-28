package errs

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return &AppError{Code: 404, Message: message}
}

func NewUnexpectedError(message string) error {
	return &AppError{Code: 500, Message: message}
}

func NewValidationError(message string) error {
	return &AppError{Code: 400, Message: message}
}
