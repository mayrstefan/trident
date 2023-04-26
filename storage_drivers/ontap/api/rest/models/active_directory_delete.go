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

// ActiveDirectoryDelete active directory delete
//
// swagger:model active_directory_delete
type ActiveDirectoryDelete struct {

	// Administrator password required for Active Directory account creation.
	// Example: testpwd
	// Min Length: 1
	Password *string `json:"password,omitempty"`

	// Administrator username required for Active Directory account creation.
	// Example: admin
	// Min Length: 1
	Username *string `json:"username,omitempty"`
}

// Validate validates this active directory delete
func (m *ActiveDirectoryDelete) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryDelete) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", *m.Password, 1); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectoryDelete) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MinLength("username", "body", *m.Username, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this active directory delete based on context it is used
func (m *ActiveDirectoryDelete) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryDelete) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryDelete) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryDelete
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
