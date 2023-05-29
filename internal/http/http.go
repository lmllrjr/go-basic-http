package http

import (
	"log"
	"net/http"

	"go-basic-http/internal/service"
	"go-basic-http/internal/utils"
)

// NewHandler returns an HTTP handler with middleware wired in.
func NewHandler(
	service service.Service,
	logger *log.Logger,
) http.Handler {
	routes := []route{
		newRoute("GET", `/`, HelloWorld(service)),
		newRoute("GET", `/([^/]+)/([0-9]+)`, TestSlugId),
		newRoute("GET", `/greet/([^/]+)`, Greet),
	}

	router := basicAuth(newRouter(routes))
	router = middlewareLogging(router, logger)

	return router
}

func middlewareLogging(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println(utils.Map2JSON(map[string]interface{}{
			"layer":           "transport",
			"transport":       "HTTP",
			"method":          r.Method,
			"url":             r.URL.Path,
			"user_agent":      r.UserAgent(),
			"x_forwarded_for": r.Header.Get("x-forwarded-for"),
			"ip":              r.RemoteAddr,
		}))

		next.ServeHTTP(w, r)
	})
}
