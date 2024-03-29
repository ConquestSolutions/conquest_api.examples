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
	"github.com/go-openapi/validate"
)

// ConquestAPIGetDocumentsQuery conquest api get documents query
//
// swagger:model conquest_apiGetDocumentsQuery
type ConquestAPIGetDocumentsQuery struct {

	// ObjectKey for supported ObjectTypes: Asset, Action, Defect, Request, DocumentContainer
	ObjectKey *ConquestAPIObjectKey `json:"ObjectKey,omitempty"`

	// WithCreateTimes is used to specify the particular document records to retrieve, by the unique document timestamp 'CreateTime'
	WithCreateTimes []strfmt.DateTime `json:"WithCreateTimes"`

	// WithDocumentIDs is used to specify the particular document records to retrieve, by the unique document timestamp 'CreateTime'
	WithDocumentIDs []int32 `json:"WithDocumentIDs"`

	// If unset, all statuses will be shown.
	// When rendering this list, only those with the Completed status are downloadable.
	WithUploadStatus []*ConquestAPIUploadStatus `json:"WithUploadStatus"`
}

// Validate validates this conquest api get documents query
func (m *ConquestAPIGetDocumentsQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithCreateTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithUploadStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetDocumentsQuery) validateObjectKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectKey) { // not required
		return nil
	}

	if m.ObjectKey != nil {
		if err := m.ObjectKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIGetDocumentsQuery) validateWithCreateTimes(formats strfmt.Registry) error {
	if swag.IsZero(m.WithCreateTimes) { // not required
		return nil
	}

	for i := 0; i < len(m.WithCreateTimes); i++ {

		if err := validate.FormatOf("WithCreateTimes"+"."+strconv.Itoa(i), "body", "date-time", m.WithCreateTimes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConquestAPIGetDocumentsQuery) validateWithUploadStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.WithUploadStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.WithUploadStatus); i++ {
		if swag.IsZero(m.WithUploadStatus[i]) { // not required
			continue
		}

		if m.WithUploadStatus[i] != nil {
			if err := m.WithUploadStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WithUploadStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("WithUploadStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this conquest api get documents query based on the context it is used
func (m *ConquestAPIGetDocumentsQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWithUploadStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIGetDocumentsQuery) contextValidateObjectKey(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectKey != nil {
		if err := m.ObjectKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIGetDocumentsQuery) contextValidateWithUploadStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WithUploadStatus); i++ {

		if m.WithUploadStatus[i] != nil {
			if err := m.WithUploadStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WithUploadStatus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("WithUploadStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIGetDocumentsQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIGetDocumentsQuery) UnmarshalBinary(b []byte) error {
	var res ConquestAPIGetDocumentsQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
