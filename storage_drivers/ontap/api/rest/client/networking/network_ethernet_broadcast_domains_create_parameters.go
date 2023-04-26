// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewNetworkEthernetBroadcastDomainsCreateParams creates a new NetworkEthernetBroadcastDomainsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkEthernetBroadcastDomainsCreateParams() *NetworkEthernetBroadcastDomainsCreateParams {
	return &NetworkEthernetBroadcastDomainsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkEthernetBroadcastDomainsCreateParamsWithTimeout creates a new NetworkEthernetBroadcastDomainsCreateParams object
// with the ability to set a timeout on a request.
func NewNetworkEthernetBroadcastDomainsCreateParamsWithTimeout(timeout time.Duration) *NetworkEthernetBroadcastDomainsCreateParams {
	return &NetworkEthernetBroadcastDomainsCreateParams{
		timeout: timeout,
	}
}

// NewNetworkEthernetBroadcastDomainsCreateParamsWithContext creates a new NetworkEthernetBroadcastDomainsCreateParams object
// with the ability to set a context for a request.
func NewNetworkEthernetBroadcastDomainsCreateParamsWithContext(ctx context.Context) *NetworkEthernetBroadcastDomainsCreateParams {
	return &NetworkEthernetBroadcastDomainsCreateParams{
		Context: ctx,
	}
}

// NewNetworkEthernetBroadcastDomainsCreateParamsWithHTTPClient creates a new NetworkEthernetBroadcastDomainsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkEthernetBroadcastDomainsCreateParamsWithHTTPClient(client *http.Client) *NetworkEthernetBroadcastDomainsCreateParams {
	return &NetworkEthernetBroadcastDomainsCreateParams{
		HTTPClient: client,
	}
}

/*
NetworkEthernetBroadcastDomainsCreateParams contains all the parameters to send to the API endpoint

	for the network ethernet broadcast domains create operation.

	Typically these are written to a http.Request.
*/
type NetworkEthernetBroadcastDomainsCreateParams struct {

	/* Info.

	   Broadcast_domain parameters
	*/
	Info *models.BroadcastDomain

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ethernet broadcast domains create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithDefaults() *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ethernet broadcast domains create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NetworkEthernetBroadcastDomainsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithTimeout(timeout time.Duration) *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithContext(ctx context.Context) *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithHTTPClient(client *http.Client) *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithInfo(info *models.BroadcastDomain) *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetInfo(info *models.BroadcastDomain) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) WithReturnRecords(returnRecords *bool) *NetworkEthernetBroadcastDomainsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ethernet broadcast domains create params
func (o *NetworkEthernetBroadcastDomainsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkEthernetBroadcastDomainsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
