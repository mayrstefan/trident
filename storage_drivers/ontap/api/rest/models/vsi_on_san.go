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

// VsiOnSan A VSI application using SAN.
//
// swagger:model vsi_on_san
type VsiOnSan struct {

	// datastore
	// Required: true
	Datastore *VsiOnSanInlineDatastore `json:"datastore"`

	// The name of the hypervisor hosting the application.
	// Required: true
	// Enum: [hyper_v vmware xen]
	Hypervisor *string `json:"hypervisor"`

	// The name of the initiator group through which the contents of this application will be accessed. Modification of this parameter is a disruptive operation. All LUNs in the application component will be unmapped from the current igroup and re-mapped to the new igroup.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	IgroupName *string `json:"igroup_name"`

	// protection type
	ProtectionType *VsiOnSanInlineProtectionType `json:"protection_type,omitempty"`

	// The list of initiator groups to create.
	// Max Items: 1
	// Min Items: 0
	VsiOnSanInlineNewIgroups []*VsiOnSanNewIgroups `json:"new_igroups,omitempty"`
}

// Validate validates this vsi on san
func (m *VsiOnSan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervisor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsiOnSanInlineNewIgroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnSan) validateDatastore(formats strfmt.Registry) error {

	if err := validate.Required("datastore", "body", m.Datastore); err != nil {
		return err
	}

	if m.Datastore != nil {
		if err := m.Datastore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore")
			}
			return err
		}
	}

	return nil
}

var vsiOnSanTypeHypervisorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hyper_v","vmware","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnSanTypeHypervisorPropEnum = append(vsiOnSanTypeHypervisorPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// vsi_on_san
	// VsiOnSan
	// hypervisor
	// Hypervisor
	// hyper_v
	// END DEBUGGING
	// VsiOnSanHypervisorHyperv captures enum value "hyper_v"
	VsiOnSanHypervisorHyperv string = "hyper_v"

	// BEGIN DEBUGGING
	// vsi_on_san
	// VsiOnSan
	// hypervisor
	// Hypervisor
	// vmware
	// END DEBUGGING
	// VsiOnSanHypervisorVmware captures enum value "vmware"
	VsiOnSanHypervisorVmware string = "vmware"

	// BEGIN DEBUGGING
	// vsi_on_san
	// VsiOnSan
	// hypervisor
	// Hypervisor
	// xen
	// END DEBUGGING
	// VsiOnSanHypervisorXen captures enum value "xen"
	VsiOnSanHypervisorXen string = "xen"
)

