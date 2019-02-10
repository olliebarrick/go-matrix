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

// SetReadMarkerParamsBody set read marker params body
// swagger:model setReadMarkerParamsBody
type SetReadMarkerParamsBody struct {

	// The event ID the read marker should be located at. The
	// event MUST belong to the room.
	// Required: true
	MFullyRead *string `json:"m.fully_read"`

	// The event ID to set the read receipt location at. This is
	// equivalent to calling ``/receipt/m.read/$elsewhere:domain.com``
	// and is provided here to save that extra call.
	MRead string `json:"m.read,omitempty"`
}

// Validate validates this set read marker params body
func (m *SetReadMarkerParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMFullyRead(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetReadMarkerParamsBody) validateMFullyRead(formats strfmt.Registry) error {

	if err := validate.Required("m.fully_read", "body", m.MFullyRead); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetReadMarkerParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetReadMarkerParamsBody) UnmarshalBinary(b []byte) error {
	var res SetReadMarkerParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
