// Code generated by go-swagger; DO NOT EDIT.

package asset_inspection_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// AssetInspectionServiceGetAssetInspectionReader is a Reader for the AssetInspectionServiceGetAssetInspection structure.
type AssetInspectionServiceGetAssetInspectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetInspectionServiceGetAssetInspectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetInspectionServiceGetAssetInspectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetInspectionServiceGetAssetInspectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetInspectionServiceGetAssetInspectionOK creates a AssetInspectionServiceGetAssetInspectionOK with default headers values
func NewAssetInspectionServiceGetAssetInspectionOK() *AssetInspectionServiceGetAssetInspectionOK {
	return &AssetInspectionServiceGetAssetInspectionOK{}
}

/* AssetInspectionServiceGetAssetInspectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetInspectionServiceGetAssetInspectionOK struct {
	Payload *models.ConquestAPIAssetInspectionEntity
}

func (o *AssetInspectionServiceGetAssetInspectionOK) Error() string {
	return fmt.Sprintf("[GET /api/asset_inspections/{value}][%d] assetInspectionServiceGetAssetInspectionOK  %+v", 200, o.Payload)
}
func (o *AssetInspectionServiceGetAssetInspectionOK) GetPayload() *models.ConquestAPIAssetInspectionEntity {
	return o.Payload
}

func (o *AssetInspectionServiceGetAssetInspectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIAssetInspectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetInspectionServiceGetAssetInspectionDefault creates a AssetInspectionServiceGetAssetInspectionDefault with default headers values
func NewAssetInspectionServiceGetAssetInspectionDefault(code int) *AssetInspectionServiceGetAssetInspectionDefault {
	return &AssetInspectionServiceGetAssetInspectionDefault{
		_statusCode: code,
	}
}

/* AssetInspectionServiceGetAssetInspectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetInspectionServiceGetAssetInspectionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset inspection service get asset inspection default response
func (o *AssetInspectionServiceGetAssetInspectionDefault) Code() int {
	return o._statusCode
}

func (o *AssetInspectionServiceGetAssetInspectionDefault) Error() string {
	return fmt.Sprintf("[GET /api/asset_inspections/{value}][%d] AssetInspectionService_GetAssetInspection default  %+v", o._statusCode, o.Payload)
}
func (o *AssetInspectionServiceGetAssetInspectionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetInspectionServiceGetAssetInspectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
