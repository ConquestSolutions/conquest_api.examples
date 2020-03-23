// Code generated by go-swagger; DO NOT EDIT.

package asset_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// ListInspectionsForAssetReader is a Reader for the ListInspectionsForAsset structure.
type ListInspectionsForAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInspectionsForAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInspectionsForAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListInspectionsForAssetOK creates a ListInspectionsForAssetOK with default headers values
func NewListInspectionsForAssetOK() *ListInspectionsForAssetOK {
	return &ListInspectionsForAssetOK{}
}

/*ListInspectionsForAssetOK handles this case with default header values.

A successful response.
*/
type ListInspectionsForAssetOK struct {
	Payload *models.ConquestAPIInspectionsForAssetResponse
}

func (o *ListInspectionsForAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/list_inspections][%d] listInspectionsForAssetOK  %+v", 200, o.Payload)
}

func (o *ListInspectionsForAssetOK) GetPayload() *models.ConquestAPIInspectionsForAssetResponse {
	return o.Payload
}

func (o *ListInspectionsForAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIInspectionsForAssetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
