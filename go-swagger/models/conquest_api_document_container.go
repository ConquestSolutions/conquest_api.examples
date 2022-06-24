// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIDocumentContainer DocumentContainer is convertible to a HierarchyNode which does not have a ParentKey nor a FamilyCode
//
// swagger:model conquest_apiDocumentContainer
type ConquestAPIDocumentContainer struct {

	// Name is display name. It is the same as the HierarchyNode.Name
	Name string `json:"Name,omitempty"`

	// Prefix is a Uri, that always ends with a '/'. It is the same as the HierarchyNode.ObjectKey.StringValue
	Prefix string `json:"Prefix,omitempty"`
}

// Validate validates this conquest api document container
func (m *ConquestAPIDocumentContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api document container based on context it is used
func (m *ConquestAPIDocumentContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDocumentContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDocumentContainer) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDocumentContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
