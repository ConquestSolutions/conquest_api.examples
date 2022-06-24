// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ConquestAPIUploadStatus conquest api upload status
//
// swagger:model conquest_apiUploadStatus
type ConquestAPIUploadStatus string

func NewConquestAPIUploadStatus(value ConquestAPIUploadStatus) *ConquestAPIUploadStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConquestAPIUploadStatus.
func (m ConquestAPIUploadStatus) Pointer() *ConquestAPIUploadStatus {
	return &m
}

const (

	// ConquestAPIUploadStatusUploadStatusNotStarted captures enum value "UploadStatus_NotStarted"
	ConquestAPIUploadStatusUploadStatusNotStarted ConquestAPIUploadStatus = "UploadStatus_NotStarted"

	// ConquestAPIUploadStatusUploadStatusError captures enum value "UploadStatus_Error"
	ConquestAPIUploadStatusUploadStatusError ConquestAPIUploadStatus = "UploadStatus_Error"

	// ConquestAPIUploadStatusUploadStatusUploading captures enum value "UploadStatus_Uploading"
	ConquestAPIUploadStatusUploadStatusUploading ConquestAPIUploadStatus = "UploadStatus_Uploading"

	// ConquestAPIUploadStatusUploadStatusReceived captures enum value "UploadStatus_Received"
	ConquestAPIUploadStatusUploadStatusReceived ConquestAPIUploadStatus = "UploadStatus_Received"

	// ConquestAPIUploadStatusUploadStatusCompleted captures enum value "UploadStatus_Completed"
	ConquestAPIUploadStatusUploadStatusCompleted ConquestAPIUploadStatus = "UploadStatus_Completed"
)

// for schema
var conquestApiUploadStatusEnum []interface{}

func init() {
	var res []ConquestAPIUploadStatus
	if err := json.Unmarshal([]byte(`["UploadStatus_NotStarted","UploadStatus_Error","UploadStatus_Uploading","UploadStatus_Received","UploadStatus_Completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conquestApiUploadStatusEnum = append(conquestApiUploadStatusEnum, v)
	}
}

func (m ConquestAPIUploadStatus) validateConquestAPIUploadStatusEnum(path, location string, value ConquestAPIUploadStatus) error {
	if err := validate.EnumCase(path, location, value, conquestApiUploadStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this conquest api upload status
func (m ConquestAPIUploadStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConquestAPIUploadStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this conquest api upload status based on context it is used
func (m ConquestAPIUploadStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
