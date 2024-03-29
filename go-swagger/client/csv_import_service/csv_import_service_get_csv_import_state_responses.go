// Code generated by go-swagger; DO NOT EDIT.

package csv_import_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// CsvImportServiceGetCsvImportStateReader is a Reader for the CsvImportServiceGetCsvImportState structure.
type CsvImportServiceGetCsvImportStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CsvImportServiceGetCsvImportStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCsvImportServiceGetCsvImportStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCsvImportServiceGetCsvImportStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCsvImportServiceGetCsvImportStateOK creates a CsvImportServiceGetCsvImportStateOK with default headers values
func NewCsvImportServiceGetCsvImportStateOK() *CsvImportServiceGetCsvImportStateOK {
	return &CsvImportServiceGetCsvImportStateOK{}
}

/* CsvImportServiceGetCsvImportStateOK describes a response with status code 200, with default header values.

A successful response.
*/
type CsvImportServiceGetCsvImportStateOK struct {
	Payload *models.ConquestAPICsvImportStateResponse
}

func (o *CsvImportServiceGetCsvImportStateOK) Error() string {
	return fmt.Sprintf("[GET /api/import/state/{JobID}][%d] csvImportServiceGetCsvImportStateOK  %+v", 200, o.Payload)
}
func (o *CsvImportServiceGetCsvImportStateOK) GetPayload() *models.ConquestAPICsvImportStateResponse {
	return o.Payload
}

func (o *CsvImportServiceGetCsvImportStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPICsvImportStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCsvImportServiceGetCsvImportStateDefault creates a CsvImportServiceGetCsvImportStateDefault with default headers values
func NewCsvImportServiceGetCsvImportStateDefault(code int) *CsvImportServiceGetCsvImportStateDefault {
	return &CsvImportServiceGetCsvImportStateDefault{
		_statusCode: code,
	}
}

/* CsvImportServiceGetCsvImportStateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CsvImportServiceGetCsvImportStateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the csv import service get csv import state default response
func (o *CsvImportServiceGetCsvImportStateDefault) Code() int {
	return o._statusCode
}

func (o *CsvImportServiceGetCsvImportStateDefault) Error() string {
	return fmt.Sprintf("[GET /api/import/state/{JobID}][%d] CsvImportService_GetCsvImportState default  %+v", o._statusCode, o.Payload)
}
func (o *CsvImportServiceGetCsvImportStateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CsvImportServiceGetCsvImportStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
