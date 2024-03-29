// Code generated by go-swagger; DO NOT EDIT.

package document_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// DocumentServiceGetDocumentsReader is a Reader for the DocumentServiceGetDocuments structure.
type DocumentServiceGetDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentServiceGetDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentServiceGetDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDocumentServiceGetDocumentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDocumentServiceGetDocumentsOK creates a DocumentServiceGetDocumentsOK with default headers values
func NewDocumentServiceGetDocumentsOK() *DocumentServiceGetDocumentsOK {
	return &DocumentServiceGetDocumentsOK{}
}

/* DocumentServiceGetDocumentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DocumentServiceGetDocumentsOK struct {
	Payload *models.ConquestAPIGetDocumentsResult
}

func (o *DocumentServiceGetDocumentsOK) Error() string {
	return fmt.Sprintf("[POST /api/documents/list][%d] documentServiceGetDocumentsOK  %+v", 200, o.Payload)
}
func (o *DocumentServiceGetDocumentsOK) GetPayload() *models.ConquestAPIGetDocumentsResult {
	return o.Payload
}

func (o *DocumentServiceGetDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIGetDocumentsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentServiceGetDocumentsDefault creates a DocumentServiceGetDocumentsDefault with default headers values
func NewDocumentServiceGetDocumentsDefault(code int) *DocumentServiceGetDocumentsDefault {
	return &DocumentServiceGetDocumentsDefault{
		_statusCode: code,
	}
}

/* DocumentServiceGetDocumentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DocumentServiceGetDocumentsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the document service get documents default response
func (o *DocumentServiceGetDocumentsDefault) Code() int {
	return o._statusCode
}

func (o *DocumentServiceGetDocumentsDefault) Error() string {
	return fmt.Sprintf("[POST /api/documents/list][%d] DocumentService_GetDocuments default  %+v", o._statusCode, o.Payload)
}
func (o *DocumentServiceGetDocumentsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DocumentServiceGetDocumentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
