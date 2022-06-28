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

// DefectServiceCalculateDefectResponseDateReader is a Reader for the DefectServiceCalculateDefectResponseDate structure.
type DefectServiceCalculateDefectResponseDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DefectServiceCalculateDefectResponseDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDefectServiceCalculateDefectResponseDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDefectServiceCalculateDefectResponseDateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDefectServiceCalculateDefectResponseDateOK creates a DefectServiceCalculateDefectResponseDateOK with default headers values
func NewDefectServiceCalculateDefectResponseDateOK() *DefectServiceCalculateDefectResponseDateOK {
	return &DefectServiceCalculateDefectResponseDateOK{}
}

/* DefectServiceCalculateDefectResponseDateOK describes a response with status code 200, with default header values.

A successful response.
*/
type DefectServiceCalculateDefectResponseDateOK struct {
	Payload *models.ConquestAPICalculateDefectResponseDateResponse
}

func (o *DefectServiceCalculateDefectResponseDateOK) Error() string {
	return fmt.Sprintf("[POST /api/defects/calculate_defect_response_date][%d] defectServiceCalculateDefectResponseDateOK  %+v", 200, o.Payload)
}
func (o *DefectServiceCalculateDefectResponseDateOK) GetPayload() *models.ConquestAPICalculateDefectResponseDateResponse {
	return o.Payload
}

func (o *DefectServiceCalculateDefectResponseDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPICalculateDefectResponseDateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDefectServiceCalculateDefectResponseDateDefault creates a DefectServiceCalculateDefectResponseDateDefault with default headers values
func NewDefectServiceCalculateDefectResponseDateDefault(code int) *DefectServiceCalculateDefectResponseDateDefault {
	return &DefectServiceCalculateDefectResponseDateDefault{
		_statusCode: code,
	}
}

/* DefectServiceCalculateDefectResponseDateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DefectServiceCalculateDefectResponseDateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the defect service calculate defect response date default response
func (o *DefectServiceCalculateDefectResponseDateDefault) Code() int {
	return o._statusCode
}

func (o *DefectServiceCalculateDefectResponseDateDefault) Error() string {
	return fmt.Sprintf("[POST /api/defects/calculate_defect_response_date][%d] DefectService_CalculateDefectResponseDate default  %+v", o._statusCode, o.Payload)
}
func (o *DefectServiceCalculateDefectResponseDateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DefectServiceCalculateDefectResponseDateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}