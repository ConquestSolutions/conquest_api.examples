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

// ActionServiceUpdateActionReader is a Reader for the ActionServiceUpdateAction structure.
type ActionServiceUpdateActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionServiceUpdateActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionServiceUpdateActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionServiceUpdateActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionServiceUpdateActionOK creates a ActionServiceUpdateActionOK with default headers values
func NewActionServiceUpdateActionOK() *ActionServiceUpdateActionOK {
	return &ActionServiceUpdateActionOK{}
}

/* ActionServiceUpdateActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActionServiceUpdateActionOK struct {
	Payload *models.ConquestAPIChangeSetResult
}

func (o *ActionServiceUpdateActionOK) Error() string {
	return fmt.Sprintf("[POST /api/actions/update_action][%d] actionServiceUpdateActionOK  %+v", 200, o.Payload)
}
func (o *ActionServiceUpdateActionOK) GetPayload() *models.ConquestAPIChangeSetResult {
	return o.Payload
}

func (o *ActionServiceUpdateActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIChangeSetResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionServiceUpdateActionDefault creates a ActionServiceUpdateActionDefault with default headers values
func NewActionServiceUpdateActionDefault(code int) *ActionServiceUpdateActionDefault {
	return &ActionServiceUpdateActionDefault{
		_statusCode: code,
	}
}

/* ActionServiceUpdateActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActionServiceUpdateActionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the action service update action default response
func (o *ActionServiceUpdateActionDefault) Code() int {
	return o._statusCode
}

func (o *ActionServiceUpdateActionDefault) Error() string {
	return fmt.Sprintf("[POST /api/actions/update_action][%d] ActionService_UpdateAction default  %+v", o._statusCode, o.Payload)
}
func (o *ActionServiceUpdateActionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ActionServiceUpdateActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
