// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewScheduleCollectionGetParams creates a new ScheduleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleCollectionGetParams() *ScheduleCollectionGetParams {
	return &ScheduleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleCollectionGetParamsWithTimeout creates a new ScheduleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewScheduleCollectionGetParamsWithTimeout(timeout time.Duration) *ScheduleCollectionGetParams {
	return &ScheduleCollectionGetParams{
		timeout: timeout,
	}
}

// NewScheduleCollectionGetParamsWithContext creates a new ScheduleCollectionGetParams object
// with the ability to set a context for a request.
func NewScheduleCollectionGetParamsWithContext(ctx context.Context) *ScheduleCollectionGetParams {
	return &ScheduleCollectionGetParams{
		Context: ctx,
	}
}

// NewScheduleCollectionGetParamsWithHTTPClient creates a new ScheduleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleCollectionGetParamsWithHTTPClient(client *http.Client) *ScheduleCollectionGetParams {
	return &ScheduleCollectionGetParams{
		HTTPClient: client,
	}
}

/* ScheduleCollectionGetParams contains all the parameters to send to the API endpoint
   for the schedule collection get operation.

   Typically these are written to a http.Request.
*/
type ScheduleCollectionGetParams struct {

	/* ClusterName.

	   Filter by cluster.name
	*/
	ClusterNameQueryParameter *string

	/* ClusterUUID.

	   Filter by cluster.uuid
	*/
	ClusterUUIDQueryParameter *string

	/* CronDays.

	   Filter by cron.days
	*/
	CronDaysQueryParameter *int64

	/* CronHours.

	   Filter by cron.hours
	*/
	CronHoursQueryParameter *int64

	/* CronMinutes.

	   Filter by cron.minutes
	*/
	CronMinutesQueryParameter *int64

	/* CronMonths.

	   Filter by cron.months
	*/
	CronMonthsQueryParameter *int64

	/* CronWeekdays.

	   Filter by cron.weekdays
	*/
	CronWeekdaysQueryParameter *int64

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Interval.

	   Filter by interval
	*/
	IntervalQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

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

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleCollectionGetParams) WithDefaults() *ScheduleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := ScheduleCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithTimeout(timeout time.Duration) *ScheduleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithContext(ctx context.Context) *ScheduleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithHTTPClient(client *http.Client) *ScheduleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNameQueryParameter adds the clusterName to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithClusterNameQueryParameter(clusterName *string) *ScheduleCollectionGetParams {
	o.SetClusterNameQueryParameter(clusterName)
	return o
}

// SetClusterNameQueryParameter adds the clusterName to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetClusterNameQueryParameter(clusterName *string) {
	o.ClusterNameQueryParameter = clusterName
}

// WithClusterUUIDQueryParameter adds the clusterUUID to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithClusterUUIDQueryParameter(clusterUUID *string) *ScheduleCollectionGetParams {
	o.SetClusterUUIDQueryParameter(clusterUUID)
	return o
}

// SetClusterUUIDQueryParameter adds the clusterUuid to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetClusterUUIDQueryParameter(clusterUUID *string) {
	o.ClusterUUIDQueryParameter = clusterUUID
}

// WithCronDaysQueryParameter adds the cronDays to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithCronDaysQueryParameter(cronDays *int64) *ScheduleCollectionGetParams {
	o.SetCronDaysQueryParameter(cronDays)
	return o
}

// SetCronDaysQueryParameter adds the cronDays to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetCronDaysQueryParameter(cronDays *int64) {
	o.CronDaysQueryParameter = cronDays
}

// WithCronHoursQueryParameter adds the cronHours to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithCronHoursQueryParameter(cronHours *int64) *ScheduleCollectionGetParams {
	o.SetCronHoursQueryParameter(cronHours)
	return o
}

// SetCronHoursQueryParameter adds the cronHours to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetCronHoursQueryParameter(cronHours *int64) {
	o.CronHoursQueryParameter = cronHours
}

// WithCronMinutesQueryParameter adds the cronMinutes to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithCronMinutesQueryParameter(cronMinutes *int64) *ScheduleCollectionGetParams {
	o.SetCronMinutesQueryParameter(cronMinutes)
	return o
}

// SetCronMinutesQueryParameter adds the cronMinutes to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetCronMinutesQueryParameter(cronMinutes *int64) {
	o.CronMinutesQueryParameter = cronMinutes
}

// WithCronMonthsQueryParameter adds the cronMonths to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithCronMonthsQueryParameter(cronMonths *int64) *ScheduleCollectionGetParams {
	o.SetCronMonthsQueryParameter(cronMonths)
	return o
}

// SetCronMonthsQueryParameter adds the cronMonths to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetCronMonthsQueryParameter(cronMonths *int64) {
	o.CronMonthsQueryParameter = cronMonths
}

