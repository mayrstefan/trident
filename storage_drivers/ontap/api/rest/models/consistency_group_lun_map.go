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

// ConsistencyGroupLunMap A LUN map is an association between a LUN and an initiator group.<br/>
// When a LUN is mapped to an initiator group, the initiator group's initiators are granted access to the LUN. The relationship between a LUN and an initiator group is many LUNs to many initiator groups.
//
// swagger:model consistency_group_lun_map
type ConsistencyGroupLunMap struct {

	// igroup
	Igroup *ConsistencyGroupLunMapInlineIgroup `json:"igroup,omitempty"`

	// The logical unit number assigned to the LUN when mapped to the specified initiator group. The number is used to identify the LUN to initiators in the initiator group when communicating through the Fibre Channel Protocol or iSCSI. Optional in POST; if no value is provided, ONTAP assigns the lowest available value.
	//
	LogicalUnitNumber *int64 `json:"logical_unit_number,omitempty"`
}

// Validate validates this consistency group lun map
func (m *ConsistencyGroupLunMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIgroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMap) validateIgroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroup) { // not required
		return nil
	}

	if m.Igroup != nil {
		if err := m.Igroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group lun map based on the context it is used
func (m *ConsistencyGroupLunMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIgroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMap) contextValidateIgroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Igroup != nil {
		if err := m.Igroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunMap) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupLunMapInlineIgroup The initiator group that directly owns the initiator, which is where modification of the initiator is supported. This property will only be populated when the initiator is a member of a nested initiator group.
//
// swagger:model consistency_group_lun_map_inline_igroup
type ConsistencyGroupLunMapInlineIgroup struct {

	// A comment available for use by the administrator. Valid in POST and PATCH.
	//
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Separate igroup definitions to include in this igroup.
	//
	Igroups []*ConsistencyGroupLunMapIgroupIgroupsItems0 `json:"igroups,omitempty"`

	// The initiators that are members of the group.
	//
	Initiators []*ConsistencyGroupLunMapIgroupInitiatorsItems0 `json:"initiators,omitempty"`

	// The name of the initiator group. Required in POST; optional in PATCH.
	//
	// Example: igroup1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The host operating system of the initiator group. All initiators in the group should be hosts of the same operating system. Required in POST; optional in PATCH.
	//
	// Enum: [aix hpux hyper_v linux netware openvms solaris vmware windows xen]
	OsType *string `json:"os_type,omitempty"`

	// The protocols supported by the initiator group. This restricts the type of initiators that can be added to the initiator group. Optional in POST; if not supplied, this defaults to _mixed_.<br/>
	// The protocol of an initiator group cannot be changed after creation of the group.
	//
	// Enum: [fcp iscsi mixed]
	Protocol *string `json:"protocol,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group lun map inline igroup
func (m *ConsistencyGroupLunMapInlineIgroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("igroup"+"."+"comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup"+"."+"comment", "body", *m.Comment, 254); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroups) { // not required
		return nil
	}

	for i := 0; i < len(m.Igroups); i++ {
		if swag.IsZero(m.Igroups[i]) { // not required
			continue
		}

		if m.Igroups[i] != nil {
			if err := m.Igroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroup" + "." + "igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateInitiators(formats strfmt.Registry) error {
	if swag.IsZero(m.Initiators) { // not required
		return nil
	}

	for i := 0; i < len(m.Initiators); i++ {
		if swag.IsZero(m.Initiators[i]) { // not required
			continue
		}

		if m.Initiators[i] != nil {
			if err := m.Initiators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroup" + "." + "initiators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("igroup"+"."+"name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup"+"."+"name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

var consistencyGroupLunMapInlineIgroupTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","hpux","hyper_v","linux","netware","openvms","solaris","vmware","windows","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupLunMapInlineIgroupTypeOsTypePropEnum = append(consistencyGroupLunMapInlineIgroupTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// aix
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeAix captures enum value "aix"
	ConsistencyGroupLunMapInlineIgroupOsTypeAix string = "aix"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// hpux
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeHpux captures enum value "hpux"
	ConsistencyGroupLunMapInlineIgroupOsTypeHpux string = "hpux"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// hyper_v
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeHyperv captures enum value "hyper_v"
	ConsistencyGroupLunMapInlineIgroupOsTypeHyperv string = "hyper_v"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// linux
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeLinux captures enum value "linux"
	ConsistencyGroupLunMapInlineIgroupOsTypeLinux string = "linux"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// netware
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeNetware captures enum value "netware"
	ConsistencyGroupLunMapInlineIgroupOsTypeNetware string = "netware"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// openvms
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeOpenvms captures enum value "openvms"
	ConsistencyGroupLunMapInlineIgroupOsTypeOpenvms string = "openvms"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// solaris
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeSolaris captures enum value "solaris"
	ConsistencyGroupLunMapInlineIgroupOsTypeSolaris string = "solaris"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// vmware
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeVmware captures enum value "vmware"
	ConsistencyGroupLunMapInlineIgroupOsTypeVmware string = "vmware"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// windows
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeWindows captures enum value "windows"
	ConsistencyGroupLunMapInlineIgroupOsTypeWindows string = "windows"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// os_type
	// OsType
	// xen
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupOsTypeXen captures enum value "xen"
	ConsistencyGroupLunMapInlineIgroupOsTypeXen string = "xen"
)

// prop value enum
func (m *ConsistencyGroupLunMapInlineIgroup) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupLunMapInlineIgroupTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("igroup"+"."+"os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

var consistencyGroupLunMapInlineIgroupTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi","mixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupLunMapInlineIgroupTypeProtocolPropEnum = append(consistencyGroupLunMapInlineIgroupTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// protocol
	// Protocol
	// fcp
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupProtocolFcp captures enum value "fcp"
	ConsistencyGroupLunMapInlineIgroupProtocolFcp string = "fcp"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// protocol
	// Protocol
	// iscsi
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupProtocolIscsi captures enum value "iscsi"
	ConsistencyGroupLunMapInlineIgroupProtocolIscsi string = "iscsi"

	// BEGIN DEBUGGING
	// consistency_group_lun_map_inline_igroup
	// ConsistencyGroupLunMapInlineIgroup
	// protocol
	// Protocol
	// mixed
	// END DEBUGGING
	// ConsistencyGroupLunMapInlineIgroupProtocolMixed captures enum value "mixed"
	ConsistencyGroupLunMapInlineIgroupProtocolMixed string = "mixed"
)

// prop value enum
func (m *ConsistencyGroupLunMapInlineIgroup) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupLunMapInlineIgroupTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("igroup"+"."+"protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group lun map inline igroup based on the context it is used
func (m *ConsistencyGroupLunMapInlineIgroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) contextValidateIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Igroups); i++ {

		if m.Igroups[i] != nil {
			if err := m.Igroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroup" + "." + "igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) contextValidateInitiators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Initiators); i++ {

		if m.Initiators[i] != nil {
			if err := m.Initiators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroup" + "." + "initiators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupLunMapInlineIgroup) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "igroup"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunMapInlineIgroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunMapInlineIgroup) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunMapInlineIgroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupLunMapIgroupIgroupsItems0 consistency group lun map igroup igroups items0
//
// swagger:model ConsistencyGroupLunMapIgroupIgroupsItems0
type ConsistencyGroupLunMapIgroupIgroupsItems0 struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The name of the initiator group.
	//
	// Example: igroup1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group lun map igroup igroups items0
func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) validateLinks(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consistency group lun map igroup igroups items0 based on the context it is used
func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunMapIgroupIgroupsItems0) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunMapIgroupIgroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupLunMapIgroupInitiatorsItems0 The initiators that are members of the initiator group.
//
// swagger:model ConsistencyGroupLunMapIgroupInitiatorsItems0
type ConsistencyGroupLunMapIgroupInitiatorsItems0 struct {

	// A comment available for use by the administrator.
	//
	// Example: my comment
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Name of initiator that is a member of the initiator group.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	Name *string `json:"name,omitempty"`
}

// Validate validates this consistency group lun map igroup initiators items0
func (m *ConsistencyGroupLunMapIgroupInitiatorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupLunMapIgroupInitiatorsItems0) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 254); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group lun map igroup initiators items0 based on context it is used
func (m *ConsistencyGroupLunMapIgroupInitiatorsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupLunMapIgroupInitiatorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupLunMapIgroupInitiatorsItems0) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupLunMapIgroupInitiatorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
