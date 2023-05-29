package none_test

import (
	"context"
	"testing"

	"go-basic-http/internal/database"
	"go-basic-http/internal/database/none"

	"github.com/stretchr/testify/assert"
)

func Test_HelloWorldRepository_HelloWorld(t *testing.T) {
	r := buildHelloWorldRepository(t)
	ctx := context.Background()

	{
		s := r.HelloWorld(ctx)
		assert.Equal(t, "hello world", s)
	}
}

func buildHelloWorldRepository(t *testing.T) database.HelloWorldRepository {
	return none.NewHelloWorldRepository(dbLogger)
}
