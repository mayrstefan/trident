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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewShadowcopyModifyParams creates a new ShadowcopyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShadowcopyModifyParams() *ShadowcopyModifyParams {
	return &ShadowcopyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShadowcopyModifyParamsWithTimeout creates a new ShadowcopyModifyParams object
// with the ability to set a timeout on a request.
func NewShadowcopyModifyParamsWithTimeout(timeout time.Duration) *ShadowcopyModifyParams {
	return &ShadowcopyModifyParams{
		timeout: timeout,
	}
}

// NewShadowcopyModifyParamsWithContext creates a new ShadowcopyModifyParams object
// with the ability to set a context for a request.
func NewShadowcopyModifyParamsWithContext(ctx context.Context) *ShadowcopyModifyParams {
	return &ShadowcopyModifyParams{
		Context: ctx,
	}
}

// NewShadowcopyModifyParamsWithHTTPClient creates a new ShadowcopyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewShadowcopyModifyParamsWithHTTPClient(client *http.Client) *ShadowcopyModifyParams {
	return &ShadowcopyModifyParams{
		HTTPClient: client,
	}
}

/*
ShadowcopyModifyParams contains all the parameters to send to the API endpoint

	for the shadowcopy modify operation.

	Typically these are written to a http.Request.
*/
type ShadowcopyModifyParams struct {

	/* ClientUUID.

	   client shadowcopy ID
	*/
	ClientUUID string

	/* Info.

	   Info Specification
	*/
	Info *models.Shadowcopy

	/* Restore.

	   Indicates whether to restore a directory
	*/
	Restore *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the shadowcopy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopyModifyParams) WithDefaults() *ShadowcopyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shadowcopy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithTimeout(timeout time.Duration) *ShadowcopyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithContext(ctx context.Context) *ShadowcopyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithHTTPClient(client *http.Client) *ShadowcopyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientUUID adds the clientUUID to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithClientUUID(clientUUID string) *ShadowcopyModifyParams {
	o.SetClientUUID(clientUUID)
	return o
}

// SetClientUUID adds the clientUuid to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetClientUUID(clientUUID string) {
	o.ClientUUID = clientUUID
}

// WithInfo adds the info to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithInfo(info *models.Shadowcopy) *ShadowcopyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetInfo(info *models.Shadowcopy) {
	o.Info = info
}

// WithRestore adds the restore to the shadowcopy modify params
func (o *ShadowcopyModifyParams) WithRestore(restore *bool) *ShadowcopyModifyParams {
	o.SetRestore(restore)
	return o
}

// SetRestore adds the restore to the shadowcopy modify params
func (o *ShadowcopyModifyParams) SetRestore(restore *bool) {
	o.Restore = restore
}

// WriteToRequest writes these params to a swagger request
func (o *ShadowcopyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_uuid
	if err := r.SetPathParam("client_uuid", o.ClientUUID); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.Restore != nil {

		// query param restore
		var qrRestore bool

		if o.Restore != nil {
			qrRestore = *o.Restore
		}
		qRestore := swag.FormatBool(qrRestore)
		if qRestore != "" {

			if err := r.SetQueryParam("restore", qRestore); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
