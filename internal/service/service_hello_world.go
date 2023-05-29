package service

import (
	"context"
	"time"

	"go-basic-http/internal/utils"
)

// HelloWorld returns the string `hello world`.
func (s *service) HelloWorld(ctx context.Context) string {
	return s.helloWorldRepo.HelloWorld(ctx)
}

func (mw loggingMiddleware) HelloWorld(ctx context.Context) string {
	defer func(begin time.Time) {
		mw.logger.Println(utils.Map2JSON(map[string]interface{}{
			"layer":  "service",
			"method": "HelloWorld",
			"took":   float64(time.Since(begin)) / 1e6,
		}))
	}(time.Now())
	return mw.next.HelloWorld(ctx)
}
