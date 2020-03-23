// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIGetUserPreferencesResult conquest api get user preferences result
//
// swagger:model conquest_apiGetUserPreferencesResult
type ConquestAPIGetUserPreferencesResult struct {

	// preferences
	Preferences []*ConquestAPIUserPreference `json:"Preferences"`
}

// Validate validates this conquest api get user preferences result
func (m *ConquestAPIGetUserPreferencesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetUserPreferencesResult) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	for i := 0; i < len(m.Preferences); i++ {
		if swag.IsZero(m.Preferences[i]) { // not required
			continue
		}

		if m.Preferences[i] != nil {
			if err := m.Preferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Preferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetUserPreferencesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetUserPreferencesResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetUserPreferencesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
