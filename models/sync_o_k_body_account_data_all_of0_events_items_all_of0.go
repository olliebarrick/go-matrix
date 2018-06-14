// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SyncOKBodyAccountDataAllOf0EventsItemsAllOf0 Event
// swagger:model syncOKBodyAccountDataAllOf0EventsItemsAllOf0
type SyncOKBodyAccountDataAllOf0EventsItemsAllOf0 struct {

	// EventContent
	//
	// The content of this event. The fields in this object will vary depending on the type of event.
	Content interface{} `json:"content,omitempty"`

	// The ID of this event, if applicable.
	EventID string `json:"event_id,omitempty"`

	// Timestamp in milliseconds on originating homeserver when this event was sent.
	OriginServerTs int64 `json:"origin_server_ts,omitempty"`

	// The MXID of the user who sent this event.
	Sender string `json:"sender,omitempty"`

	// Optional. This key will only be present for state events. A unique key which defines the overwriting semantics for this piece of room state.
	StateKey string `json:"state_key,omitempty"`

	// The type of event.
	Type string `json:"type,omitempty"`

	// unsigned
	Unsigned *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned `json:"unsigned,omitempty"`
}

// Validate validates this sync o k body account data all of0 events items all of0
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnsigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0) validateUnsigned(formats strfmt.Registry) error {

	if swag.IsZero(m.Unsigned) { // not required
		return nil
	}

	if m.Unsigned != nil {
		if err := m.Unsigned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unsigned")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyAccountDataAllOf0EventsItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
