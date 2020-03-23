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

// NewGetActionCompletionDetailsParams creates a new GetActionCompletionDetailsParams object
// with the default values initialized.
func NewGetActionCompletionDetailsParams() *GetActionCompletionDetailsParams {
	var ()
	return &GetActionCompletionDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionCompletionDetailsParamsWithTimeout creates a new GetActionCompletionDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionCompletionDetailsParamsWithTimeout(timeout time.Duration) *GetActionCompletionDetailsParams {
	var ()
	return &GetActionCompletionDetailsParams{

		timeout: timeout,
	}
}

// NewGetActionCompletionDetailsParamsWithContext creates a new GetActionCompletionDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionCompletionDetailsParamsWithContext(ctx context.Context) *GetActionCompletionDetailsParams {
	var ()
	return &GetActionCompletionDetailsParams{

		Context: ctx,
	}
}

// NewGetActionCompletionDetailsParamsWithHTTPClient creates a new GetActionCompletionDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionCompletionDetailsParamsWithHTTPClient(client *http.Client) *GetActionCompletionDetailsParams {
	var ()
	return &GetActionCompletionDetailsParams{
		HTTPClient: client,
	}
}

/*GetActionCompletionDetailsParams contains all the parameters to send to the API endpoint
for the get action completion details operation typically these are written to a http.Request
*/
type GetActionCompletionDetailsParams struct {

	/*Body*/
	Body *models.ConquestAPIGetActionCompletionDetailsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action completion details params
func (o *GetActionCompletionDetailsParams) WithTimeout(timeout time.Duration) *GetActionCompletionDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action completion details params
func (o *GetActionCompletionDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action completion details params
func (o *GetActionCompletionDetailsParams) WithContext(ctx context.Context) *GetActionCompletionDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action completion details params
func (o *GetActionCompletionDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action completion details params
func (o *GetActionCompletionDetailsParams) WithHTTPClient(client *http.Client) *GetActionCompletionDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action completion details params
func (o *GetActionCompletionDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get action completion details params
func (o *GetActionCompletionDetailsParams) WithBody(body *models.ConquestAPIGetActionCompletionDetailsRequest) *GetActionCompletionDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get action completion details params
func (o *GetActionCompletionDetailsParams) SetBody(body *models.ConquestAPIGetActionCompletionDetailsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionCompletionDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
