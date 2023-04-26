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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewMultiAdminVerifyRuleCreateParams creates a new MultiAdminVerifyRuleCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRuleCreateParams() *MultiAdminVerifyRuleCreateParams {
	return &MultiAdminVerifyRuleCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRuleCreateParamsWithTimeout creates a new MultiAdminVerifyRuleCreateParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRuleCreateParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRuleCreateParams {
	return &MultiAdminVerifyRuleCreateParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRuleCreateParamsWithContext creates a new MultiAdminVerifyRuleCreateParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRuleCreateParamsWithContext(ctx context.Context) *MultiAdminVerifyRuleCreateParams {
	return &MultiAdminVerifyRuleCreateParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRuleCreateParamsWithHTTPClient creates a new MultiAdminVerifyRuleCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRuleCreateParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRuleCreateParams {
	return &MultiAdminVerifyRuleCreateParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRuleCreateParams contains all the parameters to send to the API endpoint

	for the multi admin verify rule create operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRuleCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.MultiAdminVerifyRule

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRuleCreateParams) WithDefaults() *MultiAdminVerifyRuleCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRuleCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := MultiAdminVerifyRuleCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRuleCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) WithContext(ctx context.Context) *MultiAdminVerifyRuleCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRuleCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) WithInfo(info *models.MultiAdminVerifyRule) *MultiAdminVerifyRuleCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) SetInfo(info *models.MultiAdminVerifyRule) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) WithReturnRecords(returnRecords *bool) *MultiAdminVerifyRuleCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the multi admin verify rule create params
func (o *MultiAdminVerifyRuleCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRuleCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
