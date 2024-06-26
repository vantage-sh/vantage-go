// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateAnomalyNotification Update an Anomaly Notification.
//
// swagger:model updateAnomalyNotification
type UpdateAnomalyNotification struct {

	// The Slack/MS Teams channels that receive the notification.
	RecipientChannels []string `json:"recipient_channels"`

	// The threshold amount that must be met for the notification to fire.
	Threshold int32 `json:"threshold,omitempty"`

	// The tokens of the users that receive the notification.
	UserTokens []string `json:"user_tokens"`
}

// Validate validates this update anomaly notification
func (m *UpdateAnomalyNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update anomaly notification based on context it is used
func (m *UpdateAnomalyNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAnomalyNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAnomalyNotification) UnmarshalBinary(b []byte) error {
	var res UpdateAnomalyNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
