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
	"github.com/go-openapi/swag"
)

// NewScheduleGetParams creates a new ScheduleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleGetParams() *ScheduleGetParams {
	return &ScheduleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleGetParamsWithTimeout creates a new ScheduleGetParams object
// with the ability to set a timeout on a request.
func NewScheduleGetParamsWithTimeout(timeout time.Duration) *ScheduleGetParams {
	return &ScheduleGetParams{
		timeout: timeout,
	}
}

// NewScheduleGetParamsWithContext creates a new ScheduleGetParams object
// with the ability to set a context for a request.
func NewScheduleGetParamsWithContext(ctx context.Context) *ScheduleGetParams {
	return &ScheduleGetParams{
		Context: ctx,
	}
}

// NewScheduleGetParamsWithHTTPClient creates a new ScheduleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleGetParamsWithHTTPClient(client *http.Client) *ScheduleGetParams {
	return &ScheduleGetParams{
		HTTPClient: client,
	}
}

/*
ScheduleGetParams contains all the parameters to send to the API endpoint

	for the schedule get operation.

	Typically these are written to a http.Request.
*/
type ScheduleGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Schedule UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleGetParams) WithDefaults() *ScheduleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedule get params
func (o *ScheduleGetParams) WithTimeout(timeout time.Duration) *ScheduleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule get params
func (o *ScheduleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule get params
func (o *ScheduleGetParams) WithContext(ctx context.Context) *ScheduleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule get params
func (o *ScheduleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule get params
func (o *ScheduleGetParams) WithHTTPClient(client *http.Client) *ScheduleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule get params
func (o *ScheduleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the schedule get params
func (o *ScheduleGetParams) WithFields(fields []string) *ScheduleGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the schedule get params
func (o *ScheduleGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the schedule get params
func (o *ScheduleGetParams) WithUUID(uuid string) *ScheduleGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the schedule get params
func (o *ScheduleGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamScheduleGet binds the parameter fields
func (o *ScheduleGetParams) bindParamFields(formats strfmt.Registry) []string {
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
