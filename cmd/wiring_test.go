package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-basic-http/internal/database/none"
	httptransport "go-basic-http/internal/http"
	"go-basic-http/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestWiring(t *testing.T) {
	logger := log.Default()
	dbLogger := logger

	helloWorldRepo := none.NewHelloWorldRepository(dbLogger)

	// Build the service onion layers.
	svc := service.New(&service.ServiceConfig{
		HelloWorldRepo: helloWorldRepo,
		Logger:         logger,
	})
	httpHandler := httptransport.NewHandler(
		svc,
		logger,
	)

	// Start a test HTTP server.
	srv := httptest.NewServer(httpHandler)
	defer srv.Close()

	t.Run("REST HelloWorld", func(t *testing.T) {
		client := http.DefaultClient

		t.Run("ok", func(t *testing.T) {
			req, err := http.NewRequest("GET", srv.URL+"/", nil)
			assert.NoError(t, err)
			req.SetBasicAuth("007", "123")

			resp, err := client.Do(req)
			if assert.NoError(t, err) {
				defer resp.Body.Close()
				assert.Equal(t, 200, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				if assert.NoError(t, err) {
					assert.Equal(t, "hello world\n", string(body))
				}
			}
		})
	})

	t.Run("REST Greet", func(t *testing.T) {
		client := http.DefaultClient

		t.Run("ok", func(t *testing.T) {
			req, err := http.NewRequest("GET", srv.URL+"/greet/luke skywalker", nil)
			assert.NoError(t, err)
			req.SetBasicAuth("007", "123")

			resp, err := client.Do(req)
			if assert.NoError(t, err) {
				defer resp.Body.Close()
				assert.Equal(t, 200, resp.StatusCode)

				body, err := io.ReadAll(resp.Body)
				if assert.NoError(t, err) {
					assert.Equal(t, "hello, luke skywalker\n", string(body))
				}
			}
		})
	})
}
