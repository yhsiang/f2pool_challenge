package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/f2pool-challenge/restapi/handlers"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/root"
)

func TestRoot(t *testing.T) {

	response := handlers.Root(root.RootParams{
		HTTPRequest: &http.Request{},
	})

	assert.False(t, *response.(*root.RootOK).Payload.Kubernetes)
	assert.NotEmpty(t, response.(*root.RootOK).Payload.Date)
	assert.Equal(t, "DEVELOPMENT", response.(*root.RootOK).Payload.Version)
}
