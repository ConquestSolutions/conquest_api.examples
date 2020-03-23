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

// UpdateAssetReader is a Reader for the UpdateAsset structure.
type UpdateAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAssetOK creates a UpdateAssetOK with default headers values
func NewUpdateAssetOK() *UpdateAssetOK {
	return &UpdateAssetOK{}
}

/*UpdateAssetOK handles this case with default header values.

A successful response.
*/
type UpdateAssetOK struct {
	Payload *models.ConquestAPIChangeSetResult
}

func (o *UpdateAssetOK) Error() string {
	return fmt.Sprintf("[POST /api/assets/update_asset][%d] updateAssetOK  %+v", 200, o.Payload)
}

func (o *UpdateAssetOK) GetPayload() *models.ConquestAPIChangeSetResult {
	return o.Payload
}

func (o *UpdateAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIChangeSetResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
