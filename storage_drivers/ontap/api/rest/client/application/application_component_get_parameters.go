// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationComponentGetParams creates a new ApplicationComponentGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationComponentGetParams() *ApplicationComponentGetParams {
	return &ApplicationComponentGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationComponentGetParamsWithTimeout creates a new ApplicationComponentGetParams object
// with the ability to set a timeout on a request.
func NewApplicationComponentGetParamsWithTimeout(timeout time.Duration) *ApplicationComponentGetParams {
	return &ApplicationComponentGetParams{
		timeout: timeout,
	}
}

// NewApplicationComponentGetParamsWithContext creates a new ApplicationComponentGetParams object
// with the ability to set a context for a request.
func NewApplicationComponentGetParamsWithContext(ctx context.Context) *ApplicationComponentGetParams {
	return &ApplicationComponentGetParams{
		Context: ctx,
	}
}

// NewApplicationComponentGetParamsWithHTTPClient creates a new ApplicationComponentGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationComponentGetParamsWithHTTPClient(client *http.Client) *ApplicationComponentGetParams {
	return &ApplicationComponentGetParams{
		HTTPClient: client,
	}
}

/*
ApplicationComponentGetParams contains all the parameters to send to the API endpoint

	for the application component get operation.

	Typically these are written to a http.Request.
*/
type ApplicationComponentGetParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUID string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Application component UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application component get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentGetParams) WithDefaults() *ApplicationComponentGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application component get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application component get params
func (o *ApplicationComponentGetParams) WithTimeout(timeout time.Duration) *ApplicationComponentGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application component get params
func (o *ApplicationComponentGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application component get params
func (o *ApplicationComponentGetParams) WithContext(ctx context.Context) *ApplicationComponentGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application component get params
func (o *ApplicationComponentGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application component get params
func (o *ApplicationComponentGetParams) WithHTTPClient(client *http.Client) *ApplicationComponentGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application component get params
func (o *ApplicationComponentGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUID adds the applicationUUID to the application component get params
func (o *ApplicationComponentGetParams) WithApplicationUUID(applicationUUID string) *ApplicationComponentGetParams {
	o.SetApplicationUUID(applicationUUID)
	return o
}

// SetApplicationUUID adds the applicationUuid to the application component get params
func (o *ApplicationComponentGetParams) SetApplicationUUID(applicationUUID string) {
	o.ApplicationUUID = applicationUUID
}

// WithFields adds the fields to the application component get params
func (o *ApplicationComponentGetParams) WithFields(fields []string) *ApplicationComponentGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the application component get params
func (o *ApplicationComponentGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the application component get params
func (o *ApplicationComponentGetParams) WithUUID(uuid string) *ApplicationComponentGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the application component get params
func (o *ApplicationComponentGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationComponentGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUID); err != nil {
		return err
	}

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

// bindParamApplicationComponentGet binds the parameter fields
func (o *ApplicationComponentGetParams) bindParamFields(formats strfmt.Registry) []string {
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
