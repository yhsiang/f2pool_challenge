// Code generated by go-swagger; DO NOT EDIT.

package tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ValidateIPHandlerFunc turns a function with the right signature into a validate ip handler
type ValidateIPHandlerFunc func(ValidateIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ValidateIPHandlerFunc) Handle(params ValidateIPParams) middleware.Responder {
	return fn(params)
}

// ValidateIPHandler interface for that can handle valid validate ip params
type ValidateIPHandler interface {
	Handle(ValidateIPParams) middleware.Responder
}

// NewValidateIP creates a new http.Handler for the validate ip operation
func NewValidateIP(ctx *middleware.Context, handler ValidateIPHandler) *ValidateIP {
	return &ValidateIP{Context: ctx, Handler: handler}
}

/*
	ValidateIP swagger:route POST /v1/tools/validate tools validateIp

# Simple IP validation

Simple IP valication
*/
type ValidateIP struct {
	Context *middleware.Context
	Handler ValidateIPHandler
}

func (o *ValidateIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewValidateIPParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}