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

// SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0 Room Event
//
// In addition to the Event fields, Room Events have the following additional fields.
// swagger:model syncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0
type SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0 struct {
	SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0

	// The globally unique event identifier.
	// Required: true
	EventID *string `json:"event_id"`

	// Timestamp in milliseconds on originating homeserver when this event was sent.
	// Required: true
	OriginServerTs *int64 `json:"origin_server_ts"`

	// Contains the fully-qualified ID of the user who sent this event.
	// Required: true
	Sender *string `json:"sender"`

	// unsigned
	Unsigned *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0Unsigned `json:"unsigned,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sync o k body rooms leave additional properties state all of0 events items all of0 all of0
func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0
	if err := m.SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginServerTs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnsigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("event_id", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) validateOriginServerTs(formats strfmt.Registry) error {

	if err := validate.Required("origin_server_ts", "body", m.OriginServerTs); err != nil {
		return err
	}

	return nil
}

func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) validateUnsigned(formats strfmt.Registry) error {

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
func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res SyncOKBodyRoomsLeaveAdditionalPropertiesStateAllOf0EventsItemsAllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
