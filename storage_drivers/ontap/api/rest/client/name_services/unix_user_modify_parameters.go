// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewUnixUserModifyParams creates a new UnixUserModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixUserModifyParams() *UnixUserModifyParams {
	return &UnixUserModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixUserModifyParamsWithTimeout creates a new UnixUserModifyParams object
// with the ability to set a timeout on a request.
func NewUnixUserModifyParamsWithTimeout(timeout time.Duration) *UnixUserModifyParams {
	return &UnixUserModifyParams{
		timeout: timeout,
	}
}

// NewUnixUserModifyParamsWithContext creates a new UnixUserModifyParams object
// with the ability to set a context for a request.
func NewUnixUserModifyParamsWithContext(ctx context.Context) *UnixUserModifyParams {
	return &UnixUserModifyParams{
		Context: ctx,
	}
}

// NewUnixUserModifyParamsWithHTTPClient creates a new UnixUserModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixUserModifyParamsWithHTTPClient(client *http.Client) *UnixUserModifyParams {
	return &UnixUserModifyParams{
		HTTPClient: client,
	}
}

/* UnixUserModifyParams contains all the parameters to send to the API endpoint
   for the unix user modify operation.

   Typically these are written to a http.Request.
*/
type UnixUserModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.UnixUser

	/* Name.

	   UNIX user name
	*/
	NamePathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix user modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserModifyParams) WithDefaults() *UnixUserModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix user modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unix user modify params
func (o *UnixUserModifyParams) WithTimeout(timeout time.Duration) *UnixUserModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix user modify params
func (o *UnixUserModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix user modify params
func (o *UnixUserModifyParams) WithContext(ctx context.Context) *UnixUserModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix user modify params
func (o *UnixUserModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix user modify params
func (o *UnixUserModifyParams) WithHTTPClient(client *http.Client) *UnixUserModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix user modify params
func (o *UnixUserModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the unix user modify params
func (o *UnixUserModifyParams) WithInfo(info *models.UnixUser) *UnixUserModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix user modify params
func (o *UnixUserModifyParams) SetInfo(info *models.UnixUser) {
	o.Info = info
}

// WithNamePathParameter adds the name to the unix user modify params
func (o *UnixUserModifyParams) WithNamePathParameter(name string) *UnixUserModifyParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the unix user modify params
func (o *UnixUserModifyParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix user modify params
func (o *UnixUserModifyParams) WithSVMUUIDPathParameter(svmUUID string) *UnixUserModifyParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix user modify params
func (o *UnixUserModifyParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UnixUserModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}