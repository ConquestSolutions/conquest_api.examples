// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIReverseActionCompletionCommand conquest api reverse action completion command
//
// swagger:model conquest_apiReverseActionCompletionCommand
type ConquestAPIReverseActionCompletionCommand struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`
}

// Validate validates this conquest api reverse action completion command
func (m *ConquestAPIReverseActionCompletionCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api reverse action completion command based on context it is used
func (m *ConquestAPIReverseActionCompletionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIReverseActionCompletionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIReverseActionCompletionCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIReverseActionCompletionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
