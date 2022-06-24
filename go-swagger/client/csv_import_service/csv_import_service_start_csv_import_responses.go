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

// CsvImportServiceStartCsvImportReader is a Reader for the CsvImportServiceStartCsvImport structure.
type CsvImportServiceStartCsvImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CsvImportServiceStartCsvImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCsvImportServiceStartCsvImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCsvImportServiceStartCsvImportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCsvImportServiceStartCsvImportOK creates a CsvImportServiceStartCsvImportOK with default headers values
func NewCsvImportServiceStartCsvImportOK() *CsvImportServiceStartCsvImportOK {
	return &CsvImportServiceStartCsvImportOK{}
}

/* CsvImportServiceStartCsvImportOK describes a response with status code 200, with default header values.

A successful response.
*/
type CsvImportServiceStartCsvImportOK struct {
	Payload interface{}
}

func (o *CsvImportServiceStartCsvImportOK) Error() string {
	return fmt.Sprintf("[POST /api/import/start][%d] csvImportServiceStartCsvImportOK  %+v", 200, o.Payload)
}
func (o *CsvImportServiceStartCsvImportOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CsvImportServiceStartCsvImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCsvImportServiceStartCsvImportDefault creates a CsvImportServiceStartCsvImportDefault with default headers values
func NewCsvImportServiceStartCsvImportDefault(code int) *CsvImportServiceStartCsvImportDefault {
	return &CsvImportServiceStartCsvImportDefault{
		_statusCode: code,
	}
}

/* CsvImportServiceStartCsvImportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CsvImportServiceStartCsvImportDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the csv import service start csv import default response
func (o *CsvImportServiceStartCsvImportDefault) Code() int {
	return o._statusCode
}

func (o *CsvImportServiceStartCsvImportDefault) Error() string {
	return fmt.Sprintf("[POST /api/import/start][%d] CsvImportService_StartCsvImport default  %+v", o._statusCode, o.Payload)
}
func (o *CsvImportServiceStartCsvImportDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CsvImportServiceStartCsvImportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
