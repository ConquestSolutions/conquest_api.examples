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

// NewCsvImportServiceAddCsvImportParams creates a new CsvImportServiceAddCsvImportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCsvImportServiceAddCsvImportParams() *CsvImportServiceAddCsvImportParams {
	return &CsvImportServiceAddCsvImportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCsvImportServiceAddCsvImportParamsWithTimeout creates a new CsvImportServiceAddCsvImportParams object
// with the ability to set a timeout on a request.
func NewCsvImportServiceAddCsvImportParamsWithTimeout(timeout time.Duration) *CsvImportServiceAddCsvImportParams {
	return &CsvImportServiceAddCsvImportParams{
		timeout: timeout,
	}
}

// NewCsvImportServiceAddCsvImportParamsWithContext creates a new CsvImportServiceAddCsvImportParams object
// with the ability to set a context for a request.
func NewCsvImportServiceAddCsvImportParamsWithContext(ctx context.Context) *CsvImportServiceAddCsvImportParams {
	return &CsvImportServiceAddCsvImportParams{
		Context: ctx,
	}
}

// NewCsvImportServiceAddCsvImportParamsWithHTTPClient creates a new CsvImportServiceAddCsvImportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCsvImportServiceAddCsvImportParamsWithHTTPClient(client *http.Client) *CsvImportServiceAddCsvImportParams {
	return &CsvImportServiceAddCsvImportParams{
		HTTPClient: client,
	}
}

/* CsvImportServiceAddCsvImportParams contains all the parameters to send to the API endpoint
   for the csv import service add csv import operation.

   Typically these are written to a http.Request.
*/
type CsvImportServiceAddCsvImportParams struct {

	/* ImportType.

	   Action, Asset, Defect, Request, AssetInspection, RiskEvent, LogBook
	*/
	ImportType string

	/* File.

	   CSV file data
	*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the csv import service add csv import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CsvImportServiceAddCsvImportParams) WithDefaults() *CsvImportServiceAddCsvImportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the csv import service add csv import params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CsvImportServiceAddCsvImportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) WithTimeout(timeout time.Duration) *CsvImportServiceAddCsvImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) WithContext(ctx context.Context) *CsvImportServiceAddCsvImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) WithHTTPClient(client *http.Client) *CsvImportServiceAddCsvImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImportType adds the importType to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) WithImportType(importType string) *CsvImportServiceAddCsvImportParams {
	o.SetImportType(importType)
	return o
}

// SetImportType adds the importType to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) SetImportType(importType string) {
	o.ImportType = importType
}

// WithFile adds the file to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) WithFile(file runtime.NamedReadCloser) *CsvImportServiceAddCsvImportParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the csv import service add csv import params
func (o *CsvImportServiceAddCsvImportParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *CsvImportServiceAddCsvImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ImportType
	if err := r.SetPathParam("ImportType", o.ImportType); err != nil {
		return err
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
