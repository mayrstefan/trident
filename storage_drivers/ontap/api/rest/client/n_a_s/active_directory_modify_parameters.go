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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewActiveDirectoryModifyParams creates a new ActiveDirectoryModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActiveDirectoryModifyParams() *ActiveDirectoryModifyParams {
	return &ActiveDirectoryModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActiveDirectoryModifyParamsWithTimeout creates a new ActiveDirectoryModifyParams object
// with the ability to set a timeout on a request.
func NewActiveDirectoryModifyParamsWithTimeout(timeout time.Duration) *ActiveDirectoryModifyParams {
	return &ActiveDirectoryModifyParams{
		timeout: timeout,
	}
}

// NewActiveDirectoryModifyParamsWithContext creates a new ActiveDirectoryModifyParams object
// with the ability to set a context for a request.
func NewActiveDirectoryModifyParamsWithContext(ctx context.Context) *ActiveDirectoryModifyParams {
	return &ActiveDirectoryModifyParams{
		Context: ctx,
	}
}

// NewActiveDirectoryModifyParamsWithHTTPClient creates a new ActiveDirectoryModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewActiveDirectoryModifyParamsWithHTTPClient(client *http.Client) *ActiveDirectoryModifyParams {
	return &ActiveDirectoryModifyParams{
		HTTPClient: client,
	}
}

/*
ActiveDirectoryModifyParams contains all the parameters to send to the API endpoint

	for the active directory modify operation.

	Typically these are written to a http.Request.
*/
type ActiveDirectoryModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.ActiveDirectory

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the active directory modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveDirectoryModifyParams) WithDefaults() *ActiveDirectoryModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the active directory modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveDirectoryModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the active directory modify params
func (o *ActiveDirectoryModifyParams) WithTimeout(timeout time.Duration) *ActiveDirectoryModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the active directory modify params
func (o *ActiveDirectoryModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the active directory modify params
func (o *ActiveDirectoryModifyParams) WithContext(ctx context.Context) *ActiveDirectoryModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the active directory modify params
func (o *ActiveDirectoryModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the active directory modify params
func (o *ActiveDirectoryModifyParams) WithHTTPClient(client *http.Client) *ActiveDirectoryModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the active directory modify params
func (o *ActiveDirectoryModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the active directory modify params
func (o *ActiveDirectoryModifyParams) WithInfo(info *models.ActiveDirectory) *ActiveDirectoryModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the active directory modify params
func (o *ActiveDirectoryModifyParams) SetInfo(info *models.ActiveDirectory) {
	o.Info = info
}

// WithSvmUUID adds the svmUUID to the active directory modify params
func (o *ActiveDirectoryModifyParams) WithSvmUUID(svmUUID string) *ActiveDirectoryModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the active directory modify params
func (o *ActiveDirectoryModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ActiveDirectoryModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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
