// Code generated by go-swagger; DO NOT EDIT.

package security

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewMultiAdminVerifyRequestModifyParams creates a new MultiAdminVerifyRequestModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRequestModifyParams() *MultiAdminVerifyRequestModifyParams {
	return &MultiAdminVerifyRequestModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRequestModifyParamsWithTimeout creates a new MultiAdminVerifyRequestModifyParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRequestModifyParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRequestModifyParams {
	return &MultiAdminVerifyRequestModifyParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRequestModifyParamsWithContext creates a new MultiAdminVerifyRequestModifyParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRequestModifyParamsWithContext(ctx context.Context) *MultiAdminVerifyRequestModifyParams {
	return &MultiAdminVerifyRequestModifyParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRequestModifyParamsWithHTTPClient creates a new MultiAdminVerifyRequestModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRequestModifyParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRequestModifyParams {
	return &MultiAdminVerifyRequestModifyParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRequestModifyParams contains all the parameters to send to the API endpoint

	for the multi admin verify request modify operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRequestModifyParams struct {

	// Index.
	Index string

	/* Info.

	   Info specification
	*/
	Info *models.MultiAdminVerifyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify request modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestModifyParams) WithDefaults() *MultiAdminVerifyRequestModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify request modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRequestModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) WithContext(ctx context.Context) *MultiAdminVerifyRequestModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRequestModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) WithIndex(index string) *MultiAdminVerifyRequestModifyParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) SetIndex(index string) {
	o.Index = index
}

// WithInfo adds the info to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) WithInfo(info *models.MultiAdminVerifyRequest) *MultiAdminVerifyRequestModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the multi admin verify request modify params
func (o *MultiAdminVerifyRequestModifyParams) SetInfo(info *models.MultiAdminVerifyRequest) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRequestModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
