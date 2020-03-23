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

// ConquestAPIUserHistoryItem conquest api user history item
//
// swagger:model conquest_apiUserHistoryItem
type ConquestAPIUserHistoryItem struct {

	// LastUsed is the last time a form was opened
	// Format: date-time
	LastUsed strfmt.DateTime `json:"LastUsed,omitempty"`

	// string Username = 1;
	Link *ConquestAPIItemLink `json:"Link,omitempty"`

	// Order, if a history item has an order, it is a "pinned item", otherwise it is a "recent item"
	Order int32 `json:"Order,omitempty"`
}

// Validate validates this conquest api user history item
func (m *ConquestAPIUserHistoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIUserHistoryItem) validateLastUsed(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUsed) { // not required
		return nil
	}

	if err := validate.FormatOf("LastUsed", "body", "date-time", m.LastUsed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIUserHistoryItem) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIUserHistoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIUserHistoryItem) UnmarshalBinary(b []byte) error {
	var res ConquestAPIUserHistoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
