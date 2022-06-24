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

// ConquestAPIListDocumentLocationsResponse conquest api list document locations response
//
// swagger:model conquest_apiListDocumentLocationsResponse
type ConquestAPIListDocumentLocationsResponse struct {

	// document locations
	DocumentLocations []*ConquestAPIDocumentContainer `json:"DocumentLocations"`

	// The PreferredLocationPrefix is selected by default.
	// It does *not* mean production. Every configured location *is* production.
	// It "can" mean a sub folder!
	PreferredLocationPrefix string `json:"PreferredLocationPrefix,omitempty"`
}

// Validate validates this conquest api list document locations response
func (m *ConquestAPIListDocumentLocationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListDocumentLocationsResponse) validateDocumentLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.DocumentLocations) { // not required
		return nil
	}

	for i := 0; i < len(m.DocumentLocations); i++ {
		if swag.IsZero(m.DocumentLocations[i]) { // not required
			continue
		}

		if m.DocumentLocations[i] != nil {
			if err := m.DocumentLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DocumentLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DocumentLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api list document locations response based on the context it is used
func (m *ConquestAPIListDocumentLocationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListDocumentLocationsResponse) contextValidateDocumentLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DocumentLocations); i++ {

		if m.DocumentLocations[i] != nil {
			if err := m.DocumentLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DocumentLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DocumentLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIListDocumentLocationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIListDocumentLocationsResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPIListDocumentLocationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
