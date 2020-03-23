// Code generated by go-swagger; DO NOT EDIT.

package document_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// AddDocumentReader is a Reader for the AddDocument structure.
type AddDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDocumentOK creates a AddDocumentOK with default headers values
func NewAddDocumentOK() *AddDocumentOK {
	return &AddDocumentOK{}
}

/*AddDocumentOK handles this case with default header values.

A successful response.
*/
type AddDocumentOK struct {
	Payload *models.ConquestAPIAddDocumentResult
}

func (o *AddDocumentOK) Error() string {
	return fmt.Sprintf("[POST /api/documents/add_document][%d] addDocumentOK  %+v", 200, o.Payload)
}

func (o *AddDocumentOK) GetPayload() *models.ConquestAPIAddDocumentResult {
	return o.Payload
}

func (o *AddDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIAddDocumentResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
