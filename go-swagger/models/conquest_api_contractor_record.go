// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIContractorRecord conquest api contractor record
//
// swagger:model conquest_apiContractorRecord
type ConquestAPIContractorRecord struct {

	// Australian Company Name (ACN)
	ACN string `json:"ACN,omitempty"`

	// action categories
	ActionCategories *ConquestAPIActionCategoryList `json:"ActionCategories,omitempty"`

	// address
	Address string `json:"Address,omitempty"`

	// contact
	Contact string `json:"Contact,omitempty"`

	// contractor name
	ContractorName string `json:"ContractorName,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// phone1
	Phone1 string `json:"Phone1,omitempty"`

	// phone2
	Phone2 string `json:"Phone2,omitempty"`

	// post code
	PostCode string `json:"PostCode,omitempty"`

	// sub contractors
	SubContractors *ConquestAPISubContractorList `json:"SubContractors,omitempty"`

	// suburb
	Suburb string `json:"Suburb,omitempty"`

	// user text1
	UserText1 string `json:"UserText1,omitempty"`
}

// Validate validates this conquest api contractor record
func (m *ConquestAPIContractorRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubContractors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIContractorRecord) validateActionCategories(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionCategories) { // not required
		return nil
	}

	if m.ActionCategories != nil {
		if err := m.ActionCategories.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionCategories")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionCategories")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIContractorRecord) validateSubContractors(formats strfmt.Registry) error {
	if swag.IsZero(m.SubContractors) { // not required
		return nil
	}

	if m.SubContractors != nil {
		if err := m.SubContractors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubContractors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubContractors")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api contractor record based on the context it is used
func (m *ConquestAPIContractorRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubContractors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIContractorRecord) contextValidateActionCategories(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionCategories != nil {
		if err := m.ActionCategories.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionCategories")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionCategories")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIContractorRecord) contextValidateSubContractors(ctx context.Context, formats strfmt.Registry) error {

	if m.SubContractors != nil {
		if err := m.SubContractors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubContractors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubContractors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIContractorRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIContractorRecord) UnmarshalBinary(b []byte) error {
	var res ConquestAPIContractorRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
