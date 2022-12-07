package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/restapi/handlers"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

func TestQueriesHistory(t *testing.T) {
	db := database.NewDB(9)
	handler := handlers.NewQueryHandler(db)

	response := handler.QueriesHistory(history.QueriesHistoryParams{
		HTTPRequest: &http.Request{},
	})
	length := len(response.(*history.QueriesHistoryOK).Payload)
	assert.Equal(t, 0, length)

	for i := 0; i < 30; i++ {
		query, err := handlers.Lookup(&tools.LookupDomainParams{
			HTTPRequest: &http.Request{
				RemoteAddr: "3.3.3.3",
			},
			Domain: "google.com",
		})
		assert.NoError(t, err)

		err = db.SaveQuery(query)
		assert.NoError(t, err)
	}
	response = handler.QueriesHistory(history.QueriesHistoryParams{
		HTTPRequest: &http.Request{},
	})
	length = len(response.(*history.QueriesHistoryOK).Payload)
	assert.Equal(t, 20, length)
	db.Flush()
}
