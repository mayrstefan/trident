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

// NewSnapmirrorRelationshipTransfersGetParams creates a new SnapmirrorRelationshipTransfersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorRelationshipTransfersGetParams() *SnapmirrorRelationshipTransfersGetParams {
	return &SnapmirrorRelationshipTransfersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorRelationshipTransfersGetParamsWithTimeout creates a new SnapmirrorRelationshipTransfersGetParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorRelationshipTransfersGetParamsWithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransfersGetParams {
	return &SnapmirrorRelationshipTransfersGetParams{
		timeout: timeout,
	}
}

// NewSnapmirrorRelationshipTransfersGetParamsWithContext creates a new SnapmirrorRelationshipTransfersGetParams object
// with the ability to set a context for a request.
func NewSnapmirrorRelationshipTransfersGetParamsWithContext(ctx context.Context) *SnapmirrorRelationshipTransfersGetParams {
	return &SnapmirrorRelationshipTransfersGetParams{
		Context: ctx,
	}
}

// NewSnapmirrorRelationshipTransfersGetParamsWithHTTPClient creates a new SnapmirrorRelationshipTransfersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorRelationshipTransfersGetParamsWithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransfersGetParams {
	return &SnapmirrorRelationshipTransfersGetParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorRelationshipTransfersGetParams contains all the parameters to send to the API endpoint

	for the snapmirror relationship transfers get operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorRelationshipTransfersGetParams struct {

	/* BytesTransferred.

	   Filter by bytes_transferred
	*/
	BytesTransferred *int64

	/* CheckpointSize.

	   Filter by checkpoint_size
	*/
	CheckpointSize *int64

	/* EndTime.

	   Filter by end_time
	*/
	EndTime *string

	/* ErrorInfoCode.

	   Filter by error_info.code
	*/
	ErrorInfoCode *int64

	/* ErrorInfoMessage.

	   Filter by error_info.message
	*/
	ErrorInfoMessage *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* RelationshipDestinationClusterName.

	   Filter by relationship.destination.cluster.name
	*/
	RelationshipDestinationClusterName *string

	/* RelationshipDestinationClusterUUID.

	   Filter by relationship.destination.cluster.uuid
	*/
	RelationshipDestinationClusterUUID *string

	/* RelationshipDestinationConsistencyGroupVolumesName.

	   Filter by relationship.destination.consistency_group_volumes.name
	*/
	RelationshipDestinationConsistencyGroupVolumesName *string

	/* RelationshipDestinationPath.

	   Filter by relationship.destination.path
	*/
	RelationshipDestinationPath *string

	/* RelationshipDestinationSvmName.

	   Filter by relationship.destination.svm.name
	*/
	RelationshipDestinationSvmName *string

	/* RelationshipDestinationSvmUUID.

	   Filter by relationship.destination.svm.uuid
	*/
	RelationshipDestinationSvmUUID *string

	/* RelationshipDestinationUUID.

	   Filter by relationship.destination.uuid
	*/
	RelationshipDestinationUUID *string

	/* RelationshipRestore.

	   Filter by relationship.restore
	*/
	RelationshipRestore *bool

	/* RelationshipUUID.

	   Relationship UUID
	*/
	RelationshipUUID string

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

	/* Snapshot.

	   Filter by snapshot
	*/
	Snapshot *string

	/* State.

	   Filter by state
	*/
	State *string

	/* Throttle.

	   Filter by throttle
	*/
	Throttle *int64

	/* TotalDuration.

	   Filter by total_duration
	*/
	TotalDuration *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror relationship transfers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransfersGetParams) WithDefaults() *SnapmirrorRelationshipTransfersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror relationship transfers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorRelationshipTransfersGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnapmirrorRelationshipTransfersGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithTimeout(timeout time.Duration) *SnapmirrorRelationshipTransfersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithContext(ctx context.Context) *SnapmirrorRelationshipTransfersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithHTTPClient(client *http.Client) *SnapmirrorRelationshipTransfersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBytesTransferred adds the bytesTransferred to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithBytesTransferred(bytesTransferred *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetBytesTransferred(bytesTransferred)
	return o
}

// SetBytesTransferred adds the bytesTransferred to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetBytesTransferred(bytesTransferred *int64) {
	o.BytesTransferred = bytesTransferred
}

// WithCheckpointSize adds the checkpointSize to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithCheckpointSize(checkpointSize *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetCheckpointSize(checkpointSize)
	return o
}

