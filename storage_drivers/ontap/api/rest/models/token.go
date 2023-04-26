// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Token token
//
// swagger:model token
type Token struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// expiry time
	ExpiryTime *TokenInlineExpiryTime `json:"expiry_time,omitempty"`

	// node
	Node *NodeReference `json:"node,omitempty"`

	// Specifies the available reserve in the file clone split load for the given token.
	// Required: true
	ReserveSize *int64 `json:"reserve_size"`

	// Token UUID.
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReserveSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Token) validateLinks(formats strfmt.Registry) error {
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

func (m *Token) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if m.ExpiryTime != nil {
		if err := m.ExpiryTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiry_time")
			}
			return err
		}
	}

	return nil
}

func (m *Token) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *Token) validateReserveSize(formats strfmt.Registry) error {

	if err := validate.Required("reserve_size", "body", m.ReserveSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this token based on the context it is used
func (m *Token) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiryTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
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

func (m *Token) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Token) contextValidateExpiryTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ExpiryTime != nil {
		if err := m.ExpiryTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiry_time")
			}
			return err
		}
	}

	return nil
}

func (m *Token) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *Token) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TokenInlineExpiryTime token inline expiry time
//
// swagger:model token_inline_expiry_time
type TokenInlineExpiryTime struct {

	// Specifies the time remaining before the given token expires in ISO-8601 format.
	// Read Only: true
	Left *string `json:"left,omitempty"`

	// Specifies when the given token expires in ISO-8601 format.
	Limit *string `json:"limit,omitempty"`
}

// Validate validates this token inline expiry time
func (m *TokenInlineExpiryTime) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this token inline expiry time based on the context it is used
func (m *TokenInlineExpiryTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLeft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenInlineExpiryTime) contextValidateLeft(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiry_time"+"."+"left", "body", m.Left); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenInlineExpiryTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenInlineExpiryTime) UnmarshalBinary(b []byte) error {
	var res TokenInlineExpiryTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
