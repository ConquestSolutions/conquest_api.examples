// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIDescribeEnumerationRequest conquest api describe enumeration request
//
// swagger:model conquest_apiDescribeEnumerationRequest
type ConquestAPIDescribeEnumerationRequest struct {

	// EnumerationTypeNames is a list of the "original names" or "tags" of the values in the EnumerationType
	EnumerationTypeNames []string `json:"EnumerationTypeNames"`
}

// Validate validates this conquest api describe enumeration request
func (m *ConquestAPIDescribeEnumerationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api describe enumeration request based on context it is used
func (m *ConquestAPIDescribeEnumerationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDescribeEnumerationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDescribeEnumerationRequest) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDescribeEnumerationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
