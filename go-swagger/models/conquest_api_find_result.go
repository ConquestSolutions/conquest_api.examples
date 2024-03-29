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

// ConquestAPIFindResult conquest api find result
//
// swagger:model conquest_apiFindResult
type ConquestAPIFindResult struct {

	// filter name
	FilterName string `json:"FilterName,omitempty"`

	// record set
	RecordSet *ConquestAPIRecordSet `json:"RecordSet,omitempty"`

	// X_RelatedContexts is a list of ContextDescriptors that the returned objects can use
	XRelatedContexts []*ConquestAPIContextDescriptor `json:"X_RelatedContexts"`
}

// Validate validates this conquest api find result
func (m *ConquestAPIFindResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecordSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXRelatedContexts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFindResult) validateRecordSet(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordSet) { // not required
		return nil
	}

	if m.RecordSet != nil {
		if err := m.RecordSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RecordSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RecordSet")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIFindResult) validateXRelatedContexts(formats strfmt.Registry) error {
	if swag.IsZero(m.XRelatedContexts) { // not required
		return nil
	}

	for i := 0; i < len(m.XRelatedContexts); i++ {
		if swag.IsZero(m.XRelatedContexts[i]) { // not required
			continue
		}

		if m.XRelatedContexts[i] != nil {
			if err := m.XRelatedContexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("X_RelatedContexts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("X_RelatedContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api find result based on the context it is used
func (m *ConquestAPIFindResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecordSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXRelatedContexts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFindResult) contextValidateRecordSet(ctx context.Context, formats strfmt.Registry) error {

	if m.RecordSet != nil {
		if err := m.RecordSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RecordSet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RecordSet")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIFindResult) contextValidateXRelatedContexts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.XRelatedContexts); i++ {

		if m.XRelatedContexts[i] != nil {
			if err := m.XRelatedContexts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("X_RelatedContexts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("X_RelatedContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIFindResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIFindResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIFindResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
