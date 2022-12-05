package handlers

import (
	"context"
	"net"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/database"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/history"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

type QueryHandler struct {
	db interface {
		SaveQuery(*models.ModelQuery) error
		FetchQueries(*history.QueriesHistoryParams) ([]*models.ModelQuery, error)
	}
}

func NewQueryHandler(db *database.DB) *QueryHandler {
	return &QueryHandler{db}
}

func (m *QueryHandler) LookupDomain(params tools.LookupDomainParams) middleware.Responder {
	query, err := Lookup(&params)
	if err != nil {
		return tools.NewLookupDomainBadRequest().WithPayload(&models.UtilsHTTPError{
			Message: err.Error(),
		})
	}

	err = m.db.SaveQuery(query)
	if err != nil {
		return tools.NewLookupDomainBadRequest().WithPayload(&models.UtilsHTTPError{
			Message: err.Error(),
		})
	}

	return tools.NewLookupDomainOK().WithPayload(query)
}

func Lookup(params *tools.LookupDomainParams) (*models.ModelQuery, error) {
	ctx := context.Background()
	addrs, err := net.DefaultResolver.LookupIP(ctx, "ip4", params.Domain) //LookupIP()
	if err != nil {
		return nil, err
	}

	return &models.ModelQuery{
		Addresses: models.NewModelAddresses(addrs),
		// TODO: consider req.Header.Get("X-Forwarded-For")
		ClientIP:  params.HTTPRequest.RemoteAddr,
		CreatedAt: time.Now().Unix(),
	}, nil
}
