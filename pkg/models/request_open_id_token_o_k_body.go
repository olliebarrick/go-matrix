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

// RequestOpenIDTokenOKBody request open Id token o k body
// swagger:model requestOpenIdTokenOKBody
type RequestOpenIDTokenOKBody struct {

	// An access token the consumer may use to verify the identity of
	// the person who generated the token. This is given to the federation
	// API ``GET /openid/userinfo``.
	// Required: true
	AccessToken *string `json:"access_token"`

	// The number of seconds before this token expires and a new one must
	// be generated.
	// Required: true
	ExpiresIn *int64 `json:"expires_in"`

	// The homeserver domain the consumer should use when attempting to
	// verify the user's identity.
	// Required: true
	MatrixServerName *string `json:"matrix_server_name"`

	// The string ``Bearer``.
	// Required: true
	TokenType *string `json:"token_type"`
}

// Validate validates this request open Id token o k body
func (m *RequestOpenIDTokenOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatrixServerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestOpenIDTokenOKBody) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("access_token", "body", m.AccessToken); err != nil {
		return err
	}

	return nil
}

func (m *RequestOpenIDTokenOKBody) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expires_in", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *RequestOpenIDTokenOKBody) validateMatrixServerName(formats strfmt.Registry) error {

	if err := validate.Required("matrix_server_name", "body", m.MatrixServerName); err != nil {
		return err
	}

	return nil
}

func (m *RequestOpenIDTokenOKBody) validateTokenType(formats strfmt.Registry) error {

	if err := validate.Required("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestOpenIDTokenOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestOpenIDTokenOKBody) UnmarshalBinary(b []byte) error {
	var res RequestOpenIDTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
