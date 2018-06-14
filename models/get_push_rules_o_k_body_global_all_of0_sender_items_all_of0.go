// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0 PushRule
// swagger:model getPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0
type GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0 struct {

	// The actions to perform when this rule is matched.
	// Required: true
	Actions []interface{} `json:"actions"`

	// The conditions that must hold true for an event in order for a rule to be
	// applied to an event. A rule with no conditions always matches. Only
	// applicable to ``underride`` and ``override`` rules.
	Conditions []*GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0ConditionsItems `json:"conditions"`

	// Whether this is a default rule, or has been set explicitly.
	// Required: true
	Default *bool `json:"default"`

	// Whether the push rule is enabled or not.
	// Required: true
	Enabled *bool `json:"enabled"`

	// The glob-style pattern to match against.  Only applicable to ``content``
	// rules.
	Pattern string `json:"pattern,omitempty"`

	// The ID of this rule.
	// Required: true
	RuleID *string `json:"rule_id"`
}

// Validate validates this get push rules o k body global all of0 sender items all of0
func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	return nil
}

func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("rule_id", "body", m.RuleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res GetPushRulesOKBodyGlobalAllOf0SenderItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
