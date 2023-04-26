// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// GcpKmsRekeyExternalReader is a Reader for the GcpKmsRekeyExternal structure.
type GcpKmsRekeyExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GcpKmsRekeyExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGcpKmsRekeyExternalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGcpKmsRekeyExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGcpKmsRekeyExternalAccepted creates a GcpKmsRekeyExternalAccepted with default headers values
func NewGcpKmsRekeyExternalAccepted() *GcpKmsRekeyExternalAccepted {
	return &GcpKmsRekeyExternalAccepted{}
}

/*
GcpKmsRekeyExternalAccepted describes a response with status code 202, with default header values.

Accepted
*/
type GcpKmsRekeyExternalAccepted struct {
}

// IsSuccess returns true when this gcp kms rekey external accepted response has a 2xx status code
func (o *GcpKmsRekeyExternalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gcp kms rekey external accepted response has a 3xx status code
func (o *GcpKmsRekeyExternalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gcp kms rekey external accepted response has a 4xx status code
func (o *GcpKmsRekeyExternalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this gcp kms rekey external accepted response has a 5xx status code
func (o *GcpKmsRekeyExternalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this gcp kms rekey external accepted response a status code equal to that given
func (o *GcpKmsRekeyExternalAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GcpKmsRekeyExternalAccepted) Error() string {
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalAccepted ", 202)
}

func (o *GcpKmsRekeyExternalAccepted) String() string {
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcpKmsRekeyExternalAccepted ", 202)
}

func (o *GcpKmsRekeyExternalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGcpKmsRekeyExternalDefault creates a GcpKmsRekeyExternalDefault with default headers values
func NewGcpKmsRekeyExternalDefault(code int) *GcpKmsRekeyExternalDefault {
	return &GcpKmsRekeyExternalDefault{
		_statusCode: code,
	}
}

/*
	GcpKmsRekeyExternalDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
| 65537721 | Google Cloud KMS is not configured for the given SVM. |
| 65537729 | External rekey failed on one or more nodes. Use the REST API POST method \"/api/security/gcp-kms/{uuid}/rekey-external\" to try the rekey operation again. |
*/
type GcpKmsRekeyExternalDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gcp kms rekey external default response
func (o *GcpKmsRekeyExternalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this gcp kms rekey external default response has a 2xx status code
func (o *GcpKmsRekeyExternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gcp kms rekey external default response has a 3xx status code
func (o *GcpKmsRekeyExternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gcp kms rekey external default response has a 4xx status code
func (o *GcpKmsRekeyExternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gcp kms rekey external default response has a 5xx status code
func (o *GcpKmsRekeyExternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gcp kms rekey external default response a status code equal to that given
func (o *GcpKmsRekeyExternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GcpKmsRekeyExternalDefault) Error() string {
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcp_kms_rekey_external default  %+v", o._statusCode, o.Payload)
}

func (o *GcpKmsRekeyExternalDefault) String() string {
	return fmt.Sprintf("[POST /security/gcp-kms/{gcp_kms.uuid}/rekey-external][%d] gcp_kms_rekey_external default  %+v", o._statusCode, o.Payload)
}

func (o *GcpKmsRekeyExternalDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GcpKmsRekeyExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
