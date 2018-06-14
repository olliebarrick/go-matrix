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

// GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0 Event
//
// The basic set of fields all events must have.
// swagger:model getEventContextOKBodyEventsBeforeItemsAllOf0AllOf0
type GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0 struct {

	// The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Content interface{} `json:"content,omitempty"`

	// The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this get event context o k body events before items all of0 all of0
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res GetEventContextOKBodyEventsBeforeItemsAllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
