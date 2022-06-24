// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// ActionServiceGetActionReader is a Reader for the ActionServiceGetAction structure.
type ActionServiceGetActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionServiceGetActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionServiceGetActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionServiceGetActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionServiceGetActionOK creates a ActionServiceGetActionOK with default headers values
func NewActionServiceGetActionOK() *ActionServiceGetActionOK {
	return &ActionServiceGetActionOK{}
}

/* ActionServiceGetActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActionServiceGetActionOK struct {
	Payload *models.ConquestAPIActionEntity
}

func (o *ActionServiceGetActionOK) Error() string {
	return fmt.Sprintf("[GET /api/actions/{value}][%d] actionServiceGetActionOK  %+v", 200, o.Payload)
}
func (o *ActionServiceGetActionOK) GetPayload() *models.ConquestAPIActionEntity {
	return o.Payload
}

func (o *ActionServiceGetActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIActionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionServiceGetActionDefault creates a ActionServiceGetActionDefault with default headers values
func NewActionServiceGetActionDefault(code int) *ActionServiceGetActionDefault {
	return &ActionServiceGetActionDefault{
		_statusCode: code,
	}
}

/* ActionServiceGetActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActionServiceGetActionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the action service get action default response
func (o *ActionServiceGetActionDefault) Code() int {
	return o._statusCode
}

func (o *ActionServiceGetActionDefault) Error() string {
	return fmt.Sprintf("[GET /api/actions/{value}][%d] ActionService_GetAction default  %+v", o._statusCode, o.Payload)
}
func (o *ActionServiceGetActionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ActionServiceGetActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
