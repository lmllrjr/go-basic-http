package database

import "context"

// HelloWorldRepository persists hello message.
type HelloWorldRepository interface {
	// HelloWorld returns the string `hello world`.
	HelloWorld(ctx context.Context) string
}
