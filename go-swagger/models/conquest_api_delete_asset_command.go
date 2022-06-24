// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIDeleteAssetCommand conquest api delete asset command
//
// swagger:model conquest_apiDeleteAssetCommand
type ConquestAPIDeleteAssetCommand struct {

	// asset ID
	AssetID int32 `json:"AssetID,omitempty"`
}

// Validate validates this conquest api delete asset command
func (m *ConquestAPIDeleteAssetCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api delete asset command based on context it is used
func (m *ConquestAPIDeleteAssetCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDeleteAssetCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDeleteAssetCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDeleteAssetCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
