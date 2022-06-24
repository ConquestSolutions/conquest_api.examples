// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIErrorResponse conquest api error response
//
// swagger:model conquest_apiErrorResponse
type ConquestAPIErrorResponse struct {

	// error
	Error string `json:"error,omitempty"`

	// error description
	ErrorDescription string `json:"error_description,omitempty"`

	// error uri
	ErrorURI string `json:"error_uri,omitempty"`
}

// Validate validates this conquest api error response
func (m *ConquestAPIErrorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api error response based on context it is used
func (m *ConquestAPIErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIErrorResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPIErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
