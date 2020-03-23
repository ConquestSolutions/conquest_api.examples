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

// ConquestAPIGetAttributeSetFieldNamesResponse conquest api get attribute set field names response
//
// swagger:model conquest_apiGetAttributeSetFieldNamesResponse
type ConquestAPIGetAttributeSetFieldNamesResponse struct {

	// field names
	FieldNames []*ConquestAPIAttributeSetFieldNames `json:"FieldNames"`
}

// Validate validates this conquest api get attribute set field names response
func (m *ConquestAPIGetAttributeSetFieldNamesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetAttributeSetFieldNamesResponse) validateFieldNames(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldNames) { // not required
		return nil
	}

	for i := 0; i < len(m.FieldNames); i++ {
		if swag.IsZero(m.FieldNames[i]) { // not required
			continue
		}

		if m.FieldNames[i] != nil {
			if err := m.FieldNames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FieldNames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetAttributeSetFieldNamesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetAttributeSetFieldNamesResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetAttributeSetFieldNamesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
