// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIAttributeSetListValue conquest api attribute set list value
//
// swagger:model conquest_apiAttributeSet_ListValue
type ConquestAPIAttributeSetListValue struct {

	// code list description
	CodeListDescription string `json:"CodeListDescription,omitempty"`

	// code list ID
	CodeListID int32 `json:"CodeListID,omitempty"`
}

// Validate validates this conquest api attribute set list value
func (m *ConquestAPIAttributeSetListValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api attribute set list value based on context it is used
func (m *ConquestAPIAttributeSetListValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAttributeSetListValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAttributeSetListValue) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAttributeSetListValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
