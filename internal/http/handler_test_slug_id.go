package http

import (
	"net/http"

	"go-basic-http/internal/utils"
)

func TestSlugId(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// hm := ctx.Value(utils.CtxKey{}).([]string)
	// fmt.Fprintln(w, hm)
	slug := utils.GetField(r, 0)
	id := utils.GetField(r, 1)
	w.Write([]byte("slug: " + slug + ", ID: " + id + "\n"))
}
