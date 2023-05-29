package service

import (
	"context"
	"log"

	"go-basic-http/internal/database"
)

// Service provides business logic.
type Service interface {
	HelloWorldService
}

type HelloWorldService interface {
	HelloWorld(ctx context.Context) string
}

type service struct {
	helloWorldRepo database.HelloWorldRepository
}

var _ Service = (*service)(nil)

// ServiceConfig contains the configuration params of the service.
type ServiceConfig struct {
	HelloWorldRepo database.HelloWorldRepository
	Logger         *log.Logger
}

// New returns a service with middleware wired in.
func New(config *ServiceConfig) Service {
	var svc Service
	svc = &service{
		helloWorldRepo: config.HelloWorldRepo,
	}
	svc = LoggingMiddleware(config.Logger)(svc)
	return svc
}
