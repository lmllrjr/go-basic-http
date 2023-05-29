package http

import (
	"io"
	"net/http"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleware_route(t *testing.T) {
	testCases := map[string]struct {
		modifyRequest func(req *nethttp.Request)
		method        string
		path          string
		routes        bool

		expStatusCode      int
		expWWWAuthenticate string
		expResponse        string
	}{
		"ok": {
			method: http.MethodGet,
			modifyRequest: func(req *nethttp.Request) {
				req.SetBasicAuth("007", "123")
			},
			path:          "/foo",
			routes:        true,
			expStatusCode: http.StatusTeapot,
		},
		"method not allowed": {
			method: http.MethodPut,
			modifyRequest: func(req *nethttp.Request) {
				req.SetBasicAuth("007", "123")
			},
			path:          "/foo",
			routes:        true,
			expStatusCode: http.StatusMethodNotAllowed,
			expResponse:   "405 method not allowed\n",
		},
		"empty routes -> page not found": {
			method: http.MethodPut,
			modifyRequest: func(req *nethttp.Request) {
				req.SetBasicAuth("007", "123")
			},
			path:          "/foo",
			expStatusCode: http.StatusNotFound,
			expResponse:   "404 page not found\n",
		},
		"unauthorized": {
			method: http.MethodGet,
			modifyRequest: func(req *nethttp.Request) {
				req.SetBasicAuth("007", "006")
			},
			path:          "/bar",
			routes:        true,
			expStatusCode: http.StatusUnauthorized,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// create router
			teapot := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			})
			unauthorized := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
			})

			var routes []route
			if tc.routes {
				routes = []route{
					newRoute("GET", "/foo", teapot),
					newRoute("GET", "/bar", unauthorized),
				}
			}
			router := basicAuth(newRouter(routes))

			// Start an HTTP server.
			srv := httptest.NewServer(router)
			defer srv.Close()

			// Build an HTTP request that is passed to the handler, with our middleware in the call chain.
			req, err := http.NewRequest(tc.method, srv.URL+tc.path, nil)
			assert.NoError(t, err)
			if tc.modifyRequest != nil {
				tc.modifyRequest(req)
			}

			// Perform the request using a real HTTP client.
			cli := http.DefaultClient
			resp, err := cli.Do(req)
			if assert.NoError(t, err) {
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)

				// Verify the request.
				assert.Equal(t, tc.expStatusCode, resp.StatusCode, "http status code")
				assert.Equal(t, tc.expResponse, string(body))
				assert.Equal(t, tc.expWWWAuthenticate, resp.Header.Get("www-authenticate"))
			}
		})
	}
}
