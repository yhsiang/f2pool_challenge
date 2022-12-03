// Code generated by go-swagger; DO NOT EDIT.

package tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewLookupDomainParams creates a new LookupDomainParams object
//
// There are no default values defined in the spec.
func NewLookupDomainParams() LookupDomainParams {

	return LookupDomainParams{}
}

// LookupDomainParams contains all the bound params for the lookup domain operation
// typically these are obtained from a http.Request
//
// swagger:parameters lookup_domain
type LookupDomainParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Domain name
	  Required: true
	  In: query
	*/
	Domain string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLookupDomainParams() beforehand.
func (o *LookupDomainParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDomain, qhkDomain, _ := qs.GetOK("domain")
	if err := o.bindDomain(qDomain, qhkDomain, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDomain binds and validates parameter Domain from query.
func (o *LookupDomainParams) bindDomain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("domain", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("domain", "query", raw); err != nil {
		return err
	}
	o.Domain = raw

	return nil
}