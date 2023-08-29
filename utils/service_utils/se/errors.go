package se

func ClientKnownError(message string, options ...Option) *ServiceError {
	e := BuildError(495, "ClientKnownError", options...)
	e.Message = message
	return e
}

func ServerKnownError(cause string, options ...Option) *ServiceError {
	e := BuildError(595, "ServerKnownError", options...)
	e.Cause = cause
	e.Message = "server side errors that should not occur, please contact the developer"
	return e
}

func ServerUnknownError(cause string, options ...Option) *ServiceError {
	e := BuildError(596, "ServerUnknownError", options...)
	e.Cause = cause
	e.Message = "Unknown server error, please contact the developer"
	return e
}

// service errors below

func ValidationError(message string, options ...Option) *ServiceError {
	e := BuildError(496, "ValidationError", options...)
	e.Message = message
	return e
}

func NotFoundError(message string, options ...Option) *ServiceError {
	e := BuildError(496, "NotFound", options...)
	e.Message = message
	return e
}

func ThrottledError(message string, options ...Option) *ServiceError {
	e := BuildError(496, "NotFound", options...)
	e.Message = message
	return e
}

func InvalidTokenError(message string, options ...Option) *ServiceError {
	e := BuildError(496, "InvalidTokenError", options...)
	e.Message = message
	return e
}

func ForbiddenError(cause string, options ...Option) *ServiceError {
	e := BuildError(496, "ForbiddenError", options...)
	e.Message = "forbidden"
	e.Cause = cause
	return e
}
