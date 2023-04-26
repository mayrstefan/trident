// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewVscanOnAccessDeleteParams creates a new VscanOnAccessDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnAccessDeleteParams() *VscanOnAccessDeleteParams {
	return &VscanOnAccessDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnAccessDeleteParamsWithTimeout creates a new VscanOnAccessDeleteParams object
// with the ability to set a timeout on a request.
func NewVscanOnAccessDeleteParamsWithTimeout(timeout time.Duration) *VscanOnAccessDeleteParams {
	return &VscanOnAccessDeleteParams{
		timeout: timeout,
	}
}

// NewVscanOnAccessDeleteParamsWithContext creates a new VscanOnAccessDeleteParams object
// with the ability to set a context for a request.
func NewVscanOnAccessDeleteParamsWithContext(ctx context.Context) *VscanOnAccessDeleteParams {
	return &VscanOnAccessDeleteParams{
		Context: ctx,
	}
}

// NewVscanOnAccessDeleteParamsWithHTTPClient creates a new VscanOnAccessDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnAccessDeleteParamsWithHTTPClient(client *http.Client) *VscanOnAccessDeleteParams {
	return &VscanOnAccessDeleteParams{
		HTTPClient: client,
	}
}

/*
VscanOnAccessDeleteParams contains all the parameters to send to the API endpoint

	for the vscan on access delete operation.

	Typically these are written to a http.Request.
*/
type VscanOnAccessDeleteParams struct {

	// Name.
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan on access delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessDeleteParams) WithDefaults() *VscanOnAccessDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on access delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnAccessDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) WithTimeout(timeout time.Duration) *VscanOnAccessDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) WithContext(ctx context.Context) *VscanOnAccessDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) WithHTTPClient(client *http.Client) *VscanOnAccessDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) WithName(name string) *VscanOnAccessDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) WithSvmUUID(svmUUID string) *VscanOnAccessDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan on access delete params
func (o *VscanOnAccessDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnAccessDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
