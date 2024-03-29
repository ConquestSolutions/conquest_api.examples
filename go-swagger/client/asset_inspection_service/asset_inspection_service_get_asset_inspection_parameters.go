// Code generated by go-swagger; DO NOT EDIT.

package asset_inspection_service

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

// NewAssetInspectionServiceGetAssetInspectionParams creates a new AssetInspectionServiceGetAssetInspectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssetInspectionServiceGetAssetInspectionParams() *AssetInspectionServiceGetAssetInspectionParams {
	return &AssetInspectionServiceGetAssetInspectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssetInspectionServiceGetAssetInspectionParamsWithTimeout creates a new AssetInspectionServiceGetAssetInspectionParams object
// with the ability to set a timeout on a request.
func NewAssetInspectionServiceGetAssetInspectionParamsWithTimeout(timeout time.Duration) *AssetInspectionServiceGetAssetInspectionParams {
	return &AssetInspectionServiceGetAssetInspectionParams{
		timeout: timeout,
	}
}

// NewAssetInspectionServiceGetAssetInspectionParamsWithContext creates a new AssetInspectionServiceGetAssetInspectionParams object
// with the ability to set a context for a request.
func NewAssetInspectionServiceGetAssetInspectionParamsWithContext(ctx context.Context) *AssetInspectionServiceGetAssetInspectionParams {
	return &AssetInspectionServiceGetAssetInspectionParams{
		Context: ctx,
	}
}

// NewAssetInspectionServiceGetAssetInspectionParamsWithHTTPClient creates a new AssetInspectionServiceGetAssetInspectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssetInspectionServiceGetAssetInspectionParamsWithHTTPClient(client *http.Client) *AssetInspectionServiceGetAssetInspectionParams {
	return &AssetInspectionServiceGetAssetInspectionParams{
		HTTPClient: client,
	}
}

/* AssetInspectionServiceGetAssetInspectionParams contains all the parameters to send to the API endpoint
   for the asset inspection service get asset inspection operation.

   Typically these are written to a http.Request.
*/
type AssetInspectionServiceGetAssetInspectionParams struct {

	/* Value.

	   The int32 value.

	   Format: int32
	*/
	Value int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the asset inspection service get asset inspection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetInspectionServiceGetAssetInspectionParams) WithDefaults() *AssetInspectionServiceGetAssetInspectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asset inspection service get asset inspection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetInspectionServiceGetAssetInspectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) WithTimeout(timeout time.Duration) *AssetInspectionServiceGetAssetInspectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) WithContext(ctx context.Context) *AssetInspectionServiceGetAssetInspectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) WithHTTPClient(client *http.Client) *AssetInspectionServiceGetAssetInspectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValue adds the value to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) WithValue(value int32) *AssetInspectionServiceGetAssetInspectionParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the asset inspection service get asset inspection params
func (o *AssetInspectionServiceGetAssetInspectionParams) SetValue(value int32) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *AssetInspectionServiceGetAssetInspectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
