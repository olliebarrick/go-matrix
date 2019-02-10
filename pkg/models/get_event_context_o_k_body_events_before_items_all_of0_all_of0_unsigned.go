// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned UnsignedData
//
// Contains optional extra information about the event.
// swagger:model getEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned
type GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned struct {

	// The time in milliseconds that has elapsed since the event was sent. This field is generated by the local homeserver, and may be incorrect if the local time on at least one of the two servers is out of sync, which can cause the age to either be negative or greater than it actually is.
	Age int64 `json:"age,omitempty"`

	// Event
	//
	// Optional. The event that redacted this event, if any.
	RedactedBecause interface{} `json:"redacted_because,omitempty"`

	// The client-supplied transaction ID, if the client being given the event is the same one which sent it.
	TransactionID string `json:"transaction_id,omitempty"`
}

// Validate validates this get event context o k body events before items all of0 all of0 unsigned
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned) UnmarshalBinary(b []byte) error {
	var res GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0Unsigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
