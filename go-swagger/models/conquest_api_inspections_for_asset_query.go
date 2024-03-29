// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIInspectionsForAssetQuery conquest api inspections for asset query
//
// swagger:model conquest_apiInspectionsForAssetQuery
type ConquestAPIInspectionsForAssetQuery struct {

	// asset ID
	AssetID int32 `json:"AssetID,omitempty"`
}

// Validate validates this conquest api inspections for asset query
func (m *ConquestAPIInspectionsForAssetQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api inspections for asset query based on context it is used
func (m *ConquestAPIInspectionsForAssetQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIInspectionsForAssetQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIInspectionsForAssetQuery) UnmarshalBinary(b []byte) error {
	var res ConquestAPIInspectionsForAssetQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
