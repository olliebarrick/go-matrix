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

// JoinRoomParamsBodyThirdPartySigned ThirdPartySigned
//
// A signature of an ``m.third_party_invite`` token to prove that this user owns a third party identity which has been invited to the room.
// swagger:model joinRoomParamsBodyThirdPartySigned
type JoinRoomParamsBodyThirdPartySigned struct {

	// signed
	// Required: true
	Signed *JoinRoomParamsBodyThirdPartySignedSigned `json:"signed"`
}

// Validate validates this join room params body third party signed
func (m *JoinRoomParamsBodyThirdPartySigned) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JoinRoomParamsBodyThirdPartySigned) validateSigned(formats strfmt.Registry) error {

	if err := validate.Required("signed", "body", m.Signed); err != nil {
		return err
	}

	if m.Signed != nil {
		if err := m.Signed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JoinRoomParamsBodyThirdPartySigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JoinRoomParamsBodyThirdPartySigned) UnmarshalBinary(b []byte) error {
	var res JoinRoomParamsBodyThirdPartySigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
