// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned UnsignedData
//
// Contains optional extra information about the event.
// swagger:model initialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned
type InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned struct {

	// A subset of the state of the room at the time of the invite, if ``membership`` is ``invite``. Note that this state is informational, and SHOULD NOT be trusted; once the client has joined the room, it SHOULD fetch the live state from the server and discard the invite_room_state. Also, clients must not rely on any particular state being present here; they SHOULD behave properly (with possibly a degraded but not a broken experience) in the absence of any particular events here. If they are set on the room, at least the state for ``m.room.avatar``, ``m.room.canonical_alias``, ``m.room.join_rules``, and ``m.room.name`` SHOULD be included.
	InviteRoomState []*InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsignedInviteRoomStateItems `json:"invite_room_state"`
}

// Validate validates this initial sync o k body rooms items invite all of0 content unsigned
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInviteRoomState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned) validateInviteRoomState(formats strfmt.Registry) error {

	if swag.IsZero(m.InviteRoomState) { // not required
		return nil
	}

	for i := 0; i < len(m.InviteRoomState); i++ {
		if swag.IsZero(m.InviteRoomState[i]) { // not required
			continue
		}

		if m.InviteRoomState[i] != nil {
			if err := m.InviteRoomState[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invite_room_state" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned) UnmarshalBinary(b []byte) error {
	var res InitialSyncOKBodyRoomsItemsInviteAllOf0ContentUnsigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
