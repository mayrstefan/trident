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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewConsistencyGroupSnapshotCreateParams creates a new ConsistencyGroupSnapshotCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConsistencyGroupSnapshotCreateParams() *ConsistencyGroupSnapshotCreateParams {
	return &ConsistencyGroupSnapshotCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConsistencyGroupSnapshotCreateParamsWithTimeout creates a new ConsistencyGroupSnapshotCreateParams object
// with the ability to set a timeout on a request.
func NewConsistencyGroupSnapshotCreateParamsWithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotCreateParams {
	return &ConsistencyGroupSnapshotCreateParams{
		timeout: timeout,
	}
}

// NewConsistencyGroupSnapshotCreateParamsWithContext creates a new ConsistencyGroupSnapshotCreateParams object
// with the ability to set a context for a request.
func NewConsistencyGroupSnapshotCreateParamsWithContext(ctx context.Context) *ConsistencyGroupSnapshotCreateParams {
	return &ConsistencyGroupSnapshotCreateParams{
		Context: ctx,
	}
}

// NewConsistencyGroupSnapshotCreateParamsWithHTTPClient creates a new ConsistencyGroupSnapshotCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewConsistencyGroupSnapshotCreateParamsWithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotCreateParams {
	return &ConsistencyGroupSnapshotCreateParams{
		HTTPClient: client,
	}
}

/*
ConsistencyGroupSnapshotCreateParams contains all the parameters to send to the API endpoint

	for the consistency group snapshot create operation.

	Typically these are written to a http.Request.
*/
type ConsistencyGroupSnapshotCreateParams struct {

	/* Action.

	   Initiates the Snapshot copy create operation. The start of the Snapshot copy operation can optionally use a timeout value specified by "action_timeout". The Snapshot copy is commited by calling PATCH on the Snapshot copy href link with action specified as "commit".

	*/
	Action *string

	/* ActionTimeout.

	   Duration to complete the 2-phase Snapshot copy operation. This also specifies the maximum duration that the write-fence remains in effect on the volumes associated with this consistency group. Default is 7 seconds with a valid range of 1 to 90 seconds.

	*/
	ActionTimeout *int64

	/* ConsistencyGroupUUID.

	   The unique identifier of the consistency group to retrieve.

	*/
	ConsistencyGroupUUID string

	/* Info.

	   Information regarding a consistency group's Snapshot copy.

	*/
	Info *models.ConsistencyGroupSnapshot

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the consistency group snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotCreateParams) WithDefaults() *ConsistencyGroupSnapshotCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the consistency group snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConsistencyGroupSnapshotCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := ConsistencyGroupSnapshotCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithTimeout(timeout time.Duration) *ConsistencyGroupSnapshotCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithContext(ctx context.Context) *ConsistencyGroupSnapshotCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithHTTPClient(client *http.Client) *ConsistencyGroupSnapshotCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithAction(action *string) *ConsistencyGroupSnapshotCreateParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetAction(action *string) {
	o.Action = action
}

// WithActionTimeout adds the actionTimeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithActionTimeout(actionTimeout *int64) *ConsistencyGroupSnapshotCreateParams {
	o.SetActionTimeout(actionTimeout)
	return o
}

// SetActionTimeout adds the actionTimeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetActionTimeout(actionTimeout *int64) {
	o.ActionTimeout = actionTimeout
}

// WithConsistencyGroupUUID adds the consistencyGroupUUID to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithConsistencyGroupUUID(consistencyGroupUUID string) *ConsistencyGroupSnapshotCreateParams {
	o.SetConsistencyGroupUUID(consistencyGroupUUID)
	return o
}

// SetConsistencyGroupUUID adds the consistencyGroupUuid to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetConsistencyGroupUUID(consistencyGroupUUID string) {
	o.ConsistencyGroupUUID = consistencyGroupUUID
}

// WithInfo adds the info to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithInfo(info *models.ConsistencyGroupSnapshot) *ConsistencyGroupSnapshotCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetInfo(info *models.ConsistencyGroupSnapshot) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithReturnRecords(returnRecords *bool) *ConsistencyGroupSnapshotCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) WithReturnTimeout(returnTimeout *int64) *ConsistencyGroupSnapshotCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the consistency group snapshot create params
func (o *ConsistencyGroupSnapshotCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ConsistencyGroupSnapshotCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.ActionTimeout != nil {

		// query param action_timeout
		var qrActionTimeout int64

		if o.ActionTimeout != nil {
			qrActionTimeout = *o.ActionTimeout
		}
		qActionTimeout := swag.FormatInt64(qrActionTimeout)
		if qActionTimeout != "" {

			if err := r.SetQueryParam("action_timeout", qActionTimeout); err != nil {
				return err
			}
		}
	}

	// path param consistency_group.uuid
	if err := r.SetPathParam("consistency_group.uuid", o.ConsistencyGroupUUID); err != nil {
		return err
	}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
