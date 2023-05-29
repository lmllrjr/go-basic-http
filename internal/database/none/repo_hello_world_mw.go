package none

import (
	"context"
	"log"
	"time"

	"go-basic-http/internal/database"
	"go-basic-http/internal/utils"
)

// HelloWorldRepositoryMiddleware describes a HelloWorld repository middleware.
type HelloWorldRepositoryMiddleware func(database.HelloWorldRepository) database.HelloWorldRepository

// HelloWorldRepositoryLoggingMiddleware takes a logger as a dependency and returns a HelloWorldRepositoryMiddleware.
func HelloWorldRepositoryLoggingMiddleware(logger *log.Logger) HelloWorldRepositoryMiddleware {
	return func(next database.HelloWorldRepository) database.HelloWorldRepository {
		return helloWorldRepositoryLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type helloWorldRepositoryLoggingMiddleware struct {
	logger *log.Logger
	next   database.HelloWorldRepository
}

var _ database.HelloWorldRepository = &helloWorldRepositoryLoggingMiddleware{}

func (mw helloWorldRepositoryLoggingMiddleware) HelloWorld(ctx context.Context) string {
	defer func(begin time.Time) {
		mw.logger.Println(utils.Map2JSON(map[string]interface{}{
			"layer":  "database",
			"method": "HelloWorld",
			"took":   float64(time.Since(begin)) / 1e6,
		}))
	}(time.Now())
	return mw.next.HelloWorld(ctx)
}
