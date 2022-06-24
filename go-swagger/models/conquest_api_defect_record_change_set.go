// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIDefectRecordChangeSet conquest api defect record change set
//
// swagger:model conquest_apiDefectRecordChangeSet
type ConquestAPIDefectRecordChangeSet struct {

	// changes
	Changes []string `json:"Changes"`

	// defect ID
	DefectID int32 `json:"DefectID,omitempty"`

	// last edit
	// Format: date-time
	LastEdit strfmt.DateTime `json:"LastEdit,omitempty"`

	// original
	Original *ConquestAPIDefectRecord `json:"Original,omitempty"`

	// updated
	Updated *ConquestAPIDefectRecord `json:"Updated,omitempty"`
}

// Validate validates this conquest api defect record change set
func (m *ConquestAPIDefectRecordChangeSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEdit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIDefectRecordChangeSet) validateLastEdit(formats strfmt.Registry) error {
	if swag.IsZero(m.LastEdit) { // not required
		return nil
	}

	if err := validate.FormatOf("LastEdit", "body", "date-time", m.LastEdit.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIDefectRecordChangeSet) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Original")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIDefectRecordChangeSet) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if m.Updated != nil {
		if err := m.Updated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Updated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Updated")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api defect record change set based on the context it is used
func (m *ConquestAPIDefectRecordChangeSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIDefectRecordChangeSet) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {
		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Original")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIDefectRecordChangeSet) contextValidateUpdated(ctx context.Context, formats strfmt.Registry) error {

	if m.Updated != nil {
		if err := m.Updated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Updated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Updated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDefectRecordChangeSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDefectRecordChangeSet) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDefectRecordChangeSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
