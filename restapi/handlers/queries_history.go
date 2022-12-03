package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
)

func (m *QueryHandler) QueriesHistory(params history.QueriesHistoryParams) middleware.Responder {
	queries, err := m.db.FetchQueries(&params)
	if err != nil {
		return history.NewQueriesHistoryBadRequest()
	}

	return history.NewQueriesHistoryOK().WithPayload(queries)
}
