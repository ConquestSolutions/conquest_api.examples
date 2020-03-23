// Code generated by go-swagger; DO NOT EDIT.

package asset_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAssetParams creates a new GetAssetParams object
// with the default values initialized.
func NewGetAssetParams() *GetAssetParams {
	var ()
	return &GetAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssetParamsWithTimeout creates a new GetAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssetParamsWithTimeout(timeout time.Duration) *GetAssetParams {
	var ()
	return &GetAssetParams{

		timeout: timeout,
	}
}

// NewGetAssetParamsWithContext creates a new GetAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssetParamsWithContext(ctx context.Context) *GetAssetParams {
	var ()
	return &GetAssetParams{

		Context: ctx,
	}
}

// NewGetAssetParamsWithHTTPClient creates a new GetAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssetParamsWithHTTPClient(client *http.Client) *GetAssetParams {
	var ()
	return &GetAssetParams{
		HTTPClient: client,
	}
}

/*GetAssetParams contains all the parameters to send to the API endpoint
for the get asset operation typically these are written to a http.Request
*/
type GetAssetParams struct {

	/*Value
	  The int32 value.

	*/
	Value int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get asset params
func (o *GetAssetParams) WithTimeout(timeout time.Duration) *GetAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get asset params
func (o *GetAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get asset params
func (o *GetAssetParams) WithContext(ctx context.Context) *GetAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get asset params
func (o *GetAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get asset params
func (o *GetAssetParams) WithHTTPClient(client *http.Client) *GetAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get asset params
func (o *GetAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValue adds the value to the get asset params
func (o *GetAssetParams) WithValue(value int32) *GetAssetParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get asset params
func (o *GetAssetParams) SetValue(value int32) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param value
	if err := r.SetPathParam("value", swag.FormatInt32(o.Value)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
