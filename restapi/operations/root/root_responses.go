// Code generated by go-swagger; DO NOT EDIT.

package root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yhsiang/f2pool-challenge/models"
)

// RootOKCode is the HTTP code returned for type RootOK
const RootOKCode int = 200

/*
RootOK OK

swagger:response rootOK
*/
type RootOK struct {

	/*
	  In: Body
	*/
	Payload *models.HandlerRootResponse `json:"body,omitempty"`
}

// NewRootOK creates RootOK with default headers values
func NewRootOK() *RootOK {

	return &RootOK{}
}

// WithPayload adds the payload to the root o k response
func (o *RootOK) WithPayload(payload *models.HandlerRootResponse) *RootOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the root o k response
func (o *RootOK) SetPayload(payload *models.HandlerRootResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RootOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
