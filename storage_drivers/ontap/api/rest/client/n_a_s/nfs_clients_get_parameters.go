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

// NewNfsClientsGetParams creates a new NfsClientsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsClientsGetParams() *NfsClientsGetParams {
	return &NfsClientsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsClientsGetParamsWithTimeout creates a new NfsClientsGetParams object
// with the ability to set a timeout on a request.
func NewNfsClientsGetParamsWithTimeout(timeout time.Duration) *NfsClientsGetParams {
	return &NfsClientsGetParams{
		timeout: timeout,
	}
}

// NewNfsClientsGetParamsWithContext creates a new NfsClientsGetParams object
// with the ability to set a context for a request.
func NewNfsClientsGetParamsWithContext(ctx context.Context) *NfsClientsGetParams {
	return &NfsClientsGetParams{
		Context: ctx,
	}
}

// NewNfsClientsGetParamsWithHTTPClient creates a new NfsClientsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsClientsGetParamsWithHTTPClient(client *http.Client) *NfsClientsGetParams {
	return &NfsClientsGetParams{
		HTTPClient: client,
	}
}

/*
NfsClientsGetParams contains all the parameters to send to the API endpoint

	for the nfs clients get operation.

	Typically these are written to a http.Request.
*/
type NfsClientsGetParams struct {

	/* ClientIP.

	   Filter by client_ip
	*/
	ClientIP *string

	/* ExportPolicyID.

	   Filter by export_policy.id
	*/
	ExportPolicyID *int64

	/* ExportPolicyName.

	   Filter by export_policy.name
	*/
	ExportPolicyName *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IdleDuration.

	   Filter by idle_duration
	*/
	IdleDuration *string

	/* LocalRequestCount.

	   Filter by local_request_count
	*/
	LocalRequestCount *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Protocol.

	   Filter by protocol
	*/
	Protocol *string

	/* RemoteRequestCount.

	   Filter by remote_request_count
	*/
	RemoteRequestCount *int64

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

	/* ServerIP.

	   Filter by server_ip
	*/
	ServerIP *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* TrunkingEnabled.

	   Filter by trunking_enabled
	*/
	TrunkingEnabled *bool

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeName *string

	/* VolumeUUID.

	   Filter by volume.uuid
	*/
	VolumeUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs clients get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsGetParams) WithDefaults() *NfsClientsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs clients get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsClientsGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NfsClientsGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nfs clients get params
