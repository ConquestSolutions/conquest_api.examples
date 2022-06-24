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

// AssetServiceCreateDefectForAssetReader is a Reader for the AssetServiceCreateDefectForAsset structure.
type AssetServiceCreateDefectForAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetServiceCreateDefectForAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssetServiceCreateDefectForAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssetServiceCreateDefectForAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssetServiceCreateDefectForAssetOK creates a AssetServiceCreateDefectForAssetOK with default headers values
func NewAssetServiceCreateDefectForAssetOK() *AssetServiceCreateDefectForAssetOK {
	return &AssetServiceCreateDefectForAssetOK{}
}

/* AssetServiceCreateDefectForAssetOK describes a response with status code 200, with default header values.

A successful response.
*/
type AssetServiceCreateDefectForAssetOK struct {
	Payload int32
}

func (o *AssetServiceCreateDefectForAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/create_defect][%d] assetServiceCreateDefectForAssetOK  %+v", 200, o.Payload)
}
func (o *AssetServiceCreateDefectForAssetOK) GetPayload() int32 {
	return o.Payload
}

func (o *AssetServiceCreateDefectForAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssetServiceCreateDefectForAssetDefault creates a AssetServiceCreateDefectForAssetDefault with default headers values
func NewAssetServiceCreateDefectForAssetDefault(code int) *AssetServiceCreateDefectForAssetDefault {
	return &AssetServiceCreateDefectForAssetDefault{
		_statusCode: code,
	}
}

/* AssetServiceCreateDefectForAssetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AssetServiceCreateDefectForAssetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the asset service create defect for asset default response
func (o *AssetServiceCreateDefectForAssetDefault) Code() int {
	return o._statusCode
}

func (o *AssetServiceCreateDefectForAssetDefault) Error() string {
	return fmt.Sprintf("[POST /api/assets/create_defect][%d] AssetService_CreateDefectForAsset default  %+v", o._statusCode, o.Payload)
}
func (o *AssetServiceCreateDefectForAssetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AssetServiceCreateDefectForAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
