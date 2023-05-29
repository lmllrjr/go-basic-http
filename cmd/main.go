package main

import (
	"log"
	"os"

	"go-basic-http/internal/clicmd"
	"go-basic-http/internal/utils"
)

func main() {
	errors := make(chan error)

	// start the web worker at port 8080
	go func() {
		errors <- clicmd.Webworker()
	}()

	// log any errors that occur in ListenAndServe
	for err := range errors {
		log.Println(utils.Map2JSON(map[string]interface{}{
			"transport": "HTTP",
			"during":    "Listen",
			"err":       err,
		}))
		os.Exit(1)
	}
}
