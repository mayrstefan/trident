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
)

// NewFabricCollectionGetParams creates a new FabricCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFabricCollectionGetParams() *FabricCollectionGetParams {
	return &FabricCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFabricCollectionGetParamsWithTimeout creates a new FabricCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFabricCollectionGetParamsWithTimeout(timeout time.Duration) *FabricCollectionGetParams {
	return &FabricCollectionGetParams{
		timeout: timeout,
	}
}

// NewFabricCollectionGetParamsWithContext creates a new FabricCollectionGetParams object
// with the ability to set a context for a request.
func NewFabricCollectionGetParamsWithContext(ctx context.Context) *FabricCollectionGetParams {
	return &FabricCollectionGetParams{
		Context: ctx,
	}
}

// NewFabricCollectionGetParamsWithHTTPClient creates a new FabricCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFabricCollectionGetParamsWithHTTPClient(client *http.Client) *FabricCollectionGetParams {
	return &FabricCollectionGetParams{
		HTTPClient: client,
	}
}

/*
FabricCollectionGetParams contains all the parameters to send to the API endpoint

	for the fabric collection get operation.

	Typically these are written to a http.Request.
*/
type FabricCollectionGetParams struct {

	/* CacheAge.

	   Filter by cache.age
	*/
	CacheAge *string

	/* CacheIsCurrent.

	   Filter by cache.is_current
	*/
	CacheIsCurrent *bool

	/* CacheMaximumAge.

	   The maximum age of data in the Fibre Channel fabric cache before it should be refreshed from the fabric. The default is 15 minutes.

	   Format: iso8601
	   Default: "15 minutes"
	*/
	CacheMaximumAge *string

	/* CacheUpdateTime.

	   Filter by cache.update_time
	*/
	CacheUpdateTime *string

	/* ConnectionsClusterPortName.

	   Filter by connections.cluster_port.name
	*/
	ConnectionsClusterPortName *string

	/* ConnectionsClusterPortNodeName.

	   Filter by connections.cluster_port.node.name
	*/
	ConnectionsClusterPortNodeName *string

	/* ConnectionsClusterPortUUID.

	   Filter by connections.cluster_port.uuid
	*/
	ConnectionsClusterPortUUID *string

	/* ConnectionsClusterPortWwpn.

	   Filter by connections.cluster_port.wwpn
	*/
	ConnectionsClusterPortWwpn *string

	/* ConnectionsSwitchPortWwpn.

	   Filter by connections.switch.port.wwpn
	*/
	ConnectionsSwitchPortWwpn *string

	/* ConnectionsSwitchWwn.

	   Filter by connections.switch.wwn
	*/
	ConnectionsSwitchWwn *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* ZonesetName.

	   Filter by zoneset.name
	*/
	ZonesetName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fabric collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FabricCollectionGetParams) WithDefaults() *FabricCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fabric collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FabricCollectionGetParams) SetDefaults() {
	var (
		cacheMaximumAgeDefault = string("15 minutes")

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := FabricCollectionGetParams{
		CacheMaximumAge: &cacheMaximumAgeDefault,
		ReturnRecords:   &returnRecordsDefault,
		ReturnTimeout:   &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fabric collection get params
func (o *FabricCollectionGetParams) WithTimeout(timeout time.Duration) *FabricCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fabric collection get params
func (o *FabricCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fabric collection get params
func (o *FabricCollectionGetParams) WithContext(ctx context.Context) *FabricCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fabric collection get params
func (o *FabricCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fabric collection get params
func (o *FabricCollectionGetParams) WithHTTPClient(client *http.Client) *FabricCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fabric collection get params
func (o *FabricCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheAge adds the cacheAge to the fabric collection get params
func (o *FabricCollectionGetParams) WithCacheAge(cacheAge *string) *FabricCollectionGetParams {
	o.SetCacheAge(cacheAge)
	return o
}

// SetCacheAge adds the cacheAge to the fabric collection get params
func (o *FabricCollectionGetParams) SetCacheAge(cacheAge *string) {
	o.CacheAge = cacheAge
}

// WithCacheIsCurrent adds the cacheIsCurrent to the fabric collection get params
func (o *FabricCollectionGetParams) WithCacheIsCurrent(cacheIsCurrent *bool) *FabricCollectionGetParams {
	o.SetCacheIsCurrent(cacheIsCurrent)
	return o
}

// SetCacheIsCurrent adds the cacheIsCurrent to the fabric collection get params
func (o *FabricCollectionGetParams) SetCacheIsCurrent(cacheIsCurrent *bool) {
	o.CacheIsCurrent = cacheIsCurrent
}

// WithCacheMaximumAge adds the cacheMaximumAge to the fabric collection get params
func (o *FabricCollectionGetParams) WithCacheMaximumAge(cacheMaximumAge *string) *FabricCollectionGetParams {
	o.SetCacheMaximumAge(cacheMaximumAge)
	return o
}

// SetCacheMaximumAge adds the cacheMaximumAge to the fabric collection get params
func (o *FabricCollectionGetParams) SetCacheMaximumAge(cacheMaximumAge *string) {
	o.CacheMaximumAge = cacheMaximumAge
}

// WithCacheUpdateTime adds the cacheUpdateTime to the fabric collection get params
func (o *FabricCollectionGetParams) WithCacheUpdateTime(cacheUpdateTime *string) *FabricCollectionGetParams {
	o.SetCacheUpdateTime(cacheUpdateTime)
	return o
}

// SetCacheUpdateTime adds the cacheUpdateTime to the fabric collection get params
func (o *FabricCollectionGetParams) SetCacheUpdateTime(cacheUpdateTime *string) {
	o.CacheUpdateTime = cacheUpdateTime
}

// WithConnectionsClusterPortName adds the connectionsClusterPortName to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsClusterPortName(connectionsClusterPortName *string) *FabricCollectionGetParams {
	o.SetConnectionsClusterPortName(connectionsClusterPortName)
	return o
}

// SetConnectionsClusterPortName adds the connectionsClusterPortName to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsClusterPortName(connectionsClusterPortName *string) {
	o.ConnectionsClusterPortName = connectionsClusterPortName
}

// WithConnectionsClusterPortNodeName adds the connectionsClusterPortNodeName to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsClusterPortNodeName(connectionsClusterPortNodeName *string) *FabricCollectionGetParams {
	o.SetConnectionsClusterPortNodeName(connectionsClusterPortNodeName)
	return o
}

// SetConnectionsClusterPortNodeName adds the connectionsClusterPortNodeName to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsClusterPortNodeName(connectionsClusterPortNodeName *string) {
	o.ConnectionsClusterPortNodeName = connectionsClusterPortNodeName
}

// WithConnectionsClusterPortUUID adds the connectionsClusterPortUUID to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsClusterPortUUID(connectionsClusterPortUUID *string) *FabricCollectionGetParams {
	o.SetConnectionsClusterPortUUID(connectionsClusterPortUUID)
	return o
}

// SetConnectionsClusterPortUUID adds the connectionsClusterPortUuid to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsClusterPortUUID(connectionsClusterPortUUID *string) {
	o.ConnectionsClusterPortUUID = connectionsClusterPortUUID
}

