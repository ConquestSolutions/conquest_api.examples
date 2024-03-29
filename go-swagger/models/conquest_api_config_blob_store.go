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

// ConquestAPIConfigBlobStore conquest api config blob store
//
// swagger:model conquest_api_configBlobStore
type ConquestAPIConfigBlobStore struct {

	// blob store mode
	BlobStoreMode *ConquestAPIConfigBlobStoreMode `json:"blob_store_mode,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// A mount is a location. Mounting is the attaching of an additional file system to an application.
	// You refer to a file in a mount with the scheme 'blob://{mount}/{relativeFilePath}'
	Mount string `json:"mount,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// 'azure' provider will store documents in the container
	// 'local' provider will store documents in the {DataDirectory}/BlobStorage directory or
	Type string `json:"type,omitempty"`
}

// Validate validates this conquest api config blob store
func (m *ConquestAPIConfigBlobStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlobStoreMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIConfigBlobStore) validateBlobStoreMode(formats strfmt.Registry) error {
	if swag.IsZero(m.BlobStoreMode) { // not required
		return nil
	}

	if m.BlobStoreMode != nil {
		if err := m.BlobStoreMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blob_store_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blob_store_mode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api config blob store based on the context it is used
func (m *ConquestAPIConfigBlobStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlobStoreMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIConfigBlobStore) contextValidateBlobStoreMode(ctx context.Context, formats strfmt.Registry) error {

	if m.BlobStoreMode != nil {
		if err := m.BlobStoreMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blob_store_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blob_store_mode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIConfigBlobStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIConfigBlobStore) UnmarshalBinary(b []byte) error {
	var res ConquestAPIConfigBlobStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
