// Code generated by go-swagger; DO NOT EDIT.

package knowledge_base_service

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

// NewKnowledgeBaseServiceDescribeEnumerationParams creates a new KnowledgeBaseServiceDescribeEnumerationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKnowledgeBaseServiceDescribeEnumerationParams() *KnowledgeBaseServiceDescribeEnumerationParams {
	return &KnowledgeBaseServiceDescribeEnumerationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKnowledgeBaseServiceDescribeEnumerationParamsWithTimeout creates a new KnowledgeBaseServiceDescribeEnumerationParams object
// with the ability to set a timeout on a request.
func NewKnowledgeBaseServiceDescribeEnumerationParamsWithTimeout(timeout time.Duration) *KnowledgeBaseServiceDescribeEnumerationParams {
	return &KnowledgeBaseServiceDescribeEnumerationParams{
		timeout: timeout,
	}
}

// NewKnowledgeBaseServiceDescribeEnumerationParamsWithContext creates a new KnowledgeBaseServiceDescribeEnumerationParams object
// with the ability to set a context for a request.
func NewKnowledgeBaseServiceDescribeEnumerationParamsWithContext(ctx context.Context) *KnowledgeBaseServiceDescribeEnumerationParams {
	return &KnowledgeBaseServiceDescribeEnumerationParams{
		Context: ctx,
	}
}

// NewKnowledgeBaseServiceDescribeEnumerationParamsWithHTTPClient creates a new KnowledgeBaseServiceDescribeEnumerationParams object
// with the ability to set a custom HTTPClient for a request.
func NewKnowledgeBaseServiceDescribeEnumerationParamsWithHTTPClient(client *http.Client) *KnowledgeBaseServiceDescribeEnumerationParams {
	return &KnowledgeBaseServiceDescribeEnumerationParams{
		HTTPClient: client,
	}
}

/* KnowledgeBaseServiceDescribeEnumerationParams contains all the parameters to send to the API endpoint
   for the knowledge base service describe enumeration operation.

   Typically these are written to a http.Request.
*/
type KnowledgeBaseServiceDescribeEnumerationParams struct {

	// Body.
	Body *models.ConquestAPIDescribeEnumerationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the knowledge base service describe enumeration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WithDefaults() *KnowledgeBaseServiceDescribeEnumerationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the knowledge base service describe enumeration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KnowledgeBaseServiceDescribeEnumerationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WithTimeout(timeout time.Duration) *KnowledgeBaseServiceDescribeEnumerationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WithContext(ctx context.Context) *KnowledgeBaseServiceDescribeEnumerationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WithHTTPClient(client *http.Client) *KnowledgeBaseServiceDescribeEnumerationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WithBody(body *models.ConquestAPIDescribeEnumerationRequest) *KnowledgeBaseServiceDescribeEnumerationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the knowledge base service describe enumeration params
func (o *KnowledgeBaseServiceDescribeEnumerationParams) SetBody(body *models.ConquestAPIDescribeEnumerationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KnowledgeBaseServiceDescribeEnumerationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
