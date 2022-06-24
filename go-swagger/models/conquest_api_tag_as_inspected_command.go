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

// ConquestAPITagAsInspectedCommand conquest api tag as inspected command
//
// swagger:model conquest_apiTagAsInspectedCommand
type ConquestAPITagAsInspectedCommand struct {

	// asset i ds
	AssetIDs []int32 `json:"AssetIDs"`

	// inspection
	Inspection *ConquestAPIInspectionRecord `json:"Inspection,omitempty"`
}

// Validate validates this conquest api tag as inspected command
func (m *ConquestAPITagAsInspectedCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInspection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITagAsInspectedCommand) validateInspection(formats strfmt.Registry) error {
	if swag.IsZero(m.Inspection) { // not required
		return nil
	}

	if m.Inspection != nil {
		if err := m.Inspection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Inspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Inspection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api tag as inspected command based on the context it is used
func (m *ConquestAPITagAsInspectedCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInspection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITagAsInspectedCommand) contextValidateInspection(ctx context.Context, formats strfmt.Registry) error {

	if m.Inspection != nil {
		if err := m.Inspection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Inspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Inspection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPITagAsInspectedCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPITagAsInspectedCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPITagAsInspectedCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
