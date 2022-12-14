// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yhsiang/f2pool-challenge/models"
)

// QueriesHistoryOKCode is the HTTP code returned for type QueriesHistoryOK
const QueriesHistoryOKCode int = 200

/*
QueriesHistoryOK OK

swagger:response queriesHistoryOK
*/
type QueriesHistoryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ModelQuery `json:"body,omitempty"`
}

// NewQueriesHistoryOK creates QueriesHistoryOK with default headers values
func NewQueriesHistoryOK() *QueriesHistoryOK {

	return &QueriesHistoryOK{}
}

// WithPayload adds the payload to the queries history o k response
func (o *QueriesHistoryOK) WithPayload(payload []*models.ModelQuery) *QueriesHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the queries history o k response
func (o *QueriesHistoryOK) SetPayload(payload []*models.ModelQuery) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueriesHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ModelQuery, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// QueriesHistoryBadRequestCode is the HTTP code returned for type QueriesHistoryBadRequest
const QueriesHistoryBadRequestCode int = 400

/*
QueriesHistoryBadRequest Bad Request

swagger:response queriesHistoryBadRequest
*/
type QueriesHistoryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.UtilsHTTPError `json:"body,omitempty"`
}

// NewQueriesHistoryBadRequest creates QueriesHistoryBadRequest with default headers values
func NewQueriesHistoryBadRequest() *QueriesHistoryBadRequest {

	return &QueriesHistoryBadRequest{}
}

// WithPayload adds the payload to the queries history bad request response
func (o *QueriesHistoryBadRequest) WithPayload(payload *models.UtilsHTTPError) *QueriesHistoryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the queries history bad request response
func (o *QueriesHistoryBadRequest) SetPayload(payload *models.UtilsHTTPError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueriesHistoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
