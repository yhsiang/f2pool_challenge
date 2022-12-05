package handlers

import (
	"net"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/tools"
)

func ValidateIP(params tools.ValidateIPParams) middleware.Responder {
	// use pointer to solve boolean omitempty
	status := new(bool)
	ip := net.ParseIP(params.Request.IP)
	*status = ip != nil

	return tools.NewValidateIPOK().WithPayload(&models.HandlerValidateIPResponse{
		Status: status,
	})
}
