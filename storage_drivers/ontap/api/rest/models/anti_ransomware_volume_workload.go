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
	"github.com/go-openapi/validate"
)

// AntiRansomwareVolumeWorkload anti ransomware volume workload
//
// swagger:model anti_ransomware_volume_workload
type AntiRansomwareVolumeWorkload struct {

	// File extensions observed in the volume.
	// Example: ["pdf","jpeg","txt"]
	// Read Only: true
	AntiRansomwareVolumeWorkloadInlineFileExtensionsObserved []*string `json:"file_extensions_observed,omitempty"`

	// Count of types of file extensions observed in the volume.
	// Example: 3
	// Read Only: true
	FileExtensionTypesCount *int64 `json:"file_extension_types_count,omitempty"`

	// surge usage
	SurgeUsage *AntiRansomwareVolumeWorkloadInlineSurgeUsage `json:"surge_usage,omitempty"`

	// typical usage
	TypicalUsage *AntiRansomwareVolumeWorkloadInlineTypicalUsage `json:"typical_usage,omitempty"`
}

// Validate validates this anti ransomware volume workload
func (m *AntiRansomwareVolumeWorkload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSurgeUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypicalUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeWorkload) validateSurgeUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.SurgeUsage) { // not required
		return nil
	}

	if m.SurgeUsage != nil {
		if err := m.SurgeUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surge_usage")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkload) validateTypicalUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.TypicalUsage) { // not required
		return nil
	}

	if m.TypicalUsage != nil {
		if err := m.TypicalUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typical_usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this anti ransomware volume workload based on the context it is used
func (m *AntiRansomwareVolumeWorkload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntiRansomwareVolumeWorkloadInlineFileExtensionsObserved(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileExtensionTypesCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSurgeUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypicalUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeWorkload) contextValidateAntiRansomwareVolumeWorkloadInlineFileExtensionsObserved(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file_extensions_observed", "body", []*string(m.AntiRansomwareVolumeWorkloadInlineFileExtensionsObserved)); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkload) contextValidateFileExtensionTypesCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "file_extension_types_count", "body", m.FileExtensionTypesCount); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkload) contextValidateSurgeUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.SurgeUsage != nil {
		if err := m.SurgeUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surge_usage")
			}
			return err
		}
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkload) contextValidateTypicalUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.TypicalUsage != nil {
		if err := m.TypicalUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typical_usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkload) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeWorkload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeWorkloadInlineSurgeUsage Usage values of the volume's workload during surge.