func (o *NfsClientsGetParams) WithTimeout(timeout time.Duration) *NfsClientsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs clients get params
func (o *NfsClientsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs clients get params
func (o *NfsClientsGetParams) WithContext(ctx context.Context) *NfsClientsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs clients get params
func (o *NfsClientsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs clients get params
func (o *NfsClientsGetParams) WithHTTPClient(client *http.Client) *NfsClientsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs clients get params
func (o *NfsClientsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIP adds the clientIP to the nfs clients get params
func (o *NfsClientsGetParams) WithClientIP(clientIP *string) *NfsClientsGetParams {
	o.SetClientIP(clientIP)
	return o
}

// SetClientIP adds the clientIp to the nfs clients get params
func (o *NfsClientsGetParams) SetClientIP(clientIP *string) {
	o.ClientIP = clientIP
}

// WithExportPolicyID adds the exportPolicyID to the nfs clients get params
func (o *NfsClientsGetParams) WithExportPolicyID(exportPolicyID *int64) *NfsClientsGetParams {
	o.SetExportPolicyID(exportPolicyID)
	return o
}

// SetExportPolicyID adds the exportPolicyId to the nfs clients get params
func (o *NfsClientsGetParams) SetExportPolicyID(exportPolicyID *int64) {
	o.ExportPolicyID = exportPolicyID
}

// WithExportPolicyName adds the exportPolicyName to the nfs clients get params
func (o *NfsClientsGetParams) WithExportPolicyName(exportPolicyName *string) *NfsClientsGetParams {
	o.SetExportPolicyName(exportPolicyName)
	return o
}

// SetExportPolicyName adds the exportPolicyName to the nfs clients get params
func (o *NfsClientsGetParams) SetExportPolicyName(exportPolicyName *string) {
	o.ExportPolicyName = exportPolicyName
}

// WithFields adds the fields to the nfs clients get params
func (o *NfsClientsGetParams) WithFields(fields []string) *NfsClientsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nfs clients get params
func (o *NfsClientsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIdleDuration adds the idleDuration to the nfs clients get params
func (o *NfsClientsGetParams) WithIdleDuration(idleDuration *string) *NfsClientsGetParams {
	o.SetIdleDuration(idleDuration)
	return o
}

// SetIdleDuration adds the idleDuration to the nfs clients get params
func (o *NfsClientsGetParams) SetIdleDuration(idleDuration *string) {
	o.IdleDuration = idleDuration
}

// WithLocalRequestCount adds the localRequestCount to the nfs clients get params
func (o *NfsClientsGetParams) WithLocalRequestCount(localRequestCount *int64) *NfsClientsGetParams {
	o.SetLocalRequestCount(localRequestCount)
	return o
}

// SetLocalRequestCount adds the localRequestCount to the nfs clients get params
func (o *NfsClientsGetParams) SetLocalRequestCount(localRequestCount *int64) {
	o.LocalRequestCount = localRequestCount
}

// WithMaxRecords adds the maxRecords to the nfs clients get params
func (o *NfsClientsGetParams) WithMaxRecords(maxRecords *int64) *NfsClientsGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the nfs clients get params
func (o *NfsClientsGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeName adds the nodeName to the nfs clients get params
func (o *NfsClientsGetParams) WithNodeName(nodeName *string) *NfsClientsGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the nfs clients get params
func (o *NfsClientsGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the nfs clients get params
func (o *NfsClientsGetParams) WithNodeUUID(nodeUUID *string) *NfsClientsGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the nfs clients get params
func (o *NfsClientsGetParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOrderBy adds the orderBy to the nfs clients get params
func (o *NfsClientsGetParams) WithOrderBy(orderBy []string) *NfsClientsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the nfs clients get params
func (o *NfsClientsGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProtocol adds the protocol to the nfs clients get params
func (o *NfsClientsGetParams) WithProtocol(protocol *string) *NfsClientsGetParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the nfs clients get params
func (o *NfsClientsGetParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithRemoteRequestCount adds the remoteRequestCount to the nfs clients get params
func (o *NfsClientsGetParams) WithRemoteRequestCount(remoteRequestCount *int64) *NfsClientsGetParams {
	o.SetRemoteRequestCount(remoteRequestCount)
	return o
}

// SetRemoteRequestCount adds the remoteRequestCount to the nfs clients get params
func (o *NfsClientsGetParams) SetRemoteRequestCount(remoteRequestCount *int64) {
	o.RemoteRequestCount = remoteRequestCount
}

// WithReturnRecords adds the returnRecords to the nfs clients get params
func (o *NfsClientsGetParams) WithReturnRecords(returnRecords *bool) *NfsClientsGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nfs clients get params
func (o *NfsClientsGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nfs clients get params
func (o *NfsClientsGetParams) WithReturnTimeout(returnTimeout *int64) *NfsClientsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nfs clients get params
func (o *NfsClientsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithServerIP adds the serverIP to the nfs clients get params
func (o *NfsClientsGetParams) WithServerIP(serverIP *string) *NfsClientsGetParams {
	o.SetServerIP(serverIP)
	return o
}

// SetServerIP adds the serverIp to the nfs clients get params
func (o *NfsClientsGetParams) SetServerIP(serverIP *string) {
	o.ServerIP = serverIP
}

// WithSvmName adds the svmName to the nfs clients get params
func (o *NfsClientsGetParams) WithSvmName(svmName *string) *NfsClientsGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the nfs clients get params
func (o *NfsClientsGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the nfs clients get params
func (o *NfsClientsGetParams) WithSvmUUID(svmUUID *string) *NfsClientsGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the nfs clients get params
func (o *NfsClientsGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTrunkingEnabled adds the trunkingEnabled to the nfs clients get params
func (o *NfsClientsGetParams) WithTrunkingEnabled(trunkingEnabled *bool) *NfsClientsGetParams {
	o.SetTrunkingEnabled(trunkingEnabled)
	return o
}

// SetTrunkingEnabled adds the trunkingEnabled to the nfs clients get params
func (o *NfsClientsGetParams) SetTrunkingEnabled(trunkingEnabled *bool) {
	o.TrunkingEnabled = trunkingEnabled
}

// WithVolumeName adds the volumeName to the nfs clients get params
func (o *NfsClientsGetParams) WithVolumeName(volumeName *string) *NfsClientsGetParams {
	o.SetVolumeName(volumeName)
	return o
}

// SetVolumeName adds the volumeName to the nfs clients get params
func (o *NfsClientsGetParams) SetVolumeName(volumeName *string) {
	o.VolumeName = volumeName
}

// WithVolumeUUID adds the volumeUUID to the nfs clients get params
func (o *NfsClientsGetParams) WithVolumeUUID(volumeUUID *string) *NfsClientsGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the nfs clients get params
func (o *NfsClientsGetParams) SetVolumeUUID(volumeUUID *string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NfsClientsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIP != nil {

		// query param client_ip
		var qrClientIP string

		if o.ClientIP != nil {
			qrClientIP = *o.ClientIP
		}
		qClientIP := qrClientIP
		if qClientIP != "" {

			if err := r.SetQueryParam("client_ip", qClientIP); err != nil {
				return err
			}
		}
	}

	if o.ExportPolicyID != nil {

		// query param export_policy.id
		var qrExportPolicyID int64

		if o.ExportPolicyID != nil {
			qrExportPolicyID = *o.ExportPolicyID
		}
		qExportPolicyID := swag.FormatInt64(qrExportPolicyID)
		if qExportPolicyID != "" {

			if err := r.SetQueryParam("export_policy.id", qExportPolicyID); err != nil {
				return err
			}
		}
	}

	if o.ExportPolicyName != nil {

		// query param export_policy.name
		var qrExportPolicyName string

		if o.ExportPolicyName != nil {
			qrExportPolicyName = *o.ExportPolicyName
		}
		qExportPolicyName := qrExportPolicyName
		if qExportPolicyName != "" {

			if err := r.SetQueryParam("export_policy.name", qExportPolicyName); err != nil {
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

	if o.IdleDuration != nil {

		// query param idle_duration
		var qrIdleDuration string

		if o.IdleDuration != nil {
			qrIdleDuration = *o.IdleDuration
		}
		qIdleDuration := qrIdleDuration
		if qIdleDuration != "" {

			if err := r.SetQueryParam("idle_duration", qIdleDuration); err != nil {
				return err
			}
		}
	}

	if o.LocalRequestCount != nil {

		// query param local_request_count
		var qrLocalRequestCount int64

		if o.LocalRequestCount != nil {
			qrLocalRequestCount = *o.LocalRequestCount
		}
		qLocalRequestCount := swag.FormatInt64(qrLocalRequestCount)
		if qLocalRequestCount != "" {

			if err := r.SetQueryParam("local_request_count", qLocalRequestCount); err != nil {
				return err
			}
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

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
		}
	}

	if o.RemoteRequestCount != nil {

		// query param remote_request_count
		var qrRemoteRequestCount int64

		if o.RemoteRequestCount != nil {
			qrRemoteRequestCount = *o.RemoteRequestCount
		}
		qRemoteRequestCount := swag.FormatInt64(qrRemoteRequestCount)
		if qRemoteRequestCount != "" {

			if err := r.SetQueryParam("remote_request_count", qRemoteRequestCount); err != nil {
				return err
			}
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

	if o.ServerIP != nil {

		// query param server_ip
		var qrServerIP string

		if o.ServerIP != nil {
			qrServerIP = *o.ServerIP
		}
		qServerIP := qrServerIP
		if qServerIP != "" {

			if err := r.SetQueryParam("server_ip", qServerIP); err != nil {
				return err
			}
		}
	}

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.TrunkingEnabled != nil {

		// query param trunking_enabled
		var qrTrunkingEnabled bool

		if o.TrunkingEnabled != nil {
			qrTrunkingEnabled = *o.TrunkingEnabled
		}
		qTrunkingEnabled := swag.FormatBool(qrTrunkingEnabled)
		if qTrunkingEnabled != "" {

			if err := r.SetQueryParam("trunking_enabled", qTrunkingEnabled); err != nil {
				return err
			}
		}
	}

	if o.VolumeName != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeName != nil {
			qrVolumeName = *o.VolumeName
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	if o.VolumeUUID != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUID != nil {
			qrVolumeUUID = *o.VolumeUUID
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNfsClientsGet binds the parameter fields
func (o *NfsClientsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNfsClientsGet binds the parameter order_by
func (o *NfsClientsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
