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

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// NewActionServiceUpdateActionParams creates a new ActionServiceUpdateActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionServiceUpdateActionParams() *ActionServiceUpdateActionParams {
	return &ActionServiceUpdateActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionServiceUpdateActionParamsWithTimeout creates a new ActionServiceUpdateActionParams object
// with the ability to set a timeout on a request.
func NewActionServiceUpdateActionParamsWithTimeout(timeout time.Duration) *ActionServiceUpdateActionParams {
	return &ActionServiceUpdateActionParams{
		timeout: timeout,
	}
}

// NewActionServiceUpdateActionParamsWithContext creates a new ActionServiceUpdateActionParams object
// with the ability to set a context for a request.
func NewActionServiceUpdateActionParamsWithContext(ctx context.Context) *ActionServiceUpdateActionParams {
	return &ActionServiceUpdateActionParams{
		Context: ctx,
	}
}

// NewActionServiceUpdateActionParamsWithHTTPClient creates a new ActionServiceUpdateActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionServiceUpdateActionParamsWithHTTPClient(client *http.Client) *ActionServiceUpdateActionParams {
	return &ActionServiceUpdateActionParams{
		HTTPClient: client,
	}
}

/* ActionServiceUpdateActionParams contains all the parameters to send to the API endpoint
   for the action service update action operation.

   Typically these are written to a http.Request.
*/
type ActionServiceUpdateActionParams struct {

	// Body.
	Body *models.ConquestAPIActionRecordChangeSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action service update action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceUpdateActionParams) WithDefaults() *ActionServiceUpdateActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action service update action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceUpdateActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action service update action params
func (o *ActionServiceUpdateActionParams) WithTimeout(timeout time.Duration) *ActionServiceUpdateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action service update action params
func (o *ActionServiceUpdateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action service update action params
func (o *ActionServiceUpdateActionParams) WithContext(ctx context.Context) *ActionServiceUpdateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action service update action params
func (o *ActionServiceUpdateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action service update action params
func (o *ActionServiceUpdateActionParams) WithHTTPClient(client *http.Client) *ActionServiceUpdateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action service update action params
func (o *ActionServiceUpdateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the action service update action params
func (o *ActionServiceUpdateActionParams) WithBody(body *models.ConquestAPIActionRecordChangeSet) *ActionServiceUpdateActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the action service update action params
func (o *ActionServiceUpdateActionParams) SetBody(body *models.ConquestAPIActionRecordChangeSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActionServiceUpdateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
