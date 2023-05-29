package service

import (
	"context"

	"go-basic-http/internal/service"
)

// Service represents a mock implementation of service.Service.
type Service struct {
	// HelloWorld:
	HelloWorldFunc    func(ctx context.Context) string
	HelloWorldInvoked bool
}

var _ service.Service = &Service{}

// HelloWorld:

// HelloWorld invokes the mock implementation and marks the function as invoked.
func (s *Service) HelloWorld(ctx context.Context) string {
	s.HelloWorldInvoked = true
	return s.HelloWorldFunc(ctx)
}
