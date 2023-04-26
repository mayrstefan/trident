// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyGroupVdiskProvisioningOptions Options that are applied to the operation.
//
// swagger:model consistency_group_vdisk_provisioning_options
type ConsistencyGroupVdiskProvisioningOptions struct {

	// Operation to perform
	// Enum: [create]
	Action *string `json:"action,omitempty"`

	// Number of elements to perform the operation on.
	Count *int64 `json:"count,omitempty"`
}

// Validate validates this consistency group vdisk provisioning options
func (m *ConsistencyGroupVdiskProvisioningOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consistencyGroupVdiskProvisioningOptionsTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupVdiskProvisioningOptionsTypeActionPropEnum = append(consistencyGroupVdiskProvisioningOptionsTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_vdisk_provisioning_options
	// ConsistencyGroupVdiskProvisioningOptions
	// action
	// Action
	// create
	// END DEBUGGING
	// ConsistencyGroupVdiskProvisioningOptionsActionCreate captures enum value "create"
	ConsistencyGroupVdiskProvisioningOptionsActionCreate string = "create"
)

// prop value enum
func (m *ConsistencyGroupVdiskProvisioningOptions) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupVdiskProvisioningOptionsTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupVdiskProvisioningOptions) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group vdisk provisioning options based on context it is used
func (m *ConsistencyGroupVdiskProvisioningOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupVdiskProvisioningOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupVdiskProvisioningOptions) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupVdiskProvisioningOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
