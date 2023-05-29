package service_test

import (
	"context"
	"log"
	"testing"

	dbmock "go-basic-http/internal/database/mock"
	"go-basic-http/internal/service"

	"github.com/stretchr/testify/assert"
)

func Test_Service_HelloWorld(t *testing.T) {
	testCases := map[string]struct {
		helloWorldRepo       *dbmock.HelloWorldRepository
		expErr               error
		expHelloWorldInvoked bool
	}{
		"ok": {
			helloWorldRepo: &dbmock.HelloWorldRepository{
				HelloWorldFunc: func(ctx context.Context) string {
					return "hello world"
				},
			},
			expHelloWorldInvoked: true,
		},
	}
	for tname, tc := range testCases {
		t.Run(tname, func(t *testing.T) {
			s := service.New(&service.ServiceConfig{
				HelloWorldRepo: tc.helloWorldRepo,
				Logger:         log.Default(),
			})
			hws := s.HelloWorld(context.Background())

			assert.Equal(t, "hello world", hws)

			if tc.helloWorldRepo != nil {
				assert.Equal(t, tc.expHelloWorldInvoked, tc.helloWorldRepo.HelloWorldInvoked)
			}
		})
	}
}
