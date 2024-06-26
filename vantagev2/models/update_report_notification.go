// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateReportNotification Update a ReportNotification.
//
// swagger:model updateReportNotification
type UpdateReportNotification struct {

	// The type of change the ReportNotification is tracking. Possible values: percentage, dollars.
	Change string `json:"change,omitempty"`

	// The CostReport token.
	CostReportToken string `json:"cost_report_token,omitempty"`

	// The frequency the ReportNotification is sent. Possible values: daily, weekly, monthly.
	Frequency string `json:"frequency,omitempty"`

	// The Slack or Microsoft Teams channels that receive the notification.
	RecipientChannels []string `json:"recipient_channels"`

	// The title of the ReportNotification.
	Title string `json:"title,omitempty"`

	// The Users that receive the notification.
	UserTokens []string `json:"user_tokens"`
}

// Validate validates this update report notification
func (m *UpdateReportNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update report notification based on context it is used
func (m *UpdateReportNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateReportNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateReportNotification) UnmarshalBinary(b []byte) error {
	var res UpdateReportNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
