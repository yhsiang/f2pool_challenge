package database_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
)

func TestSaveQuery(t *testing.T) {
	db := database.NewDB(7)
	var addrs []*models.ModelAddress
	err := db.SaveQuery(&models.ModelQuery{
		Addresses: addrs,
		ClientIP:  "3.3.3.3",
		CreatedAt: time.Now().Unix(),
		Domain:    "google.com",
	})
	assert.NoError(t, err)
}

func TestFecthQueries(t *testing.T) {
	db := database.NewDB(6)
	queries, err := db.FetchQueries(&history.QueriesHistoryParams{
		HTTPRequest: &http.Request{},
	})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(queries))
}
