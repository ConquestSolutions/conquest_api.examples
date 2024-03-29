// Code generated by go-swagger; DO NOT EDIT.

package request_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// RequestServiceUpdateRequestReader is a Reader for the RequestServiceUpdateRequest structure.
type RequestServiceUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestServiceUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestServiceUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRequestServiceUpdateRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRequestServiceUpdateRequestOK creates a RequestServiceUpdateRequestOK with default headers values
func NewRequestServiceUpdateRequestOK() *RequestServiceUpdateRequestOK {
	return &RequestServiceUpdateRequestOK{}
}

/* RequestServiceUpdateRequestOK describes a response with status code 200, with default header values.

A successful response.
*/
type RequestServiceUpdateRequestOK struct {
	Payload *models.ConquestAPIChangeSetResult
}

func (o *RequestServiceUpdateRequestOK) Error() string {
	return fmt.Sprintf("[POST /api/requests/update_request][%d] requestServiceUpdateRequestOK  %+v", 200, o.Payload)
}
func (o *RequestServiceUpdateRequestOK) GetPayload() *models.ConquestAPIChangeSetResult {
	return o.Payload
}

func (o *RequestServiceUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIChangeSetResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestServiceUpdateRequestDefault creates a RequestServiceUpdateRequestDefault with default headers values
func NewRequestServiceUpdateRequestDefault(code int) *RequestServiceUpdateRequestDefault {
	return &RequestServiceUpdateRequestDefault{
		_statusCode: code,
	}
}

/* RequestServiceUpdateRequestDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RequestServiceUpdateRequestDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the request service update request default response
func (o *RequestServiceUpdateRequestDefault) Code() int {
	return o._statusCode
}

func (o *RequestServiceUpdateRequestDefault) Error() string {
	return fmt.Sprintf("[POST /api/requests/update_request][%d] RequestService_UpdateRequest default  %+v", o._statusCode, o.Payload)
}
func (o *RequestServiceUpdateRequestDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *RequestServiceUpdateRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
