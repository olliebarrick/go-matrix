// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UploadKeysParamsBodyDeviceKeysAllOf0 DeviceKeys
//
// Device identity keys
// swagger:model uploadKeysParamsBodyDeviceKeysAllOf0
type UploadKeysParamsBodyDeviceKeysAllOf0 struct {

	// The encryption algorithms supported by this device.
	// Required: true
	Algorithms []string `json:"algorithms"`

	// The ID of the device these keys belong to. Must match the device ID used
	// when logging in.
	// Required: true
	DeviceID *string `json:"device_id"`

	// Public identity keys. The names of the properties should be in the
	// format ``<algorithm>:<device_id>``. The keys themselves should be
	// encoded as specified by the key algorithm.
	// Required: true
	Keys map[string]string `json:"keys"`

	// Signatures for the device key object. A map from user ID, to a map from
	// ``<algorithm>:<device_id>`` to the signature.
	//
	// The signature is calculated using the process described at `Signing
	// JSON`_.
	// Required: true
	Signatures map[string]map[string]string `json:"signatures"`

	// The ID of the user the device belongs to. Must match the user ID used
	// when logging in.
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this upload keys params body device keys all of0
func (m *UploadKeysParamsBodyDeviceKeysAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadKeysParamsBodyDeviceKeysAllOf0) validateAlgorithms(formats strfmt.Registry) error {

	if err := validate.Required("algorithms", "body", m.Algorithms); err != nil {
		return err
	}

	return nil
}

func (m *UploadKeysParamsBodyDeviceKeysAllOf0) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *UploadKeysParamsBodyDeviceKeysAllOf0) validateKeys(formats strfmt.Registry) error {

	return nil
}

func (m *UploadKeysParamsBodyDeviceKeysAllOf0) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("signatures", "body", m.Signatures); err != nil {
		return err
	}

	return nil
}

func (m *UploadKeysParamsBodyDeviceKeysAllOf0) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UploadKeysParamsBodyDeviceKeysAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadKeysParamsBodyDeviceKeysAllOf0) UnmarshalBinary(b []byte) error {
	var res UploadKeysParamsBodyDeviceKeysAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
