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

// SoftwareMccReference software mcc reference
//
// swagger:model software_mcc_reference
type SoftwareMccReference struct {

	// Elapsed duration of update time (in seconds) of MetroCluster.
	// Example: 2140
	// Read Only: true
	ElapsedDuration *int64 `json:"elapsed_duration,omitempty"`

	// Estimated duration of update time (in seconds) of MetroCluster.
	// Example: 3480
	// Read Only: true
	EstimatedDuration *int64 `json:"estimated_duration,omitempty"`

	// Name of the site in MetroCluster.
	// Example: cluster_A
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Upgrade state of MetroCluster.
	// Example: in_progress
	// Read Only: true
	// Enum: [in_progress waiting paused_by_user paused_on_error completed canceled failed pause_pending cancel_pending]
	State interface{} `json:"state,omitempty"`
}

// Validate validates this software mcc reference
func (m *SoftwareMccReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software mcc reference based on the context it is used
func (m *SoftwareMccReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElapsedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEstimatedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareMccReference) contextValidateElapsedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "elapsed_duration", "body", m.ElapsedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareMccReference) contextValidateEstimatedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "estimated_duration", "body", m.EstimatedDuration); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareMccReference) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareMccReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareMccReference) UnmarshalBinary(b []byte) error {
	var res SoftwareMccReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
