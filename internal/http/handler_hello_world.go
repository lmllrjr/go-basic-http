package http

import (
	"net/http"

	"go-basic-http/internal/service"
)

func HelloWorld(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hw := s.HelloWorld(r.Context())
		w.Write([]byte(hw + "\n"))
	}
}
