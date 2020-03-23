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

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// NewListHierarchyNodesParams creates a new ListHierarchyNodesParams object
// with the default values initialized.
func NewListHierarchyNodesParams() *ListHierarchyNodesParams {
	var ()
	return &ListHierarchyNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListHierarchyNodesParamsWithTimeout creates a new ListHierarchyNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListHierarchyNodesParamsWithTimeout(timeout time.Duration) *ListHierarchyNodesParams {
	var ()
	return &ListHierarchyNodesParams{

		timeout: timeout,
	}
}

// NewListHierarchyNodesParamsWithContext creates a new ListHierarchyNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListHierarchyNodesParamsWithContext(ctx context.Context) *ListHierarchyNodesParams {
	var ()
	return &ListHierarchyNodesParams{

		Context: ctx,
	}
}

// NewListHierarchyNodesParamsWithHTTPClient creates a new ListHierarchyNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListHierarchyNodesParamsWithHTTPClient(client *http.Client) *ListHierarchyNodesParams {
	var ()
	return &ListHierarchyNodesParams{
		HTTPClient: client,
	}
}

/*ListHierarchyNodesParams contains all the parameters to send to the API endpoint
for the list hierarchy nodes operation typically these are written to a http.Request
*/
type ListHierarchyNodesParams struct {

	/*Body*/
	Body *models.ConquestAPIListHierarchyNodesQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) WithTimeout(timeout time.Duration) *ListHierarchyNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) WithContext(ctx context.Context) *ListHierarchyNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) WithHTTPClient(client *http.Client) *ListHierarchyNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) WithBody(body *models.ConquestAPIListHierarchyNodesQuery) *ListHierarchyNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list hierarchy nodes params
func (o *ListHierarchyNodesParams) SetBody(body *models.ConquestAPIListHierarchyNodesQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListHierarchyNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
