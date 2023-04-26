// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Snmp Cluster-wide SNMP configuration.
//
// swagger:model snmp
type Snmp struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Specifies whether to enable or disable SNMP authentication traps.
	// Example: true
	AuthTrapsEnabled *bool `json:"auth_traps_enabled,omitempty"`

	// Specifies whether to enable or disable SNMP.
	// Example: true
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies whether to enable or disable SNMP traps.
	// Example: true
	TrapsEnabled *bool `json:"traps_enabled,omitempty"`

	// Trigger a test SNMP trap.
	// Example: true
	TriggerTestTrap *bool `json:"trigger_test_trap,omitempty"`
}

// Validate validates this snmp
func (m *Snmp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snmp) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this snmp based on the context it is used
func (m *Snmp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snmp) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Snmp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snmp) UnmarshalBinary(b []byte) error {
	var res Snmp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
