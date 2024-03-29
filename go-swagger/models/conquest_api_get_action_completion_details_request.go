// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGetActionCompletionDetailsRequest conquest api get action completion details request
//
// swagger:model conquest_apiGetActionCompletionDetailsRequest
type ConquestAPIGetActionCompletionDetailsRequest struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// x include completed sub actions
	XIncludeCompletedSubActions bool `json:"XIncludeCompletedSubActions,omitempty"`
}

// Validate validates this conquest api get action completion details request
func (m *ConquestAPIGetActionCompletionDetailsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api get action completion details request based on context it is used
func (m *ConquestAPIGetActionCompletionDetailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetActionCompletionDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetActionCompletionDetailsRequest) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetActionCompletionDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
