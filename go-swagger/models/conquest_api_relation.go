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

// ConquestAPIRelation Relation refers to an attribute on an object
//
// swagger:model conquest_apiRelation
type ConquestAPIRelation struct {

	// object type
	ObjectType *ConquestAPIObjectType `json:"objectType,omitempty"`

	// path is the name of the field on an object, for example, given an ObjectType of Asset, it could be AssetDescription
	Path string `json:"path,omitempty"`
}

// Validate validates this conquest api relation
func (m *ConquestAPIRelation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIRelation) validateObjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	if m.ObjectType != nil {
		if err := m.ObjectType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api relation based on the context it is used
func (m *ConquestAPIRelation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIRelation) contextValidateObjectType(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectType != nil {
		if err := m.ObjectType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIRelation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIRelation) UnmarshalBinary(b []byte) error {
	var res ConquestAPIRelation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
