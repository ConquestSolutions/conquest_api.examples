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

// ConquestAPIAddDocumentResult AddDocumentResult returns the saved Document record and the location (UploadMethod & UploadUri) of where the actual file should be pushed (POST/PUT).
//
// The location may point to the API server itself, or a configured blob storage service.
//
// This location is valid for a period of time or until the upload has completed. Whichever is earliest.
//
// When a client performs an upload, it prepares a request that will write the file to the provided location. If the location *may* require credentials.
//
// swagger:model conquest_apiAddDocumentResult
type ConquestAPIAddDocumentResult struct {

	// document
	Document *ConquestAPIDocument `json:"Document,omitempty"`

	// upload expire time
	// Format: date-time
	UploadExpireTime strfmt.DateTime `json:"UploadExpireTime,omitempty"`

	// UploadHeaders returns some required headers, like "x-ms-blob-type: BlockBlob" for azure uploads
	UploadHeaders map[string]string `json:"UploadHeaders,omitempty"`

	// POST for a web form or PUT to just push the data
	UploadMethod string `json:"UploadMethod,omitempty"`

	// UploadUri is where you must post the UploadData.
	// It is a multipart HTTP upload with the field "document"
	// for which the document is uploaded to.
	//
	// For example, you can upload 'inspection-photo.png' like this for a standard web form
	// ```
	//   curl -i -X POST -H "Content-Type: multipart/form-data" -F "document=@inspection-photo.png" "${ApiHost}${UploadUri}"
	// ```
	//
	// Alternatively, use if no metadata needs to be submitted.
	// ```
	//   curl --upload-file "${ApiHost}${UploadUri}"
	// ```
	//
	// If the Uri is relative, prefix it with the api origin.
	//
	// UploadUri will be empty if:
	// - The Added Document was a link to an existing document (LinkExistingDocument=true)
	// - The Document Link is not relative to one of the Document Locations (see GetHierarchyNodesQuery{ObjectType=DocumentContainer})
	UploadURI string `json:"UploadUri,omitempty"`
}

// Validate validates this conquest api add document result
func (m *ConquestAPIAddDocumentResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadExpireTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAddDocumentResult) validateDocument(formats strfmt.Registry) error {

	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Document")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAddDocumentResult) validateUploadExpireTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadExpireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UploadExpireTime", "body", "date-time", m.UploadExpireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAddDocumentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAddDocumentResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAddDocumentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
