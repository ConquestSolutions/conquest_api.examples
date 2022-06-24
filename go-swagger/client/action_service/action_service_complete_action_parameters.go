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

// NewActionServiceCompleteActionParams creates a new ActionServiceCompleteActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionServiceCompleteActionParams() *ActionServiceCompleteActionParams {
	return &ActionServiceCompleteActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionServiceCompleteActionParamsWithTimeout creates a new ActionServiceCompleteActionParams object
// with the ability to set a timeout on a request.
func NewActionServiceCompleteActionParamsWithTimeout(timeout time.Duration) *ActionServiceCompleteActionParams {
	return &ActionServiceCompleteActionParams{
		timeout: timeout,
	}
}

// NewActionServiceCompleteActionParamsWithContext creates a new ActionServiceCompleteActionParams object
// with the ability to set a context for a request.
func NewActionServiceCompleteActionParamsWithContext(ctx context.Context) *ActionServiceCompleteActionParams {
	return &ActionServiceCompleteActionParams{
		Context: ctx,
	}
}

// NewActionServiceCompleteActionParamsWithHTTPClient creates a new ActionServiceCompleteActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionServiceCompleteActionParamsWithHTTPClient(client *http.Client) *ActionServiceCompleteActionParams {
	return &ActionServiceCompleteActionParams{
		HTTPClient: client,
	}
}

/* ActionServiceCompleteActionParams contains all the parameters to send to the API endpoint
   for the action service complete action operation.

   Typically these are written to a http.Request.
*/
type ActionServiceCompleteActionParams struct {

	// Body.
	Body *models.ConquestAPICompleteActionCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action service complete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceCompleteActionParams) WithDefaults() *ActionServiceCompleteActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action service complete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceCompleteActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action service complete action params
func (o *ActionServiceCompleteActionParams) WithTimeout(timeout time.Duration) *ActionServiceCompleteActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action service complete action params
func (o *ActionServiceCompleteActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action service complete action params
func (o *ActionServiceCompleteActionParams) WithContext(ctx context.Context) *ActionServiceCompleteActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action service complete action params
func (o *ActionServiceCompleteActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action service complete action params
func (o *ActionServiceCompleteActionParams) WithHTTPClient(client *http.Client) *ActionServiceCompleteActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action service complete action params
func (o *ActionServiceCompleteActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the action service complete action params
func (o *ActionServiceCompleteActionParams) WithBody(body *models.ConquestAPICompleteActionCommand) *ActionServiceCompleteActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the action service complete action params
func (o *ActionServiceCompleteActionParams) SetBody(body *models.ConquestAPICompleteActionCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActionServiceCompleteActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
