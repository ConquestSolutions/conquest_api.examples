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

// NewKnowledgeBaseServiceGetCodeListsParams creates a new KnowledgeBaseServiceGetCodeListsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKnowledgeBaseServiceGetCodeListsParams() *KnowledgeBaseServiceGetCodeListsParams {
	return &KnowledgeBaseServiceGetCodeListsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKnowledgeBaseServiceGetCodeListsParamsWithTimeout creates a new KnowledgeBaseServiceGetCodeListsParams object
// with the ability to set a timeout on a request.
func NewKnowledgeBaseServiceGetCodeListsParamsWithTimeout(timeout time.Duration) *KnowledgeBaseServiceGetCodeListsParams {
	return &KnowledgeBaseServiceGetCodeListsParams{
		timeout: timeout,
	}
}

// NewKnowledgeBaseServiceGetCodeListsParamsWithContext creates a new KnowledgeBaseServiceGetCodeListsParams object
// with the ability to set a context for a request.
func NewKnowledgeBaseServiceGetCodeListsParamsWithContext(ctx context.Context) *KnowledgeBaseServiceGetCodeListsParams {
	return &KnowledgeBaseServiceGetCodeListsParams{
		Context: ctx,
	}
}

// NewKnowledgeBaseServiceGetCodeListsParamsWithHTTPClient creates a new KnowledgeBaseServiceGetCodeListsParams object
// with the ability to set a custom HTTPClient for a request.
func NewKnowledgeBaseServiceGetCodeListsParamsWithHTTPClient(client *http.Client) *KnowledgeBaseServiceGetCodeListsParams {
	return &KnowledgeBaseServiceGetCodeListsParams{
		HTTPClient: client,
	}
}

/* KnowledgeBaseServiceGetCodeListsParams contains all the parameters to send to the API endpoint
   for the knowledge base service get code lists operation.

   Typically these are written to a http.Request.
*/
type KnowledgeBaseServiceGetCodeListsParams struct {

	// Body.
	Body *models.ConquestAPIGetCodeListsQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the knowledge base service get code lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KnowledgeBaseServiceGetCodeListsParams) WithDefaults() *KnowledgeBaseServiceGetCodeListsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the knowledge base service get code lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KnowledgeBaseServiceGetCodeListsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) WithTimeout(timeout time.Duration) *KnowledgeBaseServiceGetCodeListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) WithContext(ctx context.Context) *KnowledgeBaseServiceGetCodeListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) WithHTTPClient(client *http.Client) *KnowledgeBaseServiceGetCodeListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) WithBody(body *models.ConquestAPIGetCodeListsQuery) *KnowledgeBaseServiceGetCodeListsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the knowledge base service get code lists params
func (o *KnowledgeBaseServiceGetCodeListsParams) SetBody(body *models.ConquestAPIGetCodeListsQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KnowledgeBaseServiceGetCodeListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
