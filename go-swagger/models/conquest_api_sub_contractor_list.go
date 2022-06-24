// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPISubContractorList conquest api sub contractor list
//
// swagger:model conquest_apiSubContractorList
type ConquestAPISubContractorList struct {

	// items
	Items []int32 `json:"Items"`
}

// Validate validates this conquest api sub contractor list
func (m *ConquestAPISubContractorList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api sub contractor list based on context it is used
func (m *ConquestAPISubContractorList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPISubContractorList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPISubContractorList) UnmarshalBinary(b []byte) error {
	var res ConquestAPISubContractorList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
