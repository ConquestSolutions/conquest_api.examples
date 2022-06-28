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

// NewAssetServiceChangeAssetTypeParams creates a new AssetServiceChangeAssetTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssetServiceChangeAssetTypeParams() *AssetServiceChangeAssetTypeParams {
	return &AssetServiceChangeAssetTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssetServiceChangeAssetTypeParamsWithTimeout creates a new AssetServiceChangeAssetTypeParams object
// with the ability to set a timeout on a request.
func NewAssetServiceChangeAssetTypeParamsWithTimeout(timeout time.Duration) *AssetServiceChangeAssetTypeParams {
	return &AssetServiceChangeAssetTypeParams{
		timeout: timeout,
	}
}

// NewAssetServiceChangeAssetTypeParamsWithContext creates a new AssetServiceChangeAssetTypeParams object
// with the ability to set a context for a request.
func NewAssetServiceChangeAssetTypeParamsWithContext(ctx context.Context) *AssetServiceChangeAssetTypeParams {
	return &AssetServiceChangeAssetTypeParams{
		Context: ctx,
	}
}

// NewAssetServiceChangeAssetTypeParamsWithHTTPClient creates a new AssetServiceChangeAssetTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssetServiceChangeAssetTypeParamsWithHTTPClient(client *http.Client) *AssetServiceChangeAssetTypeParams {
	return &AssetServiceChangeAssetTypeParams{
		HTTPClient: client,
	}
}

/* AssetServiceChangeAssetTypeParams contains all the parameters to send to the API endpoint
   for the asset service change asset type operation.

   Typically these are written to a http.Request.
*/
type AssetServiceChangeAssetTypeParams struct {

	// Body.
	Body *models.ConquestAPIChangeAssetTypeCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the asset service change asset type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceChangeAssetTypeParams) WithDefaults() *AssetServiceChangeAssetTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the asset service change asset type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssetServiceChangeAssetTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) WithTimeout(timeout time.Duration) *AssetServiceChangeAssetTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) WithContext(ctx context.Context) *AssetServiceChangeAssetTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) WithHTTPClient(client *http.Client) *AssetServiceChangeAssetTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) WithBody(body *models.ConquestAPIChangeAssetTypeCommand) *AssetServiceChangeAssetTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the asset service change asset type params
func (o *AssetServiceChangeAssetTypeParams) SetBody(body *models.ConquestAPIChangeAssetTypeCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssetServiceChangeAssetTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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