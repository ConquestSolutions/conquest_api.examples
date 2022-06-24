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

// AssetServiceUpdateAssetReader is a Reader for the AssetServiceUpdateAsset structure.
type AssetServiceUpdateAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetServiceUpdateAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetServiceUpdateAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetServiceUpdateAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetServiceUpdateAssetOK creates a AssetServiceUpdateAssetOK with default headers values
func NewAssetServiceUpdateAssetOK() *AssetServiceUpdateAssetOK {
	return &AssetServiceUpdateAssetOK{}
}

/* AssetServiceUpdateAssetOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetServiceUpdateAssetOK struct {
	Payload *models.ConquestAPIChangeSetResult
}

func (o *AssetServiceUpdateAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/update_asset][%d] assetServiceUpdateAssetOK  %+v", 200, o.Payload)
}
func (o *AssetServiceUpdateAssetOK) GetPayload() *models.ConquestAPIChangeSetResult {
	return o.Payload
}

func (o *AssetServiceUpdateAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIChangeSetResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetServiceUpdateAssetDefault creates a AssetServiceUpdateAssetDefault with default headers values
func NewAssetServiceUpdateAssetDefault(code int) *AssetServiceUpdateAssetDefault {
	return &AssetServiceUpdateAssetDefault{
		_statusCode: code,
	}
}

/* AssetServiceUpdateAssetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetServiceUpdateAssetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset service update asset default response
func (o *AssetServiceUpdateAssetDefault) Code() int {
	return o._statusCode
}

func (o *AssetServiceUpdateAssetDefault) Error() string {
	return fmt.Sprintf("[POST /api/assets/update_asset][%d] AssetService_UpdateAsset default  %+v", o._statusCode, o.Payload)
}
func (o *AssetServiceUpdateAssetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetServiceUpdateAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
