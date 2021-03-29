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

// IscsiCredentialsCollectionGetReader is a Reader for the IscsiCredentialsCollectionGet structure.
type IscsiCredentialsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsCollectionGetOK creates a IscsiCredentialsCollectionGetOK with default headers values
func NewIscsiCredentialsCollectionGetOK() *IscsiCredentialsCollectionGetOK {
	return &IscsiCredentialsCollectionGetOK{}
}

/* IscsiCredentialsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsCollectionGetOK struct {
	Payload *models.IscsiCredentialsResponse
}

func (o *IscsiCredentialsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsiCredentialsCollectionGetOK  %+v", 200, o.Payload)
}
func (o *IscsiCredentialsCollectionGetOK) GetPayload() *models.IscsiCredentialsResponse {
	return o.Payload
}

func (o *IscsiCredentialsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiCredentialsCollectionGetDefault creates a IscsiCredentialsCollectionGetDefault with default headers values
func NewIscsiCredentialsCollectionGetDefault(code int) *IscsiCredentialsCollectionGetDefault {
	return &IscsiCredentialsCollectionGetDefault{
		_statusCode: code,
	}
}

/* IscsiCredentialsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type IscsiCredentialsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi credentials collection get default response
func (o *IscsiCredentialsCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *IscsiCredentialsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials][%d] iscsi_credentials_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *IscsiCredentialsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}