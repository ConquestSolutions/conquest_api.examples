// Code generated by go-swagger; DO NOT EDIT.

package defect_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// DefectServiceGetDefectReader is a Reader for the DefectServiceGetDefect structure.
type DefectServiceGetDefectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DefectServiceGetDefectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDefectServiceGetDefectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDefectServiceGetDefectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDefectServiceGetDefectOK creates a DefectServiceGetDefectOK with default headers values
func NewDefectServiceGetDefectOK() *DefectServiceGetDefectOK {
	return &DefectServiceGetDefectOK{}
}

/* DefectServiceGetDefectOK describes a response with status code 200, with default header values.

A successful response.
*/
type DefectServiceGetDefectOK struct {
	Payload *models.ConquestAPIDefectEntity
}

func (o *DefectServiceGetDefectOK) Error() string {
	return fmt.Sprintf("[GET /api/defects/{value}][%d] defectServiceGetDefectOK  %+v", 200, o.Payload)
}
func (o *DefectServiceGetDefectOK) GetPayload() *models.ConquestAPIDefectEntity {
	return o.Payload
}

func (o *DefectServiceGetDefectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIDefectEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDefectServiceGetDefectDefault creates a DefectServiceGetDefectDefault with default headers values
func NewDefectServiceGetDefectDefault(code int) *DefectServiceGetDefectDefault {
	return &DefectServiceGetDefectDefault{
		_statusCode: code,
	}
}

/* DefectServiceGetDefectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DefectServiceGetDefectDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the defect service get defect default response
func (o *DefectServiceGetDefectDefault) Code() int {
	return o._statusCode
}

func (o *DefectServiceGetDefectDefault) Error() string {
	return fmt.Sprintf("[GET /api/defects/{value}][%d] DefectService_GetDefect default  %+v", o._statusCode, o.Payload)
}
func (o *DefectServiceGetDefectDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DefectServiceGetDefectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}