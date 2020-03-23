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

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// NewUpdateActionParams creates a new UpdateActionParams object
// with the default values initialized.
func NewUpdateActionParams() *UpdateActionParams {
	var ()
	return &UpdateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActionParamsWithTimeout creates a new UpdateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateActionParamsWithTimeout(timeout time.Duration) *UpdateActionParams {
	var ()
	return &UpdateActionParams{

		timeout: timeout,
	}
}

// NewUpdateActionParamsWithContext creates a new UpdateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateActionParamsWithContext(ctx context.Context) *UpdateActionParams {
	var ()
	return &UpdateActionParams{

		Context: ctx,
	}
}

// NewUpdateActionParamsWithHTTPClient creates a new UpdateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateActionParamsWithHTTPClient(client *http.Client) *UpdateActionParams {
	var ()
	return &UpdateActionParams{
		HTTPClient: client,
	}
}

/*UpdateActionParams contains all the parameters to send to the API endpoint
for the update action operation typically these are written to a http.Request
*/
type UpdateActionParams struct {

	/*Body*/
	Body *models.ConquestAPIActionRecordChangeSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update action params
func (o *UpdateActionParams) WithTimeout(timeout time.Duration) *UpdateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update action params
func (o *UpdateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update action params
func (o *UpdateActionParams) WithContext(ctx context.Context) *UpdateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update action params
func (o *UpdateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update action params
func (o *UpdateActionParams) WithHTTPClient(client *http.Client) *UpdateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update action params
func (o *UpdateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update action params
func (o *UpdateActionParams) WithBody(body *models.ConquestAPIActionRecordChangeSet) *UpdateActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update action params
func (o *UpdateActionParams) SetBody(body *models.ConquestAPIActionRecordChangeSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
