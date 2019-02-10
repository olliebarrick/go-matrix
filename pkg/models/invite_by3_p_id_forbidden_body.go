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

// InviteBy3PIDForbiddenBody A Matrix-level Error
// swagger:model inviteBy3PIdForbiddenBody
type InviteBy3PIDForbiddenBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this invite by3 p Id forbidden body
func (m *InviteBy3PIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InviteBy3PIDForbiddenBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InviteBy3PIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InviteBy3PIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res InviteBy3PIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
