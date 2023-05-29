package service

import "log"

// Middleware describes a service middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency and returns a service Middleware.
func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger *log.Logger
	next   Service
}
