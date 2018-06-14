// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SendToDeviceParamsBody body
// swagger:model sendToDeviceParamsBody
type SendToDeviceParamsBody struct {

	// The messages to send. A map from user ID, to a map from
	// device ID to message body. The device ID may also be `*`,
	// meaning all known devices for the user.
	Messages map[string]map[string]interface{} `json:"messages,omitempty"`
}

// Validate validates this send to device params body
func (m *SendToDeviceParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SendToDeviceParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendToDeviceParamsBody) UnmarshalBinary(b []byte) error {
	var res SendToDeviceParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
