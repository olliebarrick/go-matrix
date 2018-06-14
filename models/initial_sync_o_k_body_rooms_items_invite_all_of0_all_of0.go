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

// InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0 State Event
//
// In addition to the Room Event fields, State Events have the following additional fields.
// swagger:model initialSyncOKBodyRoomsItemsInviteAllOf0AllOf0
type InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0 struct {
	InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0

	// EventContent
	//
	// Optional. The previous ``content`` for this event. If there is no previous content, this key will be missing.
	PrevContent interface{} `json:"prev_content,omitempty"`

	// A unique key which defines the overwriting semantics for this piece of room state. This value is often a zero-length string. The presence of this key makes this event a State Event. The key MUST NOT start with '_'.
	// Required: true
	StateKey *string `json:"state_key"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this initial sync o k body rooms items invite all of0 all of0
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0
	if err := m.InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) validateStateKey(formats strfmt.Registry) error {

	if err := validate.Required("state_key", "body", m.StateKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res InitialSyncOKBodyRoomsItemsInviteAllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