//
// swagger:model anti_ransomware_volume_workload_inline_surge_usage
type AntiRansomwareVolumeWorkloadInlineSurgeUsage struct {

	// Peak rate of file creates per minute in the workload of the volume during surge.
	// Example: 10
	// Read Only: true
	FileCreatePeakRatePerMinute *int64 `json:"file_create_peak_rate_per_minute,omitempty"`

	// Peak rate of file deletes per minute in the workload of the volume during surge.
	// Example: 50
	// Read Only: true
	FileDeletePeakRatePerMinute *int64 `json:"file_delete_peak_rate_per_minute,omitempty"`

	// Peak rate of file renames per minute in the workload of the volume during surge.
	// Example: 30
	// Read Only: true
	FileRenamePeakRatePerMinute *int64 `json:"file_rename_peak_rate_per_minute,omitempty"`

	// Peak percentage of high entropy data writes in the volume during surge.
	// Example: 30
	// Read Only: true
	HighEntropyDataWritePeakPercent *int64 `json:"high_entropy_data_write_peak_percent,omitempty"`

	// Peak high entropy data write rate in the volume during surge, in KBs per minute.
	// Example: 2500
	// Read Only: true
	HighEntropyDataWritePeakRateKbPerMinute *int64 `json:"high_entropy_data_write_peak_rate_kb_per_minute,omitempty"`

	// New file extensions observed in the volume during surge.
	// Read Only: true
	NewlyObservedFileExtensions []*AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0 `json:"newly_observed_file_extensions,omitempty"`

	// Timestamp at which the first surge in the volume's workload is observed.
	// Example: 2021-12-01T23:16:20+05:30
	// Read Only: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this anti ransomware volume workload inline surge usage
func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewlyObservedFileExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) validateNewlyObservedFileExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.NewlyObservedFileExtensions) { // not required
		return nil
	}

	for i := 0; i < len(m.NewlyObservedFileExtensions); i++ {
		if swag.IsZero(m.NewlyObservedFileExtensions[i]) { // not required
			continue
		}

		if m.NewlyObservedFileExtensions[i] != nil {
			if err := m.NewlyObservedFileExtensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surge_usage" + "." + "newly_observed_file_extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("surge_usage"+"."+"time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this anti ransomware volume workload inline surge usage based on the context it is used
func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileCreatePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileDeletePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileRenamePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHighEntropyDataWritePeakPercent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHighEntropyDataWritePeakRateKbPerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewlyObservedFileExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateFileCreatePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"file_create_peak_rate_per_minute", "body", m.FileCreatePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateFileDeletePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"file_delete_peak_rate_per_minute", "body", m.FileDeletePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateFileRenamePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"file_rename_peak_rate_per_minute", "body", m.FileRenamePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateHighEntropyDataWritePeakPercent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"high_entropy_data_write_peak_percent", "body", m.HighEntropyDataWritePeakPercent); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateHighEntropyDataWritePeakRateKbPerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"high_entropy_data_write_peak_rate_kb_per_minute", "body", m.HighEntropyDataWritePeakRateKbPerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateNewlyObservedFileExtensions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"newly_observed_file_extensions", "body", []*AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0(m.NewlyObservedFileExtensions)); err != nil {
		return err
	}

	for i := 0; i < len(m.NewlyObservedFileExtensions); i++ {

		if m.NewlyObservedFileExtensions[i] != nil {
			if err := m.NewlyObservedFileExtensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surge_usage" + "." + "newly_observed_file_extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "surge_usage"+"."+"time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadInlineSurgeUsage) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeWorkloadInlineSurgeUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0 anti ransomware volume workload surge usage newly observed file extensions items0
//
// swagger:model AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0
type AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0 struct {

	// Count of newly observed file extensions.
	// Example: ["20"]
	// Read Only: true
	Count *int64 `json:"count,omitempty"`

	// Name of the newly observed file extension.
	// Example: ["lockile"]
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this anti ransomware volume workload surge usage newly observed file extensions items0
func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this anti ransomware volume workload surge usage newly observed file extensions items0 based on the context it is used
func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCount(ctx, formats); err != nil {
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

func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) contextValidateCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeWorkloadSurgeUsageNewlyObservedFileExtensionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AntiRansomwareVolumeWorkloadInlineTypicalUsage Typical usage values of volume workload.
//
// swagger:model anti_ransomware_volume_workload_inline_typical_usage
type AntiRansomwareVolumeWorkloadInlineTypicalUsage struct {

	// Typical peak rate of file creates per minute in the workload of the volume.
	// Example: 50
	// Read Only: true
	FileCreatePeakRatePerMinute *int64 `json:"file_create_peak_rate_per_minute,omitempty"`

	// Typical peak rate of file deletes per minute in the workload of the volume.
	// Example: 10
	// Read Only: true
	FileDeletePeakRatePerMinute *int64 `json:"file_delete_peak_rate_per_minute,omitempty"`

	// Typical peak rate of file renames per minute in the workload of the volume.
	// Example: 5
	// Read Only: true
	FileRenamePeakRatePerMinute *int64 `json:"file_rename_peak_rate_per_minute,omitempty"`

	// Typical peak percentage of high entropy data writes in the volume.
	// Example: 10
	// Read Only: true
	HighEntropyDataWritePeakPercent *int64 `json:"high_entropy_data_write_peak_percent,omitempty"`

	// Typical peak high entropy data write rate in the volume, in KBs per minute.
	// Example: 1200
	// Read Only: true
	HighEntropyDataWritePeakRateKbPerMinute *int64 `json:"high_entropy_data_write_peak_rate_kb_per_minute,omitempty"`
}

// Validate validates this anti ransomware volume workload inline typical usage
func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this anti ransomware volume workload inline typical usage based on the context it is used
func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileCreatePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileDeletePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileRenamePeakRatePerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHighEntropyDataWritePeakPercent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHighEntropyDataWritePeakRateKbPerMinute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) contextValidateFileCreatePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "typical_usage"+"."+"file_create_peak_rate_per_minute", "body", m.FileCreatePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) contextValidateFileDeletePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "typical_usage"+"."+"file_delete_peak_rate_per_minute", "body", m.FileDeletePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) contextValidateFileRenamePeakRatePerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "typical_usage"+"."+"file_rename_peak_rate_per_minute", "body", m.FileRenamePeakRatePerMinute); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) contextValidateHighEntropyDataWritePeakPercent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "typical_usage"+"."+"high_entropy_data_write_peak_percent", "body", m.HighEntropyDataWritePeakPercent); err != nil {
		return err
	}

	return nil
}

func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) contextValidateHighEntropyDataWritePeakRateKbPerMinute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "typical_usage"+"."+"high_entropy_data_write_peak_rate_kb_per_minute", "body", m.HighEntropyDataWritePeakRateKbPerMinute); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntiRansomwareVolumeWorkloadInlineTypicalUsage) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareVolumeWorkloadInlineTypicalUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
