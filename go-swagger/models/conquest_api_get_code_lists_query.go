// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGetCodeListsQuery conquest api get code lists query
//
// swagger:model conquest_apiGetCodeListsQuery
type ConquestAPIGetCodeListsQuery struct {

	// if CodeID is greater than zero, the CodeList for this Code will be loaded.
	CodeID []int32 `json:"CodeID"`

	// code list ID
	CodeListID []int32 `json:"CodeListID"`

	// code list name
	CodeListName []string `json:"CodeListName"`
}

// Validate validates this conquest api get code lists query
func (m *ConquestAPIGetCodeListsQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api get code lists query based on context it is used
func (m *ConquestAPIGetCodeListsQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetCodeListsQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetCodeListsQuery) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetCodeListsQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
