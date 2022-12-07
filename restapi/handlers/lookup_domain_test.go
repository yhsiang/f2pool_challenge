package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/restapi/handlers"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

func TestLookup(t *testing.T) {
	query, err := handlers.Lookup(&tools.LookupDomainParams{
		HTTPRequest: &http.Request{
			RemoteAddr: "1.1.1.1",
		},
		Domain: "google.com",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, query.Addresses)
	assert.Equal(t, "1.1.1.1", query.ClientIP)
	assert.NotEmpty(t, query.CreatedAt)
}

func TestLookupDomain(t *testing.T) {
	db := database.NewDB(8)
	handler := handlers.NewQueryHandler(db)

	response := handler.LookupDomain(tools.LookupDomainParams{
		HTTPRequest: &http.Request{
			RemoteAddr: "2.2.2.2",
		},
		Domain: "google.com",
	})

	assert.NotEmpty(t, response.(*tools.LookupDomainOK).Payload.CreatedAt)
	db.Flush()
}
