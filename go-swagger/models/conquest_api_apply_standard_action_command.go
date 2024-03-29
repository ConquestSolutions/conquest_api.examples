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

// ConquestAPIApplyStandardActionCommand conquest api apply standard action command
//
// swagger:model conquest_apiApplyStandardActionCommand
type ConquestAPIApplyStandardActionCommand struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// overwrite
	Overwrite bool `json:"Overwrite,omitempty"`

	// overwrite options
	OverwriteOptions *ConquestAPIApplyStandardActionOverwriteOptions `json:"OverwriteOptions,omitempty"`

	// std action ID
	StdActionID int32 `json:"StdActionID,omitempty"`
}

// Validate validates this conquest api apply standard action command
func (m *ConquestAPIApplyStandardActionCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverwriteOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIApplyStandardActionCommand) validateOverwriteOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.OverwriteOptions) { // not required
		return nil
	}

	if m.OverwriteOptions != nil {
		if err := m.OverwriteOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OverwriteOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OverwriteOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api apply standard action command based on the context it is used
func (m *ConquestAPIApplyStandardActionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOverwriteOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIApplyStandardActionCommand) contextValidateOverwriteOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.OverwriteOptions != nil {
		if err := m.OverwriteOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OverwriteOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OverwriteOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIApplyStandardActionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIApplyStandardActionCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIApplyStandardActionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
