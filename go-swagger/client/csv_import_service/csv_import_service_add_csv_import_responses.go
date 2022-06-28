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

// CsvImportServiceAddCsvImportReader is a Reader for the CsvImportServiceAddCsvImport structure.
type CsvImportServiceAddCsvImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CsvImportServiceAddCsvImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCsvImportServiceAddCsvImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCsvImportServiceAddCsvImportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCsvImportServiceAddCsvImportOK creates a CsvImportServiceAddCsvImportOK with default headers values
func NewCsvImportServiceAddCsvImportOK() *CsvImportServiceAddCsvImportOK {
	return &CsvImportServiceAddCsvImportOK{}
}

/* CsvImportServiceAddCsvImportOK describes a response with status code 200, with default header values.

A successful response.
*/
type CsvImportServiceAddCsvImportOK struct {
	Payload string
}

func (o *CsvImportServiceAddCsvImportOK) Error() string {
	return fmt.Sprintf("[POST /api/import/add/{ImportType}][%d] csvImportServiceAddCsvImportOK  %+v", 200, o.Payload)
}
func (o *CsvImportServiceAddCsvImportOK) GetPayload() string {
	return o.Payload
}

func (o *CsvImportServiceAddCsvImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCsvImportServiceAddCsvImportDefault creates a CsvImportServiceAddCsvImportDefault with default headers values
func NewCsvImportServiceAddCsvImportDefault(code int) *CsvImportServiceAddCsvImportDefault {
	return &CsvImportServiceAddCsvImportDefault{
		_statusCode: code,
	}
}

/* CsvImportServiceAddCsvImportDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type CsvImportServiceAddCsvImportDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the csv import service add csv import default response
func (o *CsvImportServiceAddCsvImportDefault) Code() int {
	return o._statusCode
}

func (o *CsvImportServiceAddCsvImportDefault) Error() string {
	return fmt.Sprintf("[POST /api/import/add/{ImportType}][%d] CsvImportService_AddCsvImport default  %+v", o._statusCode, o.Payload)
}
func (o *CsvImportServiceAddCsvImportDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CsvImportServiceAddCsvImportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}