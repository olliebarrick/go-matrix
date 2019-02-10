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

// InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems StrippedState
//
// A stripped down state event, with only the ``type``, ``state_key`` and ``content`` keys.
// swagger:model initialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems
type InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems struct {

	// EventContent
	//
	// The ``content`` for the event.
	// Required: true
	Content interface{} `json:"content"`

	// The ``state_key`` for the event.
	// Required: true
	StateKey *string `json:"state_key"`

	// The ``type`` for the event.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this initial sync o k body rooms items invite all of0 content unsigned invite room state items
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) validateStateKey(formats strfmt.Registry) error {

	if err := validate.Required("state_key", "body", m.StateKey); err != nil {
		return err
	}

	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems) UnmarshalBinary(b []byte) error {
	var res InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
