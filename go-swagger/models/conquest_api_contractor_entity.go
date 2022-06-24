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

// ConquestAPIContractorEntity conquest api contractor entity
//
// swagger:model conquest_apiContractorEntity
type ConquestAPIContractorEntity struct {

	// contractor ID
	ContractorID int32 `json:"ContractorID,omitempty"`

	// document location
	DocumentLocation string `json:"DocumentLocation,omitempty"`

	// record
	Record *ConquestAPIContractorRecord `json:"Record,omitempty"`
}

// Validate validates this conquest api contractor entity
func (m *ConquestAPIContractorEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIContractorEntity) validateRecord(formats strfmt.Registry) error {
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

// ContextValidate validate this conquest api contractor entity based on the context it is used
func (m *ConquestAPIContractorEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecord(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIContractorEntity) contextValidateRecord(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConquestAPIContractorEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIContractorEntity) UnmarshalBinary(b []byte) error {
	var res ConquestAPIContractorEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
