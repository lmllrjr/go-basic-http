package http_test

import (
	"io"
	"log"
	nethttp "net/http"
	"net/http/httptest"
	"os"
	"testing"

	httptransport "go-basic-http/internal/http"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_TestSlugID(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	testCases := map[string]struct {
		modifyRequest func(req *nethttp.Request)

		expResponseCode             int
		expResponseContentType      string
		expResponse                 string
		expServiceHelloWorldInvoked bool
	}{
		"ok": {
			expServiceHelloWorldInvoked: true,
			expResponseCode:             nethttp.StatusOK,
			expResponse:                 "slug: slug, ID: 777\n",
		},
		"auth failed": {
			modifyRequest: func(req *nethttp.Request) {
				req.SetBasicAuth("rick sanchez", "wubbalubbadubdub")
			},
			expResponseCode: nethttp.StatusUnauthorized,
			expResponse:     "🚫 YOU SHALL NOT PASS\n",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			handler := httptransport.NewHandler(
				nil,
				logger,
			)

			srv := httptest.NewServer(handler)
			defer srv.Close()

			cli := nethttp.DefaultClient
			req, err := nethttp.NewRequest(nethttp.MethodGet, srv.URL+"/slug/777", nil)
			assert.NoError(t, err)
			req.SetBasicAuth("007", "123")
			req.Header.Set("X-Forwarded-For", "192.168.12.13")
			if tc.modifyRequest != nil {
				tc.modifyRequest(req)
			}

			resp, err := cli.Do(req)
			assert.NoError(t, err)
			if assert.NoError(t, err) {
				defer resp.Body.Close()
				assert.Equal(t, tc.expResponseCode, resp.StatusCode, "http status code")
				body, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				if assert.NoError(t, err) {
					if tc.expResponse != "" {
						assert.Equal(t, tc.expResponse, string(body))
					}
				}
			}
		})
	}
}
