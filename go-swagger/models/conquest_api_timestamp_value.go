// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPITimestampValue conquest api timestamp value
//
// swagger:model conquest_apiTimestampValue
type ConquestAPITimestampValue struct {

	// The Timestamp value.
	// Format: date-time
	Value strfmt.DateTime `json:"value,omitempty"`
}

// Validate validates this conquest api timestamp value
func (m *ConquestAPITimestampValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITimestampValue) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.FormatOf("value", "body", "date-time", m.Value.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conquest api timestamp value based on context it is used
func (m *ConquestAPITimestampValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPITimestampValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPITimestampValue) UnmarshalBinary(b []byte) error {
	var res ConquestAPITimestampValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
