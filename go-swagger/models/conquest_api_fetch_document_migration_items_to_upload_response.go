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

// ConquestAPIFetchDocumentMigrationItemsToUploadResponse conquest api fetch document migration items to upload response
//
// swagger:model conquest_apiFetchDocumentMigrationItemsToUploadResponse
type ConquestAPIFetchDocumentMigrationItemsToUploadResponse struct {

	// items
	Items []*ConquestAPIDocumentUploadJobItem `json:"Items"`

	// work load
	WorkLoad *ConquestAPIDocumentMigrationWorkLoad `json:"WorkLoad,omitempty"`
}

// Validate validates this conquest api fetch document migration items to upload response
func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkLoad(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) validateWorkLoad(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkLoad) { // not required
		return nil
	}

	if m.WorkLoad != nil {
		if err := m.WorkLoad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkLoad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("WorkLoad")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api fetch document migration items to upload response based on the context it is used
func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkLoad(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) contextValidateWorkLoad(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkLoad != nil {
		if err := m.WorkLoad.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkLoad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("WorkLoad")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIFetchDocumentMigrationItemsToUploadResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPIFetchDocumentMigrationItemsToUploadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
