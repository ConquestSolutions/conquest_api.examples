// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIActionAssignment conquest api action assignment
//
// swagger:model conquest_apiActionAssignment
type ConquestAPIActionAssignment struct {

	// act CPU
	ActCPU *ConquestAPIDecimalValue `json:"ActCPU,omitempty"`

	// act quantity
	ActQuantity float64 `json:"ActQuantity,omitempty"`

	// act rate
	ActRate *ConquestAPIDecimalValue `json:"ActRate,omitempty"`

	// after completion
	AfterCompletion bool `json:"AfterCompletion,omitempty"`

	// assignment ID
	AssignmentID int32 `json:"AssignmentID,omitempty"`

	// est CPU
	EstCPU *ConquestAPIDecimalValue `json:"EstCPU,omitempty"`

	// est quantity
	EstQuantity float64 `json:"EstQuantity,omitempty"`

	// est rate
	EstRate *ConquestAPIDecimalValue `json:"EstRate,omitempty"`

	// material units
	MaterialUnits int32 `json:"MaterialUnits,omitempty"`

	// resource description
	ResourceDescription string `json:"ResourceDescription,omitempty"`

	// resource ID
	ResourceID int32 `json:"ResourceID,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"StartDate,omitempty"`

	// units
	Units float64 `json:"Units,omitempty"`
}

// Validate validates this conquest api action assignment
func (m *ConquestAPIActionAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionAssignment) validateActCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.ActCPU) { // not required
		return nil
	}

	if m.ActCPU != nil {
		if err := m.ActCPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActCPU")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActCPU")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) validateActRate(formats strfmt.Registry) error {
	if swag.IsZero(m.ActRate) { // not required
		return nil
	}

	if m.ActRate != nil {
		if err := m.ActRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActRate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) validateEstCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.EstCPU) { // not required
		return nil
	}

	if m.EstCPU != nil {
		if err := m.EstCPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EstCPU")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EstCPU")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) validateEstRate(formats strfmt.Registry) error {
	if swag.IsZero(m.EstRate) { // not required
		return nil
	}

	if m.EstRate != nil {
		if err := m.EstRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EstRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EstRate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this conquest api action assignment based on the context it is used
func (m *ConquestAPIActionAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionAssignment) contextValidateActCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.ActCPU != nil {
		if err := m.ActCPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActCPU")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActCPU")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) contextValidateActRate(ctx context.Context, formats strfmt.Registry) error {

	if m.ActRate != nil {
		if err := m.ActRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActRate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) contextValidateEstCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.EstCPU != nil {
		if err := m.EstCPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EstCPU")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EstCPU")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionAssignment) contextValidateEstRate(ctx context.Context, formats strfmt.Registry) error {

	if m.EstRate != nil {
		if err := m.EstRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EstRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EstRate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIActionAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIActionAssignment) UnmarshalBinary(b []byte) error {
	var res ConquestAPIActionAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
