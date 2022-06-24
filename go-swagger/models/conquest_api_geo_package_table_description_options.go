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

// ConquestAPIGeoPackageTableDescriptionOptions conquest api geo package table description options
//
// swagger:model conquest_apiGeoPackageTableDescriptionOptions
type ConquestAPIGeoPackageTableDescriptionOptions struct {

	// column options
	ColumnOptions []*ConquestAPIGeoPackageKeyColumnDescriptionOption `json:"ColumnOptions"`
}

// Validate validates this conquest api geo package table description options
func (m *ConquestAPIGeoPackageTableDescriptionOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumnOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGeoPackageTableDescriptionOptions) validateColumnOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ColumnOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.ColumnOptions); i++ {
		if swag.IsZero(m.ColumnOptions[i]) { // not required
			continue
		}

		if m.ColumnOptions[i] != nil {
			if err := m.ColumnOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ColumnOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ColumnOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api geo package table description options based on the context it is used
func (m *ConquestAPIGeoPackageTableDescriptionOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumnOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGeoPackageTableDescriptionOptions) contextValidateColumnOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ColumnOptions); i++ {

		if m.ColumnOptions[i] != nil {
			if err := m.ColumnOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ColumnOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ColumnOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGeoPackageTableDescriptionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGeoPackageTableDescriptionOptions) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGeoPackageTableDescriptionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
