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
	"github.com/go-openapi/swag"
)

// NewVscanScannerPoolGetParams creates a new VscanScannerPoolGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanScannerPoolGetParams() *VscanScannerPoolGetParams {
	return &VscanScannerPoolGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanScannerPoolGetParamsWithTimeout creates a new VscanScannerPoolGetParams object
// with the ability to set a timeout on a request.
func NewVscanScannerPoolGetParamsWithTimeout(timeout time.Duration) *VscanScannerPoolGetParams {
	return &VscanScannerPoolGetParams{
		timeout: timeout,
	}
}

// NewVscanScannerPoolGetParamsWithContext creates a new VscanScannerPoolGetParams object
// with the ability to set a context for a request.
func NewVscanScannerPoolGetParamsWithContext(ctx context.Context) *VscanScannerPoolGetParams {
	return &VscanScannerPoolGetParams{
		Context: ctx,
	}
}

// NewVscanScannerPoolGetParamsWithHTTPClient creates a new VscanScannerPoolGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanScannerPoolGetParamsWithHTTPClient(client *http.Client) *VscanScannerPoolGetParams {
	return &VscanScannerPoolGetParams{
		HTTPClient: client,
	}
}

/*
VscanScannerPoolGetParams contains all the parameters to send to the API endpoint

	for the vscan scanner pool get operation.

	Typically these are written to a http.Request.
*/
type VscanScannerPoolGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

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

// WithDefaults hydrates default values in the vscan scanner pool get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanScannerPoolGetParams) WithDefaults() *VscanScannerPoolGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan scanner pool get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanScannerPoolGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithTimeout(timeout time.Duration) *VscanScannerPoolGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithContext(ctx context.Context) *VscanScannerPoolGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithHTTPClient(client *http.Client) *VscanScannerPoolGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithFields(fields []string) *VscanScannerPoolGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithName(name string) *VscanScannerPoolGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) WithSvmUUID(svmUUID string) *VscanScannerPoolGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan scanner pool get params
func (o *VscanScannerPoolGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanScannerPoolGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

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

// bindParamVscanScannerPoolGet binds the parameter fields
func (o *VscanScannerPoolGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
