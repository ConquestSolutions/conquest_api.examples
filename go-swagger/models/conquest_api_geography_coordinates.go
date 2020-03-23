// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGeographyCoordinates conquest api geography coordinates
//
// swagger:model conquest_apiGeographyCoordinates
type ConquestAPIGeographyCoordinates struct {

	// Y
	Latitude float64 `json:"latitude,omitempty"`

	// X
	Longitude float64 `json:"longitude,omitempty"`

	// pin
	Pin bool `json:"pin,omitempty"`
}

// Validate validates this conquest api geography coordinates
func (m *ConquestAPIGeographyCoordinates) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGeographyCoordinates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGeographyCoordinates) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGeographyCoordinates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
