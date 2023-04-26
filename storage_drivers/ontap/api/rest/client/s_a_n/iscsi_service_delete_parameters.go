// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIscsiServiceDeleteParams creates a new IscsiServiceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiServiceDeleteParams() *IscsiServiceDeleteParams {
	return &IscsiServiceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiServiceDeleteParamsWithTimeout creates a new IscsiServiceDeleteParams object
// with the ability to set a timeout on a request.
func NewIscsiServiceDeleteParamsWithTimeout(timeout time.Duration) *IscsiServiceDeleteParams {
	return &IscsiServiceDeleteParams{
		timeout: timeout,
	}
}

// NewIscsiServiceDeleteParamsWithContext creates a new IscsiServiceDeleteParams object
// with the ability to set a context for a request.
func NewIscsiServiceDeleteParamsWithContext(ctx context.Context) *IscsiServiceDeleteParams {
	return &IscsiServiceDeleteParams{
		Context: ctx,
	}
}

// NewIscsiServiceDeleteParamsWithHTTPClient creates a new IscsiServiceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiServiceDeleteParamsWithHTTPClient(client *http.Client) *IscsiServiceDeleteParams {
	return &IscsiServiceDeleteParams{
		HTTPClient: client,
	}
}

/*
IscsiServiceDeleteParams contains all the parameters to send to the API endpoint

	for the iscsi service delete operation.

	Typically these are written to a http.Request.
*/
type IscsiServiceDeleteParams struct {

	/* SvmUUID.

	   The unique identifier of the SVM for which to delete the iSCSI service.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceDeleteParams) WithDefaults() *IscsiServiceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi service delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiServiceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iscsi service delete params
func (o *IscsiServiceDeleteParams) WithTimeout(timeout time.Duration) *IscsiServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi service delete params
func (o *IscsiServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi service delete params
func (o *IscsiServiceDeleteParams) WithContext(ctx context.Context) *IscsiServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi service delete params
func (o *IscsiServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi service delete params
func (o *IscsiServiceDeleteParams) WithHTTPClient(client *http.Client) *IscsiServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi service delete params
func (o *IscsiServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSvmUUID adds the svmUUID to the iscsi service delete params
func (o *IscsiServiceDeleteParams) WithSvmUUID(svmUUID string) *IscsiServiceDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the iscsi service delete params
func (o *IscsiServiceDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
