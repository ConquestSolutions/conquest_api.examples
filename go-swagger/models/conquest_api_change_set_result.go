// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIChangeSetResult conquest api change set result
//
// swagger:model conquest_apiChangeSetResult
type ConquestAPIChangeSetResult struct {

	// accepted
	Accepted bool `json:"Accepted,omitempty"`

	// editor
	Editor string `json:"Editor,omitempty"`

	// last edit
	// Format: date-time
	LastEdit strfmt.DateTime `json:"LastEdit,omitempty"`

	// object key
	ObjectKey *ConquestAPIObjectKey `json:"ObjectKey,omitempty"`
}

// Validate validates this conquest api change set result
func (m *ConquestAPIChangeSetResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEdit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIChangeSetResult) validateLastEdit(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEdit) { // not required
		return nil
	}

	if err := validate.FormatOf("LastEdit", "body", "date-time", m.LastEdit.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIChangeSetResult) validateObjectKey(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ConquestAPIChangeSetResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIChangeSetResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIChangeSetResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
