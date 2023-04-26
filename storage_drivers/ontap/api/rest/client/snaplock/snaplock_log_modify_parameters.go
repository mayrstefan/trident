// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSnaplockLogModifyParams creates a new SnaplockLogModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLogModifyParams() *SnaplockLogModifyParams {
	return &SnaplockLogModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLogModifyParamsWithTimeout creates a new SnaplockLogModifyParams object
// with the ability to set a timeout on a request.
func NewSnaplockLogModifyParamsWithTimeout(timeout time.Duration) *SnaplockLogModifyParams {
	return &SnaplockLogModifyParams{
		timeout: timeout,
	}
}

// NewSnaplockLogModifyParamsWithContext creates a new SnaplockLogModifyParams object
// with the ability to set a context for a request.
func NewSnaplockLogModifyParamsWithContext(ctx context.Context) *SnaplockLogModifyParams {
	return &SnaplockLogModifyParams{
		Context: ctx,
	}
}

// NewSnaplockLogModifyParamsWithHTTPClient creates a new SnaplockLogModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLogModifyParamsWithHTTPClient(client *http.Client) *SnaplockLogModifyParams {
	return &SnaplockLogModifyParams{
		HTTPClient: client,
	}
}

/*
SnaplockLogModifyParams contains all the parameters to send to the API endpoint

	for the snaplock log modify operation.

	Typically these are written to a http.Request.
*/
type SnaplockLogModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnaplockLog

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* SvmUUID.

	   SVM UUID
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock log modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogModifyParams) WithDefaults() *SnaplockLogModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock log modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLogModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := SnaplockLogModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithTimeout(timeout time.Duration) *SnaplockLogModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithContext(ctx context.Context) *SnaplockLogModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithHTTPClient(client *http.Client) *SnaplockLogModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithInfo(info *models.SnaplockLog) *SnaplockLogModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetInfo(info *models.SnaplockLog) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithReturnTimeout(returnTimeout *int64) *SnaplockLogModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmUUID adds the svmUUID to the snaplock log modify params
func (o *SnaplockLogModifyParams) WithSvmUUID(svmUUID string) *SnaplockLogModifyParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the snaplock log modify params
func (o *SnaplockLogModifyParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLogModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
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
