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

// ConquestAPICriteriaItem Criteria is an expression in a 'where' clause that evaluates to true or false
//
// swagger:model conquest_apiCriteriaItem
type ConquestAPICriteriaItem struct {

	// Use 'Fields' instead
	FieldList *ConquestAPIFieldCriterionList `json:"FieldList,omitempty"`

	// fields
	Fields *ConquestAPIFieldCriterionList `json:"Fields,omitempty"`

	// group
	Group *ConquestAPICriteria `json:"Group,omitempty"`
}

// Validate validates this conquest api criteria item
func (m *ConquestAPICriteriaItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICriteriaItem) validateFieldList(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldList) { // not required
		return nil
	}

	if m.FieldList != nil {
		if err := m.FieldList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FieldList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FieldList")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPICriteriaItem) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if m.Fields != nil {
		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fields")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPICriteriaItem) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api criteria item based on the context it is used
func (m *ConquestAPICriteriaItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFieldList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICriteriaItem) contextValidateFieldList(ctx context.Context, formats strfmt.Registry) error {

	if m.FieldList != nil {
		if err := m.FieldList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FieldList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FieldList")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPICriteriaItem) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	if m.Fields != nil {
		if err := m.Fields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fields")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPICriteriaItem) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICriteriaItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICriteriaItem) UnmarshalBinary(b []byte) error {
	var res ConquestAPICriteriaItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
