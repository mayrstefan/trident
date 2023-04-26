// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EffectivePermissionResponse effective permission response
//
// swagger:model effective_permission_response
type EffectivePermissionResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// effective permission response inline records
	EffectivePermissionResponseInlineRecords []*EffectivePermission `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this effective permission response
func (m *EffectivePermissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectivePermissionResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EffectivePermissionResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *EffectivePermissionResponse) validateEffectivePermissionResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectivePermissionResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.EffectivePermissionResponseInlineRecords); i++ {
		if swag.IsZero(m.EffectivePermissionResponseInlineRecords[i]) { // not required
			continue
		}

		if m.EffectivePermissionResponseInlineRecords[i] != nil {
			if err := m.EffectivePermissionResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this effective permission response based on the context it is used
func (m *EffectivePermissionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEffectivePermissionResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EffectivePermissionResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EffectivePermissionResponse) contextValidateEffectivePermissionResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EffectivePermissionResponseInlineRecords); i++ {

		if m.EffectivePermissionResponseInlineRecords[i] != nil {
			if err := m.EffectivePermissionResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *EffectivePermissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EffectivePermissionResponse) UnmarshalBinary(b []byte) error {
	var res EffectivePermissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
