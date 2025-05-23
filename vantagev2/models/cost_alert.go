// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CostAlert CostAlert model
//
// swagger:model CostAlert
type CostAlert struct {

	// The date and time, in UTC, for when the alert was created. ISO 8601 Formatted.
	// Example: 2023-10-01T12:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The email addresses that will receive the alert.
	EmailRecipients []string `json:"email_recipients"`

	// The period of time used to compare costs. Options are 'day', 'week', 'month', 'quarter'.
	Interval string `json:"interval,omitempty"`

	// The tokens of the reports to alert on.
	ReportTokens []string `json:"report_tokens"`

	// The Slack channels that will receive the alert. Make sure your slack integration is connected at https://console.vantage.sh/settings/slack.
	SlackChannels []string `json:"slack_channels"`

	// The Microsoft Teams channels that will receive the alert. Make sure your teams integration is connected at https://console.vantage.sh/settings/microsoft_teams.
	TeamsChannels []string `json:"teams_channels"`

	// The cost change threshold to alert on.
	Threshold float64 `json:"threshold,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The unit type used to compare costs. Options are 'currency' or 'percentage'.
	UnitType string `json:"unit_type,omitempty"`

	// The date and time, in UTC, for when the alert was last updated. ISO 8601 Formatted.
	// Example: 2023-10-01T12:00:00Z
	UpdatedAt string `json:"updated_at,omitempty"`

	// The ID of the organization that owns the CostAlert.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this cost alert
func (m *CostAlert) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cost alert based on context it is used
func (m *CostAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CostAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostAlert) UnmarshalBinary(b []byte) error {
	var res CostAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
