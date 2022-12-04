package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/restapi/handlers"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
)

func TestQueriesHistory(t *testing.T) {
	db := database.NewDB(9)
	handler := handlers.NewQueryHandler(db)

	response := handler.QueriesHistory(history.QueriesHistoryParams{
		HTTPRequest: &http.Request{},
	})
	length := len(response.(*history.QueriesHistoryOK).Payload)
	assert.Equal(t, 0, length)
}
