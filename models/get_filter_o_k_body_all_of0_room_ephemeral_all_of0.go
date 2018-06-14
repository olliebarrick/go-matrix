// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetFilterOKBodyAllOf0RoomEphemeralAllOf0 RoomEventFilter
// swagger:model getFilterOKBodyAllOf0RoomEphemeralAllOf0
type GetFilterOKBodyAllOf0RoomEphemeralAllOf0 struct {
	GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0

	// If ``true``, includes only events with a url key in their content. If ``false``, excludes those events.
	ContainsURL bool `json:"contains_url,omitempty"`

	// A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the ``'rooms'`` filter.
	NotRooms []string `json:"not_rooms"`

	// A list of room IDs to include. If this list is absent then all rooms are included.
	Rooms []string `json:"rooms"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetFilterOKBodyAllOf0RoomEphemeralAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetFilterOKBodyAllOf0RoomEphemeralAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get filter o k body all of0 room ephemeral all of0
func (m *GetFilterOKBodyAllOf0RoomEphemeralAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0
	if err := m.GetFilterOKBodyAllOf0RoomEphemeralAllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetFilterOKBodyAllOf0RoomEphemeralAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFilterOKBodyAllOf0RoomEphemeralAllOf0) UnmarshalBinary(b []byte) error {
	var res GetFilterOKBodyAllOf0RoomEphemeralAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
