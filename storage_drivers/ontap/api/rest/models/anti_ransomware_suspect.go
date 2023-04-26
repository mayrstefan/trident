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

// AntiRansomwareSuspect File suspected to be potentially attacked by ransomware.
//
// swagger:model anti_ransomware_suspect
type AntiRansomwareSuspect struct {

	// links
	Links *AntiRansomwareSuspectInlineLinks `json:"_links,omitempty"`

	// file
	File *AntiRansomwareSuspectInlineFile `json:"file,omitempty"`

	// Specifies whether the suspected ransomware activity is a false positive or not. This parameter is only used when making a DELETE call.
	IsFalsePositive *bool `json:"is_false_positive,omitempty"`

	// volume
	Volume *AntiRansomwareSuspectInlineVolume `json:"volume,omitempty"`
}

// Validate validates this anti ransomware suspect
func (m *AntiRansomwareSuspect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspect) validateLinks(formats strfmt.Registry) error {
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

func (m *AntiRansomwareSuspect) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareSuspect) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this anti ransomware suspect based on the context it is used
func (m *AntiRansomwareSuspect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspect) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AntiRansomwareSuspect) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if m.File != nil {
		if err := m.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareSuspect) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareSuspect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareSuspect) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareSuspectInlineFile anti ransomware suspect inline file
//
// swagger:model anti_ransomware_suspect_inline_file
type AntiRansomwareSuspectInlineFile struct {

	// File format of the suspected file.
	// Example: pdf
	Format *string `json:"format,omitempty"`

	// Name of the suspected file.
	// Example: test_file
	Name *string `json:"name,omitempty"`

	// Path of the suspected file.
	// Example: d1/d2/d3
	Path *string `json:"path,omitempty"`

	// Reason behind this file bieng suspected
	// Example: High Entropy
	Reason *string `json:"reason,omitempty"`

	// Time when the file was detected as a potential suspect in date-time format.
	// Example: 2021-05-12T11:00:16-04:00
	// Read Only: true
	// Format: date-time
	SuspectTime *strfmt.DateTime `json:"suspect_time,omitempty"`
}

// Validate validates this anti ransomware suspect inline file
func (m *AntiRansomwareSuspectInlineFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuspectTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineFile) validateSuspectTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SuspectTime) { // not required
		return nil
	}

	if err := validate.FormatOf("file"+"."+"suspect_time", "body", "date-time", m.SuspectTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this anti ransomware suspect inline file based on the context it is used
func (m *AntiRansomwareSuspectInlineFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSuspectTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineFile) contextValidateSuspectTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file"+"."+"suspect_time", "body", m.SuspectTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineFile) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspectInlineFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareSuspectInlineLinks anti ransomware suspect inline links
//
// swagger:model anti_ransomware_suspect_inline__links
type AntiRansomwareSuspectInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this anti ransomware suspect inline links
func (m *AntiRansomwareSuspectInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this anti ransomware suspect inline links based on the context it is used
func (m *AntiRansomwareSuspectInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AntiRansomwareSuspectInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineLinks) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspectInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareSuspectInlineVolume anti ransomware suspect inline volume
//
// swagger:model anti_ransomware_suspect_inline_volume
type AntiRansomwareSuspectInlineVolume struct {

	// links
	Links *AntiRansomwareSuspectInlineVolumeInlineLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this anti ransomware suspect inline volume
func (m *AntiRansomwareSuspectInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this anti ransomware suspect inline volume based on the context it is used
func (m *AntiRansomwareSuspectInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineVolume) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspectInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareSuspectInlineVolumeInlineLinks anti ransomware suspect inline volume inline links
//
// swagger:model anti_ransomware_suspect_inline_volume_inline__links
type AntiRansomwareSuspectInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this anti ransomware suspect inline volume inline links
func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this anti ransomware suspect inline volume inline links based on the context it is used
func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareSuspectInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspectInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
