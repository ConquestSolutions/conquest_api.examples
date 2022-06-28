// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGetAzureSearchForAddressAPIEndPointResult conquest api get azure search for address Api end point result
//
// swagger:model conquest_apiGetAzureSearchForAddressApiEndPointResult
type ConquestAPIGetAzureSearchForAddressAPIEndPointResult struct {

	// access key
	AccessKey string `json:"accessKey,omitempty"`

	// template Url
	TemplateURL string `json:"templateUrl,omitempty"`
}

// Validate validates this conquest api get azure search for address Api end point result
func (m *ConquestAPIGetAzureSearchForAddressAPIEndPointResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api get azure search for address Api end point result based on context it is used
func (m *ConquestAPIGetAzureSearchForAddressAPIEndPointResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetAzureSearchForAddressAPIEndPointResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetAzureSearchForAddressAPIEndPointResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetAzureSearchForAddressAPIEndPointResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}