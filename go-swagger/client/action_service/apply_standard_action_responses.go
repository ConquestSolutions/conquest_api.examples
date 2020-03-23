// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ApplyStandardActionReader is a Reader for the ApplyStandardAction structure.
type ApplyStandardActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyStandardActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyStandardActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplyStandardActionOK creates a ApplyStandardActionOK with default headers values
func NewApplyStandardActionOK() *ApplyStandardActionOK {
	return &ApplyStandardActionOK{}
}

/*ApplyStandardActionOK handles this case with default header values.

A successful response.
*/
type ApplyStandardActionOK struct {
	Payload interface{}
}

func (o *ApplyStandardActionOK) Error() string {
	return fmt.Sprintf("[POST /api/actions/apply_standard_action][%d] applyStandardActionOK  %+v", 200, o.Payload)
}

func (o *ApplyStandardActionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplyStandardActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
