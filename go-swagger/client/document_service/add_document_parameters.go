// Code generated by go-swagger; DO NOT EDIT.

package document_service

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

// NewAddDocumentParams creates a new AddDocumentParams object
// with the default values initialized.
func NewAddDocumentParams() *AddDocumentParams {
	var ()
	return &AddDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDocumentParamsWithTimeout creates a new AddDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDocumentParamsWithTimeout(timeout time.Duration) *AddDocumentParams {
	var ()
	return &AddDocumentParams{

		timeout: timeout,
	}
}

// NewAddDocumentParamsWithContext creates a new AddDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDocumentParamsWithContext(ctx context.Context) *AddDocumentParams {
	var ()
	return &AddDocumentParams{

		Context: ctx,
	}
}

// NewAddDocumentParamsWithHTTPClient creates a new AddDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDocumentParamsWithHTTPClient(client *http.Client) *AddDocumentParams {
	var ()
	return &AddDocumentParams{
		HTTPClient: client,
	}
}

/*AddDocumentParams contains all the parameters to send to the API endpoint
for the add document operation typically these are written to a http.Request
*/
type AddDocumentParams struct {

	/*Body*/
	Body *models.ConquestAPIAddDocumentCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add document params
func (o *AddDocumentParams) WithTimeout(timeout time.Duration) *AddDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add document params
func (o *AddDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add document params
func (o *AddDocumentParams) WithContext(ctx context.Context) *AddDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add document params
func (o *AddDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add document params
func (o *AddDocumentParams) WithHTTPClient(client *http.Client) *AddDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add document params
func (o *AddDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add document params
func (o *AddDocumentParams) WithBody(body *models.ConquestAPIAddDocumentCommand) *AddDocumentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add document params
func (o *AddDocumentParams) SetBody(body *models.ConquestAPIAddDocumentCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
