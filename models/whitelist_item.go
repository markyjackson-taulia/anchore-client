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

// WhitelistItem Identifies a specific gate and trigger match from a policy against an image and indicates it should be ignored in final policy decisions
// swagger:model WhitelistItem
type WhitelistItem struct {

	// gate
	// Required: true
	Gate *string `json:"gate"`

	// id
	ID string `json:"id,omitempty"`

	// trigger id
	// Required: true
	TriggerID *string `json:"trigger_id"`
}

// Validate validates this whitelist item
func (m *WhitelistItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTriggerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhitelistItem) validateGate(formats strfmt.Registry) error {

	if err := validate.Required("gate", "body", m.Gate); err != nil {
		return err
	}

	return nil
}

func (m *WhitelistItem) validateTriggerID(formats strfmt.Registry) error {

	if err := validate.Required("trigger_id", "body", m.TriggerID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WhitelistItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhitelistItem) UnmarshalBinary(b []byte) error {
	var res WhitelistItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}