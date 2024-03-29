// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGetCopyDownFieldNamesResponse conquest api get copy down field names response
//
// swagger:model conquest_apiGetCopyDownFieldNamesResponse
type ConquestAPIGetCopyDownFieldNamesResponse struct {

	// field definitions
	FieldDefinitions []*ConquestAPICopyDownFieldDefinition `json:"FieldDefinitions"`

	// has children
	HasChildren bool `json:"hasChildren,omitempty"`
}

// Validate validates this conquest api get copy down field names response
func (m *ConquestAPIGetCopyDownFieldNamesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetCopyDownFieldNamesResponse) validateFieldDefinitions(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldDefinitions) { // not required
		return nil
	}

	for i := 0; i < len(m.FieldDefinitions); i++ {
		if swag.IsZero(m.FieldDefinitions[i]) { // not required
			continue
		}

		if m.FieldDefinitions[i] != nil {
			if err := m.FieldDefinitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FieldDefinitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FieldDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api get copy down field names response based on the context it is used
func (m *ConquestAPIGetCopyDownFieldNamesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFieldDefinitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetCopyDownFieldNamesResponse) contextValidateFieldDefinitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FieldDefinitions); i++ {

		if m.FieldDefinitions[i] != nil {
			if err := m.FieldDefinitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FieldDefinitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FieldDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetCopyDownFieldNamesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetCopyDownFieldNamesResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetCopyDownFieldNamesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
