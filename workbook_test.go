package main_test

import (
	"codingtestgolang/routes"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	testflight.WithServer(routes.Handler(), func(r *testflight.Requester) {
		response := r.Get("/")
		assert.Equal(t, 200, response.StatusCode)
	})
}
