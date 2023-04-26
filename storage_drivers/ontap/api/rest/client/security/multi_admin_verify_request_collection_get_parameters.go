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
)

// NewMultiAdminVerifyRequestCollectionGetParams creates a new MultiAdminVerifyRequestCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRequestCollectionGetParams() *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithTimeout creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithContext creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithContext(ctx context.Context) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRequestCollectionGetParamsWithHTTPClient creates a new MultiAdminVerifyRequestCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRequestCollectionGetParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRequestCollectionGetParams {
	return &MultiAdminVerifyRequestCollectionGetParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRequestCollectionGetParams contains all the parameters to send to the API endpoint

	for the multi admin verify request collection get operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRequestCollectionGetParams struct {

	/* ApproveExpiryTime.

	   Filter by approve_expiry_time
	*/
	ApproveExpiryTime *string

	/* ApproveTime.

	   Filter by approve_time
	*/
	ApproveTime *string

	/* ApprovedUsers.

	   Filter by approved_users
	*/
	ApprovedUsers *string

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

	/* ExecutionExpiryTime.

	   Filter by execution_expiry_time
	*/
	ExecutionExpiryTime *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Filter by index
	*/
	Index *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Operation.

	   Filter by operation
	*/
	Operation *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

	/* PendingApprovers.

	   Filter by pending_approvers
	*/
	PendingApprovers *int64

	/* PermittedUsers.

	   Filter by permitted_users
	*/
	PermittedUsers *string

	/* PotentialApprovers.

	   Filter by potential_approvers
	*/
	PotentialApprovers *string

	/* Query.

	   Filter by query
	*/
	Query *string

	/* RequiredApprovers.

	   Filter by required_approvers
	*/
	RequiredApprovers *int64

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

	/* State.

	   Filter by state
	*/
	State *string

	/* UserRequested.

	   Filter by user_requested
	*/
	UserRequested *string

	/* UserVetoed.

	   Filter by user_vetoed
	*/
	UserVetoed *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify request collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestCollectionGetParams) WithDefaults() *MultiAdminVerifyRequestCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify request collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := MultiAdminVerifyRequestCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithContext(ctx context.Context) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApproveExpiryTime adds the approveExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApproveExpiryTime(approveExpiryTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApproveExpiryTime(approveExpiryTime)
	return o
}

// SetApproveExpiryTime adds the approveExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApproveExpiryTime(approveExpiryTime *string) {
	o.ApproveExpiryTime = approveExpiryTime
}

// WithApproveTime adds the approveTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApproveTime(approveTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApproveTime(approveTime)
	return o
}

// SetApproveTime adds the approveTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApproveTime(approveTime *string) {
	o.ApproveTime = approveTime
}

// WithApprovedUsers adds the approvedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithApprovedUsers(approvedUsers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetApprovedUsers(approvedUsers)
	return o
}

// SetApprovedUsers adds the approvedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetApprovedUsers(approvedUsers *string) {
	o.ApprovedUsers = approvedUsers
}

// WithComment adds the comment to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithComment(comment *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithCreateTime adds the createTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithCreateTime(createTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithExecutionExpiryTime adds the executionExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithExecutionExpiryTime(executionExpiryTime *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetExecutionExpiryTime(executionExpiryTime)
	return o
}

// SetExecutionExpiryTime adds the executionExpiryTime to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetExecutionExpiryTime(executionExpiryTime *string) {
	o.ExecutionExpiryTime = executionExpiryTime
}

// WithFields adds the fields to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithFields(fields []string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndex adds the index to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithIndex(index *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetIndex(index *int64) {
	o.Index = index
}

// WithMaxRecords adds the maxRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithMaxRecords(maxRecords *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOperation adds the operation to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOperation(operation *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOperation(operation)
	return o
}

// SetOperation adds the operation to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOperation(operation *string) {
	o.Operation = operation
}

// WithOrderBy adds the orderBy to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOrderBy(orderBy []string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwnerName adds the ownerName to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOwnerName(ownerName *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithOwnerUUID(ownerUUID *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithPendingApprovers adds the pendingApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPendingApprovers(pendingApprovers *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPendingApprovers(pendingApprovers)
	return o
}

// SetPendingApprovers adds the pendingApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPendingApprovers(pendingApprovers *int64) {
	o.PendingApprovers = pendingApprovers
}

// WithPermittedUsers adds the permittedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPermittedUsers(permittedUsers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPermittedUsers(permittedUsers)
	return o
}

// SetPermittedUsers adds the permittedUsers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPermittedUsers(permittedUsers *string) {
	o.PermittedUsers = permittedUsers
}

// WithPotentialApprovers adds the potentialApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithPotentialApprovers(potentialApprovers *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetPotentialApprovers(potentialApprovers)
	return o
}

// SetPotentialApprovers adds the potentialApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetPotentialApprovers(potentialApprovers *string) {
	o.PotentialApprovers = potentialApprovers
}

// WithQuery adds the query to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithQuery(query *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetQuery(query *string) {
	o.Query = query
}

// WithRequiredApprovers adds the requiredApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithRequiredApprovers(requiredApprovers *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetRequiredApprovers(requiredApprovers)
	return o
}

// SetRequiredApprovers adds the requiredApprovers to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetRequiredApprovers(requiredApprovers *int64) {
	o.RequiredApprovers = requiredApprovers
}

