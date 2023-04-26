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

// ConsistencyGroupSnapshotResponse consistency group snapshot response
//
// swagger:model consistency_group_snapshot_response
type ConsistencyGroupSnapshotResponse struct {

	// links
	Links *CollectionLinks `json:"_links,omitempty"`

	// consistency group snapshot response inline records
	ConsistencyGroupSnapshotResponseInlineRecords []*ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this consistency group snapshot response
func (m *ConsistencyGroupSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroupSnapshotResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotResponse) validateConsistencyGroupSnapshotResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroupSnapshotResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistencyGroupSnapshotResponseInlineRecords); i++ {
		if swag.IsZero(m.ConsistencyGroupSnapshotResponseInlineRecords[i]) { // not required
			continue
		}

		if m.ConsistencyGroupSnapshotResponseInlineRecords[i] != nil {
			if err := m.ConsistencyGroupSnapshotResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this consistency group snapshot response based on the context it is used
func (m *ConsistencyGroupSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroupSnapshotResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotResponse) contextValidateConsistencyGroupSnapshotResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsistencyGroupSnapshotResponseInlineRecords); i++ {

		if m.ConsistencyGroupSnapshotResponseInlineRecords[i] != nil {
			if err := m.ConsistencyGroupSnapshotResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *ConsistencyGroupSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem consistency group snapshot response inline records inline array item
//
// swagger:model consistency_group_snapshot_response_inline_records_inline_array_item
type ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Comment for the Snapshot copy.
	//
	// Example: My Snapshot copy comment
	Comment *string `json:"comment,omitempty"`

	// consistency group
	ConsistencyGroup *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup `json:"consistency_group,omitempty"`

	// Consistency type. This is for categorization purposes only. A Snapshot copy should not be set to 'application consistent' unless the host application is quiesced for the Snapshot copy. Valid in POST.
	//
	// Example: crash
	// Enum: [crash application]
	ConsistencyType *string `json:"consistency_type,omitempty"`

	// Time the snapshot copy was created
	//
	// Example: 2020-10-25T11:20:00Z
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// Indicates whether the Snapshot copy taken is partial or not.
	//
	// Example: false
	IsPartial *bool `json:"is_partial,omitempty"`

	// List of volumes which are not in the Snapshot copy.
	//
	MissingVolumes []*VolumeReference `json:"missing_volumes"`

	// Name of the Snapshot copy.
	//
	Name *string `json:"name,omitempty"`

	// Snapmirror Label for the Snapshot copy.
	//
	// Example: sm_label
	SnapmirrorLabel *string `json:"snapmirror_label,omitempty"`

	// List of volume and snapshot identifiers for each volume in the Snapshot copy.
	//
	// Read Only: true
	SnapshotVolumes []*ConsistencyGroupVolumeSnapshot `json:"snapshot_volumes"`

	// The SVM in which the consistency group is located.
	//
	Svm *SvmReference `json:"svm,omitempty"`

	// The unique identifier of the Snapshot copy. The UUID is generated
	// by ONTAP when the Snapshot copy is created.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group snapshot response inline records inline array item
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissingVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateConsistencyGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroup) { // not required
		return nil
	}

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

var consistencyGroupSnapshotResponseInlineRecordsInlineArrayItemTypeConsistencyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["crash","application"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consistencyGroupSnapshotResponseInlineRecordsInlineArrayItemTypeConsistencyTypePropEnum = append(consistencyGroupSnapshotResponseInlineRecordsInlineArrayItemTypeConsistencyTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// consistency_group_snapshot_response_inline_records_inline_array_item
	// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem
	// consistency_type
	// ConsistencyType
	// crash
	// END DEBUGGING
	// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemConsistencyTypeCrash captures enum value "crash"
	ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemConsistencyTypeCrash string = "crash"

	// BEGIN DEBUGGING
	// consistency_group_snapshot_response_inline_records_inline_array_item
	// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem
	// consistency_type
	// ConsistencyType
	// application
	// END DEBUGGING
	// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemConsistencyTypeApplication captures enum value "application"
	ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemConsistencyTypeApplication string = "application"
)

// prop value enum
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateConsistencyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consistencyGroupSnapshotResponseInlineRecordsInlineArrayItemTypeConsistencyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateConsistencyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConsistencyTypeEnum("consistency_type", "body", *m.ConsistencyType); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateMissingVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.MissingVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.MissingVolumes); i++ {
		if swag.IsZero(m.MissingVolumes[i]) { // not required
			continue
		}

		if m.MissingVolumes[i] != nil {
			if err := m.MissingVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateSnapshotVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotVolumes); i++ {
		if swag.IsZero(m.SnapshotVolumes[i]) { // not required
			continue
		}

		if m.SnapshotVolumes[i] != nil {
			if err := m.SnapshotVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group snapshot response inline records inline array item based on the context it is used
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMissingVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) contextValidateConsistencyGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) contextValidateMissingVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MissingVolumes); i++ {

		if m.MissingVolumes[i] != nil {
			if err := m.MissingVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("missing_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) contextValidateSnapshotVolumes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snapshot_volumes", "body", []*ConsistencyGroupVolumeSnapshot(m.SnapshotVolumes)); err != nil {
		return err
	}

	for i := 0; i < len(m.SnapshotVolumes); i++ {

		if m.SnapshotVolumes[i] != nil {
			if err := m.SnapshotVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup The consistency group of the Snapshot copy.
//
// swagger:model consistency_group_snapshot_response_inline_records_inline_array_item_inline_consistency_group
type ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The name of the consistency group.
	// Example: my_consistency_group
	Name *string `json:"name,omitempty"`

	// The unique identifier of the consistency group.
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this consistency group snapshot response inline records inline array item inline consistency group
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group snapshot response inline records inline array item inline consistency group based on the context it is used
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotResponseInlineRecordsInlineArrayItemInlineConsistencyGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
