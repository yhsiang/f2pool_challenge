package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

type QueryHandler struct {
	db interface {
		SaveQuery(*tools.LookupDomainParams) (*models.ModelQuery, error)
		FetchQueries(*history.QueriesHistoryParams) ([]*models.ModelQuery, error)
	}
}

func NewQueryHandler(db *database.DB) *QueryHandler {
	return &QueryHandler{db}
}

func (m *QueryHandler) LookupDomain(params tools.LookupDomainParams) middleware.Responder {
	query, err := m.db.SaveQuery(&params)
	if err != nil {
		return tools.NewLookupDomainBadRequest()
	}

	return tools.NewLookupDomainOK().WithPayload(query)
}
