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

// CreateDashboard Create a Dashboard.
//
// swagger:model createDashboard
type CreateDashboard struct {

	// Determines how to group costs in the Dashboard.
	// Enum: ["cumulative","day","week","month"]
	DateBin string `json:"date_bin,omitempty"`

	// Determines the date range in the Dashboard. Incompatible with 'start_date' and 'end_date' parameters.
	// Enum: ["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date"]
	DateInterval string `json:"date_interval,omitempty"`

	// The end date for the date range for costs in the Dashboard. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	// Required: true
	EndDate *string `json:"end_date"`

	// The tokens of the Saved Filters used in the Dashboard.
	SavedFilterTokens []string `json:"saved_filter_tokens"`

	// The start date for the date range for costs in the Dashboard. ISO 8601 Formatted. Incompatible with 'date_interval' parameter.
	StartDate string `json:"start_date,omitempty"`

	// The title of the Dashboard.
	// Required: true
	Title *string `json:"title"`

	// The tokens of the widgets to add to the Dashboard. Currently only supports CostReport tokens.
	WidgetTokens []string `json:"widget_tokens"`

	// The token of the Workspace to add the Dashboard to. Required if the API token is associated with multiple Workspaces.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this create dashboard
func (m *CreateDashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateBin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
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

var createDashboardTypeDateBinPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cumulative","day","week","month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDashboardTypeDateBinPropEnum = append(createDashboardTypeDateBinPropEnum, v)
	}
}

const (

	// CreateDashboardDateBinCumulative captures enum value "cumulative"
	CreateDashboardDateBinCumulative string = "cumulative"

	// CreateDashboardDateBinDay captures enum value "day"
	CreateDashboardDateBinDay string = "day"

	// CreateDashboardDateBinWeek captures enum value "week"
	CreateDashboardDateBinWeek string = "week"

	// CreateDashboardDateBinMonth captures enum value "month"
	CreateDashboardDateBinMonth string = "month"
)

// prop value enum
func (m *CreateDashboard) validateDateBinEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDashboardTypeDateBinPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateDashboard) validateDateBin(formats strfmt.Registry) error {
	if swag.IsZero(m.DateBin) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateBinEnum("date_bin", "body", m.DateBin); err != nil {
		return err
	}

	return nil
}

var createDashboardTypeDateIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["this_month","last_7_days","last_30_days","last_month","last_3_months","last_6_months","custom","last_12_months","last_24_months","last_36_months","next_month","next_3_months","next_6_months","next_12_months","year_to_date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDashboardTypeDateIntervalPropEnum = append(createDashboardTypeDateIntervalPropEnum, v)
	}
}

const (

	// CreateDashboardDateIntervalThisMonth captures enum value "this_month"
	CreateDashboardDateIntervalThisMonth string = "this_month"

	// CreateDashboardDateIntervalLast7Days captures enum value "last_7_days"
	CreateDashboardDateIntervalLast7Days string = "last_7_days"

	// CreateDashboardDateIntervalLast30Days captures enum value "last_30_days"
	CreateDashboardDateIntervalLast30Days string = "last_30_days"

	// CreateDashboardDateIntervalLastMonth captures enum value "last_month"
	CreateDashboardDateIntervalLastMonth string = "last_month"

	// CreateDashboardDateIntervalLast3Months captures enum value "last_3_months"
	CreateDashboardDateIntervalLast3Months string = "last_3_months"

	// CreateDashboardDateIntervalLast6Months captures enum value "last_6_months"
	CreateDashboardDateIntervalLast6Months string = "last_6_months"

	// CreateDashboardDateIntervalCustom captures enum value "custom"
	CreateDashboardDateIntervalCustom string = "custom"

	// CreateDashboardDateIntervalLast12Months captures enum value "last_12_months"
	CreateDashboardDateIntervalLast12Months string = "last_12_months"

	// CreateDashboardDateIntervalLast24Months captures enum value "last_24_months"
	CreateDashboardDateIntervalLast24Months string = "last_24_months"

	// CreateDashboardDateIntervalLast36Months captures enum value "last_36_months"
	CreateDashboardDateIntervalLast36Months string = "last_36_months"

	// CreateDashboardDateIntervalNextMonth captures enum value "next_month"
	CreateDashboardDateIntervalNextMonth string = "next_month"

	// CreateDashboardDateIntervalNext3Months captures enum value "next_3_months"
	CreateDashboardDateIntervalNext3Months string = "next_3_months"

	// CreateDashboardDateIntervalNext6Months captures enum value "next_6_months"
	CreateDashboardDateIntervalNext6Months string = "next_6_months"

	// CreateDashboardDateIntervalNext12Months captures enum value "next_12_months"
	CreateDashboardDateIntervalNext12Months string = "next_12_months"

	// CreateDashboardDateIntervalYearToDate captures enum value "year_to_date"
	CreateDashboardDateIntervalYearToDate string = "year_to_date"
)

// prop value enum
func (m *CreateDashboard) validateDateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDashboardTypeDateIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateDashboard) validateDateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.DateInterval) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateIntervalEnum("date_interval", "body", m.DateInterval); err != nil {
		return err
	}

	return nil
}

func (m *CreateDashboard) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("end_date", "body", m.EndDate); err != nil {
		return err
	}

	return nil
}

func (m *CreateDashboard) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create dashboard based on context it is used
func (m *CreateDashboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDashboard) UnmarshalBinary(b []byte) error {
	var res CreateDashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
