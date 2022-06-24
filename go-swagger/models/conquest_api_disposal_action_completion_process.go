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

// ConquestAPIDisposalActionCompletionProcess conquest api disposal action completion process
//
// swagger:model conquest_apiDisposalActionCompletionProcess
type ConquestAPIDisposalActionCompletionProcess struct {

	// measurement after disposal
	MeasurementAfterDisposal *ConquestAPIDecimalValue `json:"MeasurementAfterDisposal,omitempty"`
}

// Validate validates this conquest api disposal action completion process
func (m *ConquestAPIDisposalActionCompletionProcess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeasurementAfterDisposal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIDisposalActionCompletionProcess) validateMeasurementAfterDisposal(formats strfmt.Registry) error {
	if swag.IsZero(m.MeasurementAfterDisposal) { // not required
		return nil
	}

	if m.MeasurementAfterDisposal != nil {
		if err := m.MeasurementAfterDisposal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MeasurementAfterDisposal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MeasurementAfterDisposal")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api disposal action completion process based on the context it is used
func (m *ConquestAPIDisposalActionCompletionProcess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeasurementAfterDisposal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIDisposalActionCompletionProcess) contextValidateMeasurementAfterDisposal(ctx context.Context, formats strfmt.Registry) error {

	if m.MeasurementAfterDisposal != nil {
		if err := m.MeasurementAfterDisposal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MeasurementAfterDisposal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MeasurementAfterDisposal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDisposalActionCompletionProcess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDisposalActionCompletionProcess) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDisposalActionCompletionProcess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
