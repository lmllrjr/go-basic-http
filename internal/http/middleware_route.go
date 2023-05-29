package http

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"go-basic-http/internal/utils"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

// newRoutes defines new routes by given slice of routes.
func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

// NewRouter is a routing function that loops through var routes,
// calls the first route that matches both the path and the HTTP method.
//
// Examples:
//
//	func main() {
//	        router := middleware.NewRouter()
//
//	        http.ListenAndServe(":8080", router)
//	}
//
// A typical handler with path parameters looks like this:
//
//	Handles POST /api/widgets/([^/]+)/parts/([0-9]+)/update
//
//	func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
//	        slug := getField(r, 0)
//	        id, _ := strconv.Atoi(getField(r, 1))
//	        fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
//	}
//
// https://github.com/benhoyt/go-routing/blob/master/retable/route.go
func newRouter(routes []route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), utils.CtxKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.NotFound(w, r)
	})
}
