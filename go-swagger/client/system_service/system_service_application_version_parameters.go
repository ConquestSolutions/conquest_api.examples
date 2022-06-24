// Code generated by go-swagger; DO NOT EDIT.

package system_service

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
)

// NewSystemServiceApplicationVersionParams creates a new SystemServiceApplicationVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemServiceApplicationVersionParams() *SystemServiceApplicationVersionParams {
	return &SystemServiceApplicationVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemServiceApplicationVersionParamsWithTimeout creates a new SystemServiceApplicationVersionParams object
// with the ability to set a timeout on a request.
func NewSystemServiceApplicationVersionParamsWithTimeout(timeout time.Duration) *SystemServiceApplicationVersionParams {
	return &SystemServiceApplicationVersionParams{
		timeout: timeout,
	}
}

// NewSystemServiceApplicationVersionParamsWithContext creates a new SystemServiceApplicationVersionParams object
// with the ability to set a context for a request.
func NewSystemServiceApplicationVersionParamsWithContext(ctx context.Context) *SystemServiceApplicationVersionParams {
	return &SystemServiceApplicationVersionParams{
		Context: ctx,
	}
}

// NewSystemServiceApplicationVersionParamsWithHTTPClient creates a new SystemServiceApplicationVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemServiceApplicationVersionParamsWithHTTPClient(client *http.Client) *SystemServiceApplicationVersionParams {
	return &SystemServiceApplicationVersionParams{
		HTTPClient: client,
	}
}

/* SystemServiceApplicationVersionParams contains all the parameters to send to the API endpoint
   for the system service application version operation.

   Typically these are written to a http.Request.
*/
type SystemServiceApplicationVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system service application version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemServiceApplicationVersionParams) WithDefaults() *SystemServiceApplicationVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system service application version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemServiceApplicationVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system service application version params
func (o *SystemServiceApplicationVersionParams) WithTimeout(timeout time.Duration) *SystemServiceApplicationVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system service application version params
func (o *SystemServiceApplicationVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system service application version params
func (o *SystemServiceApplicationVersionParams) WithContext(ctx context.Context) *SystemServiceApplicationVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system service application version params
func (o *SystemServiceApplicationVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system service application version params
func (o *SystemServiceApplicationVersionParams) WithHTTPClient(client *http.Client) *SystemServiceApplicationVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system service application version params
func (o *SystemServiceApplicationVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemServiceApplicationVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
