// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPICreateActionForAssetCommand conquest api create action for asset command
//
// swagger:model conquest_apiCreateActionForAssetCommand
type ConquestAPICreateActionForAssetCommand struct {

	// action description
	ActionDescription string `json:"ActionDescription,omitempty"`

	// ParentID is 0 if this is a Facility
	AssetID int32 `json:"AssetID,omitempty"`

	// std action ID
	StdActionID int32 `json:"StdActionID,omitempty"`
}

// Validate validates this conquest api create action for asset command
func (m *ConquestAPICreateActionForAssetCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api create action for asset command based on context it is used
func (m *ConquestAPICreateActionForAssetCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICreateActionForAssetCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICreateActionForAssetCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPICreateActionForAssetCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