// WithCronWeekdaysQueryParameter adds the cronWeekdays to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithCronWeekdaysQueryParameter(cronWeekdays *int64) *ScheduleCollectionGetParams {
	o.SetCronWeekdaysQueryParameter(cronWeekdays)
	return o
}

// SetCronWeekdaysQueryParameter adds the cronWeekdays to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetCronWeekdaysQueryParameter(cronWeekdays *int64) {
	o.CronWeekdaysQueryParameter = cronWeekdays
}

// WithFields adds the fields to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithFields(fields []string) *ScheduleCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIntervalQueryParameter adds the interval to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithIntervalQueryParameter(interval *string) *ScheduleCollectionGetParams {
	o.SetIntervalQueryParameter(interval)
	return o
}

// SetIntervalQueryParameter adds the interval to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetIntervalQueryParameter(interval *string) {
	o.IntervalQueryParameter = interval
}

// WithMaxRecords adds the maxRecords to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithMaxRecords(maxRecords *int64) *ScheduleCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNameQueryParameter adds the name to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithNameQueryParameter(name *string) *ScheduleCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderBy adds the orderBy to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithOrderBy(orderBy []string) *ScheduleCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithReturnRecords(returnRecords *bool) *ScheduleCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *ScheduleCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithTypeQueryParameter adds the typeVar to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithTypeQueryParameter(typeVar *string) *ScheduleCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUUIDQueryParameter adds the uuid to the schedule collection get params
func (o *ScheduleCollectionGetParams) WithUUIDQueryParameter(uuid *string) *ScheduleCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the schedule collection get params
func (o *ScheduleCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterNameQueryParameter != nil {

		// query param cluster.name
		var qrClusterName string

		if o.ClusterNameQueryParameter != nil {
			qrClusterName = *o.ClusterNameQueryParameter
		}
		qClusterName := qrClusterName
		if qClusterName != "" {

			if err := r.SetQueryParam("cluster.name", qClusterName); err != nil {
				return err
			}
		}
	}

	if o.ClusterUUIDQueryParameter != nil {

		// query param cluster.uuid
		var qrClusterUUID string

		if o.ClusterUUIDQueryParameter != nil {
			qrClusterUUID = *o.ClusterUUIDQueryParameter
		}
		qClusterUUID := qrClusterUUID
		if qClusterUUID != "" {

			if err := r.SetQueryParam("cluster.uuid", qClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.CronDaysQueryParameter != nil {

		// query param cron.days
		var qrCronDays int64

		if o.CronDaysQueryParameter != nil {
			qrCronDays = *o.CronDaysQueryParameter
		}
		qCronDays := swag.FormatInt64(qrCronDays)
		if qCronDays != "" {

			if err := r.SetQueryParam("cron.days", qCronDays); err != nil {
				return err
			}
		}
	}

	if o.CronHoursQueryParameter != nil {

		// query param cron.hours
		var qrCronHours int64

		if o.CronHoursQueryParameter != nil {
			qrCronHours = *o.CronHoursQueryParameter
		}
		qCronHours := swag.FormatInt64(qrCronHours)
		if qCronHours != "" {

			if err := r.SetQueryParam("cron.hours", qCronHours); err != nil {
				return err
			}
		}
	}

	if o.CronMinutesQueryParameter != nil {

		// query param cron.minutes
		var qrCronMinutes int64

		if o.CronMinutesQueryParameter != nil {
			qrCronMinutes = *o.CronMinutesQueryParameter
		}
		qCronMinutes := swag.FormatInt64(qrCronMinutes)
		if qCronMinutes != "" {

			if err := r.SetQueryParam("cron.minutes", qCronMinutes); err != nil {
				return err
			}
		}
	}

	if o.CronMonthsQueryParameter != nil {

		// query param cron.months
		var qrCronMonths int64

		if o.CronMonthsQueryParameter != nil {
			qrCronMonths = *o.CronMonthsQueryParameter
		}
		qCronMonths := swag.FormatInt64(qrCronMonths)
		if qCronMonths != "" {

			if err := r.SetQueryParam("cron.months", qCronMonths); err != nil {
				return err
			}
		}
	}

	if o.CronWeekdaysQueryParameter != nil {

		// query param cron.weekdays
		var qrCronWeekdays int64

		if o.CronWeekdaysQueryParameter != nil {
			qrCronWeekdays = *o.CronWeekdaysQueryParameter
		}
		qCronWeekdays := swag.FormatInt64(qrCronWeekdays)
		if qCronWeekdays != "" {

			if err := r.SetQueryParam("cron.weekdays", qCronWeekdays); err != nil {
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

	if o.IntervalQueryParameter != nil {

		// query param interval
		var qrInterval string

		if o.IntervalQueryParameter != nil {
			qrInterval = *o.IntervalQueryParameter
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
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

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
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

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
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

// bindParamScheduleCollectionGet binds the parameter fields
func (o *ScheduleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamScheduleCollectionGet binds the parameter order_by
func (o *ScheduleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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