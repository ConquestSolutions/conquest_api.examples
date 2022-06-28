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

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// NewDocumentServiceRemoveDocumentParams creates a new DocumentServiceRemoveDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDocumentServiceRemoveDocumentParams() *DocumentServiceRemoveDocumentParams {
	return &DocumentServiceRemoveDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentServiceRemoveDocumentParamsWithTimeout creates a new DocumentServiceRemoveDocumentParams object
// with the ability to set a timeout on a request.
func NewDocumentServiceRemoveDocumentParamsWithTimeout(timeout time.Duration) *DocumentServiceRemoveDocumentParams {
	return &DocumentServiceRemoveDocumentParams{
		timeout: timeout,
	}
}

// NewDocumentServiceRemoveDocumentParamsWithContext creates a new DocumentServiceRemoveDocumentParams object
// with the ability to set a context for a request.
func NewDocumentServiceRemoveDocumentParamsWithContext(ctx context.Context) *DocumentServiceRemoveDocumentParams {
	return &DocumentServiceRemoveDocumentParams{
		Context: ctx,
	}
}

// NewDocumentServiceRemoveDocumentParamsWithHTTPClient creates a new DocumentServiceRemoveDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDocumentServiceRemoveDocumentParamsWithHTTPClient(client *http.Client) *DocumentServiceRemoveDocumentParams {
	return &DocumentServiceRemoveDocumentParams{
		HTTPClient: client,
	}
}

/* DocumentServiceRemoveDocumentParams contains all the parameters to send to the API endpoint
   for the document service remove document operation.

   Typically these are written to a http.Request.
*/
type DocumentServiceRemoveDocumentParams struct {

	// Body.
	Body *models.ConquestAPIRemoveDocumentCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the document service remove document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentServiceRemoveDocumentParams) WithDefaults() *DocumentServiceRemoveDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the document service remove document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentServiceRemoveDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) WithTimeout(timeout time.Duration) *DocumentServiceRemoveDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) WithContext(ctx context.Context) *DocumentServiceRemoveDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) WithHTTPClient(client *http.Client) *DocumentServiceRemoveDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) WithBody(body *models.ConquestAPIRemoveDocumentCommand) *DocumentServiceRemoveDocumentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the document service remove document params
func (o *DocumentServiceRemoveDocumentParams) SetBody(body *models.ConquestAPIRemoveDocumentCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentServiceRemoveDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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