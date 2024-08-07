// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportNotification ReportNotification model
//
// swagger:model ReportNotification
type ReportNotification struct {

	// The type of change the ReportNotification is tracking.
	// Example: percentage
	// Enum: ["percentage","dollars"]
	Change string `json:"change,omitempty"`

	// The token for a CostReport the ReportNotification is applied to.
	// Example: rprt_abcd1234
	CostReportToken string `json:"cost_report_token,omitempty"`

	// The frequency the ReportNotification is sent.
	// Example: weekly
	// Enum: ["daily","weekly","monthly"]
	Frequency string `json:"frequency,omitempty"`

	// The Slack or Microsoft Teams channels that receive the notification.
	RecipientChannels []string `json:"recipient_channels"`

	// The title of the ReportNotification.
	// Example: Acme Report Notification
	Title string `json:"title,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The Users that receive the notification.
	UserTokens []string `json:"user_tokens"`
}

// Validate validates this report notification
func (m *ReportNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportNotificationTypeChangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["percentage","dollars"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportNotificationTypeChangePropEnum = append(reportNotificationTypeChangePropEnum, v)
	}
}

const (

	// ReportNotificationChangePercentage captures enum value "percentage"
	ReportNotificationChangePercentage string = "percentage"

	// ReportNotificationChangeDollars captures enum value "dollars"
	ReportNotificationChangeDollars string = "dollars"
)

// prop value enum
func (m *ReportNotification) validateChangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportNotificationTypeChangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportNotification) validateChange(formats strfmt.Registry) error {
	if swag.IsZero(m.Change) { // not required
		return nil
	}

	// value enum
	if err := m.validateChangeEnum("change", "body", m.Change); err != nil {
		return err
	}

	return nil
}

var reportNotificationTypeFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["daily","weekly","monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportNotificationTypeFrequencyPropEnum = append(reportNotificationTypeFrequencyPropEnum, v)
	}
}

const (

	// ReportNotificationFrequencyDaily captures enum value "daily"
	ReportNotificationFrequencyDaily string = "daily"

	// ReportNotificationFrequencyWeekly captures enum value "weekly"
	ReportNotificationFrequencyWeekly string = "weekly"

	// ReportNotificationFrequencyMonthly captures enum value "monthly"
	ReportNotificationFrequencyMonthly string = "monthly"
)

// prop value enum
func (m *ReportNotification) validateFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportNotificationTypeFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportNotification) validateFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateFrequencyEnum("frequency", "body", m.Frequency); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report notification based on context it is used
func (m *ReportNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportNotification) UnmarshalBinary(b []byte) error {
	var res ReportNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
