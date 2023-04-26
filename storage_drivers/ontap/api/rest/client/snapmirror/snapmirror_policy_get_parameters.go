// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

// NewSnapmirrorPolicyGetParams creates a new SnapmirrorPolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorPolicyGetParams() *SnapmirrorPolicyGetParams {
	return &SnapmirrorPolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorPolicyGetParamsWithTimeout creates a new SnapmirrorPolicyGetParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorPolicyGetParamsWithTimeout(timeout time.Duration) *SnapmirrorPolicyGetParams {
	return &SnapmirrorPolicyGetParams{
		timeout: timeout,
	}
}

// NewSnapmirrorPolicyGetParamsWithContext creates a new SnapmirrorPolicyGetParams object
// with the ability to set a context for a request.
func NewSnapmirrorPolicyGetParamsWithContext(ctx context.Context) *SnapmirrorPolicyGetParams {
	return &SnapmirrorPolicyGetParams{
		Context: ctx,
	}
}

// NewSnapmirrorPolicyGetParamsWithHTTPClient creates a new SnapmirrorPolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorPolicyGetParamsWithHTTPClient(client *http.Client) *SnapmirrorPolicyGetParams {
	return &SnapmirrorPolicyGetParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorPolicyGetParams contains all the parameters to send to the API endpoint

	for the snapmirror policy get operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorPolicyGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Policy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyGetParams) WithDefaults() *SnapmirrorPolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) WithTimeout(timeout time.Duration) *SnapmirrorPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) WithContext(ctx context.Context) *SnapmirrorPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) WithHTTPClient(client *http.Client) *SnapmirrorPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) WithFields(fields []string) *SnapmirrorPolicyGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) WithUUID(uuid string) *SnapmirrorPolicyGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapmirror policy get params
func (o *SnapmirrorPolicyGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnapmirrorPolicyGet binds the parameter fields
func (o *SnapmirrorPolicyGetParams) bindParamFields(formats strfmt.Registry) []string {
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
