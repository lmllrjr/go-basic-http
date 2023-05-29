package http_test

import (
	"context"
	"io"
	"log"
	nethttp "net/http"
	"net/http/httptest"
	"os"
	"testing"

	basichttp "go-basic-http/internal/http"
	svcmock "go-basic-http/internal/service/mock"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_HelloWorld(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	testCases := map[string]struct {
		service *svcmock.Service

		modifyRequest func(req *nethttp.Request)

		expResponseCode             int
		expResponseContentType      string
		expResponse                 string
		expServiceHelloWorldInvoked bool
	}{
		"ok": {
			service: &svcmock.Service{
				HelloWorldFunc: func(ctx context.Context) string {
					return "hello world"
				},
			},
			expServiceHelloWorldInvoked: true,
			expResponseCode:             nethttp.StatusOK,
			expResponse:                 "hello world\n",
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
			handler := basichttp.NewHandler(
				tc.service,
				logger,
			)

			srv := httptest.NewServer(handler)
			defer srv.Close()

			cli := nethttp.DefaultClient
			req, err := nethttp.NewRequest(nethttp.MethodGet, srv.URL+"/", nil)
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

			if tc.service != nil {
				assert.Equal(t, tc.expServiceHelloWorldInvoked, tc.service.HelloWorldInvoked, "service method HelloWorld invoked")
			}
		})
	}
}
