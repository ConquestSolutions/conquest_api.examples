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

// AssetServiceDeleteAssetReader is a Reader for the AssetServiceDeleteAsset structure.
type AssetServiceDeleteAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetServiceDeleteAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetServiceDeleteAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetServiceDeleteAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetServiceDeleteAssetOK creates a AssetServiceDeleteAssetOK with default headers values
func NewAssetServiceDeleteAssetOK() *AssetServiceDeleteAssetOK {
	return &AssetServiceDeleteAssetOK{}
}

/* AssetServiceDeleteAssetOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetServiceDeleteAssetOK struct {
	Payload interface{}
}

func (o *AssetServiceDeleteAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/delete_asset][%d] assetServiceDeleteAssetOK  %+v", 200, o.Payload)
}
func (o *AssetServiceDeleteAssetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AssetServiceDeleteAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetServiceDeleteAssetDefault creates a AssetServiceDeleteAssetDefault with default headers values
func NewAssetServiceDeleteAssetDefault(code int) *AssetServiceDeleteAssetDefault {
	return &AssetServiceDeleteAssetDefault{
		_statusCode: code,
	}
}

/* AssetServiceDeleteAssetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetServiceDeleteAssetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset service delete asset default response
func (o *AssetServiceDeleteAssetDefault) Code() int {
	return o._statusCode
}

func (o *AssetServiceDeleteAssetDefault) Error() string {
	return fmt.Sprintf("[POST /api/assets/delete_asset][%d] AssetService_DeleteAsset default  %+v", o._statusCode, o.Payload)
}
func (o *AssetServiceDeleteAssetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetServiceDeleteAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