// prop value enum
func (m *VsiOnSan) validateHypervisorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnSanTypeHypervisorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnSan) validateHypervisor(formats strfmt.Registry) error {

	if err := validate.Required("hypervisor", "body", m.Hypervisor); err != nil {
		return err
	}

	// value enum
	if err := m.validateHypervisorEnum("hypervisor", "body", *m.Hypervisor); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnSan) validateIgroupName(formats strfmt.Registry) error {

	if err := validate.Required("igroup_name", "body", m.IgroupName); err != nil {
		return err
	}

	if err := validate.MinLength("igroup_name", "body", *m.IgroupName, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup_name", "body", *m.IgroupName, 96); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnSan) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnSan) validateVsiOnSanInlineNewIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.VsiOnSanInlineNewIgroups) { // not required
		return nil
	}

	iVsiOnSanInlineNewIgroupsSize := int64(len(m.VsiOnSanInlineNewIgroups))

	if err := validate.MinItems("new_igroups", "body", iVsiOnSanInlineNewIgroupsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("new_igroups", "body", iVsiOnSanInlineNewIgroupsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.VsiOnSanInlineNewIgroups); i++ {
		if swag.IsZero(m.VsiOnSanInlineNewIgroups[i]) { // not required
			continue
		}

		if m.VsiOnSanInlineNewIgroups[i] != nil {
			if err := m.VsiOnSanInlineNewIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vsi on san based on the context it is used
func (m *VsiOnSan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsiOnSanInlineNewIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnSan) contextValidateDatastore(ctx context.Context, formats strfmt.Registry) error {

	if m.Datastore != nil {
		if err := m.Datastore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnSan) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {
		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protection_type")
			}
			return err
		}
	}

	return nil
}

func (m *VsiOnSan) contextValidateVsiOnSanInlineNewIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VsiOnSanInlineNewIgroups); i++ {

		if m.VsiOnSanInlineNewIgroups[i] != nil {
			if err := m.VsiOnSanInlineNewIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new_igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnSan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnSan) UnmarshalBinary(b []byte) error {
	var res VsiOnSan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnSanInlineDatastore vsi on san inline datastore
//
// swagger:model vsi_on_san_inline_datastore
type VsiOnSanInlineDatastore struct {

	// The number of datastores to support.
	// Maximum: 16
	// Minimum: 1
	Count *int64 `json:"count,omitempty"`

	// The size of the datastore. Usage: {&lt;integer&gt;[KB|MB|GB|TB|PB]}
	// Required: true
	Size *int64 `json:"size"`

	// storage service
	StorageService *VsiOnSanInlineDatastoreInlineStorageService `json:"storage_service,omitempty"`
}

// Validate validates this vsi on san inline datastore
func (m *VsiOnSanInlineDatastore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnSanInlineDatastore) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("datastore"+"."+"count", "body", *m.Count, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("datastore"+"."+"count", "body", *m.Count, 16, false); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnSanInlineDatastore) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("datastore"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VsiOnSanInlineDatastore) validateStorageService(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageService) { // not required
		return nil
	}

	if m.StorageService != nil {
		if err := m.StorageService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vsi on san inline datastore based on the context it is used
func (m *VsiOnSanInlineDatastore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsiOnSanInlineDatastore) contextValidateStorageService(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageService != nil {
		if err := m.StorageService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastore" + "." + "storage_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnSanInlineDatastore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnSanInlineDatastore) UnmarshalBinary(b []byte) error {
	var res VsiOnSanInlineDatastore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnSanInlineDatastoreInlineStorageService vsi on san inline datastore inline storage service
//
// swagger:model vsi_on_san_inline_datastore_inline_storage_service
type VsiOnSanInlineDatastoreInlineStorageService struct {

	// The storage service of the datastore.
	// Enum: [extreme performance value]
	Name *string `json:"name,omitempty"`
}

// Validate validates this vsi on san inline datastore inline storage service
func (m *VsiOnSanInlineDatastoreInlineStorageService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vsiOnSanInlineDatastoreInlineStorageServiceTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extreme","performance","value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnSanInlineDatastoreInlineStorageServiceTypeNamePropEnum = append(vsiOnSanInlineDatastoreInlineStorageServiceTypeNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// vsi_on_san_inline_datastore_inline_storage_service
	// VsiOnSanInlineDatastoreInlineStorageService
	// name
	// Name
	// extreme
	// END DEBUGGING
	// VsiOnSanInlineDatastoreInlineStorageServiceNameExtreme captures enum value "extreme"
	VsiOnSanInlineDatastoreInlineStorageServiceNameExtreme string = "extreme"

	// BEGIN DEBUGGING
	// vsi_on_san_inline_datastore_inline_storage_service
	// VsiOnSanInlineDatastoreInlineStorageService
	// name
	// Name
	// performance
	// END DEBUGGING
	// VsiOnSanInlineDatastoreInlineStorageServiceNamePerformance captures enum value "performance"
	VsiOnSanInlineDatastoreInlineStorageServiceNamePerformance string = "performance"

	// BEGIN DEBUGGING
	// vsi_on_san_inline_datastore_inline_storage_service
	// VsiOnSanInlineDatastoreInlineStorageService
	// name
	// Name
	// value
	// END DEBUGGING
	// VsiOnSanInlineDatastoreInlineStorageServiceNameValue captures enum value "value"
	VsiOnSanInlineDatastoreInlineStorageServiceNameValue string = "value"
)

// prop value enum
func (m *VsiOnSanInlineDatastoreInlineStorageService) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnSanInlineDatastoreInlineStorageServiceTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnSanInlineDatastoreInlineStorageService) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("datastore"+"."+"storage_service"+"."+"name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on san inline datastore inline storage service based on context it is used
func (m *VsiOnSanInlineDatastoreInlineStorageService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnSanInlineDatastoreInlineStorageService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnSanInlineDatastoreInlineStorageService) UnmarshalBinary(b []byte) error {
	var res VsiOnSanInlineDatastoreInlineStorageService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsiOnSanInlineProtectionType vsi on san inline protection type
//
// swagger:model vsi_on_san_inline_protection_type
type VsiOnSanInlineProtectionType struct {

	// The local RPO of the application.
	// Enum: [hourly none]
	LocalRpo *string `json:"local_rpo,omitempty"`

	// The remote RPO of the application.
	// Enum: [none zero]
	RemoteRpo *string `json:"remote_rpo,omitempty"`
}

// Validate validates this vsi on san inline protection type
func (m *VsiOnSanInlineProtectionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalRpo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteRpo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vsiOnSanInlineProtectionTypeTypeLocalRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnSanInlineProtectionTypeTypeLocalRpoPropEnum = append(vsiOnSanInlineProtectionTypeTypeLocalRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// vsi_on_san_inline_protection_type
	// VsiOnSanInlineProtectionType
	// local_rpo
	// LocalRpo
	// hourly
	// END DEBUGGING
	// VsiOnSanInlineProtectionTypeLocalRpoHourly captures enum value "hourly"
	VsiOnSanInlineProtectionTypeLocalRpoHourly string = "hourly"

	// BEGIN DEBUGGING
	// vsi_on_san_inline_protection_type
	// VsiOnSanInlineProtectionType
	// local_rpo
	// LocalRpo
	// none
	// END DEBUGGING
	// VsiOnSanInlineProtectionTypeLocalRpoNone captures enum value "none"
	VsiOnSanInlineProtectionTypeLocalRpoNone string = "none"
)

// prop value enum
func (m *VsiOnSanInlineProtectionType) validateLocalRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnSanInlineProtectionTypeTypeLocalRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnSanInlineProtectionType) validateLocalRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocalRpoEnum("protection_type"+"."+"local_rpo", "body", *m.LocalRpo); err != nil {
		return err
	}

	return nil
}

var vsiOnSanInlineProtectionTypeTypeRemoteRpoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","zero"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vsiOnSanInlineProtectionTypeTypeRemoteRpoPropEnum = append(vsiOnSanInlineProtectionTypeTypeRemoteRpoPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// vsi_on_san_inline_protection_type
	// VsiOnSanInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// none
	// END DEBUGGING
	// VsiOnSanInlineProtectionTypeRemoteRpoNone captures enum value "none"
	VsiOnSanInlineProtectionTypeRemoteRpoNone string = "none"

	// BEGIN DEBUGGING
	// vsi_on_san_inline_protection_type
	// VsiOnSanInlineProtectionType
	// remote_rpo
	// RemoteRpo
	// zero
	// END DEBUGGING
	// VsiOnSanInlineProtectionTypeRemoteRpoZero captures enum value "zero"
	VsiOnSanInlineProtectionTypeRemoteRpoZero string = "zero"
)

// prop value enum
func (m *VsiOnSanInlineProtectionType) validateRemoteRpoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vsiOnSanInlineProtectionTypeTypeRemoteRpoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VsiOnSanInlineProtectionType) validateRemoteRpo(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteRpo) { // not required
		return nil
	}

	// value enum
	if err := m.validateRemoteRpoEnum("protection_type"+"."+"remote_rpo", "body", *m.RemoteRpo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsi on san inline protection type based on context it is used
func (m *VsiOnSanInlineProtectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsiOnSanInlineProtectionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsiOnSanInlineProtectionType) UnmarshalBinary(b []byte) error {
	var res VsiOnSanInlineProtectionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
