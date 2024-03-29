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

// ConquestAPIJobCreationDetails JobCreationDetails of the job when submitted. It doesn't change once created.
//
// swagger:model conquest_apiJobCreationDetails
type ConquestAPIJobCreationDetails struct {

	// add document
	AddDocument *ConquestAPIAddDocumentCommand `json:"add_document,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// csv import
	CsvImport *ConquestAPIAddNewCsvImportCommand `json:"csv_import,omitempty"`

	// device id
	DeviceID string `json:"device_id,omitempty"`

	// empty
	Empty interface{} `json:"empty,omitempty"`

	// generate document link
	GenerateDocumentLink *ConquestAPIGenerateDocumentLinkCommand `json:"generate_document_link,omitempty"`

	// geopackage import
	GeopackageImport *ConquestAPIAddNewGeoPackageImportCommand `json:"geopackage_import,omitempty"`

	// geoupdate
	Geoupdate *ConquestAPIAddGeoUpdateCommand `json:"geoupdate,omitempty"`

	// issue work order
	IssueWorkOrder *ConquestAPIIssueWorkOrderCommand `json:"issue_work_order,omitempty"`

	// record key
	RecordKey *ConquestAPIRecordKey `json:"record_key,omitempty"`
}

// Validate validates this conquest api job creation details
func (m *ConquestAPIJobCreationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsvImport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenerateDocumentLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeopackageImport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeoupdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueWorkOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIJobCreationDetails) validateAddDocument(formats strfmt.Registry) error {
	if swag.IsZero(m.AddDocument) { // not required
		return nil
	}

	if m.AddDocument != nil {
		if err := m.AddDocument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_document")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateCsvImport(formats strfmt.Registry) error {
	if swag.IsZero(m.CsvImport) { // not required
		return nil
	}

	if m.CsvImport != nil {
		if err := m.CsvImport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csv_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csv_import")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateGenerateDocumentLink(formats strfmt.Registry) error {
	if swag.IsZero(m.GenerateDocumentLink) { // not required
		return nil
	}

	if m.GenerateDocumentLink != nil {
		if err := m.GenerateDocumentLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generate_document_link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generate_document_link")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateGeopackageImport(formats strfmt.Registry) error {
	if swag.IsZero(m.GeopackageImport) { // not required
		return nil
	}

	if m.GeopackageImport != nil {
		if err := m.GeopackageImport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geopackage_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geopackage_import")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateGeoupdate(formats strfmt.Registry) error {
	if swag.IsZero(m.Geoupdate) { // not required
		return nil
	}

	if m.Geoupdate != nil {
		if err := m.Geoupdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geoupdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geoupdate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateIssueWorkOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.IssueWorkOrder) { // not required
		return nil
	}

	if m.IssueWorkOrder != nil {
		if err := m.IssueWorkOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue_work_order")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issue_work_order")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) validateRecordKey(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordKey) { // not required
		return nil
	}

	if m.RecordKey != nil {
		if err := m.RecordKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conquest api job creation details based on the context it is used
func (m *ConquestAPIJobCreationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsvImport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGenerateDocumentLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeopackageImport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeoupdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssueWorkOrder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecordKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateAddDocument(ctx context.Context, formats strfmt.Registry) error {

	if m.AddDocument != nil {
		if err := m.AddDocument.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("add_document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("add_document")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateCsvImport(ctx context.Context, formats strfmt.Registry) error {

	if m.CsvImport != nil {
		if err := m.CsvImport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csv_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csv_import")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateGenerateDocumentLink(ctx context.Context, formats strfmt.Registry) error {

	if m.GenerateDocumentLink != nil {
		if err := m.GenerateDocumentLink.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generate_document_link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generate_document_link")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateGeopackageImport(ctx context.Context, formats strfmt.Registry) error {

	if m.GeopackageImport != nil {
		if err := m.GeopackageImport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geopackage_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geopackage_import")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateGeoupdate(ctx context.Context, formats strfmt.Registry) error {

	if m.Geoupdate != nil {
		if err := m.Geoupdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geoupdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geoupdate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateIssueWorkOrder(ctx context.Context, formats strfmt.Registry) error {

	if m.IssueWorkOrder != nil {
		if err := m.IssueWorkOrder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue_work_order")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issue_work_order")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIJobCreationDetails) contextValidateRecordKey(ctx context.Context, formats strfmt.Registry) error {

	if m.RecordKey != nil {
		if err := m.RecordKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("record_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("record_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIJobCreationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIJobCreationDetails) UnmarshalBinary(b []byte) error {
	var res ConquestAPIJobCreationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
