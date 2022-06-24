// Code generated by go-swagger; DO NOT EDIT.

package csv_import_service

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
)

// NewCsvImportServiceDeleteCsvImportParams creates a new CsvImportServiceDeleteCsvImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCsvImportServiceDeleteCsvImportParams() *CsvImportServiceDeleteCsvImportParams {
	return &CsvImportServiceDeleteCsvImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCsvImportServiceDeleteCsvImportParamsWithTimeout creates a new CsvImportServiceDeleteCsvImportParams object
// with the ability to set a timeout on a request.
func NewCsvImportServiceDeleteCsvImportParamsWithTimeout(timeout time.Duration) *CsvImportServiceDeleteCsvImportParams {
	return &CsvImportServiceDeleteCsvImportParams{
		timeout: timeout,
	}
}

// NewCsvImportServiceDeleteCsvImportParamsWithContext creates a new CsvImportServiceDeleteCsvImportParams object
// with the ability to set a context for a request.
func NewCsvImportServiceDeleteCsvImportParamsWithContext(ctx context.Context) *CsvImportServiceDeleteCsvImportParams {
	return &CsvImportServiceDeleteCsvImportParams{
		Context: ctx,
	}
}

// NewCsvImportServiceDeleteCsvImportParamsWithHTTPClient creates a new CsvImportServiceDeleteCsvImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCsvImportServiceDeleteCsvImportParamsWithHTTPClient(client *http.Client) *CsvImportServiceDeleteCsvImportParams {
	return &CsvImportServiceDeleteCsvImportParams{
		HTTPClient: client,
	}
}

/* CsvImportServiceDeleteCsvImportParams contains all the parameters to send to the API endpoint
   for the csv import service delete csv import operation.

   Typically these are written to a http.Request.
*/
type CsvImportServiceDeleteCsvImportParams struct {

	/* JobID.

	   aka. ProcessID/BatchID
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the csv import service delete csv import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CsvImportServiceDeleteCsvImportParams) WithDefaults() *CsvImportServiceDeleteCsvImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the csv import service delete csv import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CsvImportServiceDeleteCsvImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) WithTimeout(timeout time.Duration) *CsvImportServiceDeleteCsvImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) WithContext(ctx context.Context) *CsvImportServiceDeleteCsvImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) WithHTTPClient(client *http.Client) *CsvImportServiceDeleteCsvImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) WithJobID(jobID string) *CsvImportServiceDeleteCsvImportParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the csv import service delete csv import params
func (o *CsvImportServiceDeleteCsvImportParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *CsvImportServiceDeleteCsvImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param JobID
	if err := r.SetPathParam("JobID", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
