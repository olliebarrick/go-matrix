// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned Unsigned
//
// Information about this event which was not sent by the originating homeserver
// swagger:model syncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned
type SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned struct {

	// Time in milliseconds since the event was sent.
	Age int64 `json:"age,omitempty"`

	// EventContent
	//
	// Optional. The previous ``content`` for this state. This will be present only for state events appearing in the ``timeline``. If this is not a state event, or there is no previous content, this key will be missing.
	PrevContent interface{} `json:"prev_content,omitempty"`

	// Event
	//
	// Optional. The event that redacted this event, if any.
	RedactedBecause interface{} `json:"redacted_because,omitempty"`

	// Optional. The transaction ID set when this message was sent. This key will only be present for message events sent by the device calling this API.
	TransactionID string `json:"transaction_id,omitempty"`
}

// Validate validates this sync o k body account data all of0 events items all of0 unsigned
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyAccountDataAllOf0EventsItemsAllOf0Unsigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
