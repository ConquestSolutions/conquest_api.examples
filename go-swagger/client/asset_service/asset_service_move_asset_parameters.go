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

// NewAssetServiceMoveAssetParams creates a new AssetServiceMoveAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssetServiceMoveAssetParams() *AssetServiceMoveAssetParams {
	return &AssetServiceMoveAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssetServiceMoveAssetParamsWithTimeout creates a new AssetServiceMoveAssetParams object
// with the ability to set a timeout on a request.
func NewAssetServiceMoveAssetParamsWithTimeout(timeout time.Duration) *AssetServiceMoveAssetParams {
	return &AssetServiceMoveAssetParams{
		timeout: timeout,
	}
}

// NewAssetServiceMoveAssetParamsWithContext creates a new AssetServiceMoveAssetParams object
// with the ability to set a context for a request.
func NewAssetServiceMoveAssetParamsWithContext(ctx context.Context) *AssetServiceMoveAssetParams {
	return &AssetServiceMoveAssetParams{
		Context: ctx,
	}
}

// NewAssetServiceMoveAssetParamsWithHTTPClient creates a new AssetServiceMoveAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssetServiceMoveAssetParamsWithHTTPClient(client *http.Client) *AssetServiceMoveAssetParams {
	return &AssetServiceMoveAssetParams{
		HTTPClient: client,
	}
}

/* AssetServiceMoveAssetParams contains all the parameters to send to the API endpoint
   for the asset service move asset operation.

   Typically these are written to a http.Request.
*/
type AssetServiceMoveAssetParams struct {

	// Body.
	Body *models.ConquestAPIMoveAssetCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the asset service move asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceMoveAssetParams) WithDefaults() *AssetServiceMoveAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asset service move asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceMoveAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asset service move asset params
func (o *AssetServiceMoveAssetParams) WithTimeout(timeout time.Duration) *AssetServiceMoveAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asset service move asset params
func (o *AssetServiceMoveAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asset service move asset params
func (o *AssetServiceMoveAssetParams) WithContext(ctx context.Context) *AssetServiceMoveAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asset service move asset params
func (o *AssetServiceMoveAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asset service move asset params
func (o *AssetServiceMoveAssetParams) WithHTTPClient(client *http.Client) *AssetServiceMoveAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asset service move asset params
func (o *AssetServiceMoveAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the asset service move asset params
func (o *AssetServiceMoveAssetParams) WithBody(body *models.ConquestAPIMoveAssetCommand) *AssetServiceMoveAssetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the asset service move asset params
func (o *AssetServiceMoveAssetParams) SetBody(body *models.ConquestAPIMoveAssetCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssetServiceMoveAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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