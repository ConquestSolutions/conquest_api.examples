// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIMoveAssetCommand conquest api move asset command
//
// swagger:model conquest_apiMoveAssetCommand
type ConquestAPIMoveAssetCommand struct {

	// asset ID
	AssetID int32 `json:"AssetID,omitempty"`

	// parent ID
	ParentID int32 `json:"ParentID,omitempty"`
}

// Validate validates this conquest api move asset command
func (m *ConquestAPIMoveAssetCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api move asset command based on context it is used
func (m *ConquestAPIMoveAssetCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIMoveAssetCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIMoveAssetCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIMoveAssetCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