// WithReturnRecords adds the returnRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithReturnRecords(returnRecords *bool) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithState adds the state to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithState(state *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithUserRequested adds the userRequested to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithUserRequested(userRequested *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetUserRequested(userRequested)
	return o
}

// SetUserRequested adds the userRequested to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetUserRequested(userRequested *string) {
	o.UserRequested = userRequested
}

// WithUserVetoed adds the userVetoed to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) WithUserVetoed(userVetoed *string) *MultiAdminVerifyRequestCollectionGetParams {
	o.SetUserVetoed(userVetoed)
	return o
}

// SetUserVetoed adds the userVetoed to the multi admin verify request collection get params
func (o *MultiAdminVerifyRequestCollectionGetParams) SetUserVetoed(userVetoed *string) {
	o.UserVetoed = userVetoed
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRequestCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApproveExpiryTime != nil {

		// query param approve_expiry_time
		var qrApproveExpiryTime string

		if o.ApproveExpiryTime != nil {
			qrApproveExpiryTime = *o.ApproveExpiryTime
		}
		qApproveExpiryTime := qrApproveExpiryTime
		if qApproveExpiryTime != "" {

			if err := r.SetQueryParam("approve_expiry_time", qApproveExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.ApproveTime != nil {

		// query param approve_time
		var qrApproveTime string

		if o.ApproveTime != nil {
			qrApproveTime = *o.ApproveTime
		}
		qApproveTime := qrApproveTime
		if qApproveTime != "" {

			if err := r.SetQueryParam("approve_time", qApproveTime); err != nil {
				return err
			}
		}
	}

	if o.ApprovedUsers != nil {

		// query param approved_users
		var qrApprovedUsers string

		if o.ApprovedUsers != nil {
			qrApprovedUsers = *o.ApprovedUsers
		}
		qApprovedUsers := qrApprovedUsers
		if qApprovedUsers != "" {

			if err := r.SetQueryParam("approved_users", qApprovedUsers); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.CreateTime != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTime != nil {
			qrCreateTime = *o.CreateTime
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}

	if o.ExecutionExpiryTime != nil {

		// query param execution_expiry_time
		var qrExecutionExpiryTime string

		if o.ExecutionExpiryTime != nil {
			qrExecutionExpiryTime = *o.ExecutionExpiryTime
		}
		qExecutionExpiryTime := qrExecutionExpiryTime
		if qExecutionExpiryTime != "" {

			if err := r.SetQueryParam("execution_expiry_time", qExecutionExpiryTime); err != nil {
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

	if o.Index != nil {

		// query param index
		var qrIndex int64

		if o.Index != nil {
			qrIndex = *o.Index
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
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

	if o.Operation != nil {

		// query param operation
		var qrOperation string

		if o.Operation != nil {
			qrOperation = *o.Operation
		}
		qOperation := qrOperation
		if qOperation != "" {

			if err := r.SetQueryParam("operation", qOperation); err != nil {
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

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PendingApprovers != nil {

		// query param pending_approvers
		var qrPendingApprovers int64

		if o.PendingApprovers != nil {
			qrPendingApprovers = *o.PendingApprovers
		}
		qPendingApprovers := swag.FormatInt64(qrPendingApprovers)
		if qPendingApprovers != "" {

			if err := r.SetQueryParam("pending_approvers", qPendingApprovers); err != nil {
				return err
			}
		}
	}

	if o.PermittedUsers != nil {

		// query param permitted_users
		var qrPermittedUsers string

		if o.PermittedUsers != nil {
			qrPermittedUsers = *o.PermittedUsers
		}
		qPermittedUsers := qrPermittedUsers
		if qPermittedUsers != "" {

			if err := r.SetQueryParam("permitted_users", qPermittedUsers); err != nil {
				return err
			}
		}
	}

	if o.PotentialApprovers != nil {

		// query param potential_approvers
		var qrPotentialApprovers string

		if o.PotentialApprovers != nil {
			qrPotentialApprovers = *o.PotentialApprovers
		}
		qPotentialApprovers := qrPotentialApprovers
		if qPotentialApprovers != "" {

			if err := r.SetQueryParam("potential_approvers", qPotentialApprovers); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.RequiredApprovers != nil {

		// query param required_approvers
		var qrRequiredApprovers int64

		if o.RequiredApprovers != nil {
			qrRequiredApprovers = *o.RequiredApprovers
		}
		qRequiredApprovers := swag.FormatInt64(qrRequiredApprovers)
		if qRequiredApprovers != "" {

			if err := r.SetQueryParam("required_approvers", qRequiredApprovers); err != nil {
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

	if o.UserRequested != nil {

		// query param user_requested
		var qrUserRequested string

		if o.UserRequested != nil {
			qrUserRequested = *o.UserRequested
		}
		qUserRequested := qrUserRequested
		if qUserRequested != "" {

			if err := r.SetQueryParam("user_requested", qUserRequested); err != nil {
				return err
			}
		}
	}

	if o.UserVetoed != nil {

		// query param user_vetoed
		var qrUserVetoed string

		if o.UserVetoed != nil {
			qrUserVetoed = *o.UserVetoed
		}
		qUserVetoed := qrUserVetoed
		if qUserVetoed != "" {

			if err := r.SetQueryParam("user_vetoed", qUserVetoed); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMultiAdminVerifyRequestCollectionGet binds the parameter fields
func (o *MultiAdminVerifyRequestCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamMultiAdminVerifyRequestCollectionGet binds the parameter order_by
func (o *MultiAdminVerifyRequestCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
