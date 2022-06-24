// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIFindContextsResult conquest api find contexts result
//
// swagger:model conquest_apiFindContextsResult
type ConquestAPIFindContextsResult struct {

	// context descriptors
	ContextDescriptors []*ConquestAPIContextDescriptor `json:"ContextDescriptors"`
}

// Validate validates this conquest api find contexts result
func (m *ConquestAPIFindContextsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContextDescriptors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFindContextsResult) validateContextDescriptors(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextDescriptors) { // not required
		return nil
	}

	for i := 0; i < len(m.ContextDescriptors); i++ {
		if swag.IsZero(m.ContextDescriptors[i]) { // not required
			continue
		}

		if m.ContextDescriptors[i] != nil {
			if err := m.ContextDescriptors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ContextDescriptors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ContextDescriptors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api find contexts result based on the context it is used
func (m *ConquestAPIFindContextsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContextDescriptors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFindContextsResult) contextValidateContextDescriptors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContextDescriptors); i++ {

		if m.ContextDescriptors[i] != nil {
			if err := m.ContextDescriptors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ContextDescriptors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ContextDescriptors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIFindContextsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIFindContextsResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIFindContextsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
