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

// ConquestAPIAddNewGeoPackageImportCommand conquest api add new geo package import command
//
// swagger:model conquest_apiAddNewGeoPackageImportCommand
type ConquestAPIAddNewGeoPackageImportCommand struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// table options
	TableOptions *ConquestAPIGeoPackageTableDescriptionOptions `json:"table_options,omitempty"`

	// user filename
	UserFilename string `json:"user_filename,omitempty"`
}

// Validate validates this conquest api add new geo package import command
func (m *ConquestAPIAddNewGeoPackageImportCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTableOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAddNewGeoPackageImportCommand) validateTableOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.TableOptions) { // not required
		return nil
	}

	if m.TableOptions != nil {
		if err := m.TableOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("table_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("table_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api add new geo package import command based on the context it is used
func (m *ConquestAPIAddNewGeoPackageImportCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTableOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAddNewGeoPackageImportCommand) contextValidateTableOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.TableOptions != nil {
		if err := m.TableOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("table_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("table_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAddNewGeoPackageImportCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAddNewGeoPackageImportCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAddNewGeoPackageImportCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}