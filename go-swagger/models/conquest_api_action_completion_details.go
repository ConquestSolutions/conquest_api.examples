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

// ConquestAPIActionCompletionDetails conquest api action completion details
//
// swagger:model conquest_apiActionCompletionDetails
type ConquestAPIActionCompletionDetails struct {

	// action
	Action *ConquestAPIObjectHeader `json:"Action,omitempty"`

	// action type
	ActionType *ConquestAPIActionType `json:"ActionType,omitempty"`

	// actual cost
	ActualCost ConquestAPIDecimal `json:"ActualCost,omitempty"`

	// asset
	Asset *ConquestAPIObjectHeader `json:"Asset,omitempty"`

	// If there is a database value that the user has put in manually then the CostType shows as Other and the OtherCost shows the cost stored in database.
	// The UI should default to OtherCost if it is set.
	// If there is no database value, the UI should default to the calculated ActualCost.
	CostType *ConquestAPICostType `json:"CostType,omitempty"`

	// estimated cost
	EstimatedCost ConquestAPIDecimal `json:"EstimatedCost,omitempty"`

	// is repeatable
	IsRepeatable bool `json:"IsRepeatable,omitempty"`

	// other cost
	OtherCost ConquestAPIDecimal `json:"OtherCost,omitempty"`
}

// Validate validates this conquest api action completion details
func (m *ConquestAPIActionCompletionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActualCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	if m.ActionType != nil {
		if err := m.ActionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionType")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateActualCost(formats strfmt.Registry) error {
	if swag.IsZero(m.ActualCost) { // not required
		return nil
	}

	if err := m.ActualCost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ActualCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ActualCost")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateAsset(formats strfmt.Registry) error {
	if swag.IsZero(m.Asset) { // not required
		return nil
	}

	if m.Asset != nil {
		if err := m.Asset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Asset")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateCostType(formats strfmt.Registry) error {
	if swag.IsZero(m.CostType) { // not required
		return nil
	}

	if m.CostType != nil {
		if err := m.CostType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostType")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateEstimatedCost(formats strfmt.Registry) error {
	if swag.IsZero(m.EstimatedCost) { // not required
		return nil
	}

	if err := m.EstimatedCost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EstimatedCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EstimatedCost")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateOtherCost(formats strfmt.Registry) error {
	if swag.IsZero(m.OtherCost) { // not required
		return nil
	}

	if err := m.OtherCost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OtherCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OtherCost")
		}
		return err
	}

	return nil
}

// ContextValidate validate this conquest api action completion details based on the context it is used
func (m *ConquestAPIActionCompletionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActualCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAsset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCostType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstimatedCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOtherCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateActionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionType != nil {
		if err := m.ActionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionType")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateActualCost(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActualCost.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ActualCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ActualCost")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateAsset(ctx context.Context, formats strfmt.Registry) error {

	if m.Asset != nil {
		if err := m.Asset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Asset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Asset")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateCostType(ctx context.Context, formats strfmt.Registry) error {

	if m.CostType != nil {
		if err := m.CostType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CostType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CostType")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateEstimatedCost(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EstimatedCost.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EstimatedCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EstimatedCost")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) contextValidateOtherCost(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OtherCost.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OtherCost")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OtherCost")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIActionCompletionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIActionCompletionDetails) UnmarshalBinary(b []byte) error {
	var res ConquestAPIActionCompletionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
