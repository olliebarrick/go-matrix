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

// GetLoginFlowsTooManyRequestsBody The rate limit was reached for this request
// swagger:model getLoginFlowsTooManyRequestsBody
type GetLoginFlowsTooManyRequestsBody struct {

	// The M_LIMIT_EXCEEDED error code
	// Required: true
	Errcode *string `json:"errcode"`

	// A human-readable error message.
	Error string `json:"error,omitempty"`

	// The amount of time in milliseconds the client should wait
	// before trying the request again.
	RetryAfterMs int64 `json:"retry_after_ms,omitempty"`
}

// Validate validates this get login flows too many requests body
func (m *GetLoginFlowsTooManyRequestsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLoginFlowsTooManyRequestsBody) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetLoginFlowsTooManyRequestsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLoginFlowsTooManyRequestsBody) UnmarshalBinary(b []byte) error {
	var res GetLoginFlowsTooManyRequestsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
