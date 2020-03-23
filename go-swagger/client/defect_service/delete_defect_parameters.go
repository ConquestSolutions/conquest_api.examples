// Code generated by go-swagger; DO NOT EDIT.

package defect_service

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

// NewDeleteDefectParams creates a new DeleteDefectParams object
// with the default values initialized.
func NewDeleteDefectParams() *DeleteDefectParams {
	var ()
	return &DeleteDefectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDefectParamsWithTimeout creates a new DeleteDefectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDefectParamsWithTimeout(timeout time.Duration) *DeleteDefectParams {
	var ()
	return &DeleteDefectParams{

		timeout: timeout,
	}
}

// NewDeleteDefectParamsWithContext creates a new DeleteDefectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDefectParamsWithContext(ctx context.Context) *DeleteDefectParams {
	var ()
	return &DeleteDefectParams{

		Context: ctx,
	}
}

// NewDeleteDefectParamsWithHTTPClient creates a new DeleteDefectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDefectParamsWithHTTPClient(client *http.Client) *DeleteDefectParams {
	var ()
	return &DeleteDefectParams{
		HTTPClient: client,
	}
}

/*DeleteDefectParams contains all the parameters to send to the API endpoint
for the delete defect operation typically these are written to a http.Request
*/
type DeleteDefectParams struct {

	/*Body*/
	Body *models.ConquestAPIDeleteDefectCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete defect params
func (o *DeleteDefectParams) WithTimeout(timeout time.Duration) *DeleteDefectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete defect params
func (o *DeleteDefectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete defect params
func (o *DeleteDefectParams) WithContext(ctx context.Context) *DeleteDefectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete defect params
func (o *DeleteDefectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete defect params
func (o *DeleteDefectParams) WithHTTPClient(client *http.Client) *DeleteDefectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete defect params
func (o *DeleteDefectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete defect params
func (o *DeleteDefectParams) WithBody(body *models.ConquestAPIDeleteDefectCommand) *DeleteDefectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete defect params
func (o *DeleteDefectParams) SetBody(body *models.ConquestAPIDeleteDefectCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDefectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
