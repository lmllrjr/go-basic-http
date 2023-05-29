package http

import "net/http"

// basicAuth wraps a handler with basic authentication middleware.
func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr, pw, ok := r.BasicAuth()
		if !ok {
			setUnauthorizedHeader(w)
			return
		}

		if usr != "007" && pw != "123" {
			setUnauthorizedHeader(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// setUnauthorizedHeader sets a response writers header to http.StatusUnauthorized/401.
func setUnauthorizedHeader(w http.ResponseWriter) {
	w.Header().Set("www-authenticate", `Basic Realm="restricted", charset="UTF-8"`)
	http.Error(w, "🚫 YOU SHALL NOT PASS", http.StatusUnauthorized)
}
