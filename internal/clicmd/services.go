package clicmd

import (
	"log"

	"go-basic-http/internal/database/none"
	"go-basic-http/internal/service"
)

func makeService(logger *log.Logger) (service.Service, error) {
	// Build database repositories.
	helloWorldRepo := none.NewHelloWorldRepository(logger)

	svc := service.New(&service.ServiceConfig{
		HelloWorldRepo: helloWorldRepo,
		Logger:         logger,
	})
	return svc, nil
}