// WithConnectionsClusterPortWwpn adds the connectionsClusterPortWwpn to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsClusterPortWwpn(connectionsClusterPortWwpn *string) *FabricCollectionGetParams {
	o.SetConnectionsClusterPortWwpn(connectionsClusterPortWwpn)
	return o
}

// SetConnectionsClusterPortWwpn adds the connectionsClusterPortWwpn to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsClusterPortWwpn(connectionsClusterPortWwpn *string) {
	o.ConnectionsClusterPortWwpn = connectionsClusterPortWwpn
}

// WithConnectionsSwitchPortWwpn adds the connectionsSwitchPortWwpn to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsSwitchPortWwpn(connectionsSwitchPortWwpn *string) *FabricCollectionGetParams {
	o.SetConnectionsSwitchPortWwpn(connectionsSwitchPortWwpn)
	return o
}

// SetConnectionsSwitchPortWwpn adds the connectionsSwitchPortWwpn to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsSwitchPortWwpn(connectionsSwitchPortWwpn *string) {
	o.ConnectionsSwitchPortWwpn = connectionsSwitchPortWwpn
}

// WithConnectionsSwitchWwn adds the connectionsSwitchWwn to the fabric collection get params
func (o *FabricCollectionGetParams) WithConnectionsSwitchWwn(connectionsSwitchWwn *string) *FabricCollectionGetParams {
	o.SetConnectionsSwitchWwn(connectionsSwitchWwn)
	return o
}