// SetCheckpointSize adds the checkpointSize to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetCheckpointSize(checkpointSize *int64) {
	o.CheckpointSize = checkpointSize
}

// WithEndTime adds the endTime to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithEndTime(endTime *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithErrorInfoCode adds the errorInfoCode to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithErrorInfoCode(errorInfoCode *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetErrorInfoCode(errorInfoCode)
	return o
}

// SetErrorInfoCode adds the errorInfoCode to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetErrorInfoCode(errorInfoCode *int64) {
	o.ErrorInfoCode = errorInfoCode
}

// WithErrorInfoMessage adds the errorInfoMessage to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithErrorInfoMessage(errorInfoMessage *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetErrorInfoMessage(errorInfoMessage)
	return o
}

// SetErrorInfoMessage adds the errorInfoMessage to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetErrorInfoMessage(errorInfoMessage *string) {
	o.ErrorInfoMessage = errorInfoMessage
}

// WithFields adds the fields to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithFields(fields []string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithMaxRecords(maxRecords *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithOrderBy(orderBy []string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithRelationshipDestinationClusterName adds the relationshipDestinationClusterName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationClusterName(relationshipDestinationClusterName *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationClusterName(relationshipDestinationClusterName)
	return o
}

// SetRelationshipDestinationClusterName adds the relationshipDestinationClusterName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationClusterName(relationshipDestinationClusterName *string) {
	o.RelationshipDestinationClusterName = relationshipDestinationClusterName
}

// WithRelationshipDestinationClusterUUID adds the relationshipDestinationClusterUUID to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationClusterUUID(relationshipDestinationClusterUUID *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationClusterUUID(relationshipDestinationClusterUUID)
	return o
}

// SetRelationshipDestinationClusterUUID adds the relationshipDestinationClusterUuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationClusterUUID(relationshipDestinationClusterUUID *string) {
	o.RelationshipDestinationClusterUUID = relationshipDestinationClusterUUID
}

// WithRelationshipDestinationConsistencyGroupVolumesName adds the relationshipDestinationConsistencyGroupVolumesName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationConsistencyGroupVolumesName(relationshipDestinationConsistencyGroupVolumesName *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationConsistencyGroupVolumesName(relationshipDestinationConsistencyGroupVolumesName)
	return o
}

// SetRelationshipDestinationConsistencyGroupVolumesName adds the relationshipDestinationConsistencyGroupVolumesName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationConsistencyGroupVolumesName(relationshipDestinationConsistencyGroupVolumesName *string) {
	o.RelationshipDestinationConsistencyGroupVolumesName = relationshipDestinationConsistencyGroupVolumesName
}

// WithRelationshipDestinationPath adds the relationshipDestinationPath to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationPath(relationshipDestinationPath *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationPath(relationshipDestinationPath)
	return o
}

// SetRelationshipDestinationPath adds the relationshipDestinationPath to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationPath(relationshipDestinationPath *string) {
	o.RelationshipDestinationPath = relationshipDestinationPath
}

// WithRelationshipDestinationSvmName adds the relationshipDestinationSvmName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationSvmName(relationshipDestinationSvmName *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationSvmName(relationshipDestinationSvmName)
	return o
}

// SetRelationshipDestinationSvmName adds the relationshipDestinationSvmName to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationSvmName(relationshipDestinationSvmName *string) {
	o.RelationshipDestinationSvmName = relationshipDestinationSvmName
}

// WithRelationshipDestinationSvmUUID adds the relationshipDestinationSvmUUID to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationSvmUUID(relationshipDestinationSvmUUID *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationSvmUUID(relationshipDestinationSvmUUID)
	return o
}

// SetRelationshipDestinationSvmUUID adds the relationshipDestinationSvmUuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationSvmUUID(relationshipDestinationSvmUUID *string) {
	o.RelationshipDestinationSvmUUID = relationshipDestinationSvmUUID
}

// WithRelationshipDestinationUUID adds the relationshipDestinationUUID to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipDestinationUUID(relationshipDestinationUUID *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipDestinationUUID(relationshipDestinationUUID)
	return o
}

// SetRelationshipDestinationUUID adds the relationshipDestinationUuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipDestinationUUID(relationshipDestinationUUID *string) {
	o.RelationshipDestinationUUID = relationshipDestinationUUID
}

// WithRelationshipRestore adds the relationshipRestore to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipRestore(relationshipRestore *bool) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipRestore(relationshipRestore)
	return o
}

