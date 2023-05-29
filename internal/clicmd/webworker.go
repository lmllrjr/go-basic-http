package clicmd

import (
	"log"
	"net/http"
	"os"
	"time"

	httptransport "go-basic-http/internal/http"
	"go-basic-http/internal/utils"
)

// Webworker starts a webworker which serves REST API requests over HTTP.
func Webworker() error {
	logger := log.New(os.Stdout, "", 0)

	svc, err := makeService(logger)
	if err != nil {
		logger.Fatal(err)
	}

	router := httptransport.NewHandler(svc, logger)

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println(utils.Map2JSON(map[string]interface{}{
		"status": "Listening",
		"port":   "8080",
	}))
	return server.ListenAndServe()
}
