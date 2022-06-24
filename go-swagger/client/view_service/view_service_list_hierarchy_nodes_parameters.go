// Code generated by go-swagger; DO NOT EDIT.

package view_service

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

// NewViewServiceListHierarchyNodesParams creates a new ViewServiceListHierarchyNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewViewServiceListHierarchyNodesParams() *ViewServiceListHierarchyNodesParams {
	return &ViewServiceListHierarchyNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewViewServiceListHierarchyNodesParamsWithTimeout creates a new ViewServiceListHierarchyNodesParams object
// with the ability to set a timeout on a request.
func NewViewServiceListHierarchyNodesParamsWithTimeout(timeout time.Duration) *ViewServiceListHierarchyNodesParams {
	return &ViewServiceListHierarchyNodesParams{
		timeout: timeout,
	}
}

// NewViewServiceListHierarchyNodesParamsWithContext creates a new ViewServiceListHierarchyNodesParams object
// with the ability to set a context for a request.
func NewViewServiceListHierarchyNodesParamsWithContext(ctx context.Context) *ViewServiceListHierarchyNodesParams {
	return &ViewServiceListHierarchyNodesParams{
		Context: ctx,
	}
}

// NewViewServiceListHierarchyNodesParamsWithHTTPClient creates a new ViewServiceListHierarchyNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewViewServiceListHierarchyNodesParamsWithHTTPClient(client *http.Client) *ViewServiceListHierarchyNodesParams {
	return &ViewServiceListHierarchyNodesParams{
		HTTPClient: client,
	}
}

/* ViewServiceListHierarchyNodesParams contains all the parameters to send to the API endpoint
   for the view service list hierarchy nodes operation.

   Typically these are written to a http.Request.
*/
type ViewServiceListHierarchyNodesParams struct {

	// Body.
	Body *models.ConquestAPIListHierarchyNodesQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the view service list hierarchy nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewServiceListHierarchyNodesParams) WithDefaults() *ViewServiceListHierarchyNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the view service list hierarchy nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ViewServiceListHierarchyNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) WithTimeout(timeout time.Duration) *ViewServiceListHierarchyNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) WithContext(ctx context.Context) *ViewServiceListHierarchyNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) WithHTTPClient(client *http.Client) *ViewServiceListHierarchyNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) WithBody(body *models.ConquestAPIListHierarchyNodesQuery) *ViewServiceListHierarchyNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the view service list hierarchy nodes params
func (o *ViewServiceListHierarchyNodesParams) SetBody(body *models.ConquestAPIListHierarchyNodesQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ViewServiceListHierarchyNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
