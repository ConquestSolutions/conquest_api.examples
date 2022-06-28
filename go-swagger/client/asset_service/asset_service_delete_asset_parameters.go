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

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// NewAssetServiceDeleteAssetParams creates a new AssetServiceDeleteAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssetServiceDeleteAssetParams() *AssetServiceDeleteAssetParams {
	return &AssetServiceDeleteAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssetServiceDeleteAssetParamsWithTimeout creates a new AssetServiceDeleteAssetParams object
// with the ability to set a timeout on a request.
func NewAssetServiceDeleteAssetParamsWithTimeout(timeout time.Duration) *AssetServiceDeleteAssetParams {
	return &AssetServiceDeleteAssetParams{
		timeout: timeout,
	}
}

// NewAssetServiceDeleteAssetParamsWithContext creates a new AssetServiceDeleteAssetParams object
// with the ability to set a context for a request.
func NewAssetServiceDeleteAssetParamsWithContext(ctx context.Context) *AssetServiceDeleteAssetParams {
	return &AssetServiceDeleteAssetParams{
		Context: ctx,
	}
}

// NewAssetServiceDeleteAssetParamsWithHTTPClient creates a new AssetServiceDeleteAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssetServiceDeleteAssetParamsWithHTTPClient(client *http.Client) *AssetServiceDeleteAssetParams {
	return &AssetServiceDeleteAssetParams{
		HTTPClient: client,
	}
}

/* AssetServiceDeleteAssetParams contains all the parameters to send to the API endpoint
   for the asset service delete asset operation.

   Typically these are written to a http.Request.
*/
type AssetServiceDeleteAssetParams struct {

	// Body.
	Body *models.ConquestAPIDeleteAssetCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the asset service delete asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceDeleteAssetParams) WithDefaults() *AssetServiceDeleteAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asset service delete asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceDeleteAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) WithTimeout(timeout time.Duration) *AssetServiceDeleteAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) WithContext(ctx context.Context) *AssetServiceDeleteAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) WithHTTPClient(client *http.Client) *AssetServiceDeleteAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) WithBody(body *models.ConquestAPIDeleteAssetCommand) *AssetServiceDeleteAssetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the asset service delete asset params
func (o *AssetServiceDeleteAssetParams) SetBody(body *models.ConquestAPIDeleteAssetCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssetServiceDeleteAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}