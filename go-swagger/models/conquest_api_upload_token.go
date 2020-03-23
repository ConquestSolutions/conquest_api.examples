// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIUploadToken conquest api upload token
//
// swagger:model conquest_apiUploadToken
type ConquestAPIUploadToken struct {

	// address
	Address string `json:"Address,omitempty"`

	// object key
	ObjectKey *ConquestAPIObjectKey `json:"ObjectKey,omitempty"`

	// status
	Status ConquestAPIUploadStatus `json:"Status,omitempty"`

	// token
	Token string `json:"Token,omitempty"`
}

// Validate validates this conquest api upload token
func (m *ConquestAPIUploadToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIUploadToken) validateObjectKey(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectKey) { // not required
		return nil
	}

	if m.ObjectKey != nil {
		if err := m.ObjectKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIUploadToken) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIUploadToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIUploadToken) UnmarshalBinary(b []byte) error {
	var res ConquestAPIUploadToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
