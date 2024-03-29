// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIShapeStyle conquest api shape style
//
// swagger:model conquest_apiShapeStyle
type ConquestAPIShapeStyle struct {

	// fill color
	FillColor string `json:"FillColor,omitempty"`

	// fill opacity
	FillOpacity float64 `json:"FillOpacity,omitempty"`
}

// Validate validates this conquest api shape style
func (m *ConquestAPIShapeStyle) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api shape style based on context it is used
func (m *ConquestAPIShapeStyle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIShapeStyle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIShapeStyle) UnmarshalBinary(b []byte) error {
	var res ConquestAPIShapeStyle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
