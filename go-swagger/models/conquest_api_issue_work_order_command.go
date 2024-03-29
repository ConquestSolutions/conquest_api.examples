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

// ConquestAPIIssueWorkOrderCommand conquest api issue work order command
//
// swagger:model conquest_apiIssueWorkOrderCommand
type ConquestAPIIssueWorkOrderCommand struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// contractor email metadata
	ContractorEmailMetadata *ConquestAPIContractorEmailMetadata `json:"ContractorEmailMetadata,omitempty"`

	// issue by
	IssueBy *ConquestAPIIssueBy `json:"IssueBy,omitempty"`

	// issue date
	// Format: date-time
	IssueDate strfmt.DateTime `json:"IssueDate,omitempty"`

	// related documents
	RelatedDocuments []*ConquestAPIDocument `json:"RelatedDocuments"`

	// report
	Report *ConquestAPIReport `json:"Report,omitempty"`
}

// Validate validates this conquest api issue work order command
func (m *ConquestAPIIssueWorkOrderCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractorEmailMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) validateContractorEmailMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.ContractorEmailMetadata) { // not required
		return nil
	}

	if m.ContractorEmailMetadata != nil {
		if err := m.ContractorEmailMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContractorEmailMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContractorEmailMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) validateIssueBy(formats strfmt.Registry) error {
	if swag.IsZero(m.IssueBy) { // not required
		return nil
	}

	if m.IssueBy != nil {
		if err := m.IssueBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IssueBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IssueBy")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) validateIssueDate(formats strfmt.Registry) error {
	if swag.IsZero(m.IssueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("IssueDate", "body", "date-time", m.IssueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) validateRelatedDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedDocuments) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedDocuments); i++ {
		if swag.IsZero(m.RelatedDocuments[i]) { // not required
			continue
		}

		if m.RelatedDocuments[i] != nil {
			if err := m.RelatedDocuments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RelatedDocuments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RelatedDocuments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) validateReport(formats strfmt.Registry) error {
	if swag.IsZero(m.Report) { // not required
		return nil
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Report")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api issue work order command based on the context it is used
func (m *ConquestAPIIssueWorkOrderCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContractorEmailMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssueBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) contextValidateContractorEmailMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ContractorEmailMetadata != nil {
		if err := m.ContractorEmailMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContractorEmailMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContractorEmailMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) contextValidateIssueBy(ctx context.Context, formats strfmt.Registry) error {

	if m.IssueBy != nil {
		if err := m.IssueBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IssueBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IssueBy")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) contextValidateRelatedDocuments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedDocuments); i++ {

		if m.RelatedDocuments[i] != nil {
			if err := m.RelatedDocuments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RelatedDocuments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RelatedDocuments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIIssueWorkOrderCommand) contextValidateReport(ctx context.Context, formats strfmt.Registry) error {

	if m.Report != nil {
		if err := m.Report.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIIssueWorkOrderCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIIssueWorkOrderCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIIssueWorkOrderCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
