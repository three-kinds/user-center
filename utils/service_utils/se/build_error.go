package se

type Option func(*ServiceError)

func Message(message string) Option {
	return func(e *ServiceError) {
		e.Message = message
	}
}
func Cause(cause string) Option {
	return func(e *ServiceError) {
		e.Cause = cause
	}
}

func Detail(detail map[string]string) Option {
	return func(e *ServiceError) {
		e.Detail = detail
	}
}

func BuildError(code int, status string, options ...Option) *ServiceError {
	e := &ServiceError{Code: code, Status: status}
	for _, option := range options {
		option(e)
	}
	return e
}
