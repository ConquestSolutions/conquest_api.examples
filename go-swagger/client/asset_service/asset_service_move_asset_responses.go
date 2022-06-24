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

// AssetServiceMoveAssetReader is a Reader for the AssetServiceMoveAsset structure.
type AssetServiceMoveAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetServiceMoveAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetServiceMoveAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetServiceMoveAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetServiceMoveAssetOK creates a AssetServiceMoveAssetOK with default headers values
func NewAssetServiceMoveAssetOK() *AssetServiceMoveAssetOK {
	return &AssetServiceMoveAssetOK{}
}

/* AssetServiceMoveAssetOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetServiceMoveAssetOK struct {
	Payload interface{}
}

func (o *AssetServiceMoveAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/move_asset][%d] assetServiceMoveAssetOK  %+v", 200, o.Payload)
}
func (o *AssetServiceMoveAssetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AssetServiceMoveAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetServiceMoveAssetDefault creates a AssetServiceMoveAssetDefault with default headers values
func NewAssetServiceMoveAssetDefault(code int) *AssetServiceMoveAssetDefault {
	return &AssetServiceMoveAssetDefault{
		_statusCode: code,
	}
}

/* AssetServiceMoveAssetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetServiceMoveAssetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset service move asset default response
func (o *AssetServiceMoveAssetDefault) Code() int {
	return o._statusCode
}

func (o *AssetServiceMoveAssetDefault) Error() string {
	return fmt.Sprintf("[POST /api/assets/move_asset][%d] AssetService_MoveAsset default  %+v", o._statusCode, o.Payload)
}
func (o *AssetServiceMoveAssetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetServiceMoveAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