// SetRelationshipRestore adds the relationshipRestore to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipRestore(relationshipRestore *bool) {
	o.RelationshipRestore = relationshipRestore
}

// WithRelationshipUUID adds the relationshipUUID to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithRelationshipUUID(relationshipUUID string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetRelationshipUUID(relationshipUUID)
	return o
}

// SetRelationshipUUID adds the relationshipUuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetRelationshipUUID(relationshipUUID string) {
	o.RelationshipUUID = relationshipUUID
}

// WithReturnRecords adds the returnRecords to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithReturnRecords(returnRecords *bool) *SnapmirrorRelationshipTransfersGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithReturnTimeout(returnTimeout *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSnapshot adds the snapshot to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithSnapshot(snapshot *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetSnapshot(snapshot *string) {
	o.Snapshot = snapshot
}

// WithState adds the state to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithState(state *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetState(state *string) {
	o.State = state
}

// WithThrottle adds the throttle to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithThrottle(throttle *int64) *SnapmirrorRelationshipTransfersGetParams {
	o.SetThrottle(throttle)
	return o
}

// SetThrottle adds the throttle to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetThrottle(throttle *int64) {
	o.Throttle = throttle
}

// WithTotalDuration adds the totalDuration to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithTotalDuration(totalDuration *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetTotalDuration(totalDuration)
	return o
}

// SetTotalDuration adds the totalDuration to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetTotalDuration(totalDuration *string) {
	o.TotalDuration = totalDuration
}

