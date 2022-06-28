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

// ConquestAPIActionSchedule conquest api action schedule
//
// swagger:model conquest_apiActionSchedule
type ConquestAPIActionSchedule struct {

	// action date
	// Format: date-time
	ActionDate strfmt.DateTime `json:"ActionDate,omitempty"`

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// cycles
	Cycles int32 `json:"Cycles,omitempty"`

	// duration
	Duration int32 `json:"Duration,omitempty"`

	// The OriginalActionDate is only readOnly for the client
	// Format: date-time
	OriginalActionDate strfmt.DateTime `json:"OriginalActionDate,omitempty"`

	// repeat cycle units
	RepeatCycleUnits *ConquestAPIRepeatCycleUnits `json:"RepeatCycleUnits,omitempty"`

	// schedule type
	ScheduleType *ConquestAPIActionScheduleType `json:"ScheduleType,omitempty"`
}

// Validate validates this conquest api action schedule
func (m *ConquestAPIActionSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalActionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepeatCycleUnits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionSchedule) validateActionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ActionDate", "body", "date-time", m.ActionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIActionSchedule) validateOriginalActionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalActionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("OriginalActionDate", "body", "date-time", m.OriginalActionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIActionSchedule) validateRepeatCycleUnits(formats strfmt.Registry) error {
	if swag.IsZero(m.RepeatCycleUnits) { // not required
		return nil
	}

	if m.RepeatCycleUnits != nil {
		if err := m.RepeatCycleUnits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RepeatCycleUnits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RepeatCycleUnits")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionSchedule) validateScheduleType(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduleType) { // not required
		return nil
	}

	if m.ScheduleType != nil {
		if err := m.ScheduleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ScheduleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ScheduleType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api action schedule based on the context it is used
func (m *ConquestAPIActionSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepeatCycleUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheduleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionSchedule) contextValidateRepeatCycleUnits(ctx context.Context, formats strfmt.Registry) error {

	if m.RepeatCycleUnits != nil {
		if err := m.RepeatCycleUnits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RepeatCycleUnits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RepeatCycleUnits")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionSchedule) contextValidateScheduleType(ctx context.Context, formats strfmt.Registry) error {

	if m.ScheduleType != nil {
		if err := m.ScheduleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ScheduleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ScheduleType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIActionSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIActionSchedule) UnmarshalBinary(b []byte) error {
	var res ConquestAPIActionSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}