// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostReportNotifications Create a ReportNotification.
//
// swagger:model postReportNotifications
type PostReportNotifications struct {

	// The type of change the ReportNotification is tracking. Possible values: percentage, dollars.
	// Required: true
	Change *string `json:"change"`

	// The CostReport token.
	// Required: true
	CostReportToken *string `json:"cost_report_token"`

	// The frequency the ReportNotification is sent. Possible values: daily, weekly, monthly.
	// Required: true
	Frequency *string `json:"frequency"`

	// The Slack or Microsoft Teams channels that receive the notification.
	RecipientChannels []string `json:"recipient_channels"`

	// The title of the ReportNotification.
	// Required: true
	Title *string `json:"title"`

	// The Users that receive the notification.
	UserTokens []string `json:"user_tokens"`

	// The token of the Workspace to add the ReportNotification to. Required if the API token is associated with multiple Workspaces.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this post report notifications
func (m *PostReportNotifications) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostReportToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostReportNotifications) validateChange(formats strfmt.Registry) error {

	if err := validate.Required("change", "body", m.Change); err != nil {
		return err
	}

	return nil
}

func (m *PostReportNotifications) validateCostReportToken(formats strfmt.Registry) error {

	if err := validate.Required("cost_report_token", "body", m.CostReportToken); err != nil {
		return err
	}

	return nil
}

func (m *PostReportNotifications) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

func (m *PostReportNotifications) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post report notifications based on context it is used
func (m *PostReportNotifications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostReportNotifications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostReportNotifications) UnmarshalBinary(b []byte) error {
	var res PostReportNotifications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
