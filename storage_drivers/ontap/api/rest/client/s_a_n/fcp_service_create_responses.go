// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcpServiceCreateReader is a Reader for the FcpServiceCreate structure.
type FcpServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcpServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFcpServiceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcpServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcpServiceCreateCreated creates a FcpServiceCreateCreated with default headers values
func NewFcpServiceCreateCreated() *FcpServiceCreateCreated {
	return &FcpServiceCreateCreated{}
}

/*
FcpServiceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FcpServiceCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.FcpServiceResponse
}

// IsSuccess returns true when this fcp service create created response has a 2xx status code
func (o *FcpServiceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fcp service create created response has a 3xx status code
func (o *FcpServiceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fcp service create created response has a 4xx status code
func (o *FcpServiceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fcp service create created response has a 5xx status code
func (o *FcpServiceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fcp service create created response a status code equal to that given
func (o *FcpServiceCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *FcpServiceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/san/fcp/services][%d] fcpServiceCreateCreated  %+v", 201, o.Payload)
}

func (o *FcpServiceCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/san/fcp/services][%d] fcpServiceCreateCreated  %+v", 201, o.Payload)
}

func (o *FcpServiceCreateCreated) GetPayload() *models.FcpServiceResponse {
	return o.Payload
}

func (o *FcpServiceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.FcpServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcpServiceCreateDefault creates a FcpServiceCreateDefault with default headers values
func NewFcpServiceCreateDefault(code int) *FcpServiceCreateDefault {
	return &FcpServiceCreateDefault{
		_statusCode: code,
	}
}

/*
	FcpServiceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115127 | The cluster lacks a valid FCP license. |
| 2621462 | The supplied SVM does not exist. |
| 2621507 | The Fibre Channel Protocol is not allowed for the specified SVM. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374082 | The Fibre Channel Protocol service already exists for the SVM. |
| 5374092 | The Fibre Channel Procotol is not supported on the cluster hardware configuration; there are no Fibre Channel adapters. |
| 5374893 | The SVM is stopped. The SVM must be running to create a Fibre Channel Protocol service. |
*/
type FcpServiceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fcp service create default response
func (o *FcpServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fcp service create default response has a 2xx status code
func (o *FcpServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fcp service create default response has a 3xx status code
func (o *FcpServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fcp service create default response has a 4xx status code
func (o *FcpServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fcp service create default response has a 5xx status code
func (o *FcpServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fcp service create default response a status code equal to that given
func (o *FcpServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FcpServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/san/fcp/services][%d] fcp_service_create default  %+v", o._statusCode, o.Payload)
}

func (o *FcpServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/san/fcp/services][%d] fcp_service_create default  %+v", o._statusCode, o.Payload)
}

func (o *FcpServiceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcpServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
