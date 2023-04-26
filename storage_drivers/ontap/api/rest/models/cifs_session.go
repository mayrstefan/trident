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

// CifsSession cifs session
//
// swagger:model cifs_session
type CifsSession struct {

	// links
	Links *CifsSessionInlineLinks `json:"_links,omitempty"`

	// SMB authentication over which the client accesses the share. The following values are supported:
	// * none - No authentication
	// * ntlmv1 - Ntlm version 1 mechanism
	// * ntlmv2 - Ntlm version 2 mechanism
	// * kerberos - Kerberos authentication
	// * anonymous - Anonymous mechanism
	//
	// Example: ntlmv2
	// Read Only: true
	// Enum: [none ntlmv1 ntlmv2 kerberos anonymous]
	Authentication *string `json:"authentication,omitempty"`

	// A group of volumes, the client is accessing.
	CifsSessionInlineVolumes []*CifsSessionInlineVolumesInlineArrayItem `json:"volumes,omitempty"`

	// Specifies IP address of the client.
	//
	// Example: 10.74.7.182
	// Read Only: true
	ClientIP *string `json:"client_ip,omitempty"`

	// Specifies an ISO-8601 format of date and time used to retrieve the connected time duration in hours, minutes, and seconds format.
	//
	// Example: P4DT84H30M5S
	// Read Only: true
	ConnectedDuration *string `json:"connected_duration,omitempty"`

	// A counter used to track requests that are sent to the volumes to the node.
	//
	// Example: 0
	// Read Only: true
	ConnectionCount *int64 `json:"connection_count,omitempty"`

	// A unique 32-bit unsigned number used to represent each SMB session's connection ID.
	//
	// Example: 22802
	// Read Only: true
	ConnectionID *int64 `json:"connection_id,omitempty"`

	// The level of continuous availabilty protection provided to the SMB sessions and/or files.
	// * unavailable - Open file is not continuously available. For sessions, it contains one or more open files but none of them are continuously available.
	// * available - open file is continuously available. For sessions, it contains one or more open files and all of them are continuously available.
	// * partial - Sessions only. Contains at least one continuously available open file with other files open but not continuously available.
	//
	// Example: unavailable
	// Read Only: true
	// Enum: [unavailable available partial]
	ContinuousAvailability *string `json:"continuous_availability,omitempty"`

	// A unique 64-bit unsigned number used to represent each SMB session's identifier.
	//
	// Example: 4622663542519103000
	// Read Only: true
	Identifier *int64 `json:"identifier,omitempty"`

	// Specifies an ISO-8601 format of date and time used to retrieve the idle time duration in hours, minutes, and seconds format.
	//
	// Example: P4DT84H30M5S
	// Read Only: true
	IdleDuration *string `json:"idle_duration,omitempty"`

	// Specifies whether the large MTU is enabled or not for an SMB session.
	//
	// Example: true
	// Read Only: true
	LargeMtu *bool `json:"large_mtu,omitempty"`

	// Indicated that a mapped UNIX user has logged in.
	//
	// Example: root
	// Read Only: true
	MappedUnixUser *string `json:"mapped_unix_user,omitempty"`

	// node
	Node *CifsSessionInlineNode `json:"node,omitempty"`

	// Number of files the SMB session has opened.
	//
	// Read Only: true
	OpenFiles *int64 `json:"open_files,omitempty"`

	// Number of other files the SMB session has opened.
	//
	// Read Only: true
	OpenOther *int64 `json:"open_other,omitempty"`

	// Number of shares the SMB session has opened.
	//
	// Read Only: true
	OpenShares *int64 `json:"open_shares,omitempty"`

	// The SMB protocol version over which the client accesses the volumes. The following values are supported:
	// * smb1 - SMB version 1
	// * smb2 - SMB version 2
	// * smb2_1 - SMB version 2 minor version 1
	// * smb3 - SMB version 3
	// * smb3_1 - SMB version 3 minor version 1
	//
	// Example: smb3_1
	// Read Only: true
	// Enum: [smb1 smb2 smb2_1 smb3 smb3_1]
	Protocol *string `json:"protocol,omitempty"`

	// Specifies the IP address of the SVM.
	//
	// Example: 10.140.78.248
	// Read Only: true
	ServerIP *string `json:"server_ip,omitempty"`

	// Indicates an SMB encryption state. The following values are supported:
	// * unencrypted - SMB session is not encrypted
	// * encrypted - SMB session is fully encrypted. SVM level encryption is enabled and encryption occurs for the entire session.
	// * partially_encrypted - SMB session is partially encrypted. Share level encryption is enabled and encryption is initiated when the tree-connect occurs.
	//
	// Example: unencrypted
	// Read Only: true
	// Enum: [unencrypted encrypted partially_encrypted]
	SmbEncryption *string `json:"smb_encryption,omitempty"`

	// Specifies whether or not SMB signing is enabled.
	// Example: false
	// Read Only: true
	SmbSigning *bool `json:"smb_signing,omitempty"`

	// svm
	Svm *CifsSessionInlineSvm `json:"svm,omitempty"`

	// Indicates that a Windows user has logged in.
	//
	// Example: NBCIFSQA2\\administrator
	// Read Only: true
	User *string `json:"user,omitempty"`
}

