package http

import (
	"net/http"

	"go-basic-http/internal/utils"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	name := utils.GetField(r, 0)

	w.Write([]byte("hello, " + name + "\n"))
}
