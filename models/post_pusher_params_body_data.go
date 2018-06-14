// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostPusherParamsBodyData PusherData
//
// A dictionary of information for the pusher implementation
// itself. If ``kind`` is ``http``, this should contain ``url``
// which is the URL to use to send notifications to.
// swagger:model postPusherParamsBodyData
type PostPusherParamsBodyData struct {

	// Required if ``kind`` is ``http``. The URL to use to send
	// notifications to.
	URL string `json:"url,omitempty"`
}

// Validate validates this post pusher params body data
func (m *PostPusherParamsBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostPusherParamsBodyData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostPusherParamsBodyData) UnmarshalBinary(b []byte) error {
	var res PostPusherParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
