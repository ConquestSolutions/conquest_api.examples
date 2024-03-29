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

// NewActionServiceDeleteActionParams creates a new ActionServiceDeleteActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionServiceDeleteActionParams() *ActionServiceDeleteActionParams {
	return &ActionServiceDeleteActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionServiceDeleteActionParamsWithTimeout creates a new ActionServiceDeleteActionParams object
// with the ability to set a timeout on a request.
func NewActionServiceDeleteActionParamsWithTimeout(timeout time.Duration) *ActionServiceDeleteActionParams {
	return &ActionServiceDeleteActionParams{
		timeout: timeout,
	}
}

// NewActionServiceDeleteActionParamsWithContext creates a new ActionServiceDeleteActionParams object
// with the ability to set a context for a request.
func NewActionServiceDeleteActionParamsWithContext(ctx context.Context) *ActionServiceDeleteActionParams {
	return &ActionServiceDeleteActionParams{
		Context: ctx,
	}
}

// NewActionServiceDeleteActionParamsWithHTTPClient creates a new ActionServiceDeleteActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionServiceDeleteActionParamsWithHTTPClient(client *http.Client) *ActionServiceDeleteActionParams {
	return &ActionServiceDeleteActionParams{
		HTTPClient: client,
	}
}

/* ActionServiceDeleteActionParams contains all the parameters to send to the API endpoint
   for the action service delete action operation.

   Typically these are written to a http.Request.
*/
type ActionServiceDeleteActionParams struct {

	// Body.
	Body *models.ConquestAPIDeleteActionCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action service delete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceDeleteActionParams) WithDefaults() *ActionServiceDeleteActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action service delete action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionServiceDeleteActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action service delete action params
func (o *ActionServiceDeleteActionParams) WithTimeout(timeout time.Duration) *ActionServiceDeleteActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action service delete action params
func (o *ActionServiceDeleteActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action service delete action params
func (o *ActionServiceDeleteActionParams) WithContext(ctx context.Context) *ActionServiceDeleteActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action service delete action params
func (o *ActionServiceDeleteActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action service delete action params
func (o *ActionServiceDeleteActionParams) WithHTTPClient(client *http.Client) *ActionServiceDeleteActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action service delete action params
func (o *ActionServiceDeleteActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the action service delete action params
func (o *ActionServiceDeleteActionParams) WithBody(body *models.ConquestAPIDeleteActionCommand) *ActionServiceDeleteActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the action service delete action params
func (o *ActionServiceDeleteActionParams) SetBody(body *models.ConquestAPIDeleteActionCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActionServiceDeleteActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
