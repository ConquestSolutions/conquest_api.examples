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

// ConquestAPIListHierarchyNodesQuery ListHierarchyNodesQuery gets a list of Hierarchy nodes (ObjectHeaders whose parent ObjectType is the same as the object's ObjectType
//
// WARNING: IMPLEMENTATION INCOMPLETE (API 2.0.5)
//   Some endpoints may ignore the parameters below and return more data. DO NOT ASSUME IT WILL STAY THIS WAY.
//   Your code should always set the parameters below and assume the expected, not the actual data they currently return.
//
// swagger:model conquest_apiListHierarchyNodesQuery
type ConquestAPIListHierarchyNodesQuery struct {

	// include ancestors
	IncludeAncestors bool `json:"IncludeAncestors,omitempty"`

	// include children
	IncludeChildren bool `json:"IncludeChildren,omitempty"`

	// include descendents
	IncludeDescendents int32 `json:"IncludeDescendents,omitempty"`

	// include siblings
	IncludeSiblings bool `json:"IncludeSiblings,omitempty"`

	// IncludeSubItems will include any items that are not the children but are defined for a particular object
	// Such items will have the particular object as the parent but have a different ObjectType to the parent
	IncludeSubItems bool `json:"IncludeSubItems,omitempty"`

	// object key
	ObjectKey *ConquestAPIObjectKey `json:"ObjectKey,omitempty"`
}

// Validate validates this conquest api list hierarchy nodes query
func (m *ConquestAPIListHierarchyNodesQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListHierarchyNodesQuery) validateObjectKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectKey) { // not required
		return nil
	}

	if m.ObjectKey != nil {
		if err := m.ObjectKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api list hierarchy nodes query based on the context it is used
func (m *ConquestAPIListHierarchyNodesQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListHierarchyNodesQuery) contextValidateObjectKey(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectKey != nil {
		if err := m.ObjectKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIListHierarchyNodesQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIListHierarchyNodesQuery) UnmarshalBinary(b []byte) error {
	var res ConquestAPIListHierarchyNodesQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