// SetConnectionsSwitchWwn adds the connectionsSwitchWwn to the fabric collection get params
func (o *FabricCollectionGetParams) SetConnectionsSwitchWwn(connectionsSwitchWwn *string) {
	o.ConnectionsSwitchWwn = connectionsSwitchWwn
}

// WithFields adds the fields to the fabric collection get params
func (o *FabricCollectionGetParams) WithFields(fields []string) *FabricCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the fabric collection get params
func (o *FabricCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the fabric collection get params
func (o *FabricCollectionGetParams) WithMaxRecords(maxRecords *int64) *FabricCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the fabric collection get params
func (o *FabricCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the fabric collection get params
func (o *FabricCollectionGetParams) WithName(name *string) *FabricCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the fabric collection get params
func (o *FabricCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the fabric collection get params
func (o *FabricCollectionGetParams) WithOrderBy(orderBy []string) *FabricCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the fabric collection get params
func (o *FabricCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the fabric collection get params
func (o *FabricCollectionGetParams) WithReturnRecords(returnRecords *bool) *FabricCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fabric collection get params
func (o *FabricCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the fabric collection get params
func (o *FabricCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *FabricCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the fabric collection get params
func (o *FabricCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithZonesetName adds the zonesetName to the fabric collection get params
func (o *FabricCollectionGetParams) WithZonesetName(zonesetName *string) *FabricCollectionGetParams {
	o.SetZonesetName(zonesetName)
	return o
}

// SetZonesetName adds the zonesetName to the fabric collection get params
func (o *FabricCollectionGetParams) SetZonesetName(zonesetName *string) {
	o.ZonesetName = zonesetName
}

// WriteToRequest writes these params to a swagger request
func (o *FabricCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CacheAge != nil {

		// query param cache.age
		var qrCacheAge string

		if o.CacheAge != nil {
			qrCacheAge = *o.CacheAge
		}
		qCacheAge := qrCacheAge
		if qCacheAge != "" {

			if err := r.SetQueryParam("cache.age", qCacheAge); err != nil {
				return err
			}
		}
	}

	if o.CacheIsCurrent != nil {

		// query param cache.is_current
		var qrCacheIsCurrent bool

		if o.CacheIsCurrent != nil {
			qrCacheIsCurrent = *o.CacheIsCurrent
		}
		qCacheIsCurrent := swag.FormatBool(qrCacheIsCurrent)
		if qCacheIsCurrent != "" {

			if err := r.SetQueryParam("cache.is_current", qCacheIsCurrent); err != nil {
				return err
			}
		}
	}

	if o.CacheMaximumAge != nil {

		// query param cache.maximum_age
		var qrCacheMaximumAge string

		if o.CacheMaximumAge != nil {
			qrCacheMaximumAge = *o.CacheMaximumAge
		}
		qCacheMaximumAge := qrCacheMaximumAge
		if qCacheMaximumAge != "" {

			if err := r.SetQueryParam("cache.maximum_age", qCacheMaximumAge); err != nil {
				return err
			}
		}
	}

	if o.CacheUpdateTime != nil {

		// query param cache.update_time
		var qrCacheUpdateTime string

		if o.CacheUpdateTime != nil {
			qrCacheUpdateTime = *o.CacheUpdateTime
		}
		qCacheUpdateTime := qrCacheUpdateTime
		if qCacheUpdateTime != "" {

			if err := r.SetQueryParam("cache.update_time", qCacheUpdateTime); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsClusterPortName != nil {

		// query param connections.cluster_port.name
		var qrConnectionsClusterPortName string

		if o.ConnectionsClusterPortName != nil {
			qrConnectionsClusterPortName = *o.ConnectionsClusterPortName
		}
		qConnectionsClusterPortName := qrConnectionsClusterPortName
		if qConnectionsClusterPortName != "" {

			if err := r.SetQueryParam("connections.cluster_port.name", qConnectionsClusterPortName); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsClusterPortNodeName != nil {

		// query param connections.cluster_port.node.name
		var qrConnectionsClusterPortNodeName string

		if o.ConnectionsClusterPortNodeName != nil {
			qrConnectionsClusterPortNodeName = *o.ConnectionsClusterPortNodeName
		}
		qConnectionsClusterPortNodeName := qrConnectionsClusterPortNodeName
		if qConnectionsClusterPortNodeName != "" {

			if err := r.SetQueryParam("connections.cluster_port.node.name", qConnectionsClusterPortNodeName); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsClusterPortUUID != nil {

		// query param connections.cluster_port.uuid
		var qrConnectionsClusterPortUUID string

		if o.ConnectionsClusterPortUUID != nil {
			qrConnectionsClusterPortUUID = *o.ConnectionsClusterPortUUID
		}
		qConnectionsClusterPortUUID := qrConnectionsClusterPortUUID
		if qConnectionsClusterPortUUID != "" {

			if err := r.SetQueryParam("connections.cluster_port.uuid", qConnectionsClusterPortUUID); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsClusterPortWwpn != nil {

		// query param connections.cluster_port.wwpn
		var qrConnectionsClusterPortWwpn string

		if o.ConnectionsClusterPortWwpn != nil {
			qrConnectionsClusterPortWwpn = *o.ConnectionsClusterPortWwpn
		}
		qConnectionsClusterPortWwpn := qrConnectionsClusterPortWwpn
		if qConnectionsClusterPortWwpn != "" {

			if err := r.SetQueryParam("connections.cluster_port.wwpn", qConnectionsClusterPortWwpn); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsSwitchPortWwpn != nil {

		// query param connections.switch.port.wwpn
		var qrConnectionsSwitchPortWwpn string

		if o.ConnectionsSwitchPortWwpn != nil {
			qrConnectionsSwitchPortWwpn = *o.ConnectionsSwitchPortWwpn
		}
		qConnectionsSwitchPortWwpn := qrConnectionsSwitchPortWwpn
		if qConnectionsSwitchPortWwpn != "" {

			if err := r.SetQueryParam("connections.switch.port.wwpn", qConnectionsSwitchPortWwpn); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsSwitchWwn != nil {

		// query param connections.switch.wwn
		var qrConnectionsSwitchWwn string

		if o.ConnectionsSwitchWwn != nil {
			qrConnectionsSwitchWwn = *o.ConnectionsSwitchWwn
		}
		qConnectionsSwitchWwn := qrConnectionsSwitchWwn
		if qConnectionsSwitchWwn != "" {

			if err := r.SetQueryParam("connections.switch.wwn", qConnectionsSwitchWwn); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.ZonesetName != nil {

		// query param zoneset.name
		var qrZonesetName string

		if o.ZonesetName != nil {
			qrZonesetName = *o.ZonesetName
		}
		qZonesetName := qrZonesetName
		if qZonesetName != "" {

			if err := r.SetQueryParam("zoneset.name", qZonesetName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFabricCollectionGet binds the parameter fields
func (o *FabricCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamFabricCollectionGet binds the parameter order_by
func (o *FabricCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
