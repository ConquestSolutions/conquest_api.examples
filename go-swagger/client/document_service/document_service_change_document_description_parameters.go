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

// NewDocumentServiceChangeDocumentDescriptionParams creates a new DocumentServiceChangeDocumentDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDocumentServiceChangeDocumentDescriptionParams() *DocumentServiceChangeDocumentDescriptionParams {
	return &DocumentServiceChangeDocumentDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentServiceChangeDocumentDescriptionParamsWithTimeout creates a new DocumentServiceChangeDocumentDescriptionParams object
// with the ability to set a timeout on a request.
func NewDocumentServiceChangeDocumentDescriptionParamsWithTimeout(timeout time.Duration) *DocumentServiceChangeDocumentDescriptionParams {
	return &DocumentServiceChangeDocumentDescriptionParams{
		timeout: timeout,
	}
}

// NewDocumentServiceChangeDocumentDescriptionParamsWithContext creates a new DocumentServiceChangeDocumentDescriptionParams object
// with the ability to set a context for a request.
func NewDocumentServiceChangeDocumentDescriptionParamsWithContext(ctx context.Context) *DocumentServiceChangeDocumentDescriptionParams {
	return &DocumentServiceChangeDocumentDescriptionParams{
		Context: ctx,
	}
}

// NewDocumentServiceChangeDocumentDescriptionParamsWithHTTPClient creates a new DocumentServiceChangeDocumentDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDocumentServiceChangeDocumentDescriptionParamsWithHTTPClient(client *http.Client) *DocumentServiceChangeDocumentDescriptionParams {
	return &DocumentServiceChangeDocumentDescriptionParams{
		HTTPClient: client,
	}
}

/* DocumentServiceChangeDocumentDescriptionParams contains all the parameters to send to the API endpoint
   for the document service change document description operation.

   Typically these are written to a http.Request.
*/
type DocumentServiceChangeDocumentDescriptionParams struct {

	// Body.
	Body *models.ConquestAPIChangeDocumentDescriptionCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the document service change document description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentServiceChangeDocumentDescriptionParams) WithDefaults() *DocumentServiceChangeDocumentDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the document service change document description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentServiceChangeDocumentDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) WithTimeout(timeout time.Duration) *DocumentServiceChangeDocumentDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) WithContext(ctx context.Context) *DocumentServiceChangeDocumentDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) WithHTTPClient(client *http.Client) *DocumentServiceChangeDocumentDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) WithBody(body *models.ConquestAPIChangeDocumentDescriptionCommand) *DocumentServiceChangeDocumentDescriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the document service change document description params
func (o *DocumentServiceChangeDocumentDescriptionParams) SetBody(body *models.ConquestAPIChangeDocumentDescriptionCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentServiceChangeDocumentDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
