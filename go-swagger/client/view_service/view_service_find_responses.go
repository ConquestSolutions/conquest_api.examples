// Code generated by go-swagger; DO NOT EDIT.

package view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// ViewServiceFindReader is a Reader for the ViewServiceFind structure.
type ViewServiceFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewServiceFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewViewServiceFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewViewServiceFindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewViewServiceFindOK creates a ViewServiceFindOK with default headers values
func NewViewServiceFindOK() *ViewServiceFindOK {
	return &ViewServiceFindOK{}
}

/* ViewServiceFindOK describes a response with status code 200, with default header values.

A successful response.
*/
type ViewServiceFindOK struct {
	Payload *models.ConquestAPIFindResult
}

func (o *ViewServiceFindOK) Error() string {
	return fmt.Sprintf("[POST /api/find][%d] viewServiceFindOK  %+v", 200, o.Payload)
}
func (o *ViewServiceFindOK) GetPayload() *models.ConquestAPIFindResult {
	return o.Payload
}

func (o *ViewServiceFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIFindResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewServiceFindDefault creates a ViewServiceFindDefault with default headers values
func NewViewServiceFindDefault(code int) *ViewServiceFindDefault {
	return &ViewServiceFindDefault{
		_statusCode: code,
	}
}

/* ViewServiceFindDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ViewServiceFindDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the view service find default response
func (o *ViewServiceFindDefault) Code() int {
	return o._statusCode
}

func (o *ViewServiceFindDefault) Error() string {
	return fmt.Sprintf("[POST /api/find][%d] ViewService_Find default  %+v", o._statusCode, o.Payload)
}
func (o *ViewServiceFindDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ViewServiceFindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