// WithUUID adds the uuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) WithUUID(uuid *string) *SnapmirrorRelationshipTransfersGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapmirror relationship transfers get params
func (o *SnapmirrorRelationshipTransfersGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorRelationshipTransfersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BytesTransferred != nil {

		// query param bytes_transferred
		var qrBytesTransferred int64

		if o.BytesTransferred != nil {
			qrBytesTransferred = *o.BytesTransferred
		}
		qBytesTransferred := swag.FormatInt64(qrBytesTransferred)
		if qBytesTransferred != "" {

			if err := r.SetQueryParam("bytes_transferred", qBytesTransferred); err != nil {
				return err
			}
		}
	}

	if o.CheckpointSize != nil {

		// query param checkpoint_size
		var qrCheckpointSize int64

		if o.CheckpointSize != nil {
			qrCheckpointSize = *o.CheckpointSize
		}
		qCheckpointSize := swag.FormatInt64(qrCheckpointSize)
		if qCheckpointSize != "" {

			if err := r.SetQueryParam("checkpoint_size", qCheckpointSize); err != nil {
				return err
			}
		}
	}

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.ErrorInfoCode != nil {

		// query param error_info.code
		var qrErrorInfoCode int64

		if o.ErrorInfoCode != nil {
			qrErrorInfoCode = *o.ErrorInfoCode
		}
		qErrorInfoCode := swag.FormatInt64(qrErrorInfoCode)
		if qErrorInfoCode != "" {

			if err := r.SetQueryParam("error_info.code", qErrorInfoCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorInfoMessage != nil {

		// query param error_info.message
		var qrErrorInfoMessage string

		if o.ErrorInfoMessage != nil {
			qrErrorInfoMessage = *o.ErrorInfoMessage
		}
		qErrorInfoMessage := qrErrorInfoMessage
		if qErrorInfoMessage != "" {

			if err := r.SetQueryParam("error_info.message", qErrorInfoMessage); err != nil {
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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.RelationshipDestinationClusterName != nil {

		// query param relationship.destination.cluster.name
		var qrRelationshipDestinationClusterName string

		if o.RelationshipDestinationClusterName != nil {
			qrRelationshipDestinationClusterName = *o.RelationshipDestinationClusterName
		}
		qRelationshipDestinationClusterName := qrRelationshipDestinationClusterName
		if qRelationshipDestinationClusterName != "" {

			if err := r.SetQueryParam("relationship.destination.cluster.name", qRelationshipDestinationClusterName); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationClusterUUID != nil {

		// query param relationship.destination.cluster.uuid
		var qrRelationshipDestinationClusterUUID string

		if o.RelationshipDestinationClusterUUID != nil {
			qrRelationshipDestinationClusterUUID = *o.RelationshipDestinationClusterUUID
		}
		qRelationshipDestinationClusterUUID := qrRelationshipDestinationClusterUUID
		if qRelationshipDestinationClusterUUID != "" {

			if err := r.SetQueryParam("relationship.destination.cluster.uuid", qRelationshipDestinationClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationConsistencyGroupVolumesName != nil {

		// query param relationship.destination.consistency_group_volumes.name
		var qrRelationshipDestinationConsistencyGroupVolumesName string

		if o.RelationshipDestinationConsistencyGroupVolumesName != nil {
			qrRelationshipDestinationConsistencyGroupVolumesName = *o.RelationshipDestinationConsistencyGroupVolumesName
		}
		qRelationshipDestinationConsistencyGroupVolumesName := qrRelationshipDestinationConsistencyGroupVolumesName
		if qRelationshipDestinationConsistencyGroupVolumesName != "" {

			if err := r.SetQueryParam("relationship.destination.consistency_group_volumes.name", qRelationshipDestinationConsistencyGroupVolumesName); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationPath != nil {

		// query param relationship.destination.path
		var qrRelationshipDestinationPath string

		if o.RelationshipDestinationPath != nil {
			qrRelationshipDestinationPath = *o.RelationshipDestinationPath
		}
		qRelationshipDestinationPath := qrRelationshipDestinationPath
		if qRelationshipDestinationPath != "" {

			if err := r.SetQueryParam("relationship.destination.path", qRelationshipDestinationPath); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationSvmName != nil {

		// query param relationship.destination.svm.name
		var qrRelationshipDestinationSvmName string

		if o.RelationshipDestinationSvmName != nil {
			qrRelationshipDestinationSvmName = *o.RelationshipDestinationSvmName
		}
		qRelationshipDestinationSvmName := qrRelationshipDestinationSvmName
		if qRelationshipDestinationSvmName != "" {

			if err := r.SetQueryParam("relationship.destination.svm.name", qRelationshipDestinationSvmName); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationSvmUUID != nil {

		// query param relationship.destination.svm.uuid
		var qrRelationshipDestinationSvmUUID string

		if o.RelationshipDestinationSvmUUID != nil {
			qrRelationshipDestinationSvmUUID = *o.RelationshipDestinationSvmUUID
		}
		qRelationshipDestinationSvmUUID := qrRelationshipDestinationSvmUUID
		if qRelationshipDestinationSvmUUID != "" {

			if err := r.SetQueryParam("relationship.destination.svm.uuid", qRelationshipDestinationSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.RelationshipDestinationUUID != nil {

		// query param relationship.destination.uuid
		var qrRelationshipDestinationUUID string

		if o.RelationshipDestinationUUID != nil {
			qrRelationshipDestinationUUID = *o.RelationshipDestinationUUID
		}
		qRelationshipDestinationUUID := qrRelationshipDestinationUUID
		if qRelationshipDestinationUUID != "" {

			if err := r.SetQueryParam("relationship.destination.uuid", qRelationshipDestinationUUID); err != nil {
				return err
			}
		}
	}

	if o.RelationshipRestore != nil {

		// query param relationship.restore
		var qrRelationshipRestore bool

		if o.RelationshipRestore != nil {
			qrRelationshipRestore = *o.RelationshipRestore
		}
		qRelationshipRestore := swag.FormatBool(qrRelationshipRestore)
		if qRelationshipRestore != "" {

			if err := r.SetQueryParam("relationship.restore", qRelationshipRestore); err != nil {
				return err
			}
		}
	}

	// path param relationship.uuid
	if err := r.SetPathParam("relationship.uuid", o.RelationshipUUID); err != nil {
		return err
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

	if o.Snapshot != nil {

		// query param snapshot
		var qrSnapshot string

		if o.Snapshot != nil {
			qrSnapshot = *o.Snapshot
		}
		qSnapshot := qrSnapshot
		if qSnapshot != "" {

			if err := r.SetQueryParam("snapshot", qSnapshot); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.Throttle != nil {

		// query param throttle
		var qrThrottle int64

		if o.Throttle != nil {
			qrThrottle = *o.Throttle
		}
		qThrottle := swag.FormatInt64(qrThrottle)
		if qThrottle != "" {

			if err := r.SetQueryParam("throttle", qThrottle); err != nil {
				return err
			}
		}
	}

	if o.TotalDuration != nil {

		// query param total_duration
		var qrTotalDuration string

		if o.TotalDuration != nil {
			qrTotalDuration = *o.TotalDuration
		}
		qTotalDuration := qrTotalDuration
		if qTotalDuration != "" {

			if err := r.SetQueryParam("total_duration", qTotalDuration); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnapmirrorRelationshipTransfersGet binds the parameter fields
func (o *SnapmirrorRelationshipTransfersGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnapmirrorRelationshipTransfersGet binds the parameter order_by
func (o *SnapmirrorRelationshipTransfersGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
