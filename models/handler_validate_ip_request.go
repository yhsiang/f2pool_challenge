// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HandlerValidateIPRequest handler validate IP request
//
// swagger:model handler.ValidateIPRequest
type HandlerValidateIPRequest struct {

	// ip
	IP string `json:"ip,omitempty"`
}

// Validate validates this handler validate IP request
func (m *HandlerValidateIPRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this handler validate IP request based on context it is used
func (m *HandlerValidateIPRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HandlerValidateIPRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlerValidateIPRequest) UnmarshalBinary(b []byte) error {
	var res HandlerValidateIPRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
