// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIConfigAzureSearchAddressAPI conquest api config azure search address Api
//
// swagger:model conquest_api_configAzureSearchAddressApi
type ConquestAPIConfigAzureSearchAddressAPI struct {

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// template url
	TemplateURL string `json:"template_url,omitempty"`
}

// Validate validates this conquest api config azure search address Api
func (m *ConquestAPIConfigAzureSearchAddressAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this conquest api config azure search address Api based on context it is used
func (m *ConquestAPIConfigAzureSearchAddressAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIConfigAzureSearchAddressAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIConfigAzureSearchAddressAPI) UnmarshalBinary(b []byte) error {
	var res ConquestAPIConfigAzureSearchAddressAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
