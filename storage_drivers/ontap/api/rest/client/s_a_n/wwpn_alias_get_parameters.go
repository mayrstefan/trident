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
	"github.com/go-openapi/swag"
)

// NewWwpnAliasGetParams creates a new WwpnAliasGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWwpnAliasGetParams() *WwpnAliasGetParams {
	return &WwpnAliasGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWwpnAliasGetParamsWithTimeout creates a new WwpnAliasGetParams object
// with the ability to set a timeout on a request.
func NewWwpnAliasGetParamsWithTimeout(timeout time.Duration) *WwpnAliasGetParams {
	return &WwpnAliasGetParams{
		timeout: timeout,
	}
}

// NewWwpnAliasGetParamsWithContext creates a new WwpnAliasGetParams object
// with the ability to set a context for a request.
func NewWwpnAliasGetParamsWithContext(ctx context.Context) *WwpnAliasGetParams {
	return &WwpnAliasGetParams{
		Context: ctx,
	}
}

// NewWwpnAliasGetParamsWithHTTPClient creates a new WwpnAliasGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewWwpnAliasGetParamsWithHTTPClient(client *http.Client) *WwpnAliasGetParams {
	return &WwpnAliasGetParams{
		HTTPClient: client,
	}
}

/*
WwpnAliasGetParams contains all the parameters to send to the API endpoint

	for the wwpn alias get operation.

	Typically these are written to a http.Request.
*/
type WwpnAliasGetParams struct {

	/* Alias.

	   The name of FC WWPN alias.

	*/
	Alias string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* SvmUUID.

	   The unique identifier of the SVM in which the alias is found.

	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wwpn alias get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WwpnAliasGetParams) WithDefaults() *WwpnAliasGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wwpn alias get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WwpnAliasGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wwpn alias get params
func (o *WwpnAliasGetParams) WithTimeout(timeout time.Duration) *WwpnAliasGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wwpn alias get params
func (o *WwpnAliasGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wwpn alias get params
func (o *WwpnAliasGetParams) WithContext(ctx context.Context) *WwpnAliasGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wwpn alias get params
func (o *WwpnAliasGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wwpn alias get params
func (o *WwpnAliasGetParams) WithHTTPClient(client *http.Client) *WwpnAliasGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wwpn alias get params
func (o *WwpnAliasGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlias adds the alias to the wwpn alias get params
func (o *WwpnAliasGetParams) WithAlias(alias string) *WwpnAliasGetParams {
	o.SetAlias(alias)
	return o
}

// SetAlias adds the alias to the wwpn alias get params
func (o *WwpnAliasGetParams) SetAlias(alias string) {
	o.Alias = alias
}

// WithFields adds the fields to the wwpn alias get params
func (o *WwpnAliasGetParams) WithFields(fields []string) *WwpnAliasGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the wwpn alias get params
func (o *WwpnAliasGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithSvmUUID adds the svmUUID to the wwpn alias get params
func (o *WwpnAliasGetParams) WithSvmUUID(svmUUID string) *WwpnAliasGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the wwpn alias get params
func (o *WwpnAliasGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *WwpnAliasGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alias
	if err := r.SetPathParam("alias", o.Alias); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamWwpnAliasGet binds the parameter fields
func (o *WwpnAliasGetParams) bindParamFields(formats strfmt.Registry) []string {
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
