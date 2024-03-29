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

// ConquestAPICompleteActionResult conquest api complete action result
//
// swagger:model conquest_apiCompleteActionResult
type ConquestAPICompleteActionResult struct {

	// job key
	JobKey *ConquestAPIJobKey `json:"job_key,omitempty"`
}

// Validate validates this conquest api complete action result
func (m *ConquestAPICompleteActionResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICompleteActionResult) validateJobKey(formats strfmt.Registry) error {
	if swag.IsZero(m.JobKey) { // not required
		return nil
	}

	if m.JobKey != nil {
		if err := m.JobKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api complete action result based on the context it is used
func (m *ConquestAPICompleteActionResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICompleteActionResult) contextValidateJobKey(ctx context.Context, formats strfmt.Registry) error {

	if m.JobKey != nil {
		if err := m.JobKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICompleteActionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICompleteActionResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPICompleteActionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
