// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetNotificationsOKBodyNotificationsItemsEvent Event
//
// The Event object for the event that triggered the notification.
// swagger:model getNotificationsOKBodyNotificationsItemsEvent
type GetNotificationsOKBodyNotificationsItemsEvent struct {
	GetNotificationsOKBodyNotificationsItemsEventAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetNotificationsOKBodyNotificationsItemsEvent) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetNotificationsOKBodyNotificationsItemsEventAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetNotificationsOKBodyNotificationsItemsEventAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetNotificationsOKBodyNotificationsItemsEvent) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.GetNotificationsOKBodyNotificationsItemsEventAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get notifications o k body notifications items event
func (m *GetNotificationsOKBodyNotificationsItemsEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetNotificationsOKBodyNotificationsItemsEventAllOf0
	if err := m.GetNotificationsOKBodyNotificationsItemsEventAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetNotificationsOKBodyNotificationsItemsEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNotificationsOKBodyNotificationsItemsEvent) UnmarshalBinary(b []byte) error {
	var res GetNotificationsOKBodyNotificationsItemsEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