// Validate validates this cifs session
func (m *CifsSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCifsSessionInlineVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContinuousAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmbEncryption(formats); err != nil {
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

func (m *CifsSession) validateLinks(formats strfmt.Registry) error {
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

var cifsSessionTypeAuthenticationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","ntlmv1","ntlmv2","kerberos","anonymous"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsSessionTypeAuthenticationPropEnum = append(cifsSessionTypeAuthenticationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// authentication
	// Authentication
	// none
	// END DEBUGGING
	// CifsSessionAuthenticationNone captures enum value "none"
	CifsSessionAuthenticationNone string = "none"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// authentication
	// Authentication
	// ntlmv1
	// END DEBUGGING
	// CifsSessionAuthenticationNtlmv1 captures enum value "ntlmv1"
	CifsSessionAuthenticationNtlmv1 string = "ntlmv1"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// authentication
	// Authentication
	// ntlmv2
	// END DEBUGGING
	// CifsSessionAuthenticationNtlmv2 captures enum value "ntlmv2"
	CifsSessionAuthenticationNtlmv2 string = "ntlmv2"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// authentication
	// Authentication
	// kerberos
	// END DEBUGGING
	// CifsSessionAuthenticationKerberos captures enum value "kerberos"
	CifsSessionAuthenticationKerberos string = "kerberos"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// authentication
	// Authentication
	// anonymous
	// END DEBUGGING
	// CifsSessionAuthenticationAnonymous captures enum value "anonymous"
	CifsSessionAuthenticationAnonymous string = "anonymous"
)

// prop value enum
func (m *CifsSession) validateAuthenticationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsSessionTypeAuthenticationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsSession) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationEnum("authentication", "body", *m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) validateCifsSessionInlineVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.CifsSessionInlineVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.CifsSessionInlineVolumes); i++ {
		if swag.IsZero(m.CifsSessionInlineVolumes[i]) { // not required
			continue
		}

		if m.CifsSessionInlineVolumes[i] != nil {
			if err := m.CifsSessionInlineVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var cifsSessionTypeContinuousAvailabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unavailable","available","partial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsSessionTypeContinuousAvailabilityPropEnum = append(cifsSessionTypeContinuousAvailabilityPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// continuous_availability
	// ContinuousAvailability
	// unavailable
	// END DEBUGGING
	// CifsSessionContinuousAvailabilityUnavailable captures enum value "unavailable"
	CifsSessionContinuousAvailabilityUnavailable string = "unavailable"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// continuous_availability
	// ContinuousAvailability
	// available
	// END DEBUGGING
	// CifsSessionContinuousAvailabilityAvailable captures enum value "available"
	CifsSessionContinuousAvailabilityAvailable string = "available"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// continuous_availability
	// ContinuousAvailability
	// partial
	// END DEBUGGING
	// CifsSessionContinuousAvailabilityPartial captures enum value "partial"
	CifsSessionContinuousAvailabilityPartial string = "partial"
)

// prop value enum
func (m *CifsSession) validateContinuousAvailabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsSessionTypeContinuousAvailabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsSession) validateContinuousAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.ContinuousAvailability) { // not required
		return nil
	}

	// value enum
	if err := m.validateContinuousAvailabilityEnum("continuous_availability", "body", *m.ContinuousAvailability); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) validateNode(formats strfmt.Registry) error {
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

var cifsSessionTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smb1","smb2","smb2_1","smb3","smb3_1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsSessionTypeProtocolPropEnum = append(cifsSessionTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// protocol
	// Protocol
	// smb1
	// END DEBUGGING
	// CifsSessionProtocolSmb1 captures enum value "smb1"
	CifsSessionProtocolSmb1 string = "smb1"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// protocol
	// Protocol
	// smb2
	// END DEBUGGING
	// CifsSessionProtocolSmb2 captures enum value "smb2"
	CifsSessionProtocolSmb2 string = "smb2"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// protocol
	// Protocol
	// smb2_1
	// END DEBUGGING
	// CifsSessionProtocolSmb21 captures enum value "smb2_1"
	CifsSessionProtocolSmb21 string = "smb2_1"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// protocol
	// Protocol
	// smb3
	// END DEBUGGING
	// CifsSessionProtocolSmb3 captures enum value "smb3"
	CifsSessionProtocolSmb3 string = "smb3"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// protocol
	// Protocol
	// smb3_1
	// END DEBUGGING
	// CifsSessionProtocolSmb31 captures enum value "smb3_1"
	CifsSessionProtocolSmb31 string = "smb3_1"
)

// prop value enum
func (m *CifsSession) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsSessionTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsSession) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

var cifsSessionTypeSmbEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unencrypted","encrypted","partially_encrypted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cifsSessionTypeSmbEncryptionPropEnum = append(cifsSessionTypeSmbEncryptionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// smb_encryption
	// SmbEncryption
	// unencrypted
	// END DEBUGGING
	// CifsSessionSmbEncryptionUnencrypted captures enum value "unencrypted"
	CifsSessionSmbEncryptionUnencrypted string = "unencrypted"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// smb_encryption
	// SmbEncryption
	// encrypted
	// END DEBUGGING
	// CifsSessionSmbEncryptionEncrypted captures enum value "encrypted"
	CifsSessionSmbEncryptionEncrypted string = "encrypted"

	// BEGIN DEBUGGING
	// cifs_session
	// CifsSession
	// smb_encryption
	// SmbEncryption
	// partially_encrypted
	// END DEBUGGING
	// CifsSessionSmbEncryptionPartiallyEncrypted captures enum value "partially_encrypted"
	CifsSessionSmbEncryptionPartiallyEncrypted string = "partially_encrypted"
)

// prop value enum
func (m *CifsSession) validateSmbEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cifsSessionTypeSmbEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CifsSession) validateSmbEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.SmbEncryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateSmbEncryptionEnum("smb_encryption", "body", *m.SmbEncryption); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session based on the context it is used
func (m *CifsSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCifsSessionInlineVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContinuousAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdleDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLargeMtu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMappedUnixUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenOther(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenShares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmbEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmbSigning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSession) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsSession) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "authentication", "body", m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateCifsSessionInlineVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CifsSessionInlineVolumes); i++ {

		if m.CifsSessionInlineVolumes[i] != nil {
			if err := m.CifsSessionInlineVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CifsSession) contextValidateClientIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "client_ip", "body", m.ClientIP); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateConnectedDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_duration", "body", m.ConnectedDuration); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateConnectionCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connection_count", "body", m.ConnectionCount); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateConnectionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connection_id", "body", m.ConnectionID); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateContinuousAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "continuous_availability", "body", m.ContinuousAvailability); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateIdleDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "idle_duration", "body", m.IdleDuration); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateLargeMtu(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "large_mtu", "body", m.LargeMtu); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateMappedUnixUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mapped_unix_user", "body", m.MappedUnixUser); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsSession) contextValidateOpenFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "open_files", "body", m.OpenFiles); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateOpenOther(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "open_other", "body", m.OpenOther); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateOpenShares(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "open_shares", "body", m.OpenShares); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateServerIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "server_ip", "body", m.ServerIP); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateSmbEncryption(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "smb_encryption", "body", m.SmbEncryption); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateSmbSigning(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "smb_signing", "body", m.SmbSigning); err != nil {
		return err
	}

	return nil
}

func (m *CifsSession) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CifsSession) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSession) UnmarshalBinary(b []byte) error {
	var res CifsSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineLinks cifs session inline links
//
// swagger:model cifs_session_inline__links
type CifsSessionInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs session inline links
func (m *CifsSessionInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session inline links based on the context it is used
func (m *CifsSessionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsSessionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineNode cifs session inline node
//
// swagger:model cifs_session_inline_node
type CifsSessionInlineNode struct {

	// links
	Links *CifsSessionInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs session inline node
func (m *CifsSessionInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session inline node based on the context it is used
func (m *CifsSessionInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsSessionInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineNode) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineNodeInlineLinks cifs session inline node inline links
//
// swagger:model cifs_session_inline_node_inline__links
type CifsSessionInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs session inline node inline links
func (m *CifsSessionInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session inline node inline links based on the context it is used
func (m *CifsSessionInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsSessionInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineSvm cifs session inline svm
//
// swagger:model cifs_session_inline_svm
type CifsSessionInlineSvm struct {

	// links
	Links *CifsSessionInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs session inline svm
func (m *CifsSessionInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs session inline svm based on the context it is used
func (m *CifsSessionInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsSessionInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineSvm) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineSvmInlineLinks cifs session inline svm inline links
//
// swagger:model cifs_session_inline_svm_inline__links
type CifsSessionInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs session inline svm inline links
func (m *CifsSessionInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cifs session inline svm inline links based on the context it is used
func (m *CifsSessionInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CifsSessionInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineVolumesInlineArrayItem cifs session inline volumes inline array item
//
// swagger:model cifs_session_inline_volumes_inline_array_item
type CifsSessionInlineVolumesInlineArrayItem struct {

	// links
	Links *CifsSessionInlineVolumesInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cifs session inline volumes inline array item
func (m *CifsSessionInlineVolumesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineVolumesInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session inline volumes inline array item based on the context it is used
func (m *CifsSessionInlineVolumesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineVolumesInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsSessionInlineVolumesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineVolumesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineVolumesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CifsSessionInlineVolumesInlineArrayItemInlineLinks cifs session inline volumes inline array item inline links
//
// swagger:model cifs_session_inline_volumes_inline_array_item_inline__links
type CifsSessionInlineVolumesInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this cifs session inline volumes inline array item inline links
func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this cifs session inline volumes inline array item inline links based on the context it is used
func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsSessionInlineVolumesInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res CifsSessionInlineVolumesInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
