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

// GetPushRulesOKBody get push rules o k body
// swagger:model getPushRulesOKBody
type GetPushRulesOKBody struct {

	// global
	// Required: true
	Global *GetPushRulesOKBodyGlobal `json:"global"`
}

// Validate validates this get push rules o k body
func (m *GetPushRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPushRulesOKBody) validateGlobal(formats strfmt.Registry) error {

	if err := validate.Required("global", "body", m.Global); err != nil {
		return err
	}

	if m.Global != nil {
		if err := m.Global.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPushRulesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPushRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPushRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
