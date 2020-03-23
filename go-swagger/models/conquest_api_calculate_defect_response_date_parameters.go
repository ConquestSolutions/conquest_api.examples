// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPICalculateDefectResponseDateParameters Parameters to calculate a Defect's Response Date calculated from the Current Inspection Date, Asset's Priority and Severity.
//
// swagger:model conquest_apiCalculateDefectResponseDateParameters
type ConquestAPICalculateDefectResponseDateParameters struct {

	// The Defect for which Severities and their the Response Times are configured (on it's Standard Defect)
	DefectID int32 `json:"DefectID,omitempty"`

	// The Severity chosen for the defect
	SeverityID int32 `json:"SeverityID,omitempty"`
}

// Validate validates this conquest api calculate defect response date parameters
func (m *ConquestAPICalculateDefectResponseDateParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICalculateDefectResponseDateParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICalculateDefectResponseDateParameters) UnmarshalBinary(b []byte) error {
	var res ConquestAPICalculateDefectResponseDateParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
