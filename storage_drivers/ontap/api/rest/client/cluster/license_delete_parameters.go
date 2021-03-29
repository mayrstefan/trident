// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewLicenseDeleteParams creates a new LicenseDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLicenseDeleteParams() *LicenseDeleteParams {
	return &LicenseDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLicenseDeleteParamsWithTimeout creates a new LicenseDeleteParams object
// with the ability to set a timeout on a request.
func NewLicenseDeleteParamsWithTimeout(timeout time.Duration) *LicenseDeleteParams {
	return &LicenseDeleteParams{
		timeout: timeout,
	}
}

// NewLicenseDeleteParamsWithContext creates a new LicenseDeleteParams object
// with the ability to set a context for a request.
func NewLicenseDeleteParamsWithContext(ctx context.Context) *LicenseDeleteParams {
	return &LicenseDeleteParams{
		Context: ctx,
	}
}

// NewLicenseDeleteParamsWithHTTPClient creates a new LicenseDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLicenseDeleteParamsWithHTTPClient(client *http.Client) *LicenseDeleteParams {
	return &LicenseDeleteParams{
		HTTPClient: client,
	}
}

/* LicenseDeleteParams contains all the parameters to send to the API endpoint
   for the license delete operation.

   Typically these are written to a http.Request.
*/
type LicenseDeleteParams struct {

	/* Name.

	   Name of the license package to delete.
	*/
	NamePathParameter string

	/* SerialNumber.

	   Serial number of the license to delete.
	*/
	SerialNumberQueryParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the license delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseDeleteParams) WithDefaults() *LicenseDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the license delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LicenseDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the license delete params
func (o *LicenseDeleteParams) WithTimeout(timeout time.Duration) *LicenseDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the license delete params
func (o *LicenseDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the license delete params
func (o *LicenseDeleteParams) WithContext(ctx context.Context) *LicenseDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the license delete params
func (o *LicenseDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the license delete params
func (o *LicenseDeleteParams) WithHTTPClient(client *http.Client) *LicenseDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the license delete params
func (o *LicenseDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the license delete params
func (o *LicenseDeleteParams) WithNamePathParameter(name string) *LicenseDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the license delete params
func (o *LicenseDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSerialNumberQueryParameter adds the serialNumber to the license delete params
func (o *LicenseDeleteParams) WithSerialNumberQueryParameter(serialNumber string) *LicenseDeleteParams {
	o.SetSerialNumberQueryParameter(serialNumber)
	return o
}

// SetSerialNumberQueryParameter adds the serialNumber to the license delete params
func (o *LicenseDeleteParams) SetSerialNumberQueryParameter(serialNumber string) {
	o.SerialNumberQueryParameter = serialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *LicenseDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// query param serial_number
	qrSerialNumber := o.SerialNumberQueryParameter
	qSerialNumber := qrSerialNumber
	if qSerialNumber != "" {

		if err := r.SetQueryParam("serial_number", qSerialNumber); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}