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

// PostPusherBadRequestBody A Matrix-level Error
// swagger:model postPusherBadRequestBody
type PostPusherBadRequestBody struct {

	// An error code.
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`
}

// Validate validates this post pusher bad request body
func (m *PostPusherBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostPusherBadRequestBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostPusherBadRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostPusherBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostPusherBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
