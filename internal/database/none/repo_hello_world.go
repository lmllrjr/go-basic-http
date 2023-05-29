package none

import (
	"context"
	"log"

	"go-basic-http/internal/database"
)

// HelloWorldRepository persists hello messages.
type HelloWorldRepository struct{}

// NewHelloWorldRepository constructs a hello world repository with middleware wired in.
func NewHelloWorldRepository(logger *log.Logger) database.HelloWorldRepository {
	var repo database.HelloWorldRepository
	repo = &HelloWorldRepository{}
	repo = HelloWorldRepositoryLoggingMiddleware(logger)(repo)
	return repo
}

var _ database.HelloWorldRepository = &HelloWorldRepository{}

// HelloWorld returns the string `hello world`.
func (r *HelloWorldRepository) HelloWorld(ctx context.Context) string {
	return "hello world"
}
