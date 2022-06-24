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

// ConquestAPIMapObjectsRecordSet conquest api map objects record set
//
// swagger:model conquest_apiMapObjectsRecordSet
type ConquestAPIMapObjectsRecordSet struct {

	// objects
	Objects []*ConquestAPIMapObject `json:"Objects"`

	// context descriptor
	ContextDescriptor *ConquestAPIContextDescriptor `json:"contextDescriptor,omitempty"`

	// context descriptors for groups
	ContextDescriptorsForGroups []*ConquestAPIContextDescriptor `json:"contextDescriptorsForGroups"`
}

// Validate validates this conquest api map objects record set
func (m *ConquestAPIMapObjectsRecordSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextDescriptorsForGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) validateContextDescriptor(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextDescriptor) { // not required
		return nil
	}

	if m.ContextDescriptor != nil {
		if err := m.ContextDescriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextDescriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contextDescriptor")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) validateContextDescriptorsForGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextDescriptorsForGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ContextDescriptorsForGroups); i++ {
		if swag.IsZero(m.ContextDescriptorsForGroups[i]) { // not required
			continue
		}

		if m.ContextDescriptorsForGroups[i] != nil {
			if err := m.ContextDescriptorsForGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contextDescriptorsForGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contextDescriptorsForGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api map objects record set based on the context it is used
func (m *ConquestAPIMapObjectsRecordSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContextDescriptor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContextDescriptorsForGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {
			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) contextValidateContextDescriptor(ctx context.Context, formats strfmt.Registry) error {

	if m.ContextDescriptor != nil {
		if err := m.ContextDescriptor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextDescriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contextDescriptor")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIMapObjectsRecordSet) contextValidateContextDescriptorsForGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContextDescriptorsForGroups); i++ {

		if m.ContextDescriptorsForGroups[i] != nil {
			if err := m.ContextDescriptorsForGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contextDescriptorsForGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contextDescriptorsForGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIMapObjectsRecordSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIMapObjectsRecordSet) UnmarshalBinary(b []byte) error {
	var res ConquestAPIMapObjectsRecordSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
