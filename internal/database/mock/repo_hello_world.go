package mock

import (
	"context"

	"go-basic-http/internal/database"
)

// HelloWorldRepository represents a mock implementation of database.HelloWorldRepository.
type HelloWorldRepository struct {
	HelloWorldFunc    func(ctx context.Context) string
	HelloWorldInvoked bool
}

var _ database.HelloWorldRepository = &HelloWorldRepository{}

// HelloWorld invokes the mock implementation and marks the function as invoked.
func (r *HelloWorldRepository) HelloWorld(ctx context.Context) string {
	r.HelloWorldInvoked = true
	return r.HelloWorldFunc(ctx)
}
