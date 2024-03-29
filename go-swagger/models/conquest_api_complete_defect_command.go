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

// ConquestAPICompleteDefectCommand conquest api complete defect command
//
// swagger:model conquest_apiCompleteDefectCommand
type ConquestAPICompleteDefectCommand struct {

	// completion date
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"CompletionDate,omitempty"`

	// defect ID
	DefectID int32 `json:"DefectID,omitempty"`
}

// Validate validates this conquest api complete defect command
func (m *ConquestAPICompleteDefectCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICompleteDefectCommand) validateCompletionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CompletionDate", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conquest api complete defect command based on context it is used
func (m *ConquestAPICompleteDefectCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICompleteDefectCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICompleteDefectCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPICompleteDefectCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
