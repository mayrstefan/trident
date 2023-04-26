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

// NewCifsSessionGetParams creates a new CifsSessionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSessionGetParams() *CifsSessionGetParams {
	return &CifsSessionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSessionGetParamsWithTimeout creates a new CifsSessionGetParams object
// with the ability to set a timeout on a request.
func NewCifsSessionGetParamsWithTimeout(timeout time.Duration) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		timeout: timeout,
	}
}

// NewCifsSessionGetParamsWithContext creates a new CifsSessionGetParams object
// with the ability to set a context for a request.
func NewCifsSessionGetParamsWithContext(ctx context.Context) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		Context: ctx,
	}
}

// NewCifsSessionGetParamsWithHTTPClient creates a new CifsSessionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSessionGetParamsWithHTTPClient(client *http.Client) *CifsSessionGetParams {
	return &CifsSessionGetParams{
		HTTPClient: client,
	}
}

/*
CifsSessionGetParams contains all the parameters to send to the API endpoint

	for the cifs session get operation.

	Typically these are written to a http.Request.
*/
type CifsSessionGetParams struct {

	/* ConnectionID.

	   Unique identifier for the SMB connection.
	*/
	ConnectionID int64

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Identifier.

	   Unique identifier for the SMB session.
	*/
	Identifier int64

	/* NodeUUID.

	   Node UUID.
	*/
	NodeUUID string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs session get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionGetParams) WithDefaults() *CifsSessionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs session get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs session get params
func (o *CifsSessionGetParams) WithTimeout(timeout time.Duration) *CifsSessionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs session get params
func (o *CifsSessionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs session get params
func (o *CifsSessionGetParams) WithContext(ctx context.Context) *CifsSessionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs session get params
func (o *CifsSessionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs session get params
func (o *CifsSessionGetParams) WithHTTPClient(client *http.Client) *CifsSessionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs session get params
func (o *CifsSessionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the cifs session get params
func (o *CifsSessionGetParams) WithConnectionID(connectionID int64) *CifsSessionGetParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the cifs session get params
func (o *CifsSessionGetParams) SetConnectionID(connectionID int64) {
	o.ConnectionID = connectionID
}

// WithFields adds the fields to the cifs session get params
func (o *CifsSessionGetParams) WithFields(fields []string) *CifsSessionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs session get params
func (o *CifsSessionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIdentifier adds the identifier to the cifs session get params
func (o *CifsSessionGetParams) WithIdentifier(identifier int64) *CifsSessionGetParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the cifs session get params
func (o *CifsSessionGetParams) SetIdentifier(identifier int64) {
	o.Identifier = identifier
}

// WithNodeUUID adds the nodeUUID to the cifs session get params
func (o *CifsSessionGetParams) WithNodeUUID(nodeUUID string) *CifsSessionGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the cifs session get params
func (o *CifsSessionGetParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WithSvmUUID adds the svmUUID to the cifs session get params
func (o *CifsSessionGetParams) WithSvmUUID(svmUUID string) *CifsSessionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs session get params
func (o *CifsSessionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSessionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", swag.FormatInt64(o.ConnectionID)); err != nil {
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

	// path param identifier
	if err := r.SetPathParam("identifier", swag.FormatInt64(o.Identifier)); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUID); err != nil {
		return err
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

// bindParamCifsSessionGet binds the parameter fields
func (o *CifsSessionGetParams) bindParamFields(formats strfmt.Registry) []string {
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
