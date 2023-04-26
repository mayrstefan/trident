// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsMessageResponse ems message response
//
// swagger:model ems_message_response
type EmsMessageResponse struct {

	// links
	Links *EmsMessageResponseInlineLinks `json:"_links,omitempty"`

	// ems message response inline records
	EmsMessageResponseInlineRecords []*EmsMessageResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this ems message response
func (m *EmsMessageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmsMessageResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsMessageResponse) validateEmsMessageResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.EmsMessageResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.EmsMessageResponseInlineRecords); i++ {
		if swag.IsZero(m.EmsMessageResponseInlineRecords[i]) { // not required
			continue
		}

		if m.EmsMessageResponseInlineRecords[i] != nil {
			if err := m.EmsMessageResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems message response based on the context it is used
func (m *EmsMessageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmsMessageResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsMessageResponse) contextValidateEmsMessageResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmsMessageResponseInlineRecords); i++ {

		if m.EmsMessageResponseInlineRecords[i] != nil {
			if err := m.EmsMessageResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsMessageResponse) UnmarshalBinary(b []byte) error {
	var res EmsMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsMessageResponseInlineLinks ems message response inline links
//
// swagger:model ems_message_response_inline__links
type EmsMessageResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems message response inline links
func (m *EmsMessageResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsMessageResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems message response inline links based on the context it is used
func (m *EmsMessageResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsMessageResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsMessageResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsMessageResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsMessageResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsMessageResponseInlineRecordsInlineArrayItem ems message response inline records inline array item
//
// swagger:model ems_message_response_inline_records_inline_array_item
type EmsMessageResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// Corrective action
	// Read Only: true
	CorrectiveAction *string `json:"corrective_action,omitempty"`

	// Is deprecated?
	// Example: true
	// Read Only: true
	Deprecated *bool `json:"deprecated,omitempty"`

	// Description of the event.
	// Read Only: true
	Description *string `json:"description,omitempty"`

	// Name of the event.
	// Example: callhome.spares.low
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Severity
	// Example: error
	// Read Only: true
	// Enum: [emergency alert error notice informational debug]
	Severity *string `json:"severity,omitempty"`

	// SNMP trap type
	// Example: standard
	// Read Only: true
	// Enum: [standard built_in severity_based]
	SnmpTrapType *string `json:"snmp_trap_type,omitempty"`
}

// Validate validates this ems message response inline records inline array item
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnmpTrapType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var emsMessageResponseInlineRecordsInlineArrayItemTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emergency","alert","error","notice","informational","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsMessageResponseInlineRecordsInlineArrayItemTypeSeverityPropEnum = append(emsMessageResponseInlineRecordsInlineArrayItemTypeSeverityPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// emergency
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityEmergency captures enum value "emergency"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityEmergency string = "emergency"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// alert
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityAlert captures enum value "alert"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityAlert string = "alert"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// error
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityError captures enum value "error"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityError string = "error"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// notice
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityNotice captures enum value "notice"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityNotice string = "notice"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// informational
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityInformational captures enum value "informational"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityInformational string = "informational"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// severity
	// Severity
	// debug
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSeverityDebug captures enum value "debug"
	EmsMessageResponseInlineRecordsInlineArrayItemSeverityDebug string = "debug"
)

// prop value enum
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsMessageResponseInlineRecordsInlineArrayItemTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

var emsMessageResponseInlineRecordsInlineArrayItemTypeSnmpTrapTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["standard","built_in","severity_based"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsMessageResponseInlineRecordsInlineArrayItemTypeSnmpTrapTypePropEnum = append(emsMessageResponseInlineRecordsInlineArrayItemTypeSnmpTrapTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// snmp_trap_type
	// SnmpTrapType
	// standard
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeStandard captures enum value "standard"
	EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeStandard string = "standard"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// snmp_trap_type
	// SnmpTrapType
	// built_in
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeBuiltIn captures enum value "built_in"
	EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeBuiltIn string = "built_in"

	// BEGIN DEBUGGING
	// ems_message_response_inline_records_inline_array_item
	// EmsMessageResponseInlineRecordsInlineArrayItem
	// snmp_trap_type
	// SnmpTrapType
	// severity_based
	// END DEBUGGING
	// EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeSeverityBased captures enum value "severity_based"
	EmsMessageResponseInlineRecordsInlineArrayItemSnmpTrapTypeSeverityBased string = "severity_based"
)

// prop value enum
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) validateSnmpTrapTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsMessageResponseInlineRecordsInlineArrayItemTypeSnmpTrapTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) validateSnmpTrapType(formats strfmt.Registry) error {
	if swag.IsZero(m.SnmpTrapType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSnmpTrapTypeEnum("snmp_trap_type", "body", *m.SnmpTrapType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems message response inline records inline array item based on the context it is used
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCorrectiveAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeprecated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnmpTrapType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateCorrectiveAction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "corrective_action", "body", m.CorrectiveAction); err != nil {
		return err
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateDeprecated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deprecated", "body", m.Deprecated); err != nil {
		return err
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItem) contextValidateSnmpTrapType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snmp_trap_type", "body", m.SnmpTrapType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsMessageResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res EmsMessageResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks ems message response inline records inline array item inline links
//
// swagger:model ems_message_response_inline_records_inline_array_item_inline__links
type EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems message response inline records inline array item inline links
func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems message response inline records inline array item inline links based on the context it is used
func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsMessageResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
