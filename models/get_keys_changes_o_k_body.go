// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetKeysChangesOKBody get keys changes o k body
// swagger:model getKeysChangesOKBody
type GetKeysChangesOKBody struct {

	// The Matrix User IDs of all users who updated their device
	// identity keys.
	Changes []string `json:"changes"`
}

// Validate validates this get keys changes o k body
func (m *GetKeysChangesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetKeysChangesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKeysChangesOKBody) UnmarshalBinary(b []byte) error {
	var res GetKeysChangesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
