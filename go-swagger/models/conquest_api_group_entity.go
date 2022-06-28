// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGroupEntity conquest api group entity
//
// swagger:model conquest_apiGroupEntity
type ConquestAPIGroupEntity struct {

	// group name
	GroupName string `json:"GroupName,omitempty"`

	// record
	Record *ConquestAPIGroupRecord `json:"Record,omitempty"`
}

// Validate validates this conquest api group entity
func (m *ConquestAPIGroupEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGroupEntity) validateRecord(formats strfmt.Registry) error {
	if swag.IsZero(m.Record) { // not required
		return nil
	}

	if m.Record != nil {
		if err := m.Record.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Record")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Record")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api group entity based on the context it is used
func (m *ConquestAPIGroupEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecord(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGroupEntity) contextValidateRecord(ctx context.Context, formats strfmt.Registry) error {

	if m.Record != nil {
		if err := m.Record.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Record")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Record")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGroupEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGroupEntity) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGroupEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}