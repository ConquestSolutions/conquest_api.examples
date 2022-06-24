// Code generated by go-swagger; DO NOT EDIT.

package asset_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// AssetServiceTagAsInspectedReader is a Reader for the AssetServiceTagAsInspected structure.
type AssetServiceTagAsInspectedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetServiceTagAsInspectedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetServiceTagAsInspectedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetServiceTagAsInspectedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetServiceTagAsInspectedOK creates a AssetServiceTagAsInspectedOK with default headers values
func NewAssetServiceTagAsInspectedOK() *AssetServiceTagAsInspectedOK {
	return &AssetServiceTagAsInspectedOK{}
}

/* AssetServiceTagAsInspectedOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetServiceTagAsInspectedOK struct {
	Payload int32
}

func (o *AssetServiceTagAsInspectedOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/tag_as_inspected][%d] assetServiceTagAsInspectedOK  %+v", 200, o.Payload)
}
func (o *AssetServiceTagAsInspectedOK) GetPayload() int32 {
	return o.Payload
}

func (o *AssetServiceTagAsInspectedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetServiceTagAsInspectedDefault creates a AssetServiceTagAsInspectedDefault with default headers values
func NewAssetServiceTagAsInspectedDefault(code int) *AssetServiceTagAsInspectedDefault {
	return &AssetServiceTagAsInspectedDefault{
		_statusCode: code,
	}
}

/* AssetServiceTagAsInspectedDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetServiceTagAsInspectedDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset service tag as inspected default response
func (o *AssetServiceTagAsInspectedDefault) Code() int {
	return o._statusCode
}

func (o *AssetServiceTagAsInspectedDefault) Error() string {
	return fmt.Sprintf("[POST /api/assets/tag_as_inspected][%d] AssetService_TagAsInspected default  %+v", o._statusCode, o.Payload)
}
func (o *AssetServiceTagAsInspectedDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetServiceTagAsInspectedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
