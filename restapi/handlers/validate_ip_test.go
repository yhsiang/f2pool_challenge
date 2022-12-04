package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/handlers"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

func TestValidateIP(t *testing.T) {
	response1 := handlers.ValidateIP(tools.ValidateIPParams{
		HTTPRequest: &http.Request{},
		Request: &models.HandlerValidateIPRequest{
			IP: "1.1.1.1",
		},
	})
	assert.Equal(t, true, *response1.(*tools.ValidateIPOK).Payload.Status)

	response2 := handlers.ValidateIP(tools.ValidateIPParams{
		HTTPRequest: &http.Request{},
		Request: &models.HandlerValidateIPRequest{
			IP: "256.1.1.1",
		},
	})
	assert.Equal(t, false, *response2.(*tools.ValidateIPOK).Payload.Status)
}
