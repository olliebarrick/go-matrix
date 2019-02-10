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

// SearchUserDirectoryOKBodyResultsItems User
// swagger:model searchUserDirectoryOKBodyResultsItems
type SearchUserDirectoryOKBodyResultsItems struct {

	// The avatar url, as an MXC, if one exists.
	AvatarURL string `json:"avatar_url,omitempty"`

	// The display name of the user, if one exists.
	DisplayName string `json:"display_name,omitempty"`

	// The user's matrix user ID.
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this search user directory o k body results items
func (m *SearchUserDirectoryOKBodyResultsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchUserDirectoryOKBodyResultsItems) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchUserDirectoryOKBodyResultsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchUserDirectoryOKBodyResultsItems) UnmarshalBinary(b []byte) error {
	var res SearchUserDirectoryOKBodyResultsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
