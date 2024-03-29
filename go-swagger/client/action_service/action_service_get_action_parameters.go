// Code generated by go-swagger; DO NOT EDIT.

package action_service

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

// NewActionServiceGetActionParams creates a new ActionServiceGetActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionServiceGetActionParams() *ActionServiceGetActionParams {
	return &ActionServiceGetActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionServiceGetActionParamsWithTimeout creates a new ActionServiceGetActionParams object
// with the ability to set a timeout on a request.
func NewActionServiceGetActionParamsWithTimeout(timeout time.Duration) *ActionServiceGetActionParams {
	return &ActionServiceGetActionParams{
		timeout: timeout,
	}
}

// NewActionServiceGetActionParamsWithContext creates a new ActionServiceGetActionParams object
// with the ability to set a context for a request.
func NewActionServiceGetActionParamsWithContext(ctx context.Context) *ActionServiceGetActionParams {
	return &ActionServiceGetActionParams{
		Context: ctx,
	}
}

// NewActionServiceGetActionParamsWithHTTPClient creates a new ActionServiceGetActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionServiceGetActionParamsWithHTTPClient(client *http.Client) *ActionServiceGetActionParams {
	return &ActionServiceGetActionParams{
		HTTPClient: client,
	}
}

/* ActionServiceGetActionParams contains all the parameters to send to the API endpoint
   for the action service get action operation.

   Typically these are written to a http.Request.
*/
type ActionServiceGetActionParams struct {

	/* Value.

	   The int32 value.

	   Format: int32
	*/
	Value int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action service get action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceGetActionParams) WithDefaults() *ActionServiceGetActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action service get action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceGetActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action service get action params
func (o *ActionServiceGetActionParams) WithTimeout(timeout time.Duration) *ActionServiceGetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action service get action params
func (o *ActionServiceGetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action service get action params
func (o *ActionServiceGetActionParams) WithContext(ctx context.Context) *ActionServiceGetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action service get action params
func (o *ActionServiceGetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action service get action params
func (o *ActionServiceGetActionParams) WithHTTPClient(client *http.Client) *ActionServiceGetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action service get action params
func (o *ActionServiceGetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValue adds the value to the action service get action params
func (o *ActionServiceGetActionParams) WithValue(value int32) *ActionServiceGetActionParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the action service get action params
func (o *ActionServiceGetActionParams) SetValue(value int32) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *ActionServiceGetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
