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

// NewVscanConfigDeleteParams creates a new VscanConfigDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanConfigDeleteParams() *VscanConfigDeleteParams {
	return &VscanConfigDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanConfigDeleteParamsWithTimeout creates a new VscanConfigDeleteParams object
// with the ability to set a timeout on a request.
func NewVscanConfigDeleteParamsWithTimeout(timeout time.Duration) *VscanConfigDeleteParams {
	return &VscanConfigDeleteParams{
		timeout: timeout,
	}
}

// NewVscanConfigDeleteParamsWithContext creates a new VscanConfigDeleteParams object
// with the ability to set a context for a request.
func NewVscanConfigDeleteParamsWithContext(ctx context.Context) *VscanConfigDeleteParams {
	return &VscanConfigDeleteParams{
		Context: ctx,
	}
}

// NewVscanConfigDeleteParamsWithHTTPClient creates a new VscanConfigDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanConfigDeleteParamsWithHTTPClient(client *http.Client) *VscanConfigDeleteParams {
	return &VscanConfigDeleteParams{
		HTTPClient: client,
	}
}

/*
VscanConfigDeleteParams contains all the parameters to send to the API endpoint

	for the vscan config delete operation.

	Typically these are written to a http.Request.
*/
type VscanConfigDeleteParams struct {

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan config delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanConfigDeleteParams) WithDefaults() *VscanConfigDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan config delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanConfigDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan config delete params
func (o *VscanConfigDeleteParams) WithTimeout(timeout time.Duration) *VscanConfigDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan config delete params
func (o *VscanConfigDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan config delete params
func (o *VscanConfigDeleteParams) WithContext(ctx context.Context) *VscanConfigDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan config delete params
func (o *VscanConfigDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan config delete params
func (o *VscanConfigDeleteParams) WithHTTPClient(client *http.Client) *VscanConfigDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan config delete params
func (o *VscanConfigDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSvmUUID adds the svmUUID to the vscan config delete params
func (o *VscanConfigDeleteParams) WithSvmUUID(svmUUID string) *VscanConfigDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan config delete params
func (o *VscanConfigDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanConfigDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
