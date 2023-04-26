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

// FirmwareHistory firmware history
//
// swagger:model firmware_history
type FirmwareHistory struct {

	// links
	Links *FirmwareHistoryInlineLinks `json:"_links,omitempty"`

	// End time of this update request.
	// Example: 2019-02-02T19:00:00Z
	// Read Only: true
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// firmware history inline update status
	FirmwareHistoryInlineUpdateStatus []*FirmwareHistoryUpdateState `json:"update_status,omitempty"`

	// Name of the firmware file.
	// Example: all_disk_fw.zip
	FwFileName *string `json:"fw_file_name,omitempty"`

	// fw update state
	// Read Only: true
	// Enum: [downloading moving_firmware starting_workers waiting_on_workers completed failed]
	FwUpdateState *string `json:"fw_update_state,omitempty"`

	// job
	Job *JobLink `json:"job,omitempty"`

	// node
	Node *FirmwareHistoryInlineNode `json:"node,omitempty"`

	// Start time of this update request.
	// Example: 2019-02-02T19:00:00Z
	// Read Only: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this firmware history
func (m *FirmwareHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmwareHistoryInlineUpdateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFwUpdateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistory) validateLinks(formats strfmt.Registry) error {
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

func (m *FirmwareHistory) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareHistory) validateFirmwareHistoryInlineUpdateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.FirmwareHistoryInlineUpdateStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.FirmwareHistoryInlineUpdateStatus); i++ {
		if swag.IsZero(m.FirmwareHistoryInlineUpdateStatus[i]) { // not required
			continue
		}

		if m.FirmwareHistoryInlineUpdateStatus[i] != nil {
			if err := m.FirmwareHistoryInlineUpdateStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var firmwareHistoryTypeFwUpdateStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["downloading","moving_firmware","starting_workers","waiting_on_workers","completed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		firmwareHistoryTypeFwUpdateStatePropEnum = append(firmwareHistoryTypeFwUpdateStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// downloading
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateDownloading captures enum value "downloading"
	FirmwareHistoryFwUpdateStateDownloading string = "downloading"

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// moving_firmware
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateMovingFirmware captures enum value "moving_firmware"
	FirmwareHistoryFwUpdateStateMovingFirmware string = "moving_firmware"

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// starting_workers
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateStartingWorkers captures enum value "starting_workers"
	FirmwareHistoryFwUpdateStateStartingWorkers string = "starting_workers"

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// waiting_on_workers
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateWaitingOnWorkers captures enum value "waiting_on_workers"
	FirmwareHistoryFwUpdateStateWaitingOnWorkers string = "waiting_on_workers"

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// completed
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateCompleted captures enum value "completed"
	FirmwareHistoryFwUpdateStateCompleted string = "completed"

	// BEGIN DEBUGGING
	// firmware_history
	// FirmwareHistory
	// fw_update_state
	// FwUpdateState
	// failed
	// END DEBUGGING
	// FirmwareHistoryFwUpdateStateFailed captures enum value "failed"
	FirmwareHistoryFwUpdateStateFailed string = "failed"
)

// prop value enum
func (m *FirmwareHistory) validateFwUpdateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, firmwareHistoryTypeFwUpdateStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FirmwareHistory) validateFwUpdateState(formats strfmt.Registry) error {
	if swag.IsZero(m.FwUpdateState) { // not required
		return nil
	}

	// value enum
	if err := m.validateFwUpdateStateEnum("fw_update_state", "body", *m.FwUpdateState); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareHistory) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareHistory) validateNode(formats strfmt.Registry) error {
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

func (m *FirmwareHistory) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this firmware history based on the context it is used
func (m *FirmwareHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmwareHistoryInlineUpdateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFwUpdateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistory) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FirmwareHistory) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "end_time", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareHistory) contextValidateFirmwareHistoryInlineUpdateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirmwareHistoryInlineUpdateStatus); i++ {

		if m.FirmwareHistoryInlineUpdateStatus[i] != nil {
			if err := m.FirmwareHistoryInlineUpdateStatus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update_status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FirmwareHistory) contextValidateFwUpdateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fw_update_state", "body", m.FwUpdateState); err != nil {
		return err
	}

	return nil
}

func (m *FirmwareHistory) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

func (m *FirmwareHistory) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FirmwareHistory) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHistory) UnmarshalBinary(b []byte) error {
	var res FirmwareHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FirmwareHistoryInlineLinks firmware history inline links
//
// swagger:model firmware_history_inline__links
type FirmwareHistoryInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this firmware history inline links
func (m *FirmwareHistoryInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this firmware history inline links based on the context it is used
func (m *FirmwareHistoryInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FirmwareHistoryInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHistoryInlineLinks) UnmarshalBinary(b []byte) error {
	var res FirmwareHistoryInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FirmwareHistoryInlineNode firmware history inline node
//
// swagger:model firmware_history_inline_node
type FirmwareHistoryInlineNode struct {

	// links
	Links *FirmwareHistoryInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this firmware history inline node
func (m *FirmwareHistoryInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this firmware history inline node based on the context it is used
func (m *FirmwareHistoryInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareHistoryInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHistoryInlineNode) UnmarshalBinary(b []byte) error {
	var res FirmwareHistoryInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FirmwareHistoryInlineNodeInlineLinks firmware history inline node inline links
//
// swagger:model firmware_history_inline_node_inline__links
type FirmwareHistoryInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this firmware history inline node inline links
func (m *FirmwareHistoryInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this firmware history inline node inline links based on the context it is used
func (m *FirmwareHistoryInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirmwareHistoryInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirmwareHistoryInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirmwareHistoryInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res FirmwareHistoryInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
