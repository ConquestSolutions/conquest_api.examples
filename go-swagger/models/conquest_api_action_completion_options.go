// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIActionCompletionOptions conquest api action completion options
//
// swagger:model conquest_apiActionCompletionOptions
type ConquestAPIActionCompletionOptions struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// CompletedBy is used to choose a different user to the one completing the action.
	CompletedBy string `json:"CompletedBy,omitempty"`

	// CompletionDate is used to set a completion date other than "now"
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"CompletionDate,omitempty"`

	// completion notes
	CompletionNotes string `json:"CompletionNotes,omitempty"`

	// CopyDocumentsToSucceedingActions is used to override the preference already set on the action
	CopyDocumentsToSucceedingActions bool `json:"CopyDocumentsToSucceedingActions,omitempty"`

	// CopyUserFieldsToSucceedingActions is used to override the preference already set on the action
	CopyUserFieldsToSucceedingActions bool `json:"CopyUserFieldsToSucceedingActions,omitempty"`

	// cost
	Cost *ConquestAPIDecimalValue `json:"Cost,omitempty"`

	// cost type
	CostType ConquestAPICostType `json:"CostType,omitempty"`

	// cyclic process
	CyclicProcess *ConquestAPICyclicActionCompletionProcess `json:"CyclicProcess,omitempty"`

	// disposal process
	DisposalProcess *ConquestAPIDisposalActionCompletionProcess `json:"DisposalProcess,omitempty"`

	// log book process
	LogBookProcess *ConquestAPILogBookActionCompletionProcess `json:"LogBookProcess,omitempty"`

	// mark as completed
	MarkAsCompleted bool `json:"MarkAsCompleted,omitempty"`

	// master process
	MasterProcess *ConquestAPIMasterActionCompletionProcess `json:"MasterProcess,omitempty"`

	// response date
	// Format: date-time
	ResponseDate strfmt.DateTime `json:"ResponseDate,omitempty"`

	// simple process
	SimpleProcess *ConquestAPISimpleActionCompletionProcess `json:"SimpleProcess,omitempty"`

	// use defaults process
	UseDefaultsProcess bool `json:"UseDefaultsProcess,omitempty"`
}

// Validate validates this conquest api action completion options
func (m *ConquestAPIActionCompletionOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCyclicProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisposalProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogBookProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimpleProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CompletionDate", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateCost(formats strfmt.Registry) error {

	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cost")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateCostType(formats strfmt.Registry) error {

	if swag.IsZero(m.CostType) { // not required
		return nil
	}

	if err := m.CostType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CostType")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateCyclicProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.CyclicProcess) { // not required
		return nil
	}

	if m.CyclicProcess != nil {
		if err := m.CyclicProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CyclicProcess")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateDisposalProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.DisposalProcess) { // not required
		return nil
	}

	if m.DisposalProcess != nil {
		if err := m.DisposalProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DisposalProcess")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateLogBookProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.LogBookProcess) { // not required
		return nil
	}

	if m.LogBookProcess != nil {
		if err := m.LogBookProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LogBookProcess")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateMasterProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterProcess) { // not required
		return nil
	}

	if m.MasterProcess != nil {
		if err := m.MasterProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MasterProcess")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateResponseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ResponseDate", "body", "date-time", m.ResponseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionOptions) validateSimpleProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.SimpleProcess) { // not required
		return nil
	}

	if m.SimpleProcess != nil {
		if err := m.SimpleProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SimpleProcess")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIActionCompletionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIActionCompletionOptions) UnmarshalBinary(b []byte) error {
	var res ConquestAPIActionCompletionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
