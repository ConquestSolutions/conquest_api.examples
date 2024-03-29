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

// ConquestAPITrimDocumentsResponse conquest api trim documents response
//
// swagger:model conquest_apiTrimDocumentsResponse
type ConquestAPITrimDocumentsResponse struct {

	// trim documents
	TrimDocuments []*ConquestAPITrimDocument `json:"TrimDocuments"`
}

// Validate validates this conquest api trim documents response
func (m *ConquestAPITrimDocumentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrimDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITrimDocumentsResponse) validateTrimDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.TrimDocuments) { // not required
		return nil
	}

	for i := 0; i < len(m.TrimDocuments); i++ {
		if swag.IsZero(m.TrimDocuments[i]) { // not required
			continue
		}

		if m.TrimDocuments[i] != nil {
			if err := m.TrimDocuments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TrimDocuments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TrimDocuments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api trim documents response based on the context it is used
func (m *ConquestAPITrimDocumentsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrimDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITrimDocumentsResponse) contextValidateTrimDocuments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TrimDocuments); i++ {

		if m.TrimDocuments[i] != nil {
			if err := m.TrimDocuments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TrimDocuments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TrimDocuments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPITrimDocumentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPITrimDocumentsResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPITrimDocumentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
